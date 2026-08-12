package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type ImageInput struct {
	ID         string
	SKU        string
	Image_url  string
	Image_role string
	Sort_order string
	Alt_text   string
	Width_px   string
	Height_px  string
	Is_active  string
	Updated_at string
}

func Images(path string) ([]ImageInput, error) {
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

	images := make([]ImageInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 10

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		image := ImageInput{
			ID:         row[0],
			SKU:        row[1],
			Image_url:  row[2],
			Image_role: row[3],
			Sort_order: row[4],
			Alt_text:   row[5],
			Width_px:   row[6],
			Height_px:  row[7],
			Is_active:  row[8],
			Updated_at: row[9],
		}

		images = append(images, image)
	}

	return images, nil

}
