package signer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/a-x-a/go-metric/internal/adapter"
)

type Signer struct {
	key []byte
}

func New(key string) *Signer {
	if len(key) == 0 {
		return nil
	}

	return &Signer{[]byte(key)}
}

func (s *Signer) Hash(data []byte) ([]byte, error) {
	h := hmac.New(sha256.New, s.key)
	h.Write(data)

	return h.Sum(nil), nil
}

func (s *Signer) Verify(data []byte, hash string) (bool, error) {
	mac1, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac2, err := s.Hash(data)
	if err != nil {
		return false, err
	}

	return hmac.Equal(mac1, mac2), nil
}

func SignerMiddleware(log *zap.Logger, key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sgnr := New(key)
			if sgnr == nil {
				next.ServeHTTP(w, r)
				return
			}

			hash := r.Header.Get("HashSHA256")
			if len(hash) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Info("SIGNER", zap.String("hash received", hash))
			if len(hash) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			buf, _ := io.ReadAll(r.Body)
			rdr1 := io.NopCloser(bytes.NewBuffer(buf))
			rdr2 := io.NopCloser(bytes.NewBuffer(buf))

			data := make([]adapter.RequestMetric, 0)
			if err := json.NewDecoder(rdr1).Decode(&data); err != nil {
				log.Info("SIGNER", zap.Error(err))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if ok, err := sgnr.Verify(b, hash); !ok || err != nil {
				log.Info("SIGNER", zap.String("hash is not valid", hash))
				log.Info("SIGNER", zap.Error(err))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Info("SIGNER", zap.String("hash is valid", hash))

			r.Body = rdr2
			next.ServeHTTP(w, r)
		})
	}
}