package models

type JSONSuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type JSONErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
