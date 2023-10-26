package domain

import "github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"

type ProdukProviderDomain struct {
	Idproduk           int32  `json:"idproduk"`
	KodeProduk         string `json:"kode_produk"`
	KodeProdukProvider string `json:"kode_produk_provider"`
	NamaProduk         string `json:"nama_produk"`
	Kategori           string `json:"kategori"`
	Provider           string `json:"provider"`
	Rctype             string `json:"rctype"`
	ProviderBank       string `json:"provider_bank"`
	FeeProduk          int32  `json:"fee_produk"`
	ClassRpc           string `json:"class_rpc"`
	ClassName          string `json:"class_name"`
	ClassTipe          string `json:"class_tipe"`
	ClassIp            string `json:"class_ip"`
	ClassPort          int32  `json:"class_port"`
	ClassPath          string `json:"class_path"`
	ClassSetting       string `json:"class_setting"`
	ClassTimeout       int32  `json:"class_timeout"`
	Screenname         string `json:"screenname"`
	Status             int32  `json:"status"`
}

func (d *ProdukProviderDomain) FromMProviderProduk(pp mariadb.MProdukProvider) {
	d.Idproduk = pp.Idproduk
	d.KodeProduk = pp.KodeProduk.String
	d.KodeProdukProvider = pp.KodeProdukProvider.String
	d.NamaProduk = pp.NamaProduk.String
	d.Kategori = pp.Kategori.String
	d.Provider = pp.Provider.String
	d.Rctype = pp.Rctype.String
	d.ProviderBank = pp.ProviderBank.String
	d.FeeProduk = pp.FeeProduk.Int32
	d.ClassRpc = pp.ClassRpc.String
	d.ClassName = pp.ClassName.String
	d.ClassTipe = pp.ClassTipe.String
	d.ClassIp = pp.ClassIp.String
	d.ClassPort = pp.ClassPort.Int32
	d.ClassPath = pp.ClassPath.String
	d.ClassSetting = pp.ClassSetting.String
	d.ClassTimeout = pp.ClassTimeout.Int32
	d.Screenname = pp.Screenname.String
	d.Status = pp.Status.Int32
}
