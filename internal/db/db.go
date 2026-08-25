package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Images struct {
	ID         string
	SKU        string
	Image_url  string
	Image_role string
	Sort_order string
	Alt_text   string
	Width_px   string
	Height_px  string
	Is_active  string
	Updated_at string
}

type Iventory struct {
	Inventory_record_id string
	SKU                 string
	Warehouse_code      string
	On_hand_qty         string
	Reserved_qty        string
	Available_qty       string
	Reorder_point       string
	Backorder_allowed   string
	Inventory_status    string
	Last_counted_at     string
	Feed_updated_at     string
}

type DateInput struct {
	SKU         string
	Launch_week string
	Launch_date string
}

type ProductMetaData struct {
	Record_id         string
	SKU               string
	Locale            string
	Short_desc        string
	Long_desc         string
	Material          string
	Care_instructions string
	Bullet_1          string
	Bullet_2          string
	Seo_title         string
	Seo_description   string
	Approval_status   string
	Updated_at        string
}

type Products struct {
	SKU            string
	Name           string
	Brand          string
	Category       string
	Subcategory    string
	Colour         string
	Size           string
	Unit_price_cad string
	Cost_cad       string
	Currency       string
	Taxable        string
	Status         string
	Supplier_id    string
	Created_at     string
	Updated_at     string
}

func Db() {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from test")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var name string

		if err = rows.Scan(&id, &name); err != nil {
			fmt.Println("error")
		}

		fmt.Println(id, name)
	}

}
