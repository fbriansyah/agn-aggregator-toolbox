-- name: ListProviderProduk :many
SELECT * FROM m_produk_provider where kode_produk=?;

-- name: GetProviderByID :one
SELECT * FROM m_produk_provider WHERE idproduk = ? LIMIT 1;

-- name: CreateProvider :execresult
INSERT INTO m_produk_provider
(KODE_PRODUK, KODE_PRODUK_PROVIDER, NAMA_PRODUK, KATEGORI, PROVIDER, RCTYPE, PROVIDER_BANK, FEE_PRODUK, CLASS_RPC, CLASS_NAME, CLASS_TIPE, CLASS_IP, CLASS_PORT, CLASS_PATH, CLASS_SETTING, CLASS_TIMEOUT, TIPEMODEL, SCREENNAME, STATUS, FEE_APP_1, FEE_CA_1, FEE_BILLER_1, FEE_APP_2, FEE_CA_2, FEE_BILLER_2, FEE_APP_3, FEE_CA_3, FEE_BILLER_3, FEE_APP_4, FEE_CA_4, FEE_BILLER_4)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);

-- name: UpdateProvider :execresult
UPDATE m_produk_provider
SET 
    KODE_PRODUK=?,
    KODE_PRODUK_PROVIDER=?,
    NAMA_PRODUK=?,
    KATEGORI=?,
    PROVIDER=?,
    RCTYPE=?,
    PROVIDER_BANK=?,
    FEE_PRODUK=?,
    CLASS_RPC=?,
    CLASS_NAME=?,
    CLASS_TIPE=?,
    CLASS_IP=?,
    CLASS_PORT=?,
    CLASS_PATH=?,
    CLASS_SETTING=?,
    CLASS_TIMEOUT=?,
    TIPEMODEL=?,
    SCREENNAME=?,
    STATUS=?
WHERE IDPRODUK=?;
