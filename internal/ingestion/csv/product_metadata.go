package csv

import (
	"fmt"
)

type ProductMetaInput struct {
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

func ProductMeta(path string) ([]ProductMetaInput, error) {

	records, err := CsvReader(path)
	if err != nil {
		return nil, fmt.Errorf("read product metadata csv: %w", err)
	}

	product_metas := make([]ProductMetaInput, 0, len(records))

	for rowIndex, row := range records {

		const expectedColumns = 13

		if len(row) != expectedColumns {
			return nil, fmt.Errorf(
				"row %d: expected %d columns, received %d",
				rowIndex+1,
				expectedColumns,
				len(row),
			)
		}

		product_meta := ProductMetaInput{
			Record_id:         row[0],
			SKU:               row[1],
			Locale:            row[2],
			Short_desc:        row[3],
			Long_desc:         row[4],
			Material:          row[5],
			Care_instructions: row[6],
			Bullet_1:          row[7],
			Bullet_2:          row[8],
			Seo_title:         row[9],
			Seo_description:   row[10],
			Approval_status:   row[11],
			Updated_at:        row[12],
		}

		product_metas = append(product_metas, product_meta)
	}

	return product_metas, nil

}
