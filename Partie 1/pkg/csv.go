package csv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type JSONRecord struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Email           string `json:"email"`
	InscriptionDate string `json:"inscription_date"`
}

var layouts = []string{
	time.RFC3339,          // "2006-01-02T15:04:05Z07:00"
	"2006-01-02",          // "YYYY-MM-DD"
	"02/01/2006",          // "DD/MM/YYYY"
	"01/02/2006",          // "MM/DD/YYYY"
	"2006-01-02 15:04:05", // "YYYY-MM-DD HH:MM:SS"
	"02/01/2006 15:04:05", // "DD/MM/YYYY HH:MM:SS"
	"01/02/2006 15:04:05", // "MM/DD/YYYY HH:MM:SS"
}

// ParseCSV parses the given CSV file into a slice of slices of strings.
func ParseCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records [][]string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, strings.Split(line, ","))
	}

	return records[1:], nil
}

// AppendCSV appends the given CSV file to the given records (slice of slices
// of strings).
func AppendCSV(filename string, records [][]string) ([][]string, error) {
	newRecords, err := ParseCSV(filename)
	if err != nil {
		return nil, err
	}

	records = append(records, newRecords...)

	return records, nil
}

func ConvertCSVToJSON(CSVpath, JSONFilename string) error {
	if CSVpath == "" {
		err := errNoPathGiven
		return err
	}

	files, err := GetFiles(CSVpath)
	if err != nil {
		return err
	}

	var records [][]string

	for _, file := range files {
		records, err = AppendCSV(file, records)
		if err != nil {
			return err
		}
	}

	var jsonRecords []JSONRecord
	for _, record := range records {
		jsonRecord, err := convertRecordToJSON(record)
		if err != nil {
			return err
		}
		jsonRecords = append(jsonRecords, jsonRecord)
	}

	// write json records into a json file.

	return nil
}

func convertRecordToJSON(record []string) (JSONRecord, error) {
	var jsonRecord JSONRecord

	if len(record) < 4 {
		err := fmt.Errorf("%v : %v", errInvalidRecord, record)
		return jsonRecord, err
	}

	formattedDate, err := formatDate(record[3])
	if err != nil {
		return jsonRecord, err
	}

	jsonRecord = JSONRecord{
		FirstName:       record[0],
		LastName:        record[1],
		Email:           record[2],
		InscriptionDate: formattedDate,
	}

	return jsonRecord, nil
}

// GetFiles is a helper function that returns a list containing the files from
// the given path. If the given path lead to a file, the list only contains the
// file.
func GetFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var csvFiles []string

	if info.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".csv") {
				csvFiles = append(csvFiles, filepath.Join(path, file.Name()))
			}
		}
	} else {
		if strings.HasSuffix(path, ".csv") {
			csvFiles = append(csvFiles, path)
		}
	}

	return csvFiles, nil
}

func formatDate(date string) (string, error) {
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, date)
		if err == nil {
			return parsedTime.Format("2006-01-02 15:04:05"), nil
		}
	}

	return "", errUnsupportedDateLayout
}
