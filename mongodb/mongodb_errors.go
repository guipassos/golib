package mongodb

import "errors"

var (
	ErrCouldNotConnect = func(err error) error {
		return errors.New("could not connect to mongodb, details: " + err.Error())
	}
	ErrCantPing = func(err error) error {
		return errors.New("can't ping mongodb, details: " + err.Error())
	}
	ErrInvalidConnString = func(err error) error {
		return errors.New("invalid connection string, details: " + err.Error())
	}
)
