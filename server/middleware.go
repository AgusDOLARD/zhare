package server

import (
	"net/http"
)

func (s *Server) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &rwWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)

		s.logger.Info("Request",
			"status", wrapped.statusCode,
			"method", r.Method,
			"url", r.URL,
		)
	})
}

type rwWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *rwWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
