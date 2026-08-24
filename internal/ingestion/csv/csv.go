package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CsvReader(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %w csv: %w", path, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read %w csv: %w", path, err)
	}

	return records, nil

}
