package handlers

import (
	sa "github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func (h *Handlers) Root(ctx *sa.RequestCtx) error {
	ctx.Response.Header.SetCanonical([]byte("Location"), []byte("/index.html"))
	ctx.Response.SetStatusCode(fasthttp.StatusFound)

	return nil // ctx.RedirectResponse("/index.html", ctx.Response.StatusCode())
}
