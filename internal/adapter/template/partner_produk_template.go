package template

import "html/template"

func (t *TemplateAdapter) GetPartnerProdukCreateForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/partner-produk/create-form.html")
}

func (t *TemplateAdapter) GetListPartnerProduk() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/partner-produk/list-partner.html")
}

func (t *TemplateAdapter) GetPartnerProdukEditForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/partner-produk/edit-form.html")
}
