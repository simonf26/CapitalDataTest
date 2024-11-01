package csv

import (
	"bufio"
	"os"
	"strings"
)

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

func AppendCSV(filename string, records [][]string) ([][]string, error) {
	newRecords, err := ParseCSV(filename)
	if err != nil {
		return nil, err
	}

	records = append(records, newRecords...)

	return records, nil
}
