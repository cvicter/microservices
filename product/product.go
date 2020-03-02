package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Products []Product
}

func loadData() []byte {
	jsonFile, _ := os.Open("products.json")

	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)

	return data

}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()
	var products Products
	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.ID == vars["id"] {
			product, _ := json.Marshal(v)
			w.Write([]byte(product))
		}
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/product/{id}", GetProductById)
	http.ListenAndServe(":8081", r)
}
