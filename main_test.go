package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equalf(t, status, http.StatusOK, "expected status code: %d, got %d", http.StatusOK, status)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equalf(t, status, http.StatusOK, "expected status code: %d, got %d", http.StatusOK, status)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equalf(t, status, http.StatusBadRequest, "expected status code: %d, got %d", http.StatusBadRequest, status)
	body := responseRecorder.Body.String()
	assert.Equal(t, body, "wrong city value")
}
