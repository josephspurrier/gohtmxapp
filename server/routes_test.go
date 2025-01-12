package server_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HelloWorldTester))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

// HelloWorldTester is the tester for the Hello World hander.
func HelloWorldTester(w http.ResponseWriter, _ *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// func TestHello(t *testing.T) {
// 	s := server.NewServer()
// 	defer s.Close()
// 	//svc := httptest.NewServer(s)
// 	//defer svc.Close()

// 	resp, err := http.Get(server.URL)
// 	if err != nil {
// 		t.Fatalf("error making request to server. Err: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Assertions
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", resp.Status)
// 	}

// 	expected := "{\"message\":\"Hello World\"}"
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("error reading response body. Err: %v", err)
// 	}

// 	if expected != string(body) {
// 		t.Errorf("expected response body to be %v; got %v", expected, string(body))
// 	}
// }
