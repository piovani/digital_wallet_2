package controller

type ResponseError struct {
	Message string `json:"error_message"`
}

func NewResponseError(err error) *ResponseError {
	return &ResponseError{Message: err.Error()}
}
