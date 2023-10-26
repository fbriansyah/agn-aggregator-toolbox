package template

import "text/template"

func (t *TemplateAdapter) ProductIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/index.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

func (t *TemplateAdapter) DetailProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/editor.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

func (t *TemplateAdapter) CreateProduct() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/list-product/create-form.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}
