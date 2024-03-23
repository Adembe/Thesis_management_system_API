package dto

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Body    any    `json:"body"`
}
