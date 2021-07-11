package server

import (
	"action/models"
	"encoding/json"
	"log"
)

//var mutex sync.Mutex
var TT = make(models.TimeStore)

func addAction(s string) error {
	var aa models.ActionInput

	err := json.Unmarshal([]byte(s), &aa)
	if err != nil {
		return ErrJSONMalformed
	}
	TT[aa.Action] = TT[aa.Action] + aa.Time
	//getStats()

	return nil

}

func getStats() {
	var a []models.ActionOutput
	for k, v := range TT {
		at := models.ActionOutput{
			Action: k,
			Avg:    v,
		}
		a = append(a, at)
	}

	log.Printf("\n\n val is %v \n\n", a)
}
