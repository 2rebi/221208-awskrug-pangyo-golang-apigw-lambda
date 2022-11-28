package main

import (
	"github.com/aws-serverless-go/echolam"
	"github.com/aws-serverless-go/httplam"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"hello": "world",
		})
	})
	if httplam.IsLambdaRuntime() {
		echolam.StartLambdaWithAPIGateway(e)
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
