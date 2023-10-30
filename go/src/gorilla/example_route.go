package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Item represents a simple item in our API.
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

func main() {
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items", AddItem).Methods("POST")

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// GetItems handles the GET request to retrieve all items.
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// AddItem handles the POST request to add a new item.
func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the new item (for simplicity, we use a counter here).
	newItem.ID = "item" + string(len(items)+1)

	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}
