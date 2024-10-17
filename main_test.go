package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLFiles("view/index.html")
	router.GET("/", Hello)

	// Create a request to the root URL
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expected := "<h1>Hello Blue!</h1>"
	assert.Contains(t, w.Body.String(), expected)
}

func TestHelloEcho(t *testing.T) {
	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLFiles("view/index.html")
	router.GET("/hello/:name", HelloEcho)

	// Create a request to the /hello/{name} URL
	name := "Alice"
	req, _ := http.NewRequest(http.MethodGet, "/hello/"+name, nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expected := "<h1>Hello Alice!</h1>"
	assert.Contains(t, w.Body.String(), expected)
}
