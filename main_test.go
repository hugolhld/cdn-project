package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controller "cdn-project/Controllers"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.HealthCheck())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler renvoie un mauvais code HTTP: obtenu %v, attendu %v", status, http.StatusOK)
	}
}
