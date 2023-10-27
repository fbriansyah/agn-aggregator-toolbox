package port

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

type ServicePort interface {
	// ListProduct show all product without filter and limit
	ListProduct(ctx context.Context) ([]domain.ProductDomain, error)
	// GetProdukByCode get detail product based on kode produk
	GetProdukByCode(ctx context.Context, kodeProduk string) (domain.ProductDomain, error)
	// CreateProduct insert new produk to database
	CreateProduct(ctx context.Context, produk domain.ProductDomain) (domain.ProductDomain, error)
	// GetProductProvider get produk providers base on kode produk
	GetProductProvider(ctx context.Context, kodeProduk string) ([]domain.ProdukProviderDomain, error)
	// GetProducts get list product based on kode produk and or nama produk
	GetProducts(ctx context.Context, produk domain.ProductDomain) ([]domain.ProductDomain, error)
}
