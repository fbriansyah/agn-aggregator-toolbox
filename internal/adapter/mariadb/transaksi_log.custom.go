package mariadb

import (
	"context"
	"fmt"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain/model"
)

const listTransaksiLog = `-- name: ListTransaksiLog :many
select 
	tl.IDTRXLOG, 
	tl.BLTH, 
	tl.TGL_WAKTU, 
	tl.IDPARTNER, 
	tl.IDMERCHANT,
	tl.IDPRODUK,
	mpp.NAMA_PRODUK,
	tl.REQUEST_TIPE,
	tl.IDPEL,
	tl.PROVIDER_REQUEST,
	tl.MIDDLEWARE_RESPONSE,
	tl.STATUS
from transaksi_log tl
left join m_produk_provider mpp on tl.IDPRODUK = mpp.IDPRODUK 
where tl.KATEGORI = 'P' and %s
ORDER BY tl.TGL_WAKTU
`

func (q *Queries) GetTransaksiLog(ctx context.Context, whereQuery string) ([]model.TransaksiLog, error) {
	query := fmt.Sprintf(listTransaksiLog, whereQuery)

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []model.TransaksiLog{}
	for rows.Next() {
		var i model.TransaksiLog
		if err := rows.Scan(
			&i.Idtrxlog,
			&i.Blth,
			&i.TglWaktu,
			&i.Idpartner,
			&i.Idmerchant,
			&i.Idproduk,
			&i.NamaProduk,
			&i.RequestTipe,
			&i.Idpel,
			&i.ProviderRequest,
			&i.MiddlewareResponse,
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

func (q *Queries) GetTransaksiLogByID(ctx context.Context, idTrxLog int32) (model.TransaksiLog, error) {
	query := fmt.Sprintf(listTransaksiLog, fmt.Sprintf("tl.IDTRXLOG = %d", idTrxLog))
	query += " Limit 1"

	fmt.Println(query)

	row := q.db.QueryRowContext(ctx, query)
	var i model.TransaksiLog

	err := row.Scan(
		&i.Idtrxlog,
		&i.Blth,
		&i.TglWaktu,
		&i.Idpartner,
		&i.Idmerchant,
		&i.Idproduk,
		&i.NamaProduk,
		&i.RequestTipe,
		&i.Idpel,
		&i.ProviderRequest,
		&i.MiddlewareResponse,
		&i.Status,
	)

	return i, err
}
