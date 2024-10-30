// pkg/utils/validator.go
package utils

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate // Gunakan nama `Validator` atau `validator` sesuai dengan inisialisasi
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i) // Pastikan nama field sesuai
}
