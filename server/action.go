package server

import (
	"action/models"
	"encoding/json"
	"sync"
)

/*
  NOTE: To extend the feature in the future, we could setup API calls (CRUD operations)
        for the addAction and getAction function below.
*/

// actionMap stores the total time for each action
var actionMap = make(models.ActionMap)

// addAction accepts a json serialized string and maintains average time for
// each action
func addAction(s string) error {
	var mutex = &sync.Mutex{}
	var actionInput models.ActionInput

	// decode the input (string s)
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
	// NOTE: for future consideration: we could save the map in database.
	mutex.Lock()
	ac := actionMap[actionInput.Action]
	actionMap[actionInput.Action] = models.ActionCounter{
		TotalTime: actionInput.Time + ac.TotalTime,
		Counter:   ac.Counter + 1,
	}
	mutex.Unlock()

	return nil

}

// getStats returns a serialized json array of the average time for each action
func getStats() string {
	var mutex = &sync.Mutex{}
	var stat []models.ActionOutput
	//using mutex to protect critical section and prevent race conditions.
	mutex.Lock()
	for action, actionCounter := range actionMap {
		avgTime := actionCounter.TotalTime / (float64)(actionCounter.Counter)
		ao := models.ActionOutput{
			Action: action,
			Avg:    avgTime,
		}
		stat = append(stat, ao)
	}
	actionStats, _ := json.Marshal(stat)
	mutex.Unlock()

	//convert to serialized json string array
	return string(actionStats)

}
