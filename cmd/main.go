package main

import (
	"log"
	ingestion "product-readiness/internal/ingestion"
)

func main() {

	if err := ingestion.Ingestion("testdata"); err != nil {
		log.Fatal(err)
	}

}
