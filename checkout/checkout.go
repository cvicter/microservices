package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservices/checkout/queue"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	ProductID string `json:"product_id"`
}

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("ola checkout"))
	vars := mux.Vars(r)
	response, _ := http.Get(productsURL + "/product/" + vars["id"])

	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("finalizado"))
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.ProductID = r.FormValue("product_id")

	data, _ := json.Marshal(order)
	connection := queue.Connect()

	fmt.Println(string(data))

	queue.Notify(data, "checkout_ex", "", connection)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/finish", finish)
	r.HandleFunc("/product/{id}", displayCheckout)
	http.ListenAndServe(":8082", r)

}
