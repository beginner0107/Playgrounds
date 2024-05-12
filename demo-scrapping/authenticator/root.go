package authenticator

import (
	"demo-scrapping/config"
	"encoding/base32"
)

type authenticator struct {
	cfg          *config.Config
	secretBase32 string
}

type AuthenticatorImpl interface{}

func NewAuthenticator(cfg *config.Config) AuthenticatorImpl {
	a := &authenticator{cfg: cfg}

	authCfg := cfg.Authenticator

	var secret []byte

	for _, char := range authCfg.Secret {
		secret = append(secret, byte(char))
	}

	a.secretBase32 = base32.StdEncoding.EncodeToString(secret)
	account := authCfg.Account
	issuer := authCfg.Issuer

	return a
}

func (a *authenticator) VerifySecret(secret string) (bool, error) {
	return false, nil
}
