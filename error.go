package liquidity

func (e Error) Error() string {
	return e.Message
}

type Error struct {
	Message string
}
