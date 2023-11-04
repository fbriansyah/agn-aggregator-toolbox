package fiber

func (a *FiberAdapter) router() {
	a.app.Get("/", a.ListProduk)
	a.app.Get("/get-products", a.GetProducts)
	a.app.Get("/add-product", a.CreateProdukFormPage)
	a.app.Get("/detail-product", a.DetailProductPage)
	a.app.Get("/provider/create-form", a.GetProviderCreateForm)
	a.app.Get("/provider/list-provider", a.GetListProvider)
}
