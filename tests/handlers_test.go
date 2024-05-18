package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hryhorskyi/gin-app/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
	router := gin.Default()
	router.GET("/api/rate", handlers.GetRate)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/rate", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "rate")
}

func TestSubscribe(t *testing.T) {
	router := gin.Default()
	router.POST("/api/subscribe", handlers.Subscribe)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/subscribe", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
