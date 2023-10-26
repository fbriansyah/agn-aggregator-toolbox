package chi

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

func (a *ChiAdapter) ListProductPage(w http.ResponseWriter, req *http.Request) {
	tmplFile, err := a.template.ProductIndex()
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

func (a *ChiAdapter) DetailProductPage(w http.ResponseWriter, req *http.Request) {
	tmplFile, err := a.template.DetailProduct()
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
	tmplFile, err := a.template.CreateProduct()
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

func (a *ChiAdapter) GetProducts(w http.ResponseWriter, req *http.Request) {
	kodeProduk := req.URL.Query().Get("kode-produk")
	namaProduk := req.URL.Query().Get("nama-produk")

	products, err := a.service.GetProducts(context.Background(), domain.ProductDomain{
		KodeProduk: kodeProduk,
		NamaProduk: namaProduk,
	})
	if err != nil {
		fmt.Println(err)
	}
	tmplFile, err := a.template.ProductIndex()
	if err != nil {
		fmt.Println(err)
	}
	err = tmplFile.ExecuteTemplate(w, "list-product", M{
		"Products": products,
	})
	if err != nil {
		log.Fatalf("error execute template: %v", err)
	}
}
