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

	var resp events.APIGatewayV2HTTPResponse
	switch req.RawPath {
	case "/":
		switch req.RequestContext.HTTP.Method {
		case http.MethodGet:
			resp = helloworld()
		default:
			resp = methodNotAllowed()
		}
	case "/item":
		switch req.RequestContext.HTTP.Method {
		case http.MethodGet:
			resp = item()
		default:
			resp = methodNotAllowed()
		}
	//case "/etc":
	//	switch req.RequestContext.HTTP.Method {
	//	case http.MethodPost:
	//		//... code ...
	//	default:
	//		resp = methodNotAllowed()
	//	}
	// ... code ...
	default:
		resp = notfound()
	}
	return json.Marshal(resp)
}

func item() events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"id": 1, "name": "My Item"}`,
	}
}

func helloworld() events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"hello": "world"}`,
	}
}

func methodNotAllowed() events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "method not allowed"}`,
	}
}

func notfound() events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "not found"}`,
	}
}

func main() {
	lambda.Start(&handler{})
}
