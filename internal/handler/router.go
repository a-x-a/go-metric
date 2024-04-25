package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/a-x-a/go-metric/internal/encoder"
	"github.com/a-x-a/go-metric/internal/logger"
	"github.com/a-x-a/go-metric/internal/signer"
)

func NewRouter(s metricService, log *zap.Logger, signer *signer.Signer) http.Handler {
	metricHendlers := newMetricHandlers(s, log, signer)

	r := chi.NewRouter()

	r.Use(logger.LoggerMiddleware(log))
	r.Use(encoder.DecompressMiddleware(log))
	r.Use(encoder.CompressMiddleware(log))
	// r.Use(signer.SignerMiddleware(log))

	r.Get("/", metricHendlers.List)

	r.Post("/value", metricHendlers.GetJSON)
	r.Post("/value/", metricHendlers.GetJSON)
	r.Get("/value/{kind}/{name}", metricHendlers.Get)

	r.Post("/update", metricHendlers.UpdateJSON)
	r.Post("/update/", metricHendlers.UpdateJSON)
	r.Post("/update/{kind}/{name}/{value}", metricHendlers.Update)
	r.Post("/updates", metricHendlers.UpdateBatch)
	r.Post("/updates/", metricHendlers.UpdateBatch)

	r.Get("/ping", metricHendlers.Ping)
	r.Get("/ping/", metricHendlers.Ping)

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
