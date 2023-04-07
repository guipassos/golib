//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package entity

import (
	"time"
)

type Entity interface {
	NewID() string
	Now() time.Time
	NowMillis() int64
}

type entityImpl struct {
}

func New() Entity {
	return entityImpl{}
}
