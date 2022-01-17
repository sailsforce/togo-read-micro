package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	logKit "github.com/sailsforce/gomicro-kit/logger"
	"github.com/sailsforce/gomicro-kit/models"
	micro "github.com/sailsforce/gomicro-kit/utils"

	"github.com/sailsforce/togo-read-micro/internal/config"
)

func HeartbeatRotues() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetHeartbeat)

	return router
}

func GetHeartbeat(rw http.ResponseWriter, req *http.Request) {
	logger := logKit.GetLogEntry(req)
	dbOnline := false

	logger.Info("pinging database... ", config.RV.DatabaseURL)
	if err := micro.PingDB(config.Psql); err == nil {
		dbOnline = true
	}
	response := models.Heartbeat{
		RequestID:      middleware.GetReqID(req.Context()),
		DatabaseOnline: dbOnline,
		AppName:        config.RV.AppName,
		ReleaseDate:    config.RV.ReleaseDate,
		ReleaseVersion: config.RV.ReleaseVersion,
		Slug:           config.RV.Slug,
		Message:        "togo read service",
	}

	logger.Info("heartbeat request finished.")
	render.JSON(rw, req, response)
}
