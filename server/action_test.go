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
		input := `{"action":"jump","time":100}`
		input2 := `{"action":"run","time":250}`
		input3 := `{"action":"jump","time":200}`
		input4 := `{"action":"run","time":400}`
		input5 := `{"action":"jump","time":500}`
		//first add actions along with its time.
		// concurrent calls using go routine
		addAction(input)
		addAction(input2)
		addAction(input3)
		addAction(input4)
		addAction(input5)

		//time.Sleep(500 * time.Millisecond)

		actionOuput := models.ActionOutput{
			Action: "jump",
			Avg:    266,
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
