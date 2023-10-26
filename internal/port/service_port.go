package port

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

type ServicePort interface {
	// ListProduct show all product without filter and limit
	ListProduct(ctx context.Context) ([]domain.ProductDomain, error)
	GetProdukByCode(ctx context.Context, kodeProduk string) (domain.ProductDomain, error)
	CreateProduct(ctx context.Context, produk domain.ProductDomain) (domain.ProductDomain, error)
	GetProductProvider(ctx context.Context, kodeProduk string) ([]domain.ProdukProviderDomain, error)
	GetProducts(ctx context.Context, produk domain.ProductDomain) ([]domain.ProductDomain, error)
}
