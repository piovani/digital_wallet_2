package controller

// ResponseError when an error happens
// @Description Message error message
type ResponseError struct {
	Message string `json:"error_message"`
}

func NewResponseError(err error) *ResponseError {
	return &ResponseError{Message: err.Error()}
}
