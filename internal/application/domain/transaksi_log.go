package domain

import (
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain/model"
)

type TransaksiLog struct {
	Idtrxlog           int32     `json:"idtrxlog"`
	Blth               string    `json:"blth"`
	TglWaktu           time.Time `json:"tgl_waktu"`
	Idpartner          int32     `json:"idpartner"`
	Idproduk           int32     `json:"idproduk"`
	NamaProduk         string    `json:"nama_produk"`
	Idmerchant         int32     `json:"idmerchant"`
	RequestTipe        string    `json:"request_tipe"`
	Idpel              string    `json:"idpel"`
	MiddlewareResponse string    `json:"middleware_response"`
	ProviderRequest    string    `json:"provider_request"`
	Status             int32     `json:"status"`
}

func (t *TransaksiLog) FromModel(m model.TransaksiLog) {
	t.Idtrxlog = m.Idtrxlog
	t.Blth = m.Blth
	t.TglWaktu = m.TglWaktu.Time
	t.Idpartner = m.Idpartner.Int32
	t.Idproduk = m.Idproduk.Int32
	t.NamaProduk = m.NamaProduk.String
	t.Idmerchant = m.Idmerchant.Int32
	t.RequestTipe = m.RequestTipe.String
	t.Idpel = m.Idpel.String
	t.MiddlewareResponse = m.MiddlewareResponse.String
	t.ProviderRequest = m.ProviderRequest.String
	t.Status = m.Status.Int32
}
