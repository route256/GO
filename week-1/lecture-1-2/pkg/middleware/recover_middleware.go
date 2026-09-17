package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// middleware layer

// WithHTTPRecoverMiddleware recover panics
func WithHTTPRecoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
