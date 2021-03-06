package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello Microservice 2018"

type Handlers struct {
	logger *log.Logger
}

func (h Handlers) Home(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startime))

		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}

}
