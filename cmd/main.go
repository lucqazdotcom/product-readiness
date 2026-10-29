package main

import (
	// api "product-readiness/internal/api"
	"context"
	"fmt"
	"log"
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

	test, err := database.AddTest(context.Background(), "another row")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Printf("test id: %v, test name: %v", test.ID, test.Name)

	tests, err := database.ListTest(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Printf("tests", tests)

	// api.Api()

	// if err := ingestion.Ingestion("testdata"); err != nil {
	// 	log.Fatal(err)
	// }

}
