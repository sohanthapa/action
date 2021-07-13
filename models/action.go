package models

//ActionInput represents information needed for each action input
type ActionInput struct {
	Action string `json:"action"` // the name of the action
	Time   uint   `json:"time"`   // time taken for the action
}

//ActionOutput represents information needed for each action output
type ActionOutput struct {
	Action string `json:"action"` // name of the action
	Avg    uint   `json:"avg"`    // average time taken for each action
}

type ActionCounter struct {
	TotalTime uint
	Counter   uint
}

type ActionMap map[string]ActionCounter

// Validate validates the field for ActionInput struct
// NOTE: assumptions made: Action field cannot empty,
//                         Time field value needs to be greater than 0
func (actionInput ActionInput) Validate() error {
	if actionInput.Action == "" {
		return ErrActionFieldEmpty
	}
	if actionInput.Time == 0 {
		return ErrTimeFieldZero
	}
	return nil
}
