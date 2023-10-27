package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

func (s *Service) GetProviderDetail(ctx context.Context, idproduk int32) (domain.ProdukProviderDomain, error) {
	provider, err := s.db.GetProviderByID(ctx, idproduk)
	if err != nil {
		return domain.ProdukProviderDomain{}, err
	}
	var productProvider domain.ProdukProviderDomain

	productProvider.FromMProviderProduk(provider)

	return productProvider, nil
}
