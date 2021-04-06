package handlers

import (
	sa "github.com/savsgio/atreugo/v11"
)

func (h *Handlers) Root(ctx *sa.RequestCtx) error {
	return ctx.RedirectResponse("/index.html", ctx.Response.StatusCode())
}
