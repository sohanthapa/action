package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("error - empty action field", func(t *testing.T) {
		actionInput := ActionInput{
			Action: "",
			Time:   100,
		}
		err := actionInput.Validate()
		assert.Equal(t, ErrActionFieldEmpty, err)
	})
	t.Run("error - time field is zero", func(t *testing.T) {
		actionInput := ActionInput{
			Action: "run",
			Time:   0,
		}
		err := actionInput.Validate()
		assert.Equal(t, ErrTimeFieldZero, err)
	})

	t.Run("success", func(t *testing.T) {
		actionInput := ActionInput{
			Action: "run",
			Time:   50,
		}
		err := actionInput.Validate()
		assert.Nil(t, err)
	})
}
