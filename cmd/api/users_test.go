package main

import (
	"net/http/httptest"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	app := newTestApplication(t)

	r := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()
	app.registerUserHandler(w, r)
}

func TestActivateUserHandler(t *testing.T) {
	app := newTestApplication(t)

	r := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()
	app.activateUserHandler(w, r)
}
