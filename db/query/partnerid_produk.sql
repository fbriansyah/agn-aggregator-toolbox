-- name: ListPartnerIDProduk :many
SELECT * FROM m_partnerid_produk WHERE idproduk = ?;

-- name: CreatePartnerIDProduk :execresult
INSERT INTO m_partnerid_produk (IDPRODUK, IDPARTNER, IDMERCHANT, PRIORITAS, STATUS)
VALUES (?,?,?,?,1);

-- name: UpdatePartnerIDProduk :execresult
UPDATE m_partnerid_produk
SET 
  IDPRODUK = ?, 
  IDPARTNER = ?, 
  IDMERCHANT = ?, 
  PRIORITAS = ?, 
  STATUS = ?
WHERE IDPARTNER_PRODUK = ?;
