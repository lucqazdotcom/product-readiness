package ingestion

import (
	"context"
	"fmt"
	"path/filepath"
	db "product-readiness/internal/db"
	csv "product-readiness/internal/ingestion/csv"
	"strconv"
	"time"
)

type Feed struct {
	Name     string
	FileName string
}

type IngestionOutput struct {
	FeedName string
	Records  [][]string
}

func Ingestion(ctx context.Context, dataDir string, database *db.Pool) ([]IngestionOutput, error) {

	feeds := []Feed{
		{Name: "product", FileName: "product_data.csv"},
		{Name: "product_metas", FileName: "product_descriptions.csv"},
		{Name: "images", FileName: "images.csv"},
		{Name: "inventory", FileName: "inventory.csv"},
		{Name: "launch_dates", FileName: "launch_dates.csv"},
	}

	outputs := []IngestionOutput{}

	for _, feed := range feeds {
		records, err := csv.CsvReader(filepath.Join(dataDir, feed.FileName))
		if err != nil {
			return nil, fmt.Errorf("failed to run %s ingestion: %w", feed.Name, err)
		}

		data := IngestionOutput{
			FeedName: feed.Name,
			Records:  records,
		}

		if feed.Name == "product" {
			if len(records) == 0 {
				return nil, fmt.Errorf("product feed is empty")
			}
			for i, record := range records[1:] {
				product, err := parseProduct(record)
				if err != nil {
					return nil, fmt.Errorf("product row %d: %w", i+2, err)
				}
				if err := database.UpsertProduct(ctx, product); err != nil {
					return nil, fmt.Errorf("product row %d: %w", i+2, err)
				}
			}
		}

		outputs = append(outputs, data)

		fmt.Printf("loaded %d %s's \n", len(records), feed.Name)
	}

	return outputs, nil
}

func parseProduct(record []string) (db.Product, error) {
	if len(record) != 16 {
		return db.Product{}, fmt.Errorf("expected 16 columns, got %d", len(record))
	}

	product := db.Product{
		SKU:         record[0],
		UPC:         record[1],
		Name:        optionalString(record[2]),
		Brand:       optionalString(record[3]),
		Category:    optionalString(record[4]),
		Subcategory: optionalString(record[5]),
		Color:       optionalString(record[6]),
		Size:        optionalString(record[7]),
		Currency:    optionalString(record[10]),
		Status:      optionalString(record[12]),
		SupplierID:  optionalString(record[13]),
	}
	if product.SKU == "" || product.UPC == "" {
		return db.Product{}, fmt.Errorf("sku and upc are required")
	}

	if record[8] != "" {
		if err := product.UnitPrice.Scan(record[8]); err != nil {
			return db.Product{}, fmt.Errorf("invalid unit price %q: %w", record[8], err)
		}
	}
	if record[9] != "" {
		if err := product.CostPrice.Scan(record[9]); err != nil {
			return db.Product{}, fmt.Errorf("invalid cost price %q: %w", record[9], err)
		}
	}
	if record[11] != "" {
		taxable, err := strconv.ParseBool(record[11])
		if err != nil {
			return db.Product{}, fmt.Errorf("invalid taxable value %q: %w", record[11], err)
		}
		product.Taxable = &taxable
	}

	createdAt, err := optionalTimestamp(record[14])
	if err != nil {
		return db.Product{}, fmt.Errorf("invalid created_at: %w", err)
	}
	product.CreatedAt = createdAt

	updatedAt, err := optionalTimestamp(record[15])
	if err != nil {
		return db.Product{}, fmt.Errorf("invalid updated_at: %w", err)
	}
	product.UpdatedAt = updatedAt

	return product, nil
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func optionalTimestamp(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}
