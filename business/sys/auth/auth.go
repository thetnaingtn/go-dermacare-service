package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	method  jwt.SigningMethod
	parser  jwt.Parser
	keyFunc jwt.Keyfunc
}

func New() (*Auth, error) {
	method := jwt.GetSigningMethod("RS256")
	if method == nil {
		return nil, errors.New("Configuring algorithm RS256")
	}

	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	keyFunc := func(t *jwt.Token) (any, error) {
		publicKey, err := getPublicKey(privateKey)
		if err != nil {
			return nil, err
		}

		return publicKey, nil
	}

	parser := jwt.Parser{
		ValidMethods: []string{"RS256"},
	}

	a := &Auth{
		method:  method,
		keyFunc: keyFunc,
		parser:  parser,
	}

	return a, nil
}

func (a *Auth) GenerateToken() (string, error) {
	return "", nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	fsys := os.DirFS("/key")
	file, err := fsys.Open("private.pem")
	if err != nil {
		return nil, fmt.Errorf("open private.pem: %w", err)
	}
	defer file.Close()

	privatePEM, err := io.ReadAll(io.LimitReader(file, 1024*1024))
	if err != nil {
		return nil, fmt.Errorf("reading private.pem: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}

	return privateKey, nil
}

func getPublicKey(p *rsa.PrivateKey) (*rsa.PublicKey, error) {
	return &p.PublicKey, nil
}
