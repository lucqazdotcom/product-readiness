package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type DateInput struct {
	SKU         string
	Launch_week string
	Launch_date string
}

func Launch_dates(path string) ([]DateInput, error) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading csv: %w", err)
	}

	dates := make([]DateInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 3

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		date := DateInput{
			SKU:         row[0],
			Launch_week: row[1],
			Launch_date: row[2],
		}

		dates = append(dates, date)
	}

	return dates, nil

}
