package server

import (
	"action/models"
	"encoding/json"
	"sync"
)

// actionMap stores the total time for each action
var actionMap = make(models.ActionMap)

// addAction accepts a json serialized string and maintains average time for
// each action
func addAction(s string) error {
	var mutex = &sync.Mutex{}
	var actionInput models.ActionInput

	// decode the input s
	err := json.Unmarshal([]byte(s), &actionInput)
	if err != nil {
		return ErrJSONMalformed
	}

	// validate the input fields
	err = actionInput.Validate()
	if err != nil {
		return err
	}
	//using mutex to protect critical section and prevent race conditions.
	mutex.Lock()
	actionMap[actionInput.Action] = actionMap[actionInput.Action] + actionInput.Time
	mutex.Unlock()

	return nil

}

// getStats returns a serialized json array of the average time for each action
func getStats() string {
	var mutex = &sync.Mutex{}
	var output []models.ActionOutput
	//using mutex to protect critical section and prevent race conditions.
	mutex.Lock()
	for action, time := range actionMap {
		ao := models.ActionOutput{
			Action: action,
			Avg:    time / 2,
		}
		output = append(output, ao)
	}
	actionOutput, _ := json.Marshal(output)
	mutex.Unlock()

	return string(actionOutput)
}
