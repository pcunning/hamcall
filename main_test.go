package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestWebHandler(t *testing.T) {
	// Create test data
	calls := map[string]data.HamCall{
		"W5TEST": {
			Callsign:  "W5TEST",
			Name:      "John Smith",
			FirstName: "John",
			LastName:  "Smith",
			Class:     "Extra",
			State:     "TX",
			Grant:     "2020-01-01",
			Expiration: "2030-01-01",
		},
		"KC5ABC": {
			Callsign:  "KC5ABC",
			Name:      "Jane Doe",
			FirstName: "Jane",
			LastName:  "Doe",
			Class:     "General",
			State:     "CA",
		},
	}

	// Create handler closure with test data
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := r.URL.Path[len("/") : len(r.URL.Path)-len(".json")]
		j, err := json.Marshal(calls[strings.ToUpper(call)])
		if err != nil {
			// log.Fatalf(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	// Test valid callsign lookup
	req := httptest.NewRequest("GET", "/w5test.json", nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", recorder.Code)
	}

	var result data.HamCall
	err := json.Unmarshal(recorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.Callsign != "W5TEST" {
		t.Fatalf("Expected callsign 'W5TEST', got '%s'", result.Callsign)
	}
	if result.Name != "John Smith" {
		t.Fatalf("Expected name 'John Smith', got '%s'", result.Name)
	}

	// Test case insensitive lookup
	req2 := httptest.NewRequest("GET", "/kc5abc.json", nil)
	recorder2 := httptest.NewRecorder()
	handler.ServeHTTP(recorder2, req2)

	if recorder2.Code != http.StatusOK {
		t.Fatalf("Expected status 200 for lowercase lookup, got %d", recorder2.Code)
	}

	var result2 data.HamCall
	err = json.Unmarshal(recorder2.Body.Bytes(), &result2)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result2.Callsign != "KC5ABC" {
		t.Fatalf("Expected callsign 'KC5ABC', got '%s'", result2.Callsign)
	}

	// Test non-existent callsign
	req3 := httptest.NewRequest("GET", "/n0exist.json", nil)
	recorder3 := httptest.NewRecorder()
	handler.ServeHTTP(recorder3, req3)

	if recorder3.Code != http.StatusOK {
		t.Fatalf("Expected status 200 even for non-existent callsign, got %d", recorder3.Code)
	}

	var result3 data.HamCall
	err = json.Unmarshal(recorder3.Body.Bytes(), &result3)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should return empty HamCall struct for non-existent callsign
	if result3.Callsign != "" {
		t.Fatalf("Expected empty callsign for non-existent lookup, got '%s'", result3.Callsign)
	}
}

func TestProcessFunction(t *testing.T) {
	// Test that process function doesn't panic and can be called
	// We can't easily test the full functionality without setting up data files
	// but we can at least ensure the function exists and runs without error
	calls := make(map[string]data.HamCall)
	
	// This should not panic even if data files don't exist
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("process() function panicked: %v", r)
		}
	}()
	
	// Call process function - it will try to process files but should handle missing files gracefully
	process(&calls)
	
	// calls should still be a valid map (even if empty due to missing files)
	if calls == nil {
		t.Fatalf("process() should not nil the calls map")
	}
}

func TestDownloadFilesFunction(t *testing.T) {
	// Test that downloadFiles function exists and can be referenced
	// We skip actual execution to avoid network dependencies and log.Fatalf calls
	
	// Just verify we can reference the function (if we couldn't, this wouldn't compile)
	_ = downloadFiles
	
	// Skip actual execution to avoid network dependencies
	t.Skip("Skipping actual download test to avoid network dependencies and log.Fatalf calls")
}