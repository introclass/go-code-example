package main

import (
	"net/http"
	"testing"
)
import "net/http/httptest"
import "github.com/stretchr/testify/assert"

func TestPingRoute(t *testing.T) {
	route := setupRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	route.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
