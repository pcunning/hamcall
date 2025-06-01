package radioid

import (
	"os"
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestRadioIdProcess(t *testing.T) {
	// Create test semicolon-delimited data
	testData := `1234567;W5TEST;John;Smith;Anytown;TX;United States
2345678;KC5ABC;Jane;Doe;Somewhere;CA;United States
3456789;N0DEF;Bob;Wilson;Nowhere;FL;United States`

	// Create test file
	testFile, err := os.CreateTemp("", "radioid_test_*.dat")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(testFile.Name())
	defer testFile.Close()

	_, err = testFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	testFile.Close()

	// Change working directory temporarily and create symlink
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "radioid_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "dmrid.dat")

	// Test processing
	calls := make(map[string]data.HamCall)
	Process(&calls)

	// Verify results
	if len(calls) != 3 {
		t.Fatalf("Expected 3 calls, got %d", len(calls))
	}

	// Test W5TEST
	w5test, exists := calls["W5TEST"]
	if !exists {
		t.Fatalf("W5TEST not found in processed calls")
	}
	if len(w5test.DMRID) != 1 {
		t.Fatalf("Expected 1 DMRID for W5TEST, got %d", len(w5test.DMRID))
	}
	if w5test.DMRID[0] != 1234567 {
		t.Fatalf("Expected DMRID 1234567 for W5TEST, got %d", w5test.DMRID[0])
	}

	// Test KC5ABC
	kc5abc, exists := calls["KC5ABC"]
	if !exists {
		t.Fatalf("KC5ABC not found in processed calls")
	}
	if len(kc5abc.DMRID) != 1 {
		t.Fatalf("Expected 1 DMRID for KC5ABC, got %d", len(kc5abc.DMRID))
	}
	if kc5abc.DMRID[0] != 2345678 {
		t.Fatalf("Expected DMRID 2345678 for KC5ABC, got %d", kc5abc.DMRID[0])
	}

	// Test N0DEF
	n0def, exists := calls["N0DEF"]
	if !exists {
		t.Fatalf("N0DEF not found in processed calls")
	}
	if len(n0def.DMRID) != 1 {
		t.Fatalf("Expected 1 DMRID for N0DEF, got %d", len(n0def.DMRID))
	}
	if n0def.DMRID[0] != 3456789 {
		t.Fatalf("Expected DMRID 3456789 for N0DEF, got %d", n0def.DMRID[0])
	}
}

func TestRadioIdProcessWithExistingCalls(t *testing.T) {
	// Create test data
	testData := `1234567;W5TEST;John;Smith;Anytown;TX;United States
7654321;W5TEST;John;Smith;Anytown;TX;United States`

	// Create test file
	testFile, err := os.CreateTemp("", "radioid_test_*.dat")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(testFile.Name())
	defer testFile.Close()

	_, err = testFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	testFile.Close()

	// Change working directory temporarily and create symlink
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "radioid_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "dmrid.dat")

	// Pre-populate calls map with existing data
	calls := make(map[string]data.HamCall)
	calls["W5TEST"] = data.HamCall{
		Callsign:  "W5TEST",
		Name:      "John Smith", 
		FirstName: "John",
		LastName:  "Smith",
		State:     "TX",
		DMRID:     []int{9999999}, // existing DMRID
	}

	// Test processing
	Process(&calls)

	// Verify results - should have 1 call with multiple DMRIDs
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call, got %d", len(calls))
	}

	w5test := calls["W5TEST"]
	
	// Check that existing data is preserved
	if w5test.Name != "John Smith" {
		t.Fatalf("Expected name 'John Smith', got '%s'", w5test.Name)
	}
	if w5test.State != "TX" {
		t.Fatalf("Expected state 'TX', got '%s'", w5test.State)
	}
	
	// Check that DMRID data was added (should have 3 total: 1 existing + 2 new)
	if len(w5test.DMRID) != 3 {
		t.Fatalf("Expected 3 DMRIDs, got %d", len(w5test.DMRID))
	}
	
	// Check specific IDs (order might vary)
	expectedIDs := map[int]bool{9999999: true, 1234567: true, 7654321: true}
	for _, id := range w5test.DMRID {
		if !expectedIDs[id] {
			t.Fatalf("Unexpected DMRID %d", id)
		}
	}
}

func TestRadioIdProcessInvalidID(t *testing.T) {
	// Create test data with invalid ID
	testData := `abc123;W5TEST;John;Smith;Anytown;TX;United States
1234567;KC5ABC;Jane;Doe;Somewhere;CA;United States`

	// Create test file
	testFile, err := os.CreateTemp("", "radioid_test_*.dat")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(testFile.Name())
	defer testFile.Close()

	_, err = testFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	testFile.Close()

	// Change working directory temporarily and create symlink
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "radioid_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "dmrid.dat")

	// Test processing
	calls := make(map[string]data.HamCall)
	Process(&calls)

	// Should only process the valid record
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call (invalid ID should be skipped), got %d", len(calls))
	}

	// Should only have KC5ABC since W5TEST had invalid ID
	kc5abc, exists := calls["KC5ABC"]
	if !exists {
		t.Fatalf("KC5ABC not found - should have been processed")
	}
	if kc5abc.DMRID[0] != 1234567 {
		t.Fatalf("Expected DMRID 1234567 for KC5ABC, got %d", kc5abc.DMRID[0])
	}

	// W5TEST should not exist since its ID was invalid
	_, w5testExists := calls["W5TEST"]
	if w5testExists {
		t.Fatalf("W5TEST should not exist due to invalid DMRID")
	}
}

func TestRadioIdProcessMissingFile(t *testing.T) {
	// Change to a temporary directory where dmrid.dat doesn't exist
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "radioid_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	calls := make(map[string]data.HamCall)
	
	// This should not panic or error, just return without processing
	Process(&calls)
	
	// Should have no calls since file doesn't exist
	if len(calls) != 0 {
		t.Fatalf("Expected 0 calls when file missing, got %d", len(calls))
	}
}