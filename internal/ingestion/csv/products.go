package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type ProductInput struct {
	UPC      string
	Name     string
	Brand    string
	Category string
}

func Products(path string) ([]ProductInput, error) {

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

	products := make([]ProductInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 4

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		product := ProductInput{
			UPC:      row[0],
			Name:     row[1],
			Brand:    row[2],
			Category: row[3],
		}

		products = append(products, product)
	}

	return products, nil

}
