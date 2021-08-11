package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kittichok/go-auth/server"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

