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

	mux.Get("/", a.ListProductIndex)
	mux.Get("/detail-product", a.DetailProductIndex)
	mux.Get("/add-product", a.AddProdukPage)

	return mux
}
