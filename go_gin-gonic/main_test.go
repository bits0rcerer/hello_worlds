package main

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func testServer(t *testing.T) *httpexpect.Expect {
	engine := gin.Default()
	applyRoutes(engine)

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func Test_GetHello(t *testing.T) {
	server := testServer(t)

	// Default parameter
	server.GET("/hello").
		Expect().
		Status(http.StatusOK).Text().Equal("Hello World!")

	// custom parameter
	server.GET("/hello").WithQuery("name", "Go").
		Expect().
		Status(http.StatusOK).Text().Equal("Hello Go!")
}

func Test_GetHelloJson(t *testing.T) {
	server := testServer(t)

	// Default parameter
	server.GET("/hello_json").
		Expect().
		Status(http.StatusOK).JSON().
		Path("$.hello").String().Equal("World")

	// custom parameter
	server.GET("/hello_json").
		WithQuery("name", "Go").
		Expect().
		Status(http.StatusOK).JSON().
		Path("$.hello").String().Equal("Go")
}
