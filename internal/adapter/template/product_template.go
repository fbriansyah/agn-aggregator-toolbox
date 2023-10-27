package template

import "text/template"

// ProductIndex return template from list-product/index.html
func (t *TemplateAdapter) ProductIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/index.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

// DetailProduct return template from list-product/editor.html
func (t *TemplateAdapter) DetailProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/editor.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

// CreateProduct return template from list-product/create-form.html
func (t *TemplateAdapter) CreateProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/create-form.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}
