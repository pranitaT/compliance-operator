package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type defaultImpl struct{}

// The impl interface is borrowed from the security-profiles-operator metrics code, with the generated fake_impl.go
// file from the below generate command. Preserved in case we will ever need to regenerate (not likely).
// // go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
// // counterfeiter:generate . impl
type impl interface {
	Register(c prometheus.Collector) error
	ListenAndServe(addr string, handler http.Handler) error
}

func (d *defaultImpl) Register(c prometheus.Collector) error {
	log.Printf("Attempting to register metric: %s", c)
	err := prometheus.Register(c)
	if err != nil {
		log.Printf("Failed to register metric: %s, error: %v", c, err)
	} else {
		log.Printf("Successfully registered metric: %s", c)
	}
	return err
}

func (d *defaultImpl) ListenAndServe(addr string, handler http.Handler) error {
	log.Printf("Starting HTTP server on %s", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Printf("Failed to start HTTP server on %s, error: %v", addr, err)
	} else {
		log.Printf("HTTP server started successfully on %s", addr)
	}
	return err
}
