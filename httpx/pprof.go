package httpx

import (
	"net/http"
	"net/http/pprof"
)

//新建http pprof服务器

func NewPprofHandler() http.Handler {
	serverMux := http.NewServeMux()

	serverMux.Handle("/", http.HandlerFunc(pprof.Index))
	serverMux.Handle("/cmdline", http.HandlerFunc(pprof.Cmdline))
	serverMux.Handle("/profile", http.HandlerFunc(pprof.Profile))
	serverMux.Handle("/symbol", http.HandlerFunc(pprof.Symbol))
	serverMux.Handle("/trace", http.HandlerFunc(pprof.Trace))

	// Manually add support for paths linked to by index page at /debug/pprof/
	serverMux.Handle("/goroutine", pprof.Handler("goroutine"))
	serverMux.Handle("/heap", pprof.Handler("heap"))
	serverMux.Handle("/threadcreate", pprof.Handler("threadcreate"))
	serverMux.Handle("/block", pprof.Handler("block"))
	return serverMux
}
