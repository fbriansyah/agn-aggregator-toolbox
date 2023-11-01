package template

import "html/template"

// ProductIndex return template from list-product/index.html
func (t *TemplateAdapter) ProductIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/chi/pages/list-product/index.html",
		"templates/chi/shared/nav.html",
		"templates/chi/shared/base.html",
	)
}

// DetailProduct return template from list-product/editor.html
func (t *TemplateAdapter) DetailProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/chi/pages/list-product/editor.html",
		"templates/chi/pages/provider/list-provider.html",
		"templates/chi/shared/nav.html",
		"templates/chi/shared/base.html",
	)
}

// CreateProduct return template from list-product/create-form.html
func (t *TemplateAdapter) CreateProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/chi/pages/list-product/create-form.html",
		"templates/chi/shared/nav.html",
		"templates/chi/shared/base.html",
	)
}
