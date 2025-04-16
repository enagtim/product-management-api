package category

import (
	"net/http"
	"product-management-api/pkg/messages"
	"product-management-api/pkg/req"
	"product-management-api/pkg/res"
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

		payload, err := req.DecodedAndValidatedPayload[CategoryPayloadRequest](r)
		if err != nil {
			messages.ErrJson(w, err.Error(), http.StatusBadRequest)
			return
		}

		category, err := h.CategoryService.Create(payload)

		if err != nil {
			messages.ErrJson(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.ResponseJson(w, http.StatusCreated, &category)
	}
}
