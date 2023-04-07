package apperr

type (
	AppErr struct {
		HTTPCode int         `json:"http_code" bson:"http_code"`
		Err      error       `json:"err" bson:"err"`
		Key      string      `json:"key" bson:"key"`
		Message  string      `json:"message" bson:"message"`
		Data     interface{} `json:"data" bson:"data"`
	}
	Options struct {
		HTTPCode int
		Err      error
		Key      string
		Message  string
		Data     interface{}
		NotPrint bool
	}
)
