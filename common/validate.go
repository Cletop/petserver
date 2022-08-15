package common

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

type ValidateErrorReason struct {
	Field        string `json:"field"`
	Reason       string `json:"reason"`
	ValidateType string `json:"validate_type"`
}

type ValidateError struct {
	Status  bool                  `json:"status"`
	Message string                `json:"message"`
	Reason  []ValidateErrorReason `json:"reason"`
}
type Validate interface {
	SignalFieldValidate() (bool, ValidateError)
	StructValidate() (bool, ValidateError)
}

type ValidateSignalStruct struct {
	FieldValue string `json:"field"`
	Expected   string `json:"value"`
}

type ValidateStruct struct {
	FieldValues map[string]interface{} `json:"field"`
}

func (validate_struct *ValidateSignalStruct) SignalFieldValidate() (bool, ValidateError) {
	err := validate.Var(validate_struct.FieldValue, validate_struct.Expected)
	if err != nil {
		return false, ValidateError{
			Status:  false,
			Message: err.Error(),
			Reason: []ValidateErrorReason{
				{Field: validate_struct.FieldValue, Reason: err.Error(), ValidateType: validate_struct.Expected},
			},
		}
	}
	return true, ValidateError{Status: true, Message: "", Reason: []ValidateErrorReason{}}
}

func (validate_struct *ValidateStruct) StructValidate() (bool, ValidateError) {
	err := validate.Struct(validate_struct.FieldValues)
	if err != nil {
		return false, ValidateError{Status: false, Message: err.Error(), Reason: GetStructError(err)}
	}
	return true, ValidateError{}
}

func GetStructError(err error) []ValidateErrorReason {
	var validate_errors []ValidateErrorReason
	for _, err := range err.(validator.ValidationErrors) {
		validate_errors = append(validate_errors, ValidateErrorReason{Field: err.Field(), Reason: err.Tag(), ValidateType: err.ActualTag()})
	}
	return validate_errors
}
