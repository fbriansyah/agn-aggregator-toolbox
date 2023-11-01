package fiber

func (a *FiberAdapter) router() {
	a.app.Get("/", a.ListProduk)
	a.app.Get("/get-products", a.GetProducts)
}
