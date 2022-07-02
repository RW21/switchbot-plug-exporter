package main

import (
	"log"

	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9101"
	}

	// exporter := NewExporter()
	// prometheus.MustRegister(exporter)

	s := NewHttpServer()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+port, s))
}
