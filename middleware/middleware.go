package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/LetsFocus/goLF/logger"
)

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func AddCorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("X-CORRELATIONID") == "" {
			ctx := context.WithValue(r.Context(), "X-CORRELATION", uuid.New())
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log := logger.NewCustomLogger()

		log.Infof("Request Received at %v ,Method : %v , api: %v , X-CORRELATIONID: %v", start, r.Method,
			r.URL, r.Context().Value("X-CORRELATIONID"))

		next.ServeHTTP(w, r)

	})
}
