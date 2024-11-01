package csv

import (
	"os"
	"testing"
)

var basePath = "../"

func TestParseCSV(t *testing.T) {

	records, err := ParseCSV(basePath + "20180101_115200_contactstream.csv")
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
	records, err := ParseCSV(basePath + "20180101_115200_contactstream.csv")
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}

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
