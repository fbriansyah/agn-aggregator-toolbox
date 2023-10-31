package template

import "html/template"

// ProductProviderIndex get html template for provider index page
func (t *TemplateAdapter) ProductProviderIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/provider/index.html",
		"templates/pages/provider/detail.html",
		"templates/pages/partner-produk/list-partner.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

func (t *TemplateAdapter) GetProviderEditForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/provider/edit-form.html")
}

func (t *TemplateAdapter) GetProviderCreateForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/provider/create-form.html")
}

func (t *TemplateAdapter) GetListProvider() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/provider/list-provider.html")
}
