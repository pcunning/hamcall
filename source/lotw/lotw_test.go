package lotw

import (
	"os"
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestLotwProcess(t *testing.T) {
	// Create test CSV data
	testData := `callsign,qso_upload_date,last_upload_date
W5TEST,2023-01-15,2023-01-15
KC5ABC,2022-12-01,2023-01-10
N0DEF,2023-02-28,2023-02-28`

	// Create test file
	testFile, err := os.CreateTemp("", "lotw_test_*.csv")
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
	tempDir, _ := os.MkdirTemp("", "lotw_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "lotw.csv")

	// Test processing
	calls := make(map[string]data.HamCall)
	Process(&calls)

	// Verify results - should be 4 calls (3 real + 1 header)
	if len(calls) != 4 {
		t.Fatalf("Expected 4 calls (including header), got %d", len(calls))
	}

	// Check that header was processed as a "callsign"
	_, headerExists := calls["callsign"]
	if !headerExists {
		t.Fatalf("Header line should be processed as callsign 'callsign'")
	}

	// Test W5TEST
	w5test, exists := calls["W5TEST"]
	if !exists {
		t.Fatalf("W5TEST not found in processed calls")
	}
	expectedLOTW := "2023-01-152023-01-15"  // qso_upload_date + last_upload_date
	if w5test.LOTW != expectedLOTW {
		t.Fatalf("Expected W5TEST LOTW '%s', got '%s'", expectedLOTW, w5test.LOTW)
	}

	// Test KC5ABC
	kc5abc, exists := calls["KC5ABC"]
	if !exists {
		t.Fatalf("KC5ABC not found in processed calls")
	}
	expectedLOTW = "2022-12-012023-01-10"
	if kc5abc.LOTW != expectedLOTW {
		t.Fatalf("Expected KC5ABC LOTW '%s', got '%s'", expectedLOTW, kc5abc.LOTW)
	}

	// Test N0DEF
	n0def, exists := calls["N0DEF"]
	if !exists {
		t.Fatalf("N0DEF not found in processed calls")
	}
	expectedLOTW = "2023-02-282023-02-28"
	if n0def.LOTW != expectedLOTW {
		t.Fatalf("Expected N0DEF LOTW '%s', got '%s'", expectedLOTW, n0def.LOTW)
	}
}

func TestLotwProcessWithExistingCalls(t *testing.T) {
	// Create test CSV data
	testData := `callsign,qso_upload_date,last_upload_date
W5TEST,2023-01-15,2023-01-15`

	// Create test file
	testFile, err := os.CreateTemp("", "lotw_test_*.csv")
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
	tempDir, _ := os.MkdirTemp("", "lotw_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "lotw.csv")

	// Pre-populate calls map with existing data
	calls := make(map[string]data.HamCall)
	calls["W5TEST"] = data.HamCall{
		Callsign:  "W5TEST",
		Name:      "John Smith", 
		FirstName: "John",
		LastName:  "Smith",
		State:     "TX",
	}

	// Test processing
	Process(&calls)

	// Verify results - should have 2 calls (1 existing + 1 header)
	if len(calls) != 2 {
		t.Fatalf("Expected 2 calls (1 existing + header), got %d", len(calls))
	}

	w5test := calls["W5TEST"]
	
	// Check that existing data is preserved
	if w5test.Name != "John Smith" {
		t.Fatalf("Expected name 'John Smith', got '%s'", w5test.Name)
	}
	if w5test.State != "TX" {
		t.Fatalf("Expected state 'TX', got '%s'", w5test.State)
	}
	
	// Check that LOTW data was added
	expectedLOTW := "2023-01-152023-01-15"
	if w5test.LOTW != expectedLOTW {
		t.Fatalf("Expected LOTW '%s', got '%s'", expectedLOTW, w5test.LOTW)
	}
}

func TestLotwProcessMissingFile(t *testing.T) {
	// Change to a temporary directory where lotw.csv doesn't exist
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "lotw_test_dir_*")
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