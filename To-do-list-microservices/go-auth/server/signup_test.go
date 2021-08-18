package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/go-auth/internal/controllers"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/repository"
	"github.com/kittichok/go-auth/internal/usecase/users"
	"github.com/kittichok/go-auth/server"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestSignUpSuccess(t *testing.T) {
	d := sqlite.Open("file::memory:")
	models.ConnectDataBase(d)
	rep := repository.NewUserRepository(models.DB)
	usecase := users.NewUserUseCase(rep)
	authController := controllers.NewAuthController(usecase)
	router := server.SetupRouter(authController)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString("{\"username\":\"test\", \"password\":\"test\"}")
	req, _ := http.NewRequest("POST", "/api/v1/signup", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}


func TestSignUpDuplicateUser(t *testing.T) {
	d := sqlite.Open("file::memory:")
	models.ConnectDataBase(d)
	u := models.User{
		Username: "test",
		Password: "test"}
	models.SeedUser(u)
	rep := repository.NewUserRepository(models.DB)
	usecase := users.NewUserUseCase(rep)
	authController := controllers.NewAuthController(usecase)
	router := server.SetupRouter(authController)

	w := httptest.NewRecorder()
	body := bytes.NewBufferString("{\"username\":\"test\", \"password\":\"test\"}")
	req, _ := http.NewRequest("POST", "/api/v1/signup", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
			t.Fatal(err)
			t.FailNow()
	}
	assert.Equal(t, "user is exist", got["error"])
}