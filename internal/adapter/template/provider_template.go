package template

import "text/template"

// ProductProviderIndex get html template for provider index page
func (t *TemplateAdapter) ProductProviderIndex() (*template.Template, error) {
	return template.New("").ParseFiles(
		"templates/pages/provider/index.html",
		"templates/pages/provider/detail.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
}

func (t *TemplateAdapter) GetProviderForm() (*template.Template, error) {
	return template.New("").ParseFiles("templates/pages/provider/edit-form.html")
}
