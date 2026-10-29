package main

import (
	"context"
	"log"
	api "product-readiness/internal/api"
	db "product-readiness/internal/db"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database, err := db.NewPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// name := "Test product"
	//
	// product := db.Product{
	// 	SKU:  "TEST-SKU-001",
	// 	UPC:  "TEST-UPC-001",
	// 	Name: &name,
	// }
	//
	// if err := database.UpsertProduct(ctx, product); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println("Product saved")

	// test, err := database.AddTest(context.Background(), "another row")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()
	//
	// fmt.Printf("test id: %v, test name: %v", test.ID, test.Name)
	//
	// tests, err := database.ListTest(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()
	//
	// fmt.Printf("tests: %+v\n", tests)

	api.Api(database)

	// if err := ingestion.Ingestion("testdata"); err != nil {
	// 	log.Fatal(err)
	// }

}
