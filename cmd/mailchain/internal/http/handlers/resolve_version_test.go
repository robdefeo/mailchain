package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResolveVersion(t *testing.T) {
	assert := assert.New(t)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/version", nil)
	handler := http.HandlerFunc(GetResolveVersion())
	handler.ServeHTTP(rr, req)
	wantStatus := http.StatusOK
	wantBody := "{\"version\":\"dev\",\"commit\":\"none\",\"time\":\"unknown\"}\n"
	if !assert.Equal(wantStatus, rr.Code) {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, wantStatus)
	}
	if !assert.Equal(wantBody, rr.Body.String()) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), wantBody)
	}
}
