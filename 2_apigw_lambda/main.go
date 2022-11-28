package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

var _ lambda.Handler = (*handler)(nil)

type handler struct{}

func (*handler) Invoke(_ context.Context, payload []byte) ([]byte, error) {
	var req events.APIGatewayV2HTTPRequest
	err := json.Unmarshal(payload, &req)
	if err != nil {
		return nil, err
	}

	resp := events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		MultiValueHeaders: nil,
		Body:              `{"hello": "world"}`,
		IsBase64Encoded:   false,
		Cookies:           nil,
	}
	return json.Marshal(resp)
}

func main() {
	lambda.Start(&handler{})
}
