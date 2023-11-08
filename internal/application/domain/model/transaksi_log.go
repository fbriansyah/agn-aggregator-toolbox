package model

import "database/sql"

type TransaksiLog struct {
	Idtrxlog           int32          `json:"idtrxlog"`
	Blth               string         `json:"blth"`
	TglWaktu           sql.NullTime   `json:"tgl_waktu"`
	Idpartner          sql.NullInt32  `json:"idpartner"`
	Idproduk           sql.NullInt32  `json:"idproduk"`
	NamaProduk         sql.NullString `json:"nama_produk"`
	Idmerchant         sql.NullInt32  `json:"idmerchant"`
	RequestTipe        sql.NullString `json:"request_tipe"`
	Idpel              sql.NullString `json:"idpel"`
	MiddlewareResponse sql.NullString `json:"middleware_response"`
	ProviderRequest    sql.NullString `json:"provider_request"`
	Status             sql.NullInt32  `json:"status"`
}
