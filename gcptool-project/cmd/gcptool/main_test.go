package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCarbonquest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	Carbonquest(rr, req)

	if rr.Body.String() != "Hello from Cloud Function!\n" {
		t.Errorf("Expected response to be 'Hello from Cloud Function!', got %s", rr.Body.String())
	}
}
