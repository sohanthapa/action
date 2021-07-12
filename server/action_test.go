package server

import (
	"action/models"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddAction tests the functionality for addAction function
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

//TestGetAction tests the functionality for getAction function
func TestGetAction(t *testing.T) {
	t.Run("get stats success", func(t *testing.T) {
		input := `{"action":"jump","time":100}`
		input2 := `{"action":"run","time":250}`
		input3 := `{"action":"jump","time":200}`
		input4 := `{"action":"run","time":400}`
		//first add actions along with its time
		addAction(input)
		addAction(input2)
		addAction(input3)
		addAction(input4)
		actionOuput := models.ActionOutput{
			Action: "jump",
			Avg:    150,
		}
		actionOutput2 := models.ActionOutput{
			Action: "run",
			Avg:    325,
		}
		outputStat := []models.ActionOutput{actionOuput, actionOutput2}
		expectedBody, _ := json.Marshal(outputStat)
		output := getStats()
		assert.Equal(t, string(expectedBody), output)

	})

}
