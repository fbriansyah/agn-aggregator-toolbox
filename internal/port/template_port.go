package port

import "text/template"

type TemplatePort interface {
	// ProductIndex return template from list-product/index.html
	ProductIndex() (*template.Template, error)
	// DetailProduct return template from list-product/editor.html
	DetailProduct() (*template.Template, error)
	// CreateProduct return template from list-product/create-form.html
	CreateProduct() (*template.Template, error)
	// ProductProviderIndex get html template for provider index page
	ProductProviderIndex() (*template.Template, error)
	GetProviderForm() (*template.Template, error)
}
