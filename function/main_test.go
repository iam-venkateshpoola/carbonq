package function

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestCarbonquest(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	Carbonquest(rr, req)

	result := rr.Body.String()
	expected := "Hello from Cloud Function!"

	if !strings.Contains(result, expected) {
		t.Errorf("Expected response to contain %q, but got %q", expected, result)
	}
}
