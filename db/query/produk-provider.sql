-- name: ListProviderProduk :many
SELECT * FROM m_produk_provider where kode_produk=?;

-- name: GetProviderByID :one
SELECT * FROM m_produk_provider WHERE idproduk = ? LIMIT 1;
