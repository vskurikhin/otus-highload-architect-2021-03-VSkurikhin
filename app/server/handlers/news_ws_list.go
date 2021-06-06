package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/atreugo/websocket"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/queue"
	"strings"
)

type Method struct {
	Method string
	Offset int
	Limit  int
}

func (h *Handlers) WsNewsList(ws *websocket.Conn) error {
	ws.SetCloseHandler(func(code int, text string) error {
		delete(queue.MapBq, ws) // TODO need LOCK MapBq!!!
		logger.Debugf("close: %d, %s", code, text)
		return nil
	})

	queue.MapBq[ws] = queue.New()
	go h.readBq(ws, queue.MapBq[ws])

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Errorf("handlers.WsNewsList: %s", err)
			return err
		}
		logger.Debugf("handlers.WsNewsList: recv: %s", message)
		var m Method
		err = json.Unmarshal(message, &m)

		if err == nil {
			me := fmt.Sprintf("%v", m)
			logger.Debugf("handlers.WsNewsList: me: %s", me)
			list, _ := h.getNewsList(0, 99)
			result := "[" + strings.Join(list, ", ") + "]"
			err = ws.WriteMessage(mt, []byte(result))
			if err != nil {
				logger.Errorf("handlers.WsNewsList: %s", err)
				return err
			}
		}
	}
}

func (h *Handlers) readBq(ws *websocket.Conn, blockingQueue *queue.BlockingQueue) {
	e, ok := blockingQueue.Pop()
	for ok {
		if _, ok := queue.MapBq[ws]; !ok || h == nil { // TODO need LOCK MapBq!!!
			return
		}
		message := fmt.Sprintf("%v", e)
		logger.Debugf("handlers.readBq: message=%s", message)
		if message == "push" {
			messageCase := domain.ApiMessage{
				Code:    1,
				Message: message,
			}
			if _, ok := queue.MapBq[ws]; ok && h != nil { // TODO need LOCK MapBq!!!
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
