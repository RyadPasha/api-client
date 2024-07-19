/*
Package apiClient provides a reusable library for sending external API requests with features like debugging, retries, and request/response logging.

Author: Mohamed Riyad
Email: m@ryad.dev
Website: https://ryad.dev
*/

package apiClient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSendRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, true, 3, 2*time.Second)

	req := &APIRequest{
		Method:   http.MethodGet,
		Endpoint: "/",
		Headers:  map[string]string{"Content-Type": "application/json"},
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"message": "success"}`
	if string(resp.Body) != expectedBody {
		t.Fatalf("Expected body %s, got %s", expectedBody, resp.Body)
	}
}
