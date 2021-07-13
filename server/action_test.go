package server

import (
	"action/models"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestAddAction tests the functionality for addAction function
func TestAddAction(t *testing.T) {
	t.Run("malformed json", func(t *testing.T) {
		input := string([]byte{65, 66, 67, 226, 130, 172})
		err := addAction(input)
		assert.Equal(t, ErrJSONMalformed, err)
	})

	t.Run("error - action field empty", func(t *testing.T) {
		input := `{"action":"","time":100}`
		err := addAction(input)
		assert.Equal(t, models.ErrActionFieldEmpty, err)
	})

	t.Run("error - time field zero value", func(t *testing.T) {
		input := `{"action":"jump","time":0}`
		err := addAction(input)
		assert.Equal(t, models.ErrTimeFieldZero, err)
	})

	t.Run("add action success", func(t *testing.T) {
		input := `{"action":"jump","time":100}`
		input2 := `{"action":"run","time":200}`
		input3 := `{"action":"jump","time":1000}`
		input4 := `{"action":"run","time":200}`
		expectedResponse := make(models.ActionMap)
		expectedResponse["jump"] = models.ActionCounter{
			TotalTime: 1100,
			Counter:   2,
		}
		expectedResponse["run"] = models.ActionCounter{
			TotalTime: 400,
			Counter:   2,
		}
		err := addAction(input)
		assert.Nil(t, err)
		err = addAction(input2)
		assert.Nil(t, err)
		err = addAction(input3)
		assert.Nil(t, err)
		err = addAction(input4)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, actionMap)
		actionMap = make(models.ActionMap)
	})

}

//TestGetAction tests the functionality for getAction function
func TestGetAction(t *testing.T) {
	t.Run("get stats success", func(t *testing.T) {
		input := `{"action":"jump","time":500}`
		input2 := `{"action":"run","time":250}`
		input3 := `{"action":"jump","time":200}`
		input4 := `{"action":"run","time":400}`
		input5 := `{"action":"jump","time":500}`
		//first add actions along with its time.
		addAction(input)
		addAction(input2)
		addAction(input3)
		addAction(input4)
		addAction(input5)

		//call getStats()
		output := getStats()
		var actionOutput []models.ActionOutput
		json.Unmarshal([]byte(output), &actionOutput)
		//check if we get the correct avg time for each action
		for _, ao := range actionOutput {
			if ao.Action == "jump" {
				assert.Equal(t, ao.Avg, float64(400))
			} else if ao.Action == "run" {
				assert.Equal(t, ao.Avg, float64(325))
			}

		}

	})

}
