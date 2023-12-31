package apperr

import (
	"github.com/guipassos/golib/logger"
)

func New(opt Options) *AppErr {
	appErr := &AppErr{
		HTTPCode: opt.HTTPCode,
		Err:      opt.Err,
		Key:      opt.Key,
		Message:  opt.Message,
		Data:     opt.Data,
	}
	if opt.NotPrint {
		return appErr
	}
	logger.ErrorApp(appErr.Error())
	return appErr
}
