package port

import "html/template"

type TemplatePort interface {
	// ProductIndex return template from list-product/index.html
	ProductIndex() (*template.Template, error)
	// DetailProduct return template from list-product/editor.html
	DetailProduct() (*template.Template, error)
	// CreateProduct return template from list-product/create-form.html
	CreateProduct() (*template.Template, error)
	// ProductProviderIndex get html template from provider provider/index.html
	ProductProviderIndex() (*template.Template, error)
	// GetProviderEditForm get html template from provider/edit-form.html
	GetProviderEditForm() (*template.Template, error)
	// GetProviderCreateForm get html template from provider/create-form.html
	GetProviderCreateForm() (*template.Template, error)
	// GetListProvider get html template from provider/list-provider.html
	GetListProvider() (*template.Template, error)
	// GetPartnerProdukCreateForm get html template from partner-produk/create-form.html
	GetPartnerProdukCreateForm() (*template.Template, error)
	// GetListPartnerProduk get html template from partner-produk/list-product.html
	GetListPartnerProduk() (*template.Template, error)
	// GetPartnerProdukEditForm get html template from partner-produk/edit-form.html
	GetPartnerProdukEditForm() (*template.Template, error)
}
