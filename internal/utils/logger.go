package utils

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sailsforce/togo-read-micro/internal/config"
	"github.com/sirupsen/logrus"
)

func LogRequestId(req *http.Request, reqName string) *logrus.Logger {
	logger := config.Logger
	requestId := middleware.GetReqID(req.Context())
	logger.WithField("req_id", requestId)
	logger.Info(reqName, " -- starting request...")
	return logger
}
