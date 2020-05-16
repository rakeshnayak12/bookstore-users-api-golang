package error

// CustomError hangels errors
type CustomError struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	StatusCode   int    `json:"statusCode"`
}

// HandleError handels errors
func HandleError(errorDef string, message string, code int) *CustomError {
	return &CustomError{
		Error:        errorDef,
		ErrorMessage: message,
		StatusCode:   code,
	}
}
