package root

import (
	"github.com/fate-lovely/phi"
	"hl.svn.su/highload-architect/app/handlers/root/home"
)

// Handler returns an http.Handler
func Handler() *phi.Mux {

	r := phi.NewRouter()
	r.Get("/", home.Index)

	return r
}
