package queue

import (
	"github.com/atreugo/websocket"
	"github.com/savsgio/go-logger/v2"
	"sync"
)

type MapBq struct {
	lock  sync.Mutex
	mapBq map[*websocket.Conn]*BlockingQueue

	notifyLock sync.Mutex
	monitor    *sync.Cond
}

func newMapBq() *MapBq {
	m := &MapBq{mapBq: make(map[*websocket.Conn]*BlockingQueue)}
	m.monitor = sync.NewCond(&m.notifyLock)
	return m
}

var WebSocketsMapBQueue = newMapBq()

func (m *MapBq) PutToMapWsBq(message string) {
	for _, ws := range m.mapBq {
		logger.Debugf("queue.PutToMapWsBq(%s) for %v", message, ws)
		ws.Put(message)
	}
}

func (m *MapBq) Add(ws *websocket.Conn) {
	m.lock.Lock()
	m.mapBq[ws] = New()
	m.lock.Unlock()
	m.notifyLock.Lock()
	m.monitor.Signal()
	m.notifyLock.Unlock()
}

func (m *MapBq) Get(ws *websocket.Conn) *BlockingQueue {
	m.lock.Lock()
	result := m.mapBq[ws]
	m.lock.Unlock()
	m.notifyLock.Lock()
	m.monitor.Signal()
	m.notifyLock.Unlock()

	return result
}

func (m *MapBq) Delete(ws *websocket.Conn) {
	m.lock.Lock()
	delete(m.mapBq, ws)
	m.lock.Unlock()
	m.notifyLock.Lock()
	m.monitor.Signal()
	m.notifyLock.Unlock()
}

func (m *MapBq) Exists(ws *websocket.Conn) bool {
	m.lock.Lock()
	_, ok := m.mapBq[ws]
	m.lock.Unlock()
	m.notifyLock.Lock()
	m.monitor.Signal()
	m.notifyLock.Unlock()

	return ok
}
