package template

import "text/template"

func (t *TemplateAdapter) GetPartnerProdukCreateForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/partner-produk/create-form.html")
}

func (t *TemplateAdapter) GetListPartnerProduk() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/partner-produk/list-partner.html")
}
