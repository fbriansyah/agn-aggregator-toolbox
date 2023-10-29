package chi

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (a *ChiAdapter) ProdukProviderIndex(w http.ResponseWriter, req *http.Request) {
	idProvider := req.URL.Query().Get("id")

	tmplFile, err := a.template.ProductProviderIndex()
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.ParseInt(idProvider, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	provider, err := a.service.GetProviderDetail(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
	}

	partners, err := a.service.GetListPartnerProduk(ctx, provider.Idproduk)
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "base", M{
		"IDProvider": idProvider,
		"PageTitle":  "Provider",
		"Provider":   provider,
		"Idproduk":   provider.Idproduk,
		"Partners":   partners,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (a *ChiAdapter) GetProviderEditor(w http.ResponseWriter, req *http.Request) {
	idProvider := req.URL.Query().Get("id")
	tmplFile, err := a.template.GetProviderEditForm()
	if err != nil {
		fmt.Println(err)
	}

	id, err := strconv.ParseInt(idProvider, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	provider, err := a.service.GetProviderDetail(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
	}

	err = tmplFile.ExecuteTemplate(w, "editor", M{
		"Provider": provider,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (a *ChiAdapter) GetProviderCreateForm(w http.ResponseWriter, req *http.Request) {
	kodeProduk := req.URL.Query().Get("kode_product")
	tmpl, err := a.template.GetProviderCreateForm()
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	product, err := a.service.GetProdukByCode(ctx, kodeProduk)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "content", M{
		"KodeProduk": kodeProduk,
		"Product":    product,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (a *ChiAdapter) GetListProvider(w http.ResponseWriter, req *http.Request) {
	kodeProduk := req.URL.Query().Get("kode_product")
	tmpl, err := a.template.GetListProvider()
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	providers, err := a.service.GetProductProvider(ctx, kodeProduk)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "list-provider", M{
		"Providers": providers,
		"Product": M{
			"KodeProduk": kodeProduk,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
