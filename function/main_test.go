package function

import (
    "net/http/httptest"
    "regexp"
    "strings"
    "testing"
)

func TestCarbonquest(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    Carbonquest(rr, req)

    result := rr.Body.String()

    // Remove HTML tags
    re := regexp.MustCompile("<[^>]*>")
    strippedResult := re.ReplaceAllString(result, "")

    expected := "Hello from CarbonQuest Cloud Function!"

    if !strings.Contains(strippedResult, expected) {
        t.Errorf("Expected response to contain %q, but got %q", expected, strippedResult)
    }
}
