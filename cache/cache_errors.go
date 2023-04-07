package cache

import (
	"fmt"
)

var (
	ErrPing = func(err error) error {
		return fmt.Errorf("fail to ping cache provider, details: %v", err)
	}
	ErrInvlaidProvider = func(providerType string) error {
		return fmt.Errorf("invalid cache provider: '%s'", providerType)
	}
	ErrValueIsNotPointer = func() error {
		return fmt.Errorf("value is not pointer")
	}
)
