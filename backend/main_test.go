package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cdn-project/Configs"
	controllers "cdn-project/Controllers"
	"cdn-project/Mocks" // Assure-toi d'importer ton mock

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	// Création du mock pour la collection
	mockCollection := new(Mocks.MockCollection)

	// Remplace la vraie collection par le mock
	Configs.MemberCollection = mockCollection

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.HealthCheck()) // Utilise HealthCheck normalement

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message": "Server is running"}`, rr.Body.String())
}
