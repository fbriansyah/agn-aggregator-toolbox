package fiber

func (a *FiberAdapter) router() {
	a.app.Get("/", a.ListProduk)
}
