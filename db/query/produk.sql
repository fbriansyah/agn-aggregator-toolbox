-- name: ListProduk :many
SELECT * FROM m_produk ORDER BY NAMA_PRODUK;

-- name: GetProductByKode :one
SELECT * FROM m_produk where KODE_PRODUK = ? LIMIT 1;

-- name: CreateProduct :execresult
INSERT INTO m_produk 
(KODE_PRODUK, NAMA_PRODUK, NAMA_PRODUK_DISPLAY, nama_struk, nama_struk_singkatan, produk_alias, produk_group, label_idpel, STATUS)
vALUES
(?,?,?,?,?,?,?,?,1);