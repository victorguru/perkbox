package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Starting the server
func Router() *mux.Router {
	handlers := new(Handlers)
	handlers.setRoutes()
	return handlers.router
}

func TestServer(t *testing.T) {

	request, _ := http.NewRequest("POST", "/coupon/create", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	// 404 error is expected as the url is not valid
	assert.Equal(t, 404, response.Code, "404 (not found) is expected")
}
