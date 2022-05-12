package liquidity

func (e Error) Error() string {
	return e.Message
}

type Error struct {
	Message         string `json:"message"`
	ValidationError []struct {
		Code     string   `json:"code"`
		Expected string   `json:"expected,omitempty"`
		Received string   `json:"received,omitempty"`
		Path     []string `json:"path"`
		Message  string   `json:"message"`
	} `json:"validationError"`
}
