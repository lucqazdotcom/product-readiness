package main

import (
	"fmt"
	"log"
	csv "product-readiness/internal/ingestion/csv"
)

func main() {
	products, err := csv.Products("testdata/product_data.csv")
	images, err := csv.Images("testdata/images.csv")

	if err != nil {
		log.Fatal("failed to run csv products: %w", err)
	}

	fmt.Printf("loaded %d products \n", len(products))
	fmt.Printf("loaded %d images \n", len(images))

}
