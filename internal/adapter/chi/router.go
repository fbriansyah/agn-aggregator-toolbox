package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes setting route
func (a *ChiAdapter) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", a.ListProductPage)
	mux.Get("/detail-product", a.DetailProductPage)
	mux.Get("/add-product", a.AddProdukPage)
	mux.Get("/get-products", a.GetProducts)
	mux.Get("/provider", a.ProdukProviderIndex)
	mux.Get("/provider/editor", a.GetProviderEditor)
	mux.Get("/provider/create-form", a.GetProviderCreateForm)
	mux.Get("/provider/list-provider", a.GetListProvider)
	mux.Get("/partner-produk/create-form", a.GetPartnerProdukForm)
	mux.Get("/partner-produk", a.GetListPartnerProduk)

	return mux
}
