package queue

import (
	"github.com/atreugo/websocket"
	"github.com/savsgio/go-logger/v2"
)

var MapBq = make(map[*websocket.Conn]*BlockingQueue)

func PutToMapWsBq(message string) {
	for _, ws := range MapBq {
		logger.Debugf("queue.PutToMapWsBq(%s) for %v", message, ws)
		ws.Put(message)
	}
}
