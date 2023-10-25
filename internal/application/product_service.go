package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
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
