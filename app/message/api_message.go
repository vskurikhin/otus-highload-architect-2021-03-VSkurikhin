package message

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

type ApiMessage struct {
	Code    int
	Message string
}

var LoginRequired = ApiMessage{
	Code:    fasthttp.StatusForbidden,
	Message: "login required",
}

var YourSessionIsExpired = ApiMessage{
	Code:    fasthttp.StatusForbidden,
	Message: "your session is expired, login again please",
}

func (a *ApiMessage) String() string {
	return string(a.Marshal())
}

func (a *ApiMessage) Marshal() []byte {

	apiMessage, err := json.Marshal(*a)
	if err != nil {
		log.Println(err)
		return nil
	}
	return apiMessage
}
