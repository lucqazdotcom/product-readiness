package ingestion

import (
	"fmt"
	"path/filepath"
	csv "product-readiness/internal/ingestion/csv"
)

type Feed struct {
	Name     string
	FileName string
}

type IngestionOutput struct {
	FeedName string
	Records  [][]string
}

func Ingestion(dataDir string) ([]IngestionOutput, error) {

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
			fmt.Errorf("failed to run %s ingestion: %w", feed.Name, err)
		}

		data := IngestionOutput{
			FeedName: feed.Name,
			Records:  records,
		}

		outputs = append(outputs, data)

		fmt.Printf("loaded %d %s's \n", len(records), feed.Name)
	}

	return outputs, nil
}
