package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

var _ lambda.Handler = (*handler)(nil)

type handler struct{}

func (*handler) Invoke(_ context.Context, _ []byte) ([]byte, error) {
	return json.Marshal("hello world")
}

func main() {
	lambda.Start(&handler{})
}
