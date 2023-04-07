package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func (i entityImpl) NewID() string {
	uuid := uuid.New()
	return strings.Replace(uuid.String(), "-", "", -1)
}

func (i entityImpl) Now() time.Time {
	return time.Now()
}

func (i entityImpl) NowMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
