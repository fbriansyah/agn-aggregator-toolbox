// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package mariadb

import (
	"database/sql"
)

type MPartnerid struct {
	Idpartner int32          `json:"idpartner"`
	Nama      sql.NullString `json:"nama"`
	Alamat    sql.NullString `json:"alamat"`
	Telp      sql.NullString `json:"telp"`
	// 1 = SUBTRACT DEPOSIT OUTLET ONLY (JIKA DEPOSIT SETOR KE ADMIN), 2 = SUBTRACT DEPOSIT PARTNER AND OUTLET (JIKA DEPOSIT DISETOR KE CA), 3 = H2H DEPOSIT, 4 = H2H UNLIMITED
	Mode sql.NullInt32 `json:"mode"`
	// 1 = DEPOSIT, 2 = VIRTUAL
	UseSaldo sql.NullInt32 `json:"use_saldo"`
	Status   sql.NullInt32 `json:"status"`
}

type MPartneridProduk struct {
	IdpartnerProduk int32         `json:"idpartner_produk"`
	Idproduk        sql.NullInt32 `json:"idproduk"`
	Idpartner       sql.NullInt32 `json:"idpartner"`
	Idmerchant      sql.NullInt32 `json:"idmerchant"`
	Prioritas       sql.NullInt32 `json:"prioritas"`
	Status          sql.NullInt32 `json:"status"`
}

type MProduk struct {
	KodeProduk         string         `json:"kode_produk"`
	NamaProduk         sql.NullString `json:"nama_produk"`
	NamaProdukDisplay  sql.NullString `json:"nama_produk_display"`
	NamaStruk          sql.NullString `json:"nama_struk"`
	NamaStrukSingkatan sql.NullString `json:"nama_struk_singkatan"`
	ProdukAlias        sql.NullString `json:"produk_alias"`
	ProdukGroup        sql.NullString `json:"produk_group"`
	ProdukIndex        sql.NullInt32  `json:"produk_index"`
	ProdukDate         sql.NullTime   `json:"produk_date"`
	LabelIdpel         sql.NullString `json:"label_idpel"`
	LabelOption        sql.NullString `json:"label_option"`
	Status             sql.NullInt32  `json:"status"`
}

type MProdukProvider struct {
	Idproduk           int32          `json:"idproduk"`
	KodeProduk         sql.NullString `json:"kode_produk"`
	KodeProdukProvider sql.NullString `json:"kode_produk_provider"`
	NamaProduk         sql.NullString `json:"nama_produk"`
	Kategori           sql.NullString `json:"kategori"`
	Provider           sql.NullString `json:"provider"`
	Rctype             sql.NullString `json:"rctype"`
	ProviderBank       sql.NullString `json:"provider_bank"`
	FeeProduk          sql.NullInt32  `json:"fee_produk"`
	// N = NOMODE. R = REVERSAL, A = ADVICE
	ModeSto    sql.NullString `json:"mode_sto"`
	ModeRepeat int32          `json:"mode_repeat"`
	ClassRpc   sql.NullString `json:"class_rpc"`
	ClassName  sql.NullString `json:"class_name"`
	// tcp,http,https
	ClassTipe    sql.NullString `json:"class_tipe"`
	ClassIp      sql.NullString `json:"class_ip"`
	ClassPort    sql.NullInt32  `json:"class_port"`
	ClassPath    sql.NullString `json:"class_path"`
	ClassSetting sql.NullString `json:"class_setting"`
	ClassTimeout sql.NullInt32  `json:"class_timeout"`
	// 0 = POSTPAID, 1 = PREPAID WITH INQ SINGLE OPTION,   2 = PREPAID WITH INQ MULTIPLE OPTION, 3 = PREPAID NON INQ
	Tipemodel  sql.NullString `json:"tipemodel"`
	Screenname sql.NullString `json:"screenname"`
	Status     sql.NullInt32  `json:"status"`
	FeeApp1    sql.NullInt32  `json:"fee_app_1"`
	FeeCa1     sql.NullInt32  `json:"fee_ca_1"`
	FeeBiller1 sql.NullInt32  `json:"fee_biller_1"`
	FeeApp2    sql.NullInt32  `json:"fee_app_2"`
	FeeCa2     sql.NullInt32  `json:"fee_ca_2"`
	FeeBiller2 sql.NullInt32  `json:"fee_biller_2"`
	FeeApp3    sql.NullInt32  `json:"fee_app_3"`
	FeeCa3     sql.NullInt32  `json:"fee_ca_3"`
	FeeBiller3 sql.NullInt32  `json:"fee_biller_3"`
	FeeApp4    sql.NullInt32  `json:"fee_app_4"`
	FeeCa4     sql.NullInt32  `json:"fee_ca_4"`
	FeeBiller4 sql.NullInt32  `json:"fee_biller_4"`
}

type RekonIdproduk struct {
	Idproduk     int32          `json:"idproduk"`
	IdprodukPpob sql.NullString `json:"idproduk_ppob"`
	Nama         sql.NullString `json:"nama"`
}

type ResponseCode struct {
	Rcode       string         `json:"rcode"`
	Rctype      string         `json:"rctype"`
	RcodeClient sql.NullString `json:"rcode_client"`
	MessageID   sql.NullString `json:"message_id"`
	MessageEn   sql.NullString `json:"message_en"`
	Status      sql.NullString `json:"status"`
}

type Transaksi struct {
	Idtrx              int32          `json:"idtrx"`
	Blth               string         `json:"blth"`
	TglWaktu           sql.NullTime   `json:"tgl_waktu"`
	Idpartner          sql.NullInt32  `json:"idpartner"`
	Idoutlet           sql.NullInt32  `json:"idoutlet"`
	Iduser             sql.NullInt32  `json:"iduser"`
	IduserAdmin        sql.NullInt32  `json:"iduser_admin"`
	Idproduk           sql.NullInt32  `json:"idproduk"`
	Idmerchant         sql.NullInt32  `json:"idmerchant"`
	Kategori           sql.NullString `json:"kategori"`
	KodeBank           sql.NullString `json:"kode_bank"`
	Idpel              sql.NullString `json:"idpel"`
	SubIdpel           sql.NullString `json:"sub_idpel"`
	RpBillerPokok      sql.NullInt64  `json:"rp_biller_pokok"`
	RpBillerDenda      sql.NullInt32  `json:"rp_biller_denda"`
	RpBillerLain       sql.NullInt32  `json:"rp_biller_lain"`
	RpBillerPotongan   sql.NullInt32  `json:"rp_biller_potongan"`
	RpBillerTag        sql.NullInt64  `json:"rp_biller_tag"`
	RpFeeApp           sql.NullInt32  `json:"rp_fee_app"`
	RpFeePartner       sql.NullInt32  `json:"rp_fee_partner"`
	RpFeeBiller        sql.NullInt32  `json:"rp_fee_biller"`
	RpFeeAgregator     sql.NullInt32  `json:"rp_fee_agregator"`
	RpFeeUser          sql.NullInt32  `json:"rp_fee_user"`
	RpFeeStruk         sql.NullInt32  `json:"rp_fee_struk"`
	Provider           sql.NullString `json:"provider"`
	RpAmountStruk      sql.NullInt64  `json:"rp_amount_struk"`
	RpAmount           sql.NullInt64  `json:"rp_amount"`
	TipeSaldoPartner   sql.NullInt32  `json:"tipe_saldo_partner"`
	SaldoPartner       sql.NullInt64  `json:"saldo_partner"`
	TipeSaldoUser      sql.NullInt32  `json:"tipe_saldo_user"`
	SaldoUser          sql.NullInt64  `json:"saldo_user"`
	SubtractMode       sql.NullInt32  `json:"subtract_mode"`
	Lembar             sql.NullInt32  `json:"lembar"`
	PrivateKey         sql.NullString `json:"private_key"`
	AgnRef             sql.NullString `json:"agn_ref"`
	ClientRef          sql.NullString `json:"client_ref"`
	ClientStan         sql.NullInt64  `json:"client_stan"`
	ClientMessage      sql.NullString `json:"client_message"`
	ClientResponse     sql.NullString `json:"client_response"`
	ClientRcode        sql.NullString `json:"client_rcode"`
	ClientIdtrx        sql.NullInt32  `json:"client_idtrx"`
	MiddlewareResponse sql.NullString `json:"middleware_response"`
	ProviderRef        sql.NullString `json:"provider_ref"`
	ProviderRequest    sql.NullString `json:"provider_request"`
	ProviderResponse   sql.NullString `json:"provider_response"`
	ProviderRcode      sql.NullString `json:"provider_rcode"`
	BlthTagihan        sql.NullString `json:"blth_tagihan"`
	AccessIp           sql.NullString `json:"access_ip"`
	AccessBrowser      sql.NullString `json:"access_browser"`
	Ket                sql.NullString `json:"ket"`
	// -2 = HAPUS, -1 = GAGAL, 0 = PENDING, 1 = SUKSES, 2 = RESTITUSI
	Status         sql.NullInt32  `json:"status"`
	Idtrxdetail    sql.NullInt32  `json:"idtrxdetail"`
	Idtrxrestitusi sql.NullInt32  `json:"idtrxrestitusi"`
	Idtrxdeposit   sql.NullInt32  `json:"idtrxdeposit"`
	Idkonfirmasi   sql.NullInt32  `json:"idkonfirmasi"`
	Info           sql.NullString `json:"info"`
	Copy           sql.NullInt32  `json:"copy"`
	StatusBackup   sql.NullInt32  `json:"status_backup"`
}

type TransaksiLog struct {
	Idtrxlog         int32          `json:"idtrxlog"`
	Blth             string         `json:"blth"`
	TglWaktu         sql.NullTime   `json:"tgl_waktu"`
	Idpartner        sql.NullInt32  `json:"idpartner"`
	Idoutlet         sql.NullInt32  `json:"idoutlet"`
	Iduser           sql.NullInt32  `json:"iduser"`
	Idproduk         sql.NullInt32  `json:"idproduk"`
	Idmerchant       sql.NullInt32  `json:"idmerchant"`
	RequestTipe      sql.NullString `json:"request_tipe"`
	Kategori         sql.NullString `json:"kategori"`
	KodeBank         sql.NullString `json:"kode_bank"`
	Idpel            sql.NullString `json:"idpel"`
	SubIdpel         sql.NullString `json:"sub_idpel"`
	RpBillerPokok    sql.NullInt64  `json:"rp_biller_pokok"`
	RpBillerDenda    sql.NullInt32  `json:"rp_biller_denda"`
	RpBillerLain     sql.NullInt32  `json:"rp_biller_lain"`
	RpBillerPotongan sql.NullInt32  `json:"rp_biller_potongan"`
	RpBillerTag      sql.NullInt64  `json:"rp_biller_tag"`
	RpFeeApp         sql.NullInt32  `json:"rp_fee_app"`
	RpFeePartner     sql.NullInt32  `json:"rp_fee_partner"`
	RpFeeBiller      sql.NullInt32  `json:"rp_fee_biller"`
	RpFeeAgregator   sql.NullInt32  `json:"rp_fee_agregator"`
	RpFeeUser        sql.NullInt32  `json:"rp_fee_user"`
	RpFeeStruk       sql.NullInt32  `json:"rp_fee_struk"`
	// RP_BILLER_AMOUNT + RP_FEE_APP + RP_FEE_PARTNER
	RpAmountPartner sql.NullInt64 `json:"rp_amount_partner"`
	// RP_BILLER_AMOUNT + RP_FEE_STRUK
	RpAmountStruk      sql.NullInt64  `json:"rp_amount_struk"`
	Lembar             sql.NullInt32  `json:"lembar"`
	PrivateKey         sql.NullString `json:"private_key"`
	AgnRef             sql.NullString `json:"agn_ref"`
	ClientRef          sql.NullString `json:"client_ref"`
	ClientStan         sql.NullInt64  `json:"client_stan"`
	ClientMessage      sql.NullString `json:"client_message"`
	ClientResponse     sql.NullString `json:"client_response"`
	ClientRcode        sql.NullString `json:"client_rcode"`
	MiddlewareResponse sql.NullString `json:"middleware_response"`
	ProviderRef        sql.NullString `json:"provider_ref"`
	ProviderRequest    sql.NullString `json:"provider_request"`
	ProviderResponse   sql.NullString `json:"provider_response"`
	ProviderRcode      sql.NullString `json:"provider_rcode"`
	BlthTagihan        sql.NullString `json:"blth_tagihan"`
	AccessIp           sql.NullString `json:"access_ip"`
	AccessBrowser      sql.NullString `json:"access_browser"`
	// -2 = HAPUS, -1 = GAGAL, 0 = PENDING, 1 = SUKSES
	Status sql.NullInt32 `json:"status"`
	Idtrx  sql.NullInt32 `json:"idtrx"`
}
