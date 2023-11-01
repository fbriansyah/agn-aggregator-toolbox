package template

import "html/template"

// ProductProviderIndex get html template for provider index page
func (t *TemplateAdapter) ProductProviderIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/chi/pages/provider/index.html",
		"templates/chi/pages/provider/detail.html",
		"templates/chi/pages/partner-produk/list-partner.html",
		"templates/chi/shared/nav.html",
		"templates/chi/shared/base.html",
	)
}

func (t *TemplateAdapter) GetProviderEditForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/provider/edit-form.html")
}

func (t *TemplateAdapter) GetProviderCreateForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/provider/create-form.html")
}

func (t *TemplateAdapter) GetListProvider() (*template.Template, error) {
	return template.New("").ParseFiles("templates/chi/pages/provider/list-provider.html")
}
