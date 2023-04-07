package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type (
	Options struct {
		URI        string
		CtxTimeout time.Duration
		IsReader   bool
	}
)

func (o *Options) SetDefaults() {
	if o.CtxTimeout == 0 {
		o.CtxTimeout = 10 * time.Second
	}
	if o.URI == "" {
		o.URI = "mongodb://localhost:27017"
	}
}

func (o *Options) ExtractConnString() (connstring.ConnString, error) {
	return connstring.ParseAndValidate(o.URI)
}
