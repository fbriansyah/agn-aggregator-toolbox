package mariadb

import (
	"context"
	"fmt"
)

const getProducts = `-- name: ListProduk :many
SELECT 
	kode_produk, 
	nama_produk, 
	nama_produk_display, 
	nama_struk, 
	nama_struk_singkatan, 
	produk_alias, 
	produk_group, 
	produk_index, 
	produk_date, 
	label_idpel, 
	label_option, 
	status 
FROM m_produk 
WHERE %s
ORDER BY NAMA_PRODUK
`

func (q *Queries) GetProducts(ctx context.Context, whereQuery string) ([]MProduk, error) {
	query := fmt.Sprintf(getProducts, whereQuery)

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MProduk{}
	for rows.Next() {
		var i MProduk
		if err := rows.Scan(
			&i.KodeProduk,
			&i.NamaProduk,
			&i.NamaProdukDisplay,
			&i.NamaStruk,
			&i.NamaStrukSingkatan,
			&i.ProdukAlias,
			&i.ProdukGroup,
			&i.ProdukIndex,
			&i.ProdukDate,
			&i.LabelIdpel,
			&i.LabelOption,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
