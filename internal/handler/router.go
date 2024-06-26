package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/a-x-a/go-metric/internal/encoder"
	"github.com/a-x-a/go-metric/internal/logger"
)

func NewRouter(s metricService, log *zap.Logger) http.Handler {
	metricHendlers := newMetricHandlers(s, log)
	// mw := middlewarewithlogger.New(log)

	r := chi.NewRouter()

	r.Use(logger.LoggerMiddleware(log))
	r.Use(encoder.DecompressMiddleware(log))
	r.Use(encoder.CompressMiddleware(log))
	// r.Use(mw.Logger)
	// r.Use(mw.Decompress)
	// r.Use(mw.Compress)

	r.Get("/", metricHendlers.List)

	r.Post("/value/", metricHendlers.GetJSON)
	r.Get("/value/{kind}/{name}", metricHendlers.Get)

	r.Post("/update/", metricHendlers.UpdateJSON)
	r.Post("/update/{kind}/{name}/{value}", metricHendlers.Update)

	return r
}

func responseWithError(w http.ResponseWriter, code int, err error, logger *zap.Logger) {
	resp := fmt.Sprintf("%d: %s", code, err.Error())
	logger.Error(resp)
	http.Error(w, resp, code)
}

func responseWithCode(w http.ResponseWriter, code int, logger *zap.Logger) {
	resp := fmt.Sprintf("%d: %s", code, http.StatusText(code))
	logger.Debug(resp)
	w.WriteHeader(code)
}
