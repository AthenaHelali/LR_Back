package httpserver

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() {
	s := Server{}
	e := echo.New()
	e.GET("/health-check", s.healthCheck)
	go e.Start("localhost:8088")
	time.Sleep(100 * time.Millisecond)
}

func TestHealth(t *testing.T) {
	setup()
	_, err := http.Get("http://127.0.0.1:8088/health-check")
	assert.NoError(t, err)
}
