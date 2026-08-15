package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path"

	"github.com/XPD-Kasun/layout"
	"github.com/XPD-Kasun/layout/examples/default/models"
	"github.com/XPD-Kasun/layout/examples/default/viewmodels"
)

func loadProducts() []models.Product {
	productData, err := os.ReadFile(path.Join("data", "products.json"))
	if err != nil {
		panic(errors.Join(errors.New("Could not load product data"), err))
	}

	var products []models.Product
	err = json.Unmarshal(productData, &products)
	if err != nil {
		panic(err)
	}

	return products
}

func main() {

	// load the data
	products := loadProducts()
	// create the layout
	view := layout.New()

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		listingVM := viewmodels.ListingVM{Title: "Product Listing", Products: products}

		err := view.Render(w, "products/listing", listingVM)
		if err != nil {
			println(err.Error())
		}
	})

	http.ListenAndServe("localhost:3002", server)

}
