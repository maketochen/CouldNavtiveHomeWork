package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RegistryHttpRoutine() {
	http.HandleFunc("/healthz", CommonHandler(Healthz))
	http.HandleFunc("/version", CommonHandler(Version))
	http.Handle("/metrics", promhttp.Handler())
}
