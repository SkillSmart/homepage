package utils


// Public Interface to the config section of the error package
var (
	ApplicationError *applicationInterface
)

type applicationInterface struct {}
func (*applicationInterface) Info(msg string) *applicationError {
	return &applicationError{
		Message:    msg,
		StatusCode: 1,
		Code:       "Not implemented! yet...",
	}
}
func (*applicationInterface) Debug(msg string) *applicationError {
	return &applicationError{
		Message:    msg,
		StatusCode: 1,
		Code:       "Not implemented! yet...",
	}
}
func (*applicationInterface) Fatal(msg string) *applicationError {
	return &applicationError{
		Message:    msg,
		StatusCode: 1,
		Code:       "Not implemented! yet...",
	}
}
// ApplicationError reports unexpected server side errors and is for internal use only
type applicationError struct {
	Message string 	`json:"message"`
	StatusCode int `json:"status"`
	Code string `json:"code"`
}
func (a *applicationError) Error() string {
	return a.Message
}

// RequestError reports errors during request handling back to the client
type Response struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	DocumentationUrl string `json"documentation_url"`
	Errors []*error	`json:"errors`
}






