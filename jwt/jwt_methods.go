package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (j jwtImpl) NewToken(jwtSecret, issuer, subject string, duration time.Duration) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        j.entity.NewID(),
		Issuer:    issuer,
		Subject:   subject,
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		ExpiresAt: now.Add(duration).Unix(),
	})
	return token.SignedString([]byte(jwtSecret))
}

func (j jwtImpl) GetUserID(ctx context.Context) string {
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s", claims["iss"])
}

func (j jwtImpl) GetUserEmail(ctx context.Context) string {
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s", claims["sub"])
}

func (j jwtImpl) GetUserClaims(ctx context.Context) map[string]interface{} {
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		return nil
	}
	return claims
}
