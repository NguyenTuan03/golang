package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationRequired(fieldName string, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required ", fieldName)
	}
	return nil
}

// GetValidationErrorMsg dịch các lỗi từ validator sang Tiếng Việt
func GetValidationErrorMsg(err error) map[string]string {
	errors := make(map[string]string)
	if valErrs, ok := err.(validator.ValidationErrors); ok {
		for _, f := range valErrs {
			field := f.Field()
			tag := f.Tag()
			param := f.Param()

			switch tag {
			case "required":
				errors[field] = fmt.Sprintf("%s là bắt buộc", field)
			case "min":
				errors[field] = fmt.Sprintf("%s tối thiểu là %s ký tự", field, param)
			case "max":
				errors[field] = fmt.Sprintf("%s tối đa là %s ký tự", field, param)
			case "email":
				errors[field] = fmt.Sprintf("%s không đúng định dạng email", field)
			default:
				errors[field] = fmt.Sprintf("%s không hợp lệ", field)
			}
		}
	} else {
		errors["message"] = err.Error()
	}

	return errors
}
