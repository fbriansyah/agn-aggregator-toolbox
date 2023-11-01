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
	// /detail-product?kode_product=83016
	mux.Get("/detail-product", a.DetailProductPage)
	mux.Get("/add-product", a.AddProdukPage)
	mux.Get("/get-products", a.GetProducts)
	// /provider?id={{.Provider.Idproduk}}
	mux.Get("/provider", a.ProdukProviderIndex)
	// /provider/editor?id={{.IDProvider}}
	mux.Get("/provider/editor", a.GetProviderEditor)
	mux.Get("/provider/create-form", a.GetProviderCreateForm)
	// /provider/list-provider?kode_product={{.KodeProduk}}
	mux.Get("/provider/list-provider", a.GetListProvider)
	mux.Get("/partner-produk/create-form", a.GetPartnerProdukForm)
	// /partner-produk?idproduk={{.Idproduk}}
	mux.Get("/partner-produk", a.GetListPartnerProduk)
	// /partner-produk/edit-form?id={{.IdpartnerProduk}}
	mux.Get("/partner-produk/edit-form", a.GetPartnerProdukEditForm)

	return mux
}
