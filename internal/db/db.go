package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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

type Pool struct {
	pool *pgxpool.Pool
}

type Test struct {
	ID   int
	Name string
}

func NewPool(ctx context.Context) (*Pool, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Create pool error %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("Connect to database %w", err)
	}

	return &Pool{pool: pool}, nil
}

func (p *Pool) Close() {
	p.pool.Close()
}

func (p *Pool) AddTest(ctx context.Context, name string) (Test, error) {
	var test Test

	err := p.pool.QueryRow(
		ctx,
		`INSERT INTO test (name)
		VALUES ($1)
		RETURNING id, name`,
		name,
	).Scan(&test.ID, &test.Name)

	if err != nil {
		return test, fmt.Errorf("Insert test %w", err)
	}

	return test, nil

}

func (p *Pool) ListTest(ctx context.Context) ([]Test, error) {
	rows, err := p.pool.Query(
		ctx,
		`SELECT * FROM test`,
	)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("List test %w", err)
	}

	var tests []Test

	for rows.Next() {
		var test Test
		if err := rows.Scan(&test.ID, &test.Name); err != nil {
			return nil, fmt.Errorf("scanning rows %w", err)
		}

		tests = append(tests, test)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate rows %w", err)
	}

	return tests, nil
}
