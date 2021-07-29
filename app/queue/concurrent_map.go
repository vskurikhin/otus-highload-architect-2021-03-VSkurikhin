package queue

import (
	"encoding/json"
	"github.com/atreugo/websocket"
	"github.com/savsgio/go-logger/v2"
	"sync"
	"unsafe"
)

var SHARD_COUNT = 32

// A "thread" safe карта типа *websocket.Conn: *BlockingQueue.
// Чтобы избежать узких мест в блокировке, эта карта разделена на несколько сегментов (SHARD_COUNT).
type ConcurrentMap []*ConcurrentMapShared

// A "thread" safe *websocket.Conn: to *BlockingQueue.
type ConcurrentMapShared struct {
	items        map[*websocket.Conn]*BlockingQueue
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

var WebSocketsMapBQueue = NewConcurrentMap()

// NewConcurrentMap Создает новую параллельную карту.
func NewConcurrentMap() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[*websocket.Conn]*BlockingQueue)}
	}
	return m
}

// GetShard возвращает shard под заданным ключом
func (m ConcurrentMap) GetShard(key *websocket.Conn) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConcurrentMap) MSet(data map[*websocket.Conn]*BlockingQueue) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Set данное значение под указанным ключом.
func (m ConcurrentMap) Set(key *websocket.Conn, value *BlockingQueue) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

//Callback для возврата нового элемента для вставки на карту
//Он вызывается при удержании блокировки, поэтому НЕ ДОЛЖЕН
//попробуйте получить доступ к другим ключам на той же карте,
//так как это может привести к тупиковой ситуации, так как
//Go sync.RWLock не повторяется
type UpsertCb func(exist bool, valueInMap *BlockingQueue, newValue *BlockingQueue) *BlockingQueue

// Upsert Insert or Update - обновляет существующий элемент или вставляет новый с помощью UpsertCb.
func (m ConcurrentMap) Upsert(key *websocket.Conn, value *BlockingQueue, cb UpsertCb) (res *BlockingQueue) {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	res = cb(ok, v, value)
	shard.items[key] = res
	shard.Unlock()
	return res
}

// SetIfAbsent Устанавливает заданное значение под указанным ключом, если с ним не было связано никакого значения.
func (m ConcurrentMap) SetIfAbsent(key *websocket.Conn, value *BlockingQueue) bool {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
	}
	shard.Unlock()
	return !ok
}

// Get извлекает элемент с карты под заданным ключом.
func (m ConcurrentMap) Get(key *websocket.Conn) (*BlockingQueue, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Count возвращает количество элементов в карте.
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has Ищет элемент по указанному ключу.
func (m ConcurrentMap) Has(key *websocket.Conn) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Remove удаляет элемент из карты.
func (m ConcurrentMap) Remove(key *websocket.Conn) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// RemoveCb is a callback executed in a map.RemoveCb() call, while Lock is held
// If returns true, the element will be removed from the map
type RemoveCb func(key *websocket.Conn, v *BlockingQueue, exists bool) bool

// RemoveCb locks the shard containing the key, retrieves its current value and calls the callback with those params
// If callback returns true and element exists, it will remove it from the map
// Returns the value returned by the callback (even if element was not present in the map)
func (m ConcurrentMap) RemoveCb(key *websocket.Conn, cb RemoveCb) bool {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	remove := cb(key, v, ok)
	if remove && ok {
		delete(shard.items, key)
	}
	shard.Unlock()
	return remove
}

// Pop удаляет элемент из карты и возвращает его.
func (m ConcurrentMap) Pop(key *websocket.Conn) (v *BlockingQueue, exists bool) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, exists = shard.items[key]
	delete(shard.items, key)
	shard.Unlock()
	return v, exists
}

// IsEmpty проверяет, пуста ли карта.
func (m ConcurrentMap) IsEmpty() bool {
	return m.Count() == 0
}

// Используется функцией IterBuffered для объединения двух переменных в канал,
type Tuple struct {
	Key *websocket.Conn
	Val *BlockingQueue
}

// IterBuffered возвращает буферизованный итератор, который можно использовать в цикле for range.
func (m ConcurrentMap) IterBuffered() <-chan Tuple {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan Tuple, total)
	go fanIn(chans, ch)
	return ch
}

// Clear удаляет все элементы из карты.
func (m ConcurrentMap) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshot(m ConcurrentMap) (chans []chan Tuple) {
	chans = make([]chan Tuple, SHARD_COUNT)
	wg := sync.WaitGroup{}
	wg.Add(SHARD_COUNT)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *ConcurrentMapShared) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan Tuple, len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- Tuple{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanIn reads elements from channels `chans` into channel `out`
func fanIn(chans []chan Tuple, out chan Tuple) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan Tuple) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

// Items returns all items as map[string]interface{}
func (m ConcurrentMap) Items() map[*websocket.Conn]*BlockingQueue {
	tmp := make(map[*websocket.Conn]*BlockingQueue)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return tmp
}

// Iterator callback,called for every key,value found in
// maps. RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCb func(key *websocket.Conn, v *BlockingQueue)

// Callback based iterator, cheapest way to read
// all elements in a map.
func (m ConcurrentMap) IterCb(fn IterCb) {
	for idx := range m {
		shard := (m)[idx]
		shard.RLock()
		for key, value := range shard.items {
			fn(key, value)
		}
		shard.RUnlock()
	}
}

// Keys returns all keys as []string
func (m ConcurrentMap) Keys() []*websocket.Conn {
	count := m.Count()
	ch := make(chan *websocket.Conn, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(SHARD_COUNT)
		for _, shard := range m {
			go func(shard *ConcurrentMapShared) {
				// Foreach key, value pair.
				shard.RLock()
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	// Generate keys
	keys := make([]*websocket.Conn, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

//Reviles ConcurrentMap "private" variables to json marshal.
func (m ConcurrentMap) MarshalJSON() ([]byte, error) {
	// Create a temporary map, which will hold all item spread across shards.
	tmp := make(map[*websocket.Conn]*BlockingQueue)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}
	return json.Marshal(tmp)
}

func fnv32(key *websocket.Conn) uint32 {

	const a = 658812287170766751
	const b = 85406431439286701
	const c = 5874916716030321
	uint64key := (*uint64)(unsafe.Pointer(key))
	low := (*uint64key)
	high := (*uint64key) >> 32
	result := uint32((a*low + b*high + c) >> 32)
	logger.Debugf("fnv32: %d", result)
	return result
}
