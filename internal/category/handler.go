package category

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	CategoryService *CategoryService
}

func NewCategoryHandler(router *http.ServeMux, categoryService *CategoryService) {
	handler := &CategoryHandler{
		CategoryService: categoryService,
	}
	router.HandleFunc("POST /categories", handler.CreateCategory())
}

func (h *CategoryHandler) CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CategoryPayloadRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)

		if err != nil {
			http.Error(w, "ошибка валидации", http.StatusBadRequest)
			return
		}

		category, err := h.CategoryService.Create(&payload)

		if err != nil {
			http.Error(w, "Ошибка создания сущности в базе данных", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&category)

	}
}
