package uls

import (
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestBuildCallRedirections(t *testing.T) {
	calls := make(map[string]data.HamCall)
	redirections := data.NewCallRedirections()

	// Set up test data - same FRN with two different callsigns and grant dates
	testFRN := "0012345678"
	
	// KO4JZT is the older call (granted 2010-01-01)
	calls["KO4JZT"] = data.HamCall{
		Callsign: "KO4JZT",
		FRN:      testFRN,
		Grant:    "01/01/2010",
		Name:     "John Smith",
	}
	
	// WW0CJ is the newer call (granted 2020-01-01) - should be the current one
	calls["WW0CJ"] = data.HamCall{
		Callsign: "WW0CJ",
		FRN:      testFRN,
		Grant:    "01/01/2020",
		Name:     "John Smith",
	}

	// Add a call with different FRN to make sure it's not affected
	calls["N0TEST"] = data.HamCall{
		Callsign: "N0TEST",
		FRN:      "0087654321",
		Grant:    "01/01/2015",
		Name:     "Jane Doe",
	}

	// Build redirections
	BuildCallRedirections(&calls, redirections)

	// Test that the FRN maps to the most recent call
	currentCall := redirections.FRNToCurrentCall[testFRN]
	if currentCall != "WW0CJ" {
		t.Errorf("Expected FRN %s to map to current call WW0CJ, got %s", testFRN, currentCall)
	}

	// Test that the former call maps to the FRN
	formerFRN := redirections.FormerCallToFRN["KO4JZT"]
	if formerFRN != testFRN {
		t.Errorf("Expected former call KO4JZT to map to FRN %s, got %s", testFRN, formerFRN)
	}

	// Test that the current call is NOT in the former call map
	if _, exists := redirections.FormerCallToFRN["WW0CJ"]; exists {
		t.Error("Current call WW0CJ should not be in the former call map")
	}

	// Test that a call with different FRN is not affected
	if _, exists := redirections.FormerCallToFRN["N0TEST"]; exists {
		t.Error("N0TEST should not be in former call map as it has unique FRN")
	}

	// Test redirection functionality
	resolved := redirections.ResolveCallsign("KO4JZT")
	if resolved != "WW0CJ" {
		t.Errorf("Expected KO4JZT to resolve to WW0CJ, got %s", resolved)
	}

	resolved = redirections.ResolveCallsign("WW0CJ")
	if resolved != "WW0CJ" {
		t.Errorf("Expected WW0CJ to resolve to itself, got %s", resolved)
	}

	resolved = redirections.ResolveCallsign("N0TEST")
	if resolved != "N0TEST" {
		t.Errorf("Expected N0TEST to resolve to itself, got %s", resolved)
	}
}

func TestBuildCallRedirectionsEmptyFRN(t *testing.T) {
	calls := make(map[string]data.HamCall)
	redirections := data.NewCallRedirections()

	// Add calls with empty FRN - should not create redirections
	calls["W5TEST"] = data.HamCall{
		Callsign: "W5TEST",
		FRN:      "",
		Grant:    "01/01/2020",
	}

	BuildCallRedirections(&calls, redirections)

	// Should not create any redirections
	if len(redirections.FRNToCurrentCall) != 0 {
		t.Error("Expected no redirections for calls with empty FRN")
	}

	if len(redirections.FormerCallToFRN) != 0 {
		t.Error("Expected no former call mappings for calls with empty FRN")
	}
}

func TestBuildCallRedirectionsSingleCall(t *testing.T) {
	calls := make(map[string]data.HamCall)
	redirections := data.NewCallRedirections()

	// Add single call with FRN - should not create redirections
	testFRN := "0012345678"
	calls["W5TEST"] = data.HamCall{
		Callsign: "W5TEST",
		FRN:      testFRN,
		Grant:    "01/01/2020",
	}

	BuildCallRedirections(&calls, redirections)

	// Should not create redirections for single call per FRN
	if len(redirections.FRNToCurrentCall) != 0 {
		t.Error("Expected no redirections for single call per FRN")
	}

	if len(redirections.FormerCallToFRN) != 0 {
		t.Error("Expected no former call mappings for single call per FRN")
	}
}