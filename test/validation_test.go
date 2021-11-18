package test

import (
	"base-project/helper"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string `validate:"required"`
}

func TestStructValidation(t *testing.T) {
	v := validator.New()
	t.Run("success none empty struct field", func(t *testing.T) {
		data := User{Name: "Fariz"}
		err := v.Struct(data)
		assert.Nil(t,err)
	})
	t.Run("success validate empty struct field", func(t *testing.T) {
		data := User{}
		err := v.Struct(data)
		if err != nil {
			helper.Panic(err)
		}
	})
}
