package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/?q=testmessage", nil)
	r.ServeHTTP(w, req)

	// Check if the status code is correct
	if w.Code != 200 {
		t.Error("Received invalid response code. Expected:", 200, "Received:", w.Code)
	}
	// Check if the response from the server is same as the q value we sent
	if w.Body.String() != "\"testmessage\"" {
		t.Error("Received invalid response code. Expected:", "testmessage", "Received:", w.Body.String())
	}
}
