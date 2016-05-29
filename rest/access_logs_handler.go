package rest

import (
	"net/http"
	"time"

	"github.com/op/go-logging"
)

func AccessLogsHandler(inner http.Handler, name string) http.Handler {
	var log = logging.MustGetLogger("access_log")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Debugf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
