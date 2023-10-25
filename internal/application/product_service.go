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
