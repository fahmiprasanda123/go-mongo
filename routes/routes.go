package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// Item represents a simple data structure
type Item struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

var items []Item

// GetItems handles GET requests and returns all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// CreateItem handles POST requests to add a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

// UpdateItem handles PUT requests to update an item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var updatedItem Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)

	for index, item := range items {
		if item.ID == id {
			items[index].Name = updatedItem.Name
			json.NewEncoder(w).Encode(items[index])
			return
		}
	}
	json.NewEncoder(w).Encode("No item found with given ID.")
}

// DeleteItem handles DELETE requests to delete an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for index, item := range items {
		if item.ID == id {
			items = append(items[:index], items[index+1:]...)
			json.NewEncoder(w).Encode("Item deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No item found with given ID.")
}

// SetupRoutes sets up the router
func SetupRoutes(db *mongo.Database) *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")    // Update route
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE") // Delete route

	return router
}
