package chi

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (a *ChiAdapter) GetPartnerProdukForm(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("idproduk")
	tmpl, err := a.template.GetPartnerProdukCreateForm()
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "content", M{
		"Idproduk": id,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (a *ChiAdapter) GetListPartnerProduk(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("idproduk")
	idproduk, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	partnerProduks, err := a.service.GetListPartnerProduk(ctx, int32(idproduk))

	tmpl, err := a.template.GetListPartnerProduk()
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "list-partner", M{
		"Idproduk": id,
		"Partners": partnerProduks,
	})

	if err != nil {
		fmt.Println(err)
	}
}
