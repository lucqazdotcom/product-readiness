package ingestion

import (
	"fmt"
	"log"
	"path/filepath"
	csv "product-readiness/internal/ingestion/csv"
)

type IngestionOutput struct {
	data []string
}

func Ingestion(dataDir string) error {

	products, err := csv.Products(filepath.Join(dataDir, "product_data.csv"))
	if err != nil {
		log.Fatal("failed to run products ingestion: %w", err)
	}

	product_metas, err := csv.ProductMeta(filepath.Join(dataDir, "product_descriptions.csv"))
	if err != nil {
		log.Fatal("failed to run product metadata ingestion: %w", err)
	}

	images, err := csv.Images(filepath.Join(dataDir, "images.csv"))
	if err != nil {
		log.Fatal("failed to run images ingestion: %w", err)
	}

	inventory, err := csv.Inventory(filepath.Join(dataDir, "inventory.csv"))
	if err != nil {
		log.Fatal("failed to run inventory ingestion: %w", err)
	}

	launch_dates, err := csv.Launch_dates(filepath.Join(dataDir, "launch_dates.csv"))
	if err != nil {
		log.Fatal("failed to run launch dates ingestion: %w", err)
	}

	fmt.Printf("loaded %d products \n", len(products))
	fmt.Printf("loaded %d product metadata \n", len(product_metas))
	fmt.Printf("loaded %d images \n", len(images))
	fmt.Printf("loaded %d inventory \n", len(inventory))
	fmt.Printf("loaded %d launch_dates \n", len(launch_dates))

	return nil
}
