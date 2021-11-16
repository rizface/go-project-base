package test

import (
	"base-project/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_,err := helper.GeneratePassword("rahasia")
		assert.Equal(t, nil,err)
	})
	t.Run("emptyPassword", func(t *testing.T) {
		_,err := helper.GeneratePassword("")
		assert.Nil(t, err)
	})
}