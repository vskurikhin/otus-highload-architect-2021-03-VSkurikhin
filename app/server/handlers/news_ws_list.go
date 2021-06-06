package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/atreugo/websocket"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/queue"
	"strings"
	"time"
)

type Method struct {
	Method string
	Offset int
	Limit  int
}

func (h *Handlers) WsNewsList(ws *websocket.Conn) error {
	ws.SetCloseHandler(func(code int, text string) error {
		queue.WebSocketsMapBQueue.Delete(ws)
		logger.Debugf("close: %d, %s", code, text)
		return nil
	})
	ws.SetPongHandler(func(string) error {
		err := ws.SetWriteDeadline(time.Now().Add(3 * time.Second))
		if err != nil {
			fmt.Println("error first setwrite deadline ", err)
		}
		return nil
	})

	queue.WebSocketsMapBQueue.Add(ws)
	go h.readBq(ws, queue.WebSocketsMapBQueue.Get(ws))

	for {
		err := ws.SetReadDeadline(time.Now().Add(55 * time.Second))
		if err != nil {
			logger.Errorf("handlers.WsNewsList SetReadDeadline: %s", err)
			return err
		}
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Errorf("handlers.WsNewsList ReadMessage: %s", err)
			return err
		} else {
			err = h.WsNewsListSwitchMethod(ws, mt, message)
			if err != nil {
				logger.Errorf("handlers.WsNewsList WsNewsListSwitchMethod: %s", err)
				return err
			}
		}
	}
}

func (h *Handlers) WsNewsListSwitchMethod(ws *websocket.Conn, mt int, message []byte) error {
	logger.Debugf("handlers.WsNewsList: recv: %s", message)
	var m Method
	err := json.Unmarshal(message, &m)
	if err == nil {
		switch m.Method {
		case "fetch":
			return h.WsNewsListFetch(ws, mt, m)
		case "heartbeat":
			return h.WsNewsListHeartbeat(ws, mt)
		}
	}
	return err
}

func (h *Handlers) WsNewsListFetch(ws *websocket.Conn, mt int, m Method) error {
	list, _ := h.getNewsList(m.Offset, m.Limit)
	result := "[" + strings.Join(list, ", ") + "]"
	return ws.WriteMessage(mt, []byte(result))
}

func (h *Handlers) WsNewsListHeartbeat(ws *websocket.Conn, mt int) error {
	return ws.WriteMessage(mt, []byte(`{"Code": 200, "Message": "ok"}`))
}

func (h *Handlers) readBq(ws *websocket.Conn, blockingQueue *queue.BlockingQueue) {
	e, ok := blockingQueue.Pop()
	for ok {
		if !queue.WebSocketsMapBQueue.Exists(ws) {
			return
		}
		message := fmt.Sprintf("%v", e)
		logger.Debugf("handlers.readBq: message=%s", message)
		if message == "push" {
			messageCase := domain.ApiMessage{
				Code:    1,
				Message: message,
			}
			if queue.WebSocketsMapBQueue.Exists(ws) {
				err := ws.WriteMessage(1, []byte(messageCase.String()))
				if err != nil {
					logger.Errorf("handlers.readBq: %s", err)
				}
			} else {
				return
			}
		}
		e, ok = blockingQueue.Pop()
	}
}
