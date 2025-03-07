package errorutils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Error modes (how should the application behave when the CheckError function is invoked):
type OnErrorHandler func(error) error

var CheckError = func(err error) error {
	return err
}

func CheckErrorf(format string, a ...interface{}) error {
	if len(a) > 0 {
		return CheckError(fmt.Errorf(format, a...))
	}
	return CheckError(errors.New(format))
}

// Check expected status codes and return error if needed
func CheckResponseStatus(resp *http.Response, expectedStatusCodes ...int) error {
	for _, statusCode := range expectedStatusCodes {
		if statusCode == resp.StatusCode {
			return nil
		}
	}

	errorBody, _ := ioutil.ReadAll(resp.Body)
	return GenerateResponseError(resp.Status, string(errorBody))
}

func GenerateResponseError(status, body string) error {
	return errors.New("Server response: " + status + "\n" + body)
}
