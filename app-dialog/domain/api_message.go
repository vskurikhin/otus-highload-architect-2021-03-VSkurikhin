package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
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
		logger.Errorf("%v", err)
		return nil
	}
	return apiMessage
}
