package httphelper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func GetBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	if bodyReader, ok := body.(strings.Reader); ok {
		return &bodyReader, nil
	}
	if bodyReader, ok := body.(io.Reader); ok {
		return bodyReader, nil
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bodyBytes), nil
}

func DecodeResponse(response *http.Response, dest interface{}) *ResponseFail {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		var decodeData map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &decodeData); err != nil {
			return &ResponseFail{
				Err:        err,
				StatusCode: response.StatusCode,
				Header:     response.Header,
				Body:       bodyBytes,
			}
		}
		return &ResponseFail{
			Err:        ErrInvalidStatusCode,
			StatusCode: response.StatusCode,
			Data:       decodeData,
			Header:     response.Header,
			Body:       bodyBytes,
		}
	}
	if dest != nil {
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return &ResponseFail{
				Err:        err,
				StatusCode: response.StatusCode,
				Header:     response.Header,
				Body:       bodyBytes,
			}
		}
	}
	return nil
}
