package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type InventoryInput struct {
	Inventory_record_id string
	SKU                 string
	Warehouse_code      string
	On_hand_qty         string
	Reserved_qty        string
	Available_qty       string
	Reorder_point       string
	Backorder_allowed   string
	Inventory_status    string
	Last_counted_at     string
	Feed_updated_at     string
}

func Inventory(path string) ([]InventoryInput, error) {
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

	inventories := make([]InventoryInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 11

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		inventory := InventoryInput{
			Inventory_record_id: row[0],
			SKU:                 row[1],
			Warehouse_code:      row[2],
			On_hand_qty:         row[3],
			Reserved_qty:        row[4],
			Available_qty:       row[5],
			Reorder_point:       row[6],
			Backorder_allowed:   row[7],
			Inventory_status:    row[8],
			Last_counted_at:     row[9],
			Feed_updated_at:     row[10],
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil

}
