package lib

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("updateTableData", updateTableDataValidator)
	_ = validate.RegisterValidation("notEmpty", notEmtpy)
}

func ValidateStruct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ValidateVar(value interface{}, tags string) error {

	errs := validate.Var(value, tags)

	if errs != nil {
		return errs
	}
	return nil
}

func notEmtpy(fl validator.FieldLevel) bool {
	return fl.Field().Len() != 0
}

func updateTableDataValidator(fl validator.FieldLevel) bool {
	typeParam := fl.Parent().FieldByName("Type")
	switch typeParam.String() {
	case "add":
		if fl.Field().Kind() != reflect.Map {
			return false
		}
		return true
	case "modify":
		if fl.Field().Kind() != reflect.Map {
			return false
		}
		return true
	case "delete":
		if fl.Field().Kind() != reflect.Slice {
			return false
		}
		return true
	}
	return true
}
