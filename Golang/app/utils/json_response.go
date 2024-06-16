package utils

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error     bool        `json:"error"`
	FieldName string      `json:"field_name"`
	Value     interface{} `json:"value"`
	Message   string      `json:"message"`
}
