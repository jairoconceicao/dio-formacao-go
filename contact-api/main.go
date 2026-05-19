package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Contact struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Address *Address `json:"address"`
}

var contacts []Contact = []Contact{
	{ID: "1", Name: "John Doe", Email: "john.doe@example.com", Address: &Address{Street: "123 Main St", City: "Anytown", State: "CA", ZipCode: "12345"}},
	{ID: "2", Name: "Jane Smith", Email: "jane.smith@example.com", Address: &Address{Street: "456 Oak Ave", City: "Somewhere", State: "NY", ZipCode: "67890"}},
}

func ListContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range contacts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = fmt.Sprintf("%d", len(contacts)+1)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contacts)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	// Implementation to update an existing contact
	params := mux.Vars(r)
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	for i, item := range contacts {
		if item.ID == params["id"] {
			contacts[i] = contact
			json.NewEncoder(w).Encode(contacts)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	// Implementation to delete a contact by ID
	params := mux.Vars(r)
	for i, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:i], contacts[i+1:]...)
			json.NewEncoder(w).Encode(contacts)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contacts", ListContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contacts", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
