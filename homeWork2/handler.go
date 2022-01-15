package main

import (
	"net/http"
	"strings"
)

// Response 响应内容结构
type Response struct {
	Message string `json:"message"`
}

//健康检查
func (a *App) healthz(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, Response{Message: "healthz success"}, nil)
}

// 获取版本号
func (a *App) version(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{"VERSION": getEnv().Version}
	for name, value := range r.Header {
		headers[name] = strings.Join(value, ",")
	}
	respondWithJSON(w, http.StatusOK, Response{Message: "version success"}, headers)

}
