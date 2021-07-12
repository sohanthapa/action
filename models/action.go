package models

type ActionInput struct {
	Action string `json:"action"` // the name of the action
	Time   uint   `json:"time"`   // time taken for the action
}

type ActionOutput struct {
	Action string // name of the action
	Avg    uint   // average time taken for each action
}

type ActionMap map[string]uint

// Validate validates the field for ActionInput struct
func (actionInput ActionInput) Validate() error {
	if actionInput.Action == "" {
		return ErrActionFieldEmpty
	}
	if actionInput.Time == 0 {
		return ErrTimeFieldZero
	}
	return nil
}
