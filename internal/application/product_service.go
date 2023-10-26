package application

import (
	"context"
	"errors"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
)

// ListProduct show all product without filter and limit
func (s *Service) ListProduct(ctx context.Context) ([]domain.ProductDomain, error) {
	products, err := s.db.ListProduk(ctx)
	if err != nil {
		return []domain.ProductDomain{}, err
	}

	listProduct := []domain.ProductDomain{}

	for _, product := range products {
		var p domain.ProductDomain
		p.FromMProduct(product)
		listProduct = append(listProduct, p)
	}

	return listProduct, nil
}

func (s *Service) GetProdukByCode(ctx context.Context, kodeProduk string) (domain.ProductDomain, error) {
	mProduct, err := s.db.GetProductByKode(ctx, kodeProduk)
	if err != nil {
		return domain.ProductDomain{}, err
	}

	produk := domain.ProductDomain{}
	produk.FromMProduct(mProduct)

	return produk, nil
}

func (s *Service) CreatProduct(ctx context.Context, produk domain.ProductDomain) (domain.ProductDomain, error) {
	arg := mariadb.CreateProductParams{
		KodeProduk:         produk.KodeProduk,
		NamaProduk:         util.ValidNullString(produk.NamaProduk),
		NamaProdukDisplay:  util.ValidNullString(produk.NamaProdukDisplay),
		NamaStruk:          util.ValidNullString(produk.NamaStruk),
		NamaStrukSingkatan: util.ValidNullString(produk.NamaStrukSingkatan),
		ProdukAlias:        util.ValidNullString(produk.ProdukAlias),
		ProdukGroup:        util.ValidNullString(produk.ProdukGroup),
		LabelIdpel:         util.ValidNullString(produk.LabelIdpel),
	}

	result, err := s.db.CreateProduct(ctx, arg)
	if err != nil {
		return domain.ProductDomain{}, nil
	}

	lastId, err := result.RowsAffected()
	if err != nil {
		return domain.ProductDomain{}, err
	}

	if lastId < int64(0) {
		return domain.ProductDomain{}, errors.New("error creating product")
	}

	return produk, nil
}
