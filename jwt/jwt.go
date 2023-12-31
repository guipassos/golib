//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package jwt

import (
	"context"
	"time"

	"github.com/guipassos/golib/entity"
)

type (
	JWT interface {
		NewToken(secret, issuer, subject string, duration time.Duration) (string, error)
		GetUserID(ctx context.Context) string
		GetUserEmail(ctx context.Context) string
		GetUserClaims(ctx context.Context) map[string]interface{}
	}
	jwtImpl struct {
		entity entity.Entity
	}
)

func New() JWT {
	return jwtImpl{
		entity: entity.New(),
	}
}
