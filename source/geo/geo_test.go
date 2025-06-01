package geo

import (
	"os"
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestGeoProcess(t *testing.T) {
	// Create test CSV data with lat/lon at positions 9 and 10
	testData := `callsign,name,first,last,addr,city,state,zip,country,lat,lon,grid
W5TEST,John Smith,John,Smith,123 Main St,Anytown,TX,12345,US,32.5,-97.0,EM12ab
KC5ABC,Jane Doe,Jane,Doe,456 Oak Ave,Somewhere,CA,54321,US,34.0,-118.25,DM04aa
N0DEF,Bob Wilson,Bob,Wilson,789 Pine Rd,Nowhere,FL,98765,US,25.5,-80.25,EL95ab`

	// Create test file
	testFile, err := os.CreateTemp("", "geo_test_*.csv")
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
	tempDir, _ := os.MkdirTemp("", "geo_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "ham-stations.csv")

	// Test processing
	calls := make(map[string]data.HamCall)
	Process(&calls)

	// Verify results - should be 3 calls (header is skipped due to invalid coordinates)
	if len(calls) != 3 {
		t.Fatalf("Expected 3 calls (header skipped), got %d", len(calls))
	}

	// Header should NOT exist since lat/lon are invalid
	_, headerExists := calls["callsign"]
	if headerExists {
		t.Fatalf("Header line should NOT be processed due to invalid coordinates")
	}

	// Test W5TEST location
	w5test, exists := calls["W5TEST"]
	if !exists {
		t.Fatalf("W5TEST not found in processed calls")
	}
	if w5test.Location == nil {
		t.Fatalf("W5TEST location should not be nil")
	}
	if w5test.Location.Latitude != 32.5 {
		t.Fatalf("Expected W5TEST latitude 32.5, got %f", w5test.Location.Latitude)
	}
	if w5test.Location.Longitude != -97.0 {
		t.Fatalf("Expected W5TEST longitude -97.0, got %f", w5test.Location.Longitude)
	}

	// Test KC5ABC location
	kc5abc, exists := calls["KC5ABC"]
	if !exists {
		t.Fatalf("KC5ABC not found in processed calls")
	}
	if kc5abc.Location == nil {
		t.Fatalf("KC5ABC location should not be nil")
	}
	if kc5abc.Location.Latitude != 34.0 {
		t.Fatalf("Expected KC5ABC latitude 34.0, got %f", kc5abc.Location.Latitude)
	}
	if kc5abc.Location.Longitude != -118.25 {
		t.Fatalf("Expected KC5ABC longitude -118.25, got %f", kc5abc.Location.Longitude)
	}

	// Test N0DEF location
	n0def, exists := calls["N0DEF"]
	if !exists {
		t.Fatalf("N0DEF not found in processed calls")
	}
	if n0def.Location == nil {
		t.Fatalf("N0DEF location should not be nil")
	}
	if n0def.Location.Latitude != 25.5 {
		t.Fatalf("Expected N0DEF latitude 25.5, got %f", n0def.Location.Latitude)
	}
	if n0def.Location.Longitude != -80.25 {
		t.Fatalf("Expected N0DEF longitude -80.25, got %f", n0def.Location.Longitude)
	}
}

func TestGeoProcessWithExistingCalls(t *testing.T) {
	// Create test CSV data
	testData := `callsign,name,first,last,addr,city,state,zip,country,lat,lon,grid
W5TEST,John Smith,John,Smith,123 Main St,Anytown,TX,12345,US,32.5,-97.0,EM12ab`

	// Create test file
	testFile, err := os.CreateTemp("", "geo_test_*.csv")
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
	tempDir, _ := os.MkdirTemp("", "geo_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "ham-stations.csv")

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

	// Verify results - should have 1 call (header skipped due to invalid coordinates)
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call (header skipped), got %d", len(calls))
	}

	w5test := calls["W5TEST"]
	
	// Check that existing data is preserved
	if w5test.Name != "John Smith" {
		t.Fatalf("Expected name 'John Smith', got '%s'", w5test.Name)
	}
	if w5test.State != "TX" {
		t.Fatalf("Expected state 'TX', got '%s'", w5test.State)
	}
	
	// Check that location data was added
	if w5test.Location == nil {
		t.Fatalf("Location should not be nil")
	}
	if w5test.Location.Latitude != 32.5 {
		t.Fatalf("Expected latitude 32.5, got %f", w5test.Location.Latitude)
	}
	if w5test.Location.Longitude != -97.0 {
		t.Fatalf("Expected longitude -97.0, got %f", w5test.Location.Longitude)
	}
}

func TestGeoProcessInvalidCoordinates(t *testing.T) {
	// Create test data with invalid coordinates
	testData := `callsign,name,first,last,addr,city,state,zip,country,lat,lon,grid
W5TEST,John Smith,John,Smith,123 Main St,Anytown,TX,12345,US,invalid,-97.0,EM12ab
KC5ABC,Jane Doe,Jane,Doe,456 Oak Ave,Somewhere,CA,54321,US,34.0,invalid,DM04aa
N0DEF,Bob Wilson,Bob,Wilson,789 Pine Rd,Nowhere,FL,98765,US,25.5,-80.25,EL95ab`

	// Create test file
	testFile, err := os.CreateTemp("", "geo_test_*.csv")
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
	tempDir, _ := os.MkdirTemp("", "geo_test_dir_*")
	defer os.RemoveAll(tempDir)
	
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	// Create symlink to test file
	os.Symlink(testFile.Name(), "ham-stations.csv")

	// Test processing
	calls := make(map[string]data.HamCall)
	Process(&calls)

	// Should only process valid records - only N0DEF (header, W5TEST and KC5ABC have invalid coords)
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call (only N0DEF valid), got %d", len(calls))
	}

	// N0DEF should exist with valid coordinates
	n0def, exists := calls["N0DEF"]
	if !exists {
		t.Fatalf("N0DEF should exist - has valid coordinates")
	}
	if n0def.Location == nil {
		t.Fatalf("N0DEF location should not be nil")
	}
	if n0def.Location.Latitude != 25.5 {
		t.Fatalf("Expected N0DEF latitude 25.5, got %f", n0def.Location.Latitude)
	}

	// W5TEST and KC5ABC should not exist due to invalid coordinates
	_, w5testExists := calls["W5TEST"]
	if w5testExists {
		t.Fatalf("W5TEST should not exist due to invalid latitude")
	}

	_, kc5abcExists := calls["KC5ABC"]
	if kc5abcExists {
		t.Fatalf("KC5ABC should not exist due to invalid longitude")
	}
}

func TestGeoProcessMissingFile(t *testing.T) {
	// Change to a temporary directory where ham-stations.csv doesn't exist
	originalDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "geo_test_dir_*")
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