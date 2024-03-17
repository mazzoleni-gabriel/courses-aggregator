package modules

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
	"go.uber.org/fx"

	"github.com/mazzoleni-gabriel/courses-aggregator/internal/server"
)

var internalModule = fx.Options(
	fx.Provide(
		newRouter,
	),
	fx.Invoke(
		server.StartHTTPServer,
	),
)

func newRouter() *chi.Mux {
	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: true,
	})

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return router
}
