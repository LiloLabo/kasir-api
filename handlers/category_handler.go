package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategorys - GET /api/produk
func (h *CategoryHandler) HandleCategorys(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Categorys, err := h.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusInternalServerError,
			Message: "General error",
			Data:    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Categories list",
		Data:    Categorys,
	})
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Category models.Category
	err := json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	err = h.service.Create(&Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "New Category is added successfully",
		Data:    Category,
	})
}

// HandleCategoryByID - GET/PUT/DELETE /api/produk/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetByID - GET /api/produk/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/v2/categories/")
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

	Category, err := h.service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusNotFound,
			Message: "Category not found",
			Data:    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Category details",
		Data:    Category,
	})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/v2/categories/")
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

	var Category models.Category
	err = json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	Category.ID = id
	err = h.service.Update(&Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Category ID = " + idStr + " is updated successfully",
		Data:    Category,
	})
}

// Delete - DELETE /api/produk/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/v2/categories/")
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

	err = h.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Category ID = " + idStr + " is deleted successfully",
		Data:    nil,
	})
}
