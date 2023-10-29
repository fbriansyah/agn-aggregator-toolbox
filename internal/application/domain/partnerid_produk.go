package domain

import "github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"

type PartnerProdukDomain struct {
	IdpartnerProduk int32 `json:"idpartner_produk"`
	Idproduk        int32 `json:"idproduk"`
	Idpartner       int32 `json:"idpartner"`
	Idmerchant      int32 `json:"idmerchant"`
	Prioritas       int32 `json:"prioritas"`
	Status          int32 `json:"status"`
}

// FromMPartneridProduk fill domain with data from mariadb.MPartneridProduk
func (d *PartnerProdukDomain) FromMPartneridProduk(p mariadb.MPartneridProduk) {
	d.IdpartnerProduk = p.IdpartnerProduk
	d.Idproduk = p.Idproduk.Int32
	d.Idpartner = p.Idpartner.Int32
	d.Idmerchant = p.Idmerchant.Int32
	d.Prioritas = p.Prioritas.Int32
	d.Status = p.Status.Int32
}
