package main

import (
	"greenlight.bcc/internal/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {

	app := newTestApplication(t)

	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	user := &data.User{Name: "testuser"}
	req = app.contextSetUser(req, user)

	rr := httptest.NewRecorder()

	http.HandlerFunc(app.healthcheckHandler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
