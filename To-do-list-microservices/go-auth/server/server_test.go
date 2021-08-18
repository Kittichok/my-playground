package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kittichok/go-auth/internal/controllers"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/repository"
	"github.com/kittichok/go-auth/internal/usecase/users"
	"github.com/kittichok/go-auth/server"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	rep := repository.NewUserRepository(models.DB)
	u := users.NewUserUseCase(rep)
	authController := controllers.NewAuthController(u)
	router := server.SetupRouter(authController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

