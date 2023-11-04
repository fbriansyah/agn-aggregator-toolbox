package fiber

func (a *FiberAdapter) router() {
	a.app.Get("/", a.ListProduk)
	a.app.Get("/get-products", a.GetProducts)
	a.app.Get("/add-product", a.CreateProdukFormPage)
	a.app.Get("/detail-product", a.DetailProductPage)
	a.app.Post("/produk/save", a.SaveProduk)
	a.app.Post("/produk/update", a.UpdateProduk)

	a.app.Get("/provider/create-form", a.GetProviderCreateForm)
	a.app.Get("/provider/list-provider", a.GetListProvider)
	a.app.Get("/provider", a.ProdukProviderIndex)
	a.app.Get("/provider/editor", a.GetProviderEditor)
	a.app.Post("/provider/save", a.SaveProvider)
	a.app.Post("/provider/update", a.UpdateProvider)

	a.app.Get("/partner-produk/create-form", a.GetPartnerProdukForm)
	a.app.Get("/partner-produk", a.GetListPartnerProduk)
	a.app.Get("/partner-produk/edit-form", a.GetPartnerProdukEditForm)
}
