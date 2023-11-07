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
	// GetProviderDetail get detail from provider by its id
	GetProviderDetail(ctx context.Context, idproduk int32) (domain.ProdukProviderDomain, error)
	// GetListPartnerProduk get partner and produk data based on idproduk
	GetListPartnerProduk(ctx context.Context, idproduk int32) ([]domain.PartnerProdukDomain, error)
	// CreatePartnerProduk insert partner produk to database
	CreatePartnerProduk(ctx context.Context, data domain.PartnerProdukDomain) (domain.PartnerProdukDomain, error)
	// UpdatePartnerProduk update data partner
	UpdatePartnerProduk(ctx context.Context, data domain.PartnerProdukDomain) (domain.PartnerProdukDomain, error)
	// GetPartnerProduk get one partner produk base on idpartnerproduk
	GetPartnerProduk(ctx context.Context, idpartnerproduk int32) (domain.PartnerProdukDomain, error)
	// UpdateProduct save updated produk to database
	UpdateProduct(ctx context.Context, produk domain.ProductDomain) (domain.ProductDomain, error)
	// CreateProvider insert provider data to database
	CreateProvider(ctx context.Context, provider domain.ProdukProviderDomain) (domain.ProdukProviderDomain, error)
	// UpdateProvider update provider data
	UpdateProvider(ctx context.Context, provider domain.ProdukProviderDomain) (domain.ProdukProviderDomain, error)
	// GetPartners get list all partners
	GetPartners(ctx context.Context) ([]domain.PartnerDomain, error)
}
