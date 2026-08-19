package main

import (
	"fmt"
	"log"
	csv "product-readiness/internal/ingestion/csv"
)

func main() {
	products, err := csv.Products("testdata/product_data.csv")
	product_metas, err := csv.ProductMeta("testdata/product_descriptions.csv")
	images, err := csv.Images("testdata/images.csv")
	inventory, err := csv.Inventory("testdata/inventory.csv")
	launch_dates, err := csv.Launch_dates("testdata/launch_dates.csv")

	if err != nil {
		log.Fatal("failed to run csv products: %w", err)
	}

	fmt.Printf("loaded %d products \n", len(products))
	fmt.Printf("loaded %d product metadata \n", len(product_metas))
	fmt.Printf("loaded %d images \n", len(images))
	fmt.Printf("loaded %d inventory \n", len(inventory))
	fmt.Printf("loaded %d launch_dates \n", len(launch_dates))

}
