package utils

import "github.com/savsgio/go-logger/v2"

func PanicCheck(e error, messages ...string) {

	if e != nil {
		for _, msg := range messages {
			logger.Error(msg)
		}
		panic(e)
	}
}

func PanicNotOk(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}
