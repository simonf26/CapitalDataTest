package csv

import (
	"os"
	"testing"
)

var basePath = "../"

func TestParseCSV(t *testing.T) {
	// Test case: unexisting file
	_, err := ParseCSV("unexisting_file.csv")
	if !os.IsNotExist(err) {
		t.Fatal("Parsing unexisting file should return an error")
	}

	records, err := ParseCSV(basePath + "20180101_115200_contactstream.csv")
	// Ensure the given file exists
	if os.IsNotExist(err) {
		t.Fatalf("Error reading file: %v", err)
	}
	// Ensure the parsing succeeded
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}
	// Ensure the parsing result is not empty
	if len(records) == 0 {
		t.Fatal("Expected non-empty CSV records")
	}

	t.Log("ParseCSV parsing successfully completed")

	// Run the same tests on different files.
	records, err = ParseCSV(basePath + "20180101_132200_contactstream2.csv")
	if os.IsNotExist(err) {
		t.Fatalf("Error reading file: %v", err)
	}
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}
	if len(records) == 0 {
		t.Fatal("Expected non-empty CSV records")
	}

	t.Log("ParseCSV parsing successfully completed")

	records, err = ParseCSV(basePath + "20180102_140045_contactstream3.csv")
	if os.IsNotExist(err) {
		t.Fatalf("Error reading file: %v", err)
	}
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}
	if len(records) == 0 {
		t.Fatal("Expected non-empty CSV records")
	}

	t.Log("ParseCSV parsing successfully completed")

	records, err = ParseCSV(basePath + "20180204_120204_contactstream4.csv")
	if os.IsNotExist(err) {
		t.Fatalf("Error reading file: %v", err)
	}
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}
	if len(records) == 0 {
		t.Fatal("Expected non-empty CSV records")
	}

	t.Log("ParseCSV parsing successfully completed")
}

func TestAppendCSV(t *testing.T) {
	// Get records to append
	records, err := ParseCSV(basePath + "20180101_115200_contactstream.csv")
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}

	// Test case: unexisting file
	_, err = AppendCSV("unexisting_file.csv", records)
	if !os.IsNotExist(err) {
		t.Fatal("Appending unexisting file should return an error")
	}

	// Test case: nominal
	appendedRecords, err := AppendCSV(basePath+"20180101_132200_contactstream2.csv", records)
	if err != nil {
		t.Fatalf("Error appending CSV: %v", err)
	}
	if len(appendedRecords) == 0 {
		t.Fatal("Expected non-empty CSV records after appending")
	}

	// get new records separatly for comparison purposes
	newRecords, err := ParseCSV(basePath + "20180101_132200_contactstream2.csv")
	if err != nil {
		t.Fatalf("Couldn't parse CSV file: %v", err)
		return
	}

	if len(appendedRecords) != len(records)+len(newRecords) {
		t.Fatalf("Expected %d records after appending, got %d", len(records)+len(newRecords), len(appendedRecords))
	}
	if !sliceEqual(appendedRecords[:len(records)-1], records) {
		t.Fatal("Original CSV records are not all in the resulting table")
	}
	if !sliceEqual(appendedRecords[len(records):], newRecords) {
		t.Fatal("Resulting table doesn't match the expected appended records.")
	}

	t.Log("AppendCSV appending successfully completed")
}

func TestConvertCSVToJSON(t *testing.T) {
	// Test case: unexisting file
	err := ConvertCSVToJSON(
		"unexisting_file.csv",
		basePath+"output.json",
	)
	if err == nil {
		t.Fatal("Converting unexisting file should return an error")
	}

	// Test case: unexisting folder
	err = ConvertCSVToJSON(
		"unexisting_folder",
		basePath+"output.json",
	)
	if err == nil {
		t.Fatal("Converting on unexisting folder should return an error")
	}

	// Test case: nominal on CSV file
	err = ConvertCSVToJSON(
		basePath+"20180101_115200_contactstream.csv",
		basePath+"output.json",
	)
	if err != nil {
		t.Fatalf("Error converting CSV to JSON: %v", err)
	}

	// Test case: nominal on folder
	err = ConvertCSVToJSON(
		basePath,
		basePath+"output_folder",
	)
	if err != nil {
		t.Fatalf("Error converting folder to JSON: %v", err)
	}

	t.Log("ConvertCSVToJSON convertion was successful")
}

// helper function to check that 2 slices are equal
func sliceEqual(a, b [][]string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
