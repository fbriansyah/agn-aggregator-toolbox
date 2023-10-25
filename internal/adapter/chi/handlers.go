package chi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (a *ChiAdapter) ListProductIndex(w http.ResponseWriter, req *http.Request) {

	tmplFile, err := template.New("").ParseFiles(
		"templates/pages/list-product/index.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)
	if err != nil {
		log.Fatalf("error parse template: %v", err)
	}

	products, err := a.service.ListProduct(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "base", M{
		"PageTitle": "Product",
		"Products":  products,
	})

	if err != nil {
		log.Fatalf("error execute template: %v", err)
	}
}
