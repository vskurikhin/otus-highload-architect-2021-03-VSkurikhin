package queue

import (
	"container/list"
	"errors"
	"sync"
)

// BlockingQueue это очередь FIFO, в которой операция Pop() блокируется,
// если элементов не существует
type BlockingQueue struct {
	closed bool
	lock   sync.Mutex
	queue  *list.List

	notifyLock sync.Mutex
	monitor    *sync.Cond
}

// New создать экземпляр очереди FIFO
func New() *BlockingQueue {
	bq := &BlockingQueue{queue: list.New()}
	bq.monitor = sync.NewCond(&bq.notifyLock)
	return bq
}

// Put положить любое значение в очередь. Возвращает false, если очередь закрыта
func (bq *BlockingQueue) Put(value interface{}) bool {
	if bq.closed {
		return false
	}
	bq.lock.Lock()
	if bq.closed {
		return false
	}
	bq.queue.PushBack(value)
	bq.lock.Unlock()

	bq.notifyLock.Lock()
	bq.monitor.Signal()
	bq.notifyLock.Unlock()
	return true
}

// PutOrDrop положить значение в очередь или удалить, если очередь заполнена.
// Возвращает false, если очередь закрыта или она заполнена
func (bq *BlockingQueue) PutOrDrop(value interface{}, limit int) bool {
	if bq.closed {
		return false
	}
	ok := false
	bq.lock.Lock()
	if bq.closed {
		return false
	}
	if bq.queue.Len() < limit {
		bq.queue.PushBack(value)
		ok = true
	}
	bq.lock.Unlock()
	if ok {
		bq.monitor.Signal()
	}
	return ok
}

// Pop вытащить переднее значение из очереди.
// Возвращает nil и false, если очередь закрыта
func (bq *BlockingQueue) Pop() (interface{}, bool) {
	if bq.closed {
		return nil, false
	}
	val, ok := bq.getUnblock()
	if ok {
		return val, ok
	}
	for !bq.closed {
		bq.notifyLock.Lock()
		bq.monitor.Wait()
		val, ok = bq.getUnblock()
		bq.notifyLock.Unlock()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

// Size размер очереди. Производительность O(1)
func (bq *BlockingQueue) Size() int {
	bq.lock.Lock()
	defer bq.lock.Unlock()
	return bq.queue.Len()
}

// Closed flag
func (bq *BlockingQueue) Closed() bool {
	bq.lock.Lock()
	defer bq.lock.Unlock()
	return bq.closed
}

// Close закрыть очередь и явно удалите каждый элемент из очереди.
// Также уведомляет всех читателей (они вернут nil и false).
// Возвращает ошибку, если очередь уже закрыта.
func (bq *BlockingQueue) Close() error {
	if bq.closed {
		return errors.New("Already closed")
	}
	bq.closed = true
	bq.lock.Lock()
	//Clear
	for bq.queue.Len() > 0 {
		bq.queue.Remove(bq.queue.Front())
	}
	bq.lock.Unlock()
	bq.monitor.Broadcast()
	return nil
}

func (bq *BlockingQueue) getUnblock() (interface{}, bool) {
	bq.lock.Lock()
	defer bq.lock.Unlock()
	if bq.closed {
		return nil, false
	}
	if bq.queue.Len() > 0 {
		elem := bq.queue.Front()
		bq.queue.Remove(elem)
		return elem.Value, true
	}
	return nil, false
}
