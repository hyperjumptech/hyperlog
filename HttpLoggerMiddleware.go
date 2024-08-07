package hyperlog

import (
	"net/http"
	"time"
)

func HttpLoggerMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		responseWrapper := DummyWrapper{
			originalRequest: r,
			originalWriter:  w,
			start:           time.Now(),
		}

		h.ServeHTTP(responseWrapper, r)

	}
	return http.HandlerFunc(fn)
}

type DummyWrapper struct {
	originalRequest *http.Request
	originalWriter  http.ResponseWriter
	start           time.Time
}

func (ww DummyWrapper) Header() http.Header {
	return ww.originalWriter.Header()
}

func (ww DummyWrapper) Write(byts []byte) (int, error) {
	return ww.originalWriter.Write(byts)
}

func (ww DummyWrapper) WriteHeader(statusCode int) {
	dur := time.Since(ww.start)
	Infof("%s %d %s [%d ms]", ww.originalRequest.Method, statusCode, ww.originalRequest.URL.Path, dur/time.Millisecond)
	ww.originalWriter.WriteHeader(statusCode)
}
