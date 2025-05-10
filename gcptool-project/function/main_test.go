package function

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rr := httptest.NewRecorder()

    HelloWorld(rr, req)

    if rr.Body.String() != "Hello from Cloud Function!\n" {
        t.Errorf("Expected response to be 'Hello from Cloud Function!', got %s", rr.Body.String())
    }
}
