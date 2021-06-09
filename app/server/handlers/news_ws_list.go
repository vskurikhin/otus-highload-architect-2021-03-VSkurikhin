package handlers

import (
	"encoding/json"
	"errors"
	"github.com/atreugo/websocket"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/queue"
	"strings"
	"time"
)

const READ_DEADLINE = 55

func (h *Handlers) WsNewsList(ws *websocket.Conn) error {
	defer func() { _ = ws.Close() }()

	p, err := h.wsNewsListGetProfile(ws)
	if err != nil {
		return err
	}

	wsNewsListSetCloseHandler(ws)
	wsNewsListSetPongHandler(ws)
	queue.WebSocketsMapBQueue.Add(ws)
	// TODO Subscribe for Friend PubSub queue go h.readBq(ws, queue.WebSocketsMapBQueue.Get(ws), p)

	for {
		err := ws.SetReadDeadline(time.Now().Add(READ_DEADLINE * time.Second))
		if err != nil {
			logger.Errorf("handlers.WsNewsList SetReadDeadline: %s", err)
			return err
		}
		err = h.wsNewsListReadMessage(ws, p)
		if err != nil {
			return err
		}
	}
}

func (h *Handlers) wsNewsListReadMessage(ws *websocket.Conn, p *domain.Profile) error {

	mt, message, err := ws.ReadMessage()
	if err != nil {
		logger.Errorf("handlers.wsNewsListReadMessage ReadMessage: %s", err)
		return err
	}
	logger.Debugf("handlers.wsNewsListReadMessage ReadMessage: mt: %d, message: %s", mt, message)
	err = h.wsNewsListSwitchMethod(ws, p, mt, message)
	if err != nil {
		logger.Errorf("handlers.wsNewsListReadMessage wsNewsListSwitchMethod: %s", err)
	}
	return nil
}

func (h *Handlers) wsNewsListSwitchMethod(ws *websocket.Conn, p *domain.Profile, mt int, message []byte) error {

	logger.Debugf("handlers.wsNewsListSwitchMethod: recv: %s", message)
	var m domain.Method
	err := json.Unmarshal(message, &m)
	if err == nil {
		switch m.Method {
		case "fetch":
			return h.wsNewsListFetch(ws, p, mt, m)
		case "heartbeat":
			return h.wsNewsListHeartbeat(ws, mt)
		}
	}
	return err
}

func (h *Handlers) wsNewsListFetch(ws *websocket.Conn, p *domain.Profile, mt int, m domain.Method) error {
	list, err := h.getNewsList(p, m.Offset, m.Limit)
	if err != nil {
		logger.Errorf("wsNewsListFetch error %v", err)
	}
	result := "[" + strings.Join(list, ", ") + "]"
	return ws.WriteMessage(mt, []byte(result))
}

func (h *Handlers) wsNewsListHeartbeat(ws *websocket.Conn, mt int) error {
	return ws.WriteMessage(mt, []byte(`{"Code": 200, "Message": "ok"}`))
}

func (h *Handlers) wsNewsListGetProfile(ws *websocket.Conn) (*domain.Profile, error) {

	jwtCookie := string(ws.UserValue("jwtCookie").([]byte))
	logger.Debugf("jwtCookie: %s", jwtCookie)

	if len(jwtCookie) == 0 {
		return nil, errors.New(" JWT is empty ")
	}
	psid, err := h.Server.JWT.SessionIdFromToken(string(jwtCookie))
	if err != nil {
		return nil, err
	}
	sessionId, err := uuid.Parse(*psid)
	if err != nil {
		return nil, err
	}
	p, err := h.GetProfile(sessionId)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func wsNewsListSetCloseHandler(ws *websocket.Conn) {
	ws.SetCloseHandler(func(code int, text string) error {
		queue.WebSocketsMapBQueue.Delete(ws)
		logger.Debugf("close: %d, %s", code, text)
		return nil
	})
}

func wsNewsListSetPongHandler(ws *websocket.Conn) {
	ws.SetPongHandler(func(string) error {
		err := ws.SetWriteDeadline(time.Now().Add(15 * time.Second))
		if err != nil {
			logger.Errorf("error first setwrite deadline {}", err)
		}
		return nil
	})
}
