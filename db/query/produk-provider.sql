-- name: ListProviderProduk :many
SELECT * FROM m_produk_provider where kode_produk=?;