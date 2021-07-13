package models

//ActionInput represents information needed for each action input
type ActionInput struct {
	Action string  `json:"action"` // the name of the action
	Time   float64 `json:"time"`   // time taken for the action
}

//ActionOutput represents information needed for each action output
type ActionOutput struct {
	Action string  `json:"action"` // name of the action
	Avg    float64 `json:"avg"`    // average time taken for each action
}

//ActionCounter represents information needed to keep track how many times
// an action is called
type ActionCounter struct {
	TotalTime float64 // total time for each action
	Counter   uint    // counter for the number of times an action is called.
}

//ActionMap stores the action along with its time and counter
// key - action: value - ActionCounter struct
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
