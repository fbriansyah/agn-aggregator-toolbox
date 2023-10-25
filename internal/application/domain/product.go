package domain

import (
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
)

type ProductDomain struct {
	KodeProduk         string    `json:"kode_produk"`
	NamaProduk         string    `json:"nama_produk"`
	NamaProdukDisplay  string    `json:"nama_produk_display"`
	NamaStruk          string    `json:"nama_struk"`
	NamaStrukSingkatan string    `json:"nama_struk_singkatan"`
	ProdukAlias        string    `json:"produk_alias"`
	ProdukGroup        string    `json:"produk_group"`
	ProdukIndex        int32     `json:"produk_index"`
	ProdukDate         time.Time `json:"produk_date"`
	LabelIdpel         string    `json:"label_idpel"`
	LabelOption        string    `json:"label_option"`
	Status             int32     `json:"status"`
}

// FromMProduct fill product domain from MProduct data
func (p *ProductDomain) FromMProduct(product mariadb.MProduk) {
	p.KodeProduk = product.KodeProduk
	p.NamaProduk = product.NamaProduk.String
	p.NamaProdukDisplay = product.NamaProdukDisplay.String
	p.NamaStruk = product.NamaStruk.String
	p.NamaStrukSingkatan = product.NamaStrukSingkatan.String
	p.ProdukAlias = product.ProdukAlias.String
	p.ProdukGroup = product.ProdukGroup.String
	p.ProdukIndex = product.ProdukIndex.Int32
	p.ProdukDate = product.ProdukDate.Time
	p.LabelIdpel = product.LabelIdpel.String
	p.LabelOption = product.LabelOption.String
	p.Status = product.Status.Int32
}
