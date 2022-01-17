package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	logKit "github.com/sailsforce/gomicro-kit/logger"
	middlewareKit "github.com/sailsforce/gomicro-kit/middleware"
	"github.com/sailsforce/togo-read-micro/internal/config"
	"github.com/sailsforce/togo-read-micro/internal/rest"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.RequestID,
		logKit.NewStructuredLogger(config.Logger),
		middlewareKit.Headers,
		middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/heartbeat", rest.HeartbeatRotues())
	})

	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		config.Logger.Debug(rw.Write([]byte("online")))
	})

	return r
}
