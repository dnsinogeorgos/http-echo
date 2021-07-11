package http_echo

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func Run() {
	addr := flag.String("a", ":8080", "listen address")
	rate := flag.Int("r", 3, "number of requests allowed")
	window := flag.Int("w", 10, "rate window in seconds")
	flag.Parse()

	r := NewRouter(*rate, *window)

	log.Printf("starting http-echo on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}

func NewRouter(rate int, window int) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(httprate.LimitByIP(rate, time.Duration(window)*time.Second))
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.HandleFunc("/*", DumpRequest)

	return r
}
