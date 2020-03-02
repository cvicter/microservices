package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservices/order/db"
	"microservices/order/queue"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ProductID string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAT time.Time `json:"created_at"`
}

var productsURL string

func createOrder(payload []byte) {
	var order Order
	json.Unmarshal(payload, &order)

	id := uuid.New()
	order.ID = id.String()
	order.Status = "pendente"
	order.CreatedAT = time.Now()
	saveOrder(order)
}

func getProductById(id string) Product {
	response, _ := http.Get(productsURL + "/product/" + id)
	data, _ := ioutil.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)
	return product

}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)
	connection := db.Connect()

	err := connection.Set(order.ID, string(json), 0).Err()

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming(connection, in)

	for payload := range in {
		createOrder(payload)
		fmt.Println(string(payload))

	}

}
