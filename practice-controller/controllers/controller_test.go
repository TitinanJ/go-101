package controllers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"httptesting/controllers"
)

func TestGetUser(t *testing.T) {
    userID := "123"
    req := httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
    req.SetPathValue("id", userID)
    w := httptest.NewRecorder()

    controllers.HelloController(w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }
    defer resp.Body.Close()
    respondedBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }
    expectedText := "Hello, user " + userID + "!"
    if string(respondedBytes) != expectedText {
        t.Fatalf("Expected response body '%s', got '%s'", expectedText, string(respondedBytes))
    }
}