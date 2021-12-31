package main

import (
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

func (a *App) healthz(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, Response{Message: "success"}, nil)
}

func (a *App) version(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	headers := map[string]string{"VERSION": getEnv().Version}
	for name, value := range header {
		headers[name] = strings.Join(value, ",")
	}
	respondWithJSON(w, http.StatusOK, Response{Message: "success"}, headers)


}
