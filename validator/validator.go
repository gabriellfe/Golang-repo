package validator

import (
	"fmt"
	"gabriellfe/dto"
	"reflect"
	"regexp"
	"strings"
)

// Name of the struct tag used in examples.
const tagName = "validate"

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}, reflect.Value) (bool, string)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}, kind reflect.Value) (bool, string) {
	return true, ""
}

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	Min int
	Max int
}

func (v StringValidator) Validate(val interface{}, kind reflect.Value) (bool, string) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Sprintf("cannot be blank or null")
	}

	if l < v.Min {
		return false, fmt.Sprintf("should be at least %v chars long", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Sprintf("should be less than %v chars long", v.Max)
	}

	return true, ""
}

// NumberValidator performs numerical value validation.
// Its limited to int type for simplicity.
type NumberValidator struct {
	Min int
	Max int
}

func (v NumberValidator) Validate(val interface{}, value reflect.Value) (bool, string) {
	kind := value.Kind()

	switch kind {
	case reflect.Int:
		if float64(val.(int)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Int8:
		if float64(val.(int8)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Int16:
		if float64(val.(int16)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Int32:
		if float64(val.(int32)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Int64:
		if float64(val.(int64)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Uint:
		if float64(val.(uint)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Uint8:
		if float64(val.(uint8)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Uint16:
		if float64(val.(uint16)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Uint32:
		if float64(val.(uint32)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Uint64:
		if float64(val.(uint64)) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Float32:
		if val.(float32) < float32(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	case reflect.Float64:
		if val.(float64) < float64(v.Min) {
			return false, fmt.Sprintf("field must be greater or equal than %d", v.Min)
		}
	}

	switch kind {
	case reflect.Int:
		if float64(val.(int)) > float64(v.Max) {
			return false, fmt.Sprintf("Field should be less or equal than %d", v.Max)
		}
	case reflect.Int8:
		if float64(val.(int8)) > float64(v.Max) {
			return false, fmt.Sprintf("Field should be less or equal than %d", v.Max)
		}
	case reflect.Int16:
		if float64(val.(int16)) > float64(v.Max) {
			return false, fmt.Sprintf("Field should be less or equal than %d", v.Max)
		}
	case reflect.Int32:
		if float64(val.(int32)) > float64(v.Max) {
			return false, fmt.Sprintf("Field should be less or equal than %d", v.Max)
		}
	case reflect.Int64:
		if float64(val.(int64)) > float64(v.Max) {
			return false, fmt.Sprintf("Field should be less or equal than %d", v.Max)
		}
	case reflect.Uint:
		if float64(val.(uint)) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Uint8:
		if float64(val.(uint8)) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Uint16:
		if float64(val.(uint16)) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Uint32:
		if float64(val.(uint32)) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Uint64:
		if float64(val.(uint64)) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Float32:
		if val.(float32) > float32(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	case reflect.Float64:
		if val.(float64) > float64(v.Max) {
			return false, fmt.Sprintf("field must be less or equal than %d", v.Max)
		}
	}
	return true, ""
}

// EmailValidator checks if string is a valid email address.
type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}, kind reflect.Value) (bool, string) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Sprintf("field cannot be blank or null")
	}

	if !mailRe.MatchString(val.(string)) {
		return false, fmt.Sprintf("field is not a valid email address")
	}
	return true, ""
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		return EmailValidator{}
	}

	return DefaultValidator{}
}

// Performs actual data validation using validator definitions on the struct
func ValidateStruct(s interface{}) dto.ErrorSchemaDto {
	errs := []string{}
	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)
		value := v.Field(i)
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface(), value)

		// Append error to results
		if !valid && err != "" {
			errs = append(errs, fmt.Sprintf("%s %s", v.Type().Field(i).Name, err))
		}
	}

	return dto.ErrorSchemaDto{Errors: errs}
}
