package port

import "text/template"

type TemplatePort interface {
	ProductIndex() (*template.Template, error)
	DetailProduct() (*template.Template, error)
	CreateProduct() (*template.Template, error)
}
