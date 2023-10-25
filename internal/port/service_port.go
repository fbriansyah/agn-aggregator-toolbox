package port

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

type ServicePort interface {
	// ListProduct show all product without filter and limit
	ListProduct(ctx context.Context) ([]domain.ProductDomain, error)
}
