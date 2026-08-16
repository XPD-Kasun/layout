package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

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

	server.HandleFunc("/static/{file}", func(w http.ResponseWriter, r *http.Request) {
		fileName := path.Join("wwwroot", r.PathValue("file"))
		f, err := os.Stat(fileName)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if f.IsDir() {
			http.NotFound(w, r)
			return
		}

		reader, err := os.Open(fileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.Copy(w, reader)
	})

	server.HandleFunc("/products/{id}/{productName}", func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid product id.", err.Error())
			http.NotFound(w, r)
			return
		}
		for _, pro := range products {
			if pro.Id == id {
				vm := viewmodels.ProductVM{
					Title:   fmt.Sprintf("%s product", pro.Name),
					Product: pro,
				}

				err := view.Render(w, "products/product", &vm)
				if err != nil {
					fmt.Println("Error at rendering product : ", err.Error())
				}
				break
			}
		}

	})

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		listingVM := viewmodels.ListingVM{Title: "Product Listing", Products: products}

		err := view.Render(w, "products/listing", listingVM)
		if err != nil {
			println(err.Error())
		}
	})

	http.ListenAndServe("localhost:3002", server)

}
