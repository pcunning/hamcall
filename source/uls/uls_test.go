package uls

import (
	"os"
	"strings"
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestProcessAM(t *testing.T) {
	// Create test directory structure
	testDir := "/tmp/uls_test"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create test AM.dat file
	amContent := `AM|1234567890|KNTF000001||W5TEST|T||0|||||||||||
AM|1234567891|KNTF000002||KC5ABC|G||0|||||||||||
AM|1234567892|KNTF000003||N0DEF|A||0|||||||||||`

	amFile, err := os.Create(testDir + "/AM.dat")
	if err != nil {
		t.Fatalf("Failed to create test AM.dat: %v", err)
	}
	defer amFile.Close()
	amFile.WriteString(amContent)

	// Change working directory temporarily to test directory
	originalDir, _ := os.Getwd()
	os.Chdir("/tmp")
	defer os.Chdir(originalDir)

	// Create l_amat directory structure
	os.MkdirAll("l_amat", 0755)
	defer os.RemoveAll("l_amat")
	os.Rename(testDir+"/AM.dat", "l_amat/AM.dat")

	calls := make(map[string]data.HamCall)
	ProcessAM(&calls)

	// Test that callsigns were processed
	if len(calls) != 3 {
		t.Fatalf("Expected 3 calls, got %d", len(calls))
	}

	// Test specific callsigns and classes
	if calls["W5TEST"].Class != "T" {
		t.Fatalf("Expected W5TEST class 'T', got '%s'", calls["W5TEST"].Class)
	}

	if calls["KC5ABC"].Class != "G" {
		t.Fatalf("Expected KC5ABC class 'G', got '%s'", calls["KC5ABC"].Class)
	}

	if calls["N0DEF"].Class != "A" {
		t.Fatalf("Expected N0DEF class 'A', got '%s'", calls["N0DEF"].Class)
	}
}

func TestProcessEN(t *testing.T) {
	// Create test directory structure
	testDir := "/tmp/uls_test"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create test EN.dat file (needs at least 25 fields, FRN at position 22)
	enContent := `EN|1234567890|KNTF000001||W5TEST|L||John|Smith|M|Johnson|||||123 Main St|Anytown|TX|12345||||0123456789|I||A
EN|1234567891|KNTF000002||KC5ABC|L||Jane|A|B|Doe|||||456 Oak Ave|Somewhere|CA|54321||||0987654321|I||A`

	enFile, err := os.Create(testDir + "/EN.dat")
	if err != nil {
		t.Fatalf("Failed to create test EN.dat: %v", err)
	}
	defer enFile.Close()
	enFile.WriteString(enContent)

	// Change working directory temporarily to test directory
	originalDir, _ := os.Getwd()
	os.Chdir("/tmp")
	defer os.Chdir(originalDir)

	// Create l_amat directory structure
	os.MkdirAll("l_amat", 0755)
	defer os.RemoveAll("l_amat")
	os.Rename(testDir+"/EN.dat", "l_amat/EN.dat")

	calls := make(map[string]data.HamCall)
	ProcessEN(&calls)

	// Test that callsigns were processed
	if len(calls) != 2 {
		t.Fatalf("Expected 2 calls, got %d", len(calls))
	}

	// Test W5TEST details
	w5test := calls["W5TEST"]
	if w5test.Name != "John" {
		t.Fatalf("Expected W5TEST name 'John', got '%s'", w5test.Name)
	}
	if w5test.FirstName != "Smith" {
		t.Fatalf("Expected W5TEST first name 'Smith', got '%s'", w5test.FirstName)
	}
	if w5test.LastName != "Johnson" {
		t.Fatalf("Expected W5TEST last name 'Johnson', got '%s'", w5test.LastName)
	}
	if w5test.State != "TX" {
		t.Fatalf("Expected W5TEST state 'TX', got '%s'", w5test.State)
	}
	if w5test.FRN != "0123456789" {
		t.Fatalf("Expected W5TEST FRN '0123456789', got '%s'", w5test.FRN)
	}
}

func TestProcessHD(t *testing.T) {
	// Create test directory structure
	testDir := "/tmp/uls_test"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create test HD.dat file (needs at least 43 fields for effective date at position 42)
	// Fields: 0=HD,1=unique_sys_id,2=file_num,3=ebf,4=callsign,5=status,6=service,7=grant,8=expiration,...,42=effective
	emptyFields := strings.Repeat("|", 33)  // Fields 9-41 (33 fields)
	hdContent := `HD|1234567890|KNTF000001||W5TEST|A|HA|01/01/2020|01/01/2030` + emptyFields + `|01/01/2020
HD|1234567891|KNTF000002||KC5ABC|A|HA|01/01/2020|01/01/2030` + emptyFields + `|01/01/2020`

	hdFile, err := os.Create(testDir + "/HD.dat")
	if err != nil {
		t.Fatalf("Failed to create test HD.dat: %v", err)
	}
	defer hdFile.Close()
	hdFile.WriteString(hdContent)

	// Change working directory temporarily to test directory
	originalDir, _ := os.Getwd()
	os.Chdir("/tmp")
	defer os.Chdir(originalDir)

	// Create l_amat directory structure
	os.MkdirAll("l_amat", 0755)
	defer os.RemoveAll("l_amat")
	os.Rename(testDir+"/HD.dat", "l_amat/HD.dat")

	calls := make(map[string]data.HamCall)
	ProcessHD(&calls)

	// Test that callsigns were processed
	if len(calls) != 2 {
		t.Fatalf("Expected 2 calls, got %d", len(calls))
	}

	// Test W5TEST details
	w5test := calls["W5TEST"]
	if w5test.Grant != "01/01/2020" {
		t.Fatalf("Expected W5TEST grant '01/01/2020', got '%s'", w5test.Grant)
	}
	if w5test.Expiration != "01/01/2030" {
		t.Fatalf("Expected W5TEST expiration '01/01/2030', got '%s'", w5test.Expiration)
	}
	if w5test.FileNumber != "KNTF000001" {
		t.Fatalf("Expected W5TEST file number 'KNTF000001', got '%s'", w5test.FileNumber)
	}
	if w5test.Effective != "01/01/2020" {
		t.Fatalf("Expected W5TEST effective '01/01/2020', got '%s'", w5test.Effective)
	}
}

func TestProcessIntegration(t *testing.T) {
	// Create test directory structure
	testDir := "/tmp/uls_test"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create test files
	amContent := `AM|1234567890|KNTF000001||W5TEST|T||0|||||||||||`
	enContent := `EN|1234567890|KNTF000001||W5TEST|L||John|Smith|M|Johnson|||||123 Main St|Anytown|TX|12345||||0123456789|I||A`
	// Create HD data with proper field count
	emptyFields := strings.Repeat("|", 33)  // 33 empty fields to get to position 42
	hdContent := `HD|1234567890|KNTF000001||W5TEST|A|HA|01/01/2020|01/01/2030` + emptyFields + `|01/01/2020`

	files := map[string]string{
		"AM.dat": amContent,
		"EN.dat": enContent,
		"HD.dat": hdContent,
	}

	for filename, content := range files {
		file, err := os.Create(testDir + "/" + filename)
		if err != nil {
			t.Fatalf("Failed to create test %s: %v", filename, err)
		}
		file.WriteString(content)
		file.Close()
	}

	// Change working directory temporarily to test directory
	originalDir, _ := os.Getwd()
	os.Chdir("/tmp")
	defer os.Chdir(originalDir)

	// Create l_amat directory structure
	os.MkdirAll("l_amat", 0755)
	defer os.RemoveAll("l_amat")
	for filename := range files {
		os.Rename(testDir+"/"+filename, "l_amat/"+filename)
	}

	calls := make(map[string]data.HamCall)
	
	// Process all three data types
	ProcessAM(&calls)
	ProcessEN(&calls)
	ProcessHD(&calls)

	// Test that the single callsign has data from all sources
	if len(calls) != 1 {
		t.Fatalf("Expected 1 call, got %d", len(calls))
	}

	w5test := calls["W5TEST"]
	
	// Test AM data
	if w5test.Class != "T" {
		t.Fatalf("Expected class 'T', got '%s'", w5test.Class)
	}
	
	// Test EN data
	if w5test.Name != "John" {
		t.Fatalf("Expected name 'John', got '%s'", w5test.Name)
	}
	if w5test.State != "TX" {
		t.Fatalf("Expected state 'TX', got '%s'", w5test.State)
	}
	
	// Test HD data
	if w5test.Grant != "01/01/2020" {
		t.Fatalf("Expected grant '01/01/2020', got '%s'", w5test.Grant)
	}
	if w5test.Expiration != "01/01/2030" {
		t.Fatalf("Expected expiration '01/01/2030', got '%s'", w5test.Expiration)
	}
}