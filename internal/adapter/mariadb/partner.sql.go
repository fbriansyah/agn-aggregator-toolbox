// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: partner.sql

package mariadb

import (
	"context"
)

const getPartners = `-- name: GetPartners :many
Select idpartner, nama, alamat, telp, mode, use_saldo, status from m_partnerid
`

func (q *Queries) GetPartners(ctx context.Context) ([]MPartnerid, error) {
	rows, err := q.db.QueryContext(ctx, getPartners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MPartnerid{}
	for rows.Next() {
		var i MPartnerid
		if err := rows.Scan(
			&i.Idpartner,
			&i.Nama,
			&i.Alamat,
			&i.Telp,
			&i.Mode,
			&i.UseSaldo,
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
