package req

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodedAndValidatedPayload[T any](r *http.Request) (*T, error) {
	var payload T

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, errors.New("ошибка декодирования тела запроса")
	}

	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		return nil, errors.New("ошибка валидации тела запроса")
	}
	return &payload, nil
}
