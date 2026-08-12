package main

import (
	"fmt"
	"log"
	csv "product-readiness/internal/ingestion/csv"
)

func main() {
	fmt.Println("hello")
	products, err := csv.Products("testdata/product_data.csv")

	if err != nil {
		log.Fatal("failed to run csv products: %w", err)
	}

	fmt.Printf("loaded %d products \n", len(products))

}
