package apperr

import (
	"fmt"
	"strings"
)

func (a AppErr) Error() string {
	var elements = make([]string, 0, 4)
	if a.HTTPCode > 0 {
		elements = append(elements, fmt.Sprintf("http: %d", a.HTTPCode))
	}
	if a.Err != nil {
		elements = append(elements, fmt.Sprintf("err: %v", a.Err))
	}
	if a.Key != "" {
		elements = append(elements, fmt.Sprintf("key: %s", a.Key))
	}
	if a.Message != "" {
		elements = append(elements, fmt.Sprintf("msg: %s", a.Message))
	}
	return strings.Join(elements, ", ")
}
