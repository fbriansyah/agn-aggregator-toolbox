package chi

import (
	"fmt"
	"net/http"
)

func (a *ChiAdapter) ProdukProviderIndex(w http.ResponseWriter, req *http.Request) {
	idProvider := req.URL.Query().Get("id")

	tmplFile, err := a.template.ProductProviderIndex()
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "base", M{
		"IDProvider": idProvider,
		"PageTitle":  "Provider",
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (a *ChiAdapter) GetProviderEditor(w http.ResponseWriter, req *http.Request) {
	tmplFile, err := a.template.GetProviderForm()
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "editor", nil)
	if err != nil {
		fmt.Println(err)
	}
}
