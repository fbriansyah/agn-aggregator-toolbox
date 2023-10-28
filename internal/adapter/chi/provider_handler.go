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

	err = tmplFile.ExecuteTemplate(w, "base", M{
		"IDProvider": idProvider,
		"PageTitle":  "Provider",
		"Provider":   provider,
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
