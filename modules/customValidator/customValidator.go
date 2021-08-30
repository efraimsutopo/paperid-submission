package customValidator

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func New() *CustomValidator {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &CustomValidator{
		Validator: validate,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func TransformValidatorError(err error) structs.ErrorResponse {
	var result = structs.ErrorResponse{
		Code: http.StatusBadRequest,
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				result.Message = fmt.Sprintf("field '%s' is required",
					err.Field())
			case "oneof":
				result.Message = fmt.Sprintf("field '%s' must be one of (%s)",
					err.Field(), err.Param())
			default:
				result.Message = fmt.Sprintf("field '%s' failed to satisfy validation %s", err.Field(), err.ActualTag())
			}

			break
		}
	}

	return result
}
