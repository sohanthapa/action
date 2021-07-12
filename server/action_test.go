package server

import (
	"action/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAction(t *testing.T) {
	t.Run("malformed json", func(t *testing.T) {
		input := string([]byte{65, 66, 67, 226, 130, 172})
		err := addAction(input)
		assert.Equal(t, ErrJSONMalformed, err)
	})

	t.Run("add action success", func(t *testing.T) {
		input := `{"action":"jump","time":100}`
		input2 := `{"action":"run","time":200}`
		expectedResponse := make(models.ActionMap)
		expectedResponse["jump"] = 100
		expectedResponse["run"] = 200
		err := addAction(input)
		assert.Nil(t, err)
		err = addAction(input2)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, actionMap)
		actionMap = make(models.ActionMap)
	})

}

func TestGetAction(t *testing.T) {
	t.Run("malformed json", func(t *testing.T) {
		input := string([]byte{65, 66, 67, 226, 130, 172})
		err := addAction(input)
		assert.Equal(t, ErrJSONMalformed, err)
	})

	t.Run("add action success", func(t *testing.T) {
		input := `{"action":"jump","time":200}`
		input2 := `{"action":"run","time":200}`
		input3 := `{"action":"jump","time":300}`
		input4 := `{"action":"run","time":400}`
		addAction(input)
		addAction(input2)
		addAction(input3)
		addAction(input4)
		expectedBody := "[{\"Action\":\"jump\",\"Avg\":250},{\"Action\":\"run\",\"Avg\":300}]"
		output := getStats()
		assert.Equal(t, expectedBody, output)
	})

}
