-- name: ListProduk :many
SELECT * FROM m_produk;

-- name: GetProductByKode :one
SELECT * FROM m_produk where KODE_PRODUK = ? LIMIT 1;