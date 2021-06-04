package handlers

import (
	"github.com/atreugo/websocket"
	"log"
	"strings"
)

func (h *Handlers) WsNewsList(ws *websocket.Conn) error {
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("err: %s", err)
			return err
		}

		// name := ws.UserValue("name").(string)

		log.Printf("recv: %s", message)

		list, _ := h.getNewsList(0, 99)
		result := "[" + strings.Join(list, ", ") + "]"

		err = ws.WriteMessage(mt, []byte(result))
		if err != nil {
			log.Printf("err: %s", err)
			return err
		}
	}
}
