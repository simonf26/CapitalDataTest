package csv

import (
	"os"
	"testing"
)

func TestParseCSV(t *testing.T) {
	records, err := ParseCSV("20180101_115200_contactstream.csv")
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

	records, err = ParseCSV("20180101_132200_contactstream2.csv")
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

	records, err = ParseCSV("20180102_140045_contactstream3.csv")
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

	records, err = ParseCSV("20180204_120204_contactstream4.csv")
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
