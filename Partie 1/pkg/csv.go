package csv

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

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

	return nil
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
				csvFiles = append(csvFiles, file.Name())
			}
		}
	} else {
		if strings.HasSuffix(path, ".csv") {
			csvFiles = append(csvFiles, filepath.Base(path))
		}
	}

	return csvFiles, nil
}
