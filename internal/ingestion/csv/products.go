package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type ProductInput struct {
	SKU            string
	Name           string
	Brand          string
	Category       string
	Subcategory    string
	Colour         string
	Size           string
	Unit_price_cad string
	Cost_cad       string
	Currency       string
	Taxable        string
	Status         string
	Supplier_id    string
	Created_at     string
	Updated_at     string
}

func Products(path string) ([]ProductInput, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open products csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read products csv: %w", err)
	}

	products := make([]ProductInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 16

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		product := ProductInput{
			SKU:            row[0],
			Name:           row[1],
			Brand:          row[2],
			Category:       row[3],
			Subcategory:    row[4],
			Colour:         row[5],
			Size:           row[6],
			Unit_price_cad: row[7],
			Cost_cad:       row[8],
			Currency:       row[9],
			Taxable:        row[10],
			Status:         row[11],
			Supplier_id:    row[12],
			Created_at:     row[13],
			Updated_at:     row[14],
		}

		products = append(products, product)
	}

	return products, nil

}
