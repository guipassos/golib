//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package crypt

type Crypt interface {
	NewHashSha1(str string) string
	NewHashSha256(str string) string
	PasswordEncrypt(password string) ([]byte, error)
	PasswordCompare(encryptedPassword, password []byte) error
	RandomBytes(size int) ([]byte, error)
}

type cryptImpl struct {
}

func New() Crypt {
	return cryptImpl{}
}
