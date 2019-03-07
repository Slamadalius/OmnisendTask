package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Router() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", Index)

	return r
}

func TestIndexEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response   := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK 200 response is expected")
}