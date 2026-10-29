package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Image struct {
	ID        string
	SKU       string
	ImageUrl  *string
	ImageRole *string
	SortOrder *int32
	AltText   *string
	WidthPx   *int32
	HeightPx  *int32
	IsActive  *bool
	UpdatedAt *time.Time
}

type Inventory struct {
	InventoryRecordID string
	SKU               string
	WarehouseCode     string
	OnHandQty         *int32
	ReservedQty       *int32
	AvailableQty      *int32
	ReorderPoint      *int32
	BackorderAllowed  *bool
	InventoryStatus   *string
	LastCountedAt     *time.Time
	FeedUpdatedAt     *time.Time
}

type LaunchDate struct {
	UPC        string
	LaunchWeek *string
	LaunchDate *time.Time
}

type ProductMetaData struct {
	RecordID         string
	SKU              string
	Locale           string
	ShortDesc        *string
	LongDesc         *string
	Material         *string
	CareInstructions *string
	Bullet1          *string
	Bullet2          *string
	SEOTitle         *string
	SEODescription   *string
	ApprovalStatus   *string
	UpdatedAt        *time.Time
}

type Product struct {
	SKU         string
	UPC         string
	Name        *string
	Brand       *string
	Category    *string
	Subcategory *string
	Color       *string
	Size        *string
	UnitPrice   pgtype.Numeric
	CostPrice   pgtype.Numeric
	Currency    *string
	Taxable     *bool
	Status      *string
	SupplierID  *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
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

func (p *Pool) UpsertProduct(ctx context.Context, product Product) error {
	_, err := p.pool.Exec(
		ctx,
		`INSERT INTO products (
			sku, upc, product_name, brand, category, subcategory,
			color, size, unit_price, cost_price, currency,
			taxable, status, supplier_id, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12, $13, $14, $15, $16
		)
		ON CONFLICT (sku) DO UPDATE SET
			upc = EXCLUDED.upc,
			product_name = EXCLUDED.product_name,
			brand = EXCLUDED.brand,
			category = EXCLUDED.category,
			subcategory = EXCLUDED.subcategory,
			color = EXCLUDED.color,
			size = EXCLUDED.size,
			unit_price = EXCLUDED.unit_price,
			cost_price = EXCLUDED.cost_price,
			currency = EXCLUDED.currency,
			taxable = EXCLUDED.taxable,
			status = EXCLUDED.status,
			supplier_id = EXCLUDED.supplier_id,
			created_at = EXCLUDED.created_at,
			updated_at = EXCLUDED.updated_at`,

		product.SKU,
		product.UPC,
		product.Name,
		product.Brand,
		product.Category,
		product.Subcategory,
		product.Color,
		product.Size,
		product.UnitPrice,
		product.CostPrice,
		product.Currency,
		product.Taxable,
		product.Status,
		product.SupplierID,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("upsert product %s: %w", product.SKU, err)
	}
	return nil
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

	if err != nil {
		return nil, fmt.Errorf("List test %w", err)
	}

	defer rows.Close()

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
