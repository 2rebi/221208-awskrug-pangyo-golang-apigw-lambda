package main

import (
	"github.com/aws-serverless-go/httplam"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"hello": "world"}`))
	})
	if httplam.IsLambdaRuntime() {
		httplam.StartLambdaWithAPIGateway(m)
	} else {
		http.ListenAndServe(":3000", m)
	}
}
