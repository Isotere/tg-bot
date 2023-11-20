package telegram

type UpdateOut struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int64  `json:"update_id"`
	Message string `json:"message"`
}
