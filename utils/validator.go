package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	// Initialize the validator
	validate = validator.New()
}

func Validate(inf interface{}) error {
	err := validate.Struct(inf)
	fmt.Println(err)
	return err

}
