package main

import (
	"encoding/json"
	. "fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	ID          string  `json:"id"`
	Productcode string  `json:"productcode"`
	ProductName string  `json:"productName"`
	Seller      *Seller `json:"seller"`
}

type Seller struct {
	CompanyName string `json:"companyname"`
	Location    string `json:"location"`
}

var products []Products

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products) //return the existing movies
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Products
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(200000324))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//set the params
	params := mux.Vars(r)
	//loop over the products
	for index, item := range products {
		//delete the product with the id thst is been sent
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			// add a new product
			var product Products
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	products = append(products, Products{ID: "1", Productcode: "4534678A",
		ProductName: "Milk", Seller: &Seller{CompanyName: "REWE", Location: "Berlin"}})
	products = append(products, Products{ID: "2", Productcode: "284756362F", ProductName: "Biscuit",
		Seller: &Seller{CompanyName: "Griesson de Beukelar", Location: "Kahla"}})
	products = append(products, Products{ID: "3", ProductName: "Water", Productcode: "38475738947G",
		Seller: &Seller{CompanyName: "Ja", Location: "Erfurt"}})

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProducts).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	Printf("starting server at port 9090\n")
	log.Fatal(http.ListenAndServe(":9090", r))
}
