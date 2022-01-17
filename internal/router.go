package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sailsforce/togo-read-micro/internal/rest"
	"github.com/sailsforce/togo-read-micro/internal/utils"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.RequestID,
		middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/heartbeat", rest.HeartbeatRotues())
	})

	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		logger := utils.LogRequestId(req, "")
		logger.Debug(rw.Write([]byte("online")))
	})

	return r
}
