package csv

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var layouts = []string{
	time.RFC3339, // "2006-01-02T15:04:05Z07:00"
	"02/01/2006", // "DD/MM/YYYY"
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

func ConvertCSVToJSON(CSVpath string) error {
	stats = newStat()

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
		records, err = ParseCSV(file)
		if err != nil {
			return err
		}

		stats.CSVLines += len(records)

		jsonRecords := make(JSONRecordDictionary)
		for _, record := range records {
			jsonRecord, err := convertRecordToJSON(record)
			if err != nil {
				if errors.Is(err, errInvalidRecord) {
					stats.InvalidRecords++
				} else {
					return fmt.Errorf("%v on record: %v", err, record)
				}
			}
			if jsonRecord.isValid() {
				stats.JSONLines += 1
				jsonRecords.add(jsonRecord)
			} else {
				stats.InvalidRecords++
			}
		}

		jsonContacts := jsonRecords.toContact()
		jsonContacts.sort()

		jsonFilename := strings.TrimSuffix(filepath.Base(file), ".csv") + ".json"

		jsonFile, err := os.Create(jsonFilename)
		if err != nil {
			return err
		}
		defer jsonFile.Close()

		encoder := json.NewEncoder(jsonFile)
		err = encoder.Encode(jsonContacts)
		if err != nil {
			return err
		}

		fmt.Printf("Converted %s to JSON\n", file)

	}

	stats.Print()

	return nil
}

func convertRecordToJSON(record []string) (*JSONRecord, error) {
	var jsonRecord JSONRecord

	if len(record) < 4 {
		err := fmt.Errorf("%v : %v", errInvalidRecord, record)
		return nil, err
	}

	parsedTime, err := parseDate(record[3])
	if err != nil {
		// if an error occurred, use an empty string
		fmt.Println(fmt.Errorf("error while parsing date:%v", err))
		return nil, errInvalidRecord
	}
	if parsedTime == nil {
		return nil, errInvalidRecord
	}

	jsonRecord = JSONRecord{
		FirstName:       record[0],
		LastName:        record[1],
		Email:           record[2],
		InscriptionDate: *parsedTime,
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

func parseDate(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, date)
		if err == nil {
			return &parsedTime, nil
		}
	}
	if stats != nil {
		stats.UnsupportedDateLayout++
	}

	return nil, fmt.Errorf("%v: %v", errUnsupportedDateLayout, date)
}
