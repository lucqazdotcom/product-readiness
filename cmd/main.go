package main

import (
	// api "product-readiness/internal/api"
	db "product-readiness/internal/db"
)

func main() {

	db.Db()

	// api.Api()

	// if err := ingestion.Ingestion("testdata"); err != nil {
	// 	log.Fatal(err)
	// }

}
