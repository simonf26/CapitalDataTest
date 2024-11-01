package csv

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type JSONRecord struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Email           string `json:"email"`
	InscriptionDate string `json:"inscription_date"`
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

	return nil
}

func convertRecordToJSON(record []string) (*JSONRecord, error) {
	if len(record) < 4 {
		err := fmt.Errorf("%v : %v", errInvalidRecord, record)
		return nil, err
	}
	jsonRecord := JSONRecord{
		FirstName:       record[0],
		LastName:        record[1],
		Email:           record[2],
		InscriptionDate: record[3],
	}

	return &jsonRecord, nil
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
