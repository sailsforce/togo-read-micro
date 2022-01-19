package internal

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	logKit "github.com/sailsforce/gomicro-kit/logger"
	middlewareKit "github.com/sailsforce/gomicro-kit/middleware"
	"github.com/sailsforce/togo-read-micro/internal/config"
	"github.com/sailsforce/togo-read-micro/internal/graph"
	"github.com/sailsforce/togo-read-micro/internal/graph/generated"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/heartbeat", rest.HeartbeatRotues())
		r.Mount("/playground", playground.Handler("GraphQL Playground", "/v1/query"))
		r.Mount("/query", srv)
	})

	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		config.Logger.Debug(rw.Write([]byte("online")))
	})

	return r
}
