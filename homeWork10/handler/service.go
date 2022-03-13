package handler


import "net/http"

// Healthz 4. 当访问localhost/healthz时，应返回200
func Healthz(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusOK
	writer.WriteHeader(*statusCode)
	_, err := writer.Write([]byte("Healthz Success \n"))
	if err != nil {
		return
	}
}

func Version(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusAccepted
	writer.WriteHeader(*statusCode)
	_, err := writer.Write([]byte("Version Success\n"))
	if err != nil {
		return
	}
}
