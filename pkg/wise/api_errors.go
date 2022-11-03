package wise

import "fmt"

type ValidationErrorDetail struct {
	Code      string   `json:"code"`
	Message   string   `json:"message"`
	Arguments []string `json:"arguments"`
}

type ValidationError struct {
	Errors []ValidationErrorDetail `json:"errors"`
}

func (v ValidationError) Error() string {
	errStr := "validation errors:"
	for _, validationErr := range v.Errors {
		errStr = errStr + " " + validationErr.Message
	}
	return errStr
}

type AuthenticationError struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (a AuthenticationError) Error() string {
	return fmt.Sprintf("authentication error. error: %s. error description: %s", a.Err, a.ErrorDescription)
}

type SystemError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Err       string `json:"error"`
	Exception string `json:"exception"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

func (s SystemError) Error() string {
	return fmt.Sprintf("system error. %s. %s. %s. %s", s.Err, s.Exception, s.Message, s.Path)
}
