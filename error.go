package liquidity

import (
	"encoding/json"
)

func (e Error) Error() string {
	marshal, _ := json.Marshal(e.ValidationError)
	return e.Message + ": " + string(marshal)
}

type Error struct {
	Message         string      `json:"message"`
	ValidationError interface{} `json:"validationError"`
}
