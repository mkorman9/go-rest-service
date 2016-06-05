package rest

import (
	"net/http"
	"errors"

	"github.com/op/go-logging"
	"fmt"
	"os"
)

func RecoverHandler(h http.Handler) http.Handler {
	var log = logging.MustGetLogger("panic_log")

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				log.Errorf(
					"%s\t%s\t%s\tFatal error during request handling: %s",
					req.RemoteAddr,
					req.Method,
					req.RequestURI,
					err.Error(),
				)
				fmt.Fprintf(
					os.Stderr,
					"%s\t%s\t%s\tFatal error during request handling: %s\n",
					req.RemoteAddr,
					req.Method,
					req.RequestURI,
					err.Error(),
				)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, req)
	})
}
