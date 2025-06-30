package data

import (
	"testing"
)

func TestCallRedirections(t *testing.T) {
	redirections := NewCallRedirections()

	// Set up test data - WW0CJ is the current call, KO4JZT is the former call
	testFRN := "0012345678"
	currentCall := "WW0CJ"
	formerCall := "KO4JZT"

	redirections.FRNToCurrentCall[testFRN] = currentCall
	redirections.FormerCallToFRN[formerCall] = testFRN

	// Test resolving a current callsign returns itself
	resolved := redirections.ResolveCallsign(currentCall)
	if resolved != currentCall {
		t.Errorf("Expected current call %s to resolve to itself, got %s", currentCall, resolved)
	}

	// Test resolving a former callsign returns the current callsign
	resolved = redirections.ResolveCallsign(formerCall)
	if resolved != currentCall {
		t.Errorf("Expected former call %s to resolve to current call %s, got %s", formerCall, currentCall, resolved)
	}

	// Test resolving an unknown callsign returns itself
	unknownCall := "N0CALL"
	resolved = redirections.ResolveCallsign(unknownCall)
	if resolved != unknownCall {
		t.Errorf("Expected unknown call %s to resolve to itself, got %s", unknownCall, resolved)
	}
}

func TestNewCallRedirections(t *testing.T) {
	redirections := NewCallRedirections()

	if redirections.FRNToCurrentCall == nil {
		t.Error("FRNToCurrentCall map should be initialized")
	}

	if redirections.FormerCallToFRN == nil {
		t.Error("FormerCallToFRN map should be initialized")
	}

	if len(redirections.FRNToCurrentCall) != 0 {
		t.Error("FRNToCurrentCall map should be empty initially")
	}

	if len(redirections.FormerCallToFRN) != 0 {
		t.Error("FormerCallToFRN map should be empty initially")
	}
}