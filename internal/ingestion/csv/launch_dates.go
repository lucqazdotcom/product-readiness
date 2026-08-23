package csv

import (
	"fmt"
)

type DateInput struct {
	SKU         string
	Launch_week string
	Launch_date string
}

func Launch_dates(path string) ([]DateInput, error) {

	records, err := CsvReader(path)
	if err != nil {
		return nil, fmt.Errorf("read launch dates csv: %w", err)
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
