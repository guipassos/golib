package apperr

import (
	"errors"
	"strconv"
	"strings"
)

func GetFromString(str string) *AppErr {
	var (
		elements = strings.Split(str, ", ")
		appERR   = &AppErr{}
	)
	if len(elements) == 0 {
		return nil
	}
	for _, element := range elements {
		switch {
		case strings.HasPrefix(element, "http: "):
			appERR.HTTPCode, _ = strconv.Atoi(strings.TrimPrefix(element, "http: "))
		case strings.HasPrefix(element, "err: "):
			appERR.Err = errors.New(strings.TrimPrefix(element, "err: "))
		case strings.HasPrefix(element, "key: "):
			appERR.Key = strings.TrimPrefix(element, "key: ")
		case strings.HasPrefix(element, "msg: "):
			appERR.Message = strings.TrimPrefix(element, "msg: ")
		}
	}
	return appERR
}
