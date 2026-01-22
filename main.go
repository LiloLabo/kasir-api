package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories []Category = []Category{
	{ID: 1, Name: "Category 1", Description: "Description of Category 1"},
	{ID: 2, Name: "Category 2", Description: "Description of Category 2"},
	{ID: 3, Name: "Category 3", Description: "Description of Category 3"},
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Categories list",
		Data:    categories,
	})
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for _, p := range categories {
		if p.ID == id {
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category details",
				Data:    p,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}

func postCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCategory Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "New category is added successfully",
		Data:    newCategory,
	})
}

func putCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var updatedCategory Category
	if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for i, p := range categories {
		if p.ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category ID = " + idStr + " is updated successfully",
				Data:    updatedCategory,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for i, p := range categories {
		if p.ID == id {
			deletedCategory := p
			categories = append(categories[:i], categories[i+1:]...)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category ID = " + idStr + " is deleted successfully",
				Data:    deletedCategory,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}

var products []Product = []Product{
	{Id: 1, Name: "Product 1", Price: 10000, Stock: 10},
	{Id: 2, Name: "Product 2", Price: 20000, Stock: 20},
	{Id: 3, Name: "Product 3", Price: 30000, Stock: 30},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Products list",
		Data:    products,
	})
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for _, p := range products {
		if p.Id == id {
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Product details",
				Data:    p,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}

func postProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	newProduct.Id = len(products) + 1
	products = append(products, newProduct)

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "Product added successfully",
		Data:    newProduct,
	})
}

func putProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for i, p := range products {
		if p.Id == id {
			updatedProduct.Id = id
			products[i] = updatedProduct
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Product updated successfully",
				Data:    updatedProduct,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for i, p := range products {
		if p.Id == id {
			deletedProduct := p
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Product deleted successfully",
				Data:    deletedProduct,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}

func main() {
	fmt.Println("Server started on :8000")

	http.HandleFunc("GET /api/products", getProducts)
	http.HandleFunc("POST /api/products", postProduct)
	http.HandleFunc("GET /api/products/{id}", getProduct)
	http.HandleFunc("PUT /api/products/{id}", putProduct)
	http.HandleFunc("DELETE /api/products/{id}", deleteProduct)

	http.HandleFunc("GET /categories", getCategories)
	http.HandleFunc("POST /categories", postCategory)
	http.HandleFunc("GET /categories/{id}", getCategory)
	http.HandleFunc("PUT /categories/{id}", putCategory)
	http.HandleFunc("DELETE /categories/{id}", deleteCategory)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusOK,
			Message: "Hello world!",
			Data: map[string]interface{}{
				"app":     "Cashier API",
				"version": "1.0.0",
			},
		})
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
