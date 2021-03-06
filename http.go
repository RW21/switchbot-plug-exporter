package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HttpServer struct {
	mux *http.ServeMux
}

func NewHttpServer() *HttpServer {
	s := &HttpServer{
		mux: http.NewServeMux(),
	}
	s.mux.HandleFunc("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}).ServeHTTP)

	s.mux.HandleFunc("/scrape", s.ScrapeHandler)
	return s
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

//https://github.com/oliver006/redis_exporter
func (s *HttpServer) ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' parameter must be specified", 400)
		return
	}

	registry := prometheus.NewRegistry()
	exporter := NewExporter(target)
	registry.MustRegister(exporter)

	promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
