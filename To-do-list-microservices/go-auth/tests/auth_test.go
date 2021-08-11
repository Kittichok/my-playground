package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/server"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestSignInSuccess(t *testing.T) {
	d := sqlite.Open("file::memory:")
	models.ConnectDataBase(d)
	u := models.User{
		Username: "test",
		Password: "test"}
	models.SeedUser(u)

	router := server.SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString("{\"username\":\"test\", \"password\":\"test\"}")
	req, _ := http.NewRequest("POST", "/v1/api/signin", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
			t.Fatal(err)
			t.FailNow()
	}
}


func TestSignInInvalidUser(t *testing.T) {
	d := sqlite.Open("file::memory:")
	models.ConnectDataBase(d)

	router := server.SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBufferString("{\"username\":\"test\", \"password\":\"test\"}")
	req, _ := http.NewRequest("POST", "/v1/api/signin", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
			t.Fatal(err)
			t.FailNow()
	}

	assert.Equal(t, "invalid username or password", got["error"])
}