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
		"PageTitle":    "Product",
		"AddButtonUrl": "/add-product",
		"Products":     products,
	})

	if err != nil {
		log.Fatalf("error execute template: %v", err)
	}
}

func (a *ChiAdapter) DetailProductIndex(w http.ResponseWriter, req *http.Request) {

	tmplFile, err := template.New("").ParseFiles(
		"templates/pages/list-product/editor.html",
		"templates/shared/nav.html",
		"templates/shared/base.html",
	)

	if err != nil {
		log.Fatalf("error parse template: %v", err)
	}

	kodeProduk := req.URL.Query().Get("kode_product")

	if kodeProduk == "" {
		http.Redirect(w, req, "/", http.StatusPermanentRedirect)
		return
	}

	product, err := a.service.GetProdukByCode(context.Background(), kodeProduk)
	if err != nil {
		fmt.Println(err)
	}

	providers, err := a.service.GetProductProvider(context.Background(), kodeProduk)
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "base", M{
		"PageTitle":    "Detail Product",
		"AddButtonUrl": "",
		"Product":      product,
		"Providers":    providers,
	})

	if err != nil {
		log.Fatalf("error execute template: %v", err)
	}
}

func (a *ChiAdapter) AddProdukPage(w http.ResponseWriter, req *http.Request) {

	tmplFile, err := template.New("").ParseFiles(
		"templates/pages/list-product/create-form.html",
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
		"PageTitle":    "Product",
		"AddButtonUrl": "",
		"Products":     products,
	})

	if err != nil {
		log.Fatalf("error execute template: %v", err)
	}
}
