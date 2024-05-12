package authenticator

import (
	"demo-scrapping/config"
	"encoding/base32"
	"net/url"
	"os"

	"rsc.io/qr"
)

type authenticator struct {
	cfg          *config.Config
	secretBase32 string
}

type AuthenticatorImpl interface{}

func NewAuthenticator(cfg *config.Config) (AuthenticatorImpl, error) {
	a := &authenticator{cfg: cfg}

	authCfg := cfg.Authenticator

	var secret []byte

	for _, char := range authCfg.Secret {
		secret = append(secret, byte(char))
	}

	a.secretBase32 = base32.StdEncoding.EncodeToString(secret)
	account := authCfg.Account
	issuer := authCfg.Issuer

	if URL, err := url.Parse("otpauth://totp"); err != nil {
		return nil, err
	} else {
		URL.Path = "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
		params := url.Values{}
		params.Add("secret", a.secretBase32)
		params.Add("issuer", issuer)

		if code, err := qr.Encode(URL.String(), qr.Q); err != nil {
			return nil, err
		} else if err = os.WriteFile(authCfg.FieldName, code.PNG(), 0600); err != nil {
			return nil, err
		} else {
			return a, nil
		}
	}
}

func (a *authenticator) VerifySecret(secret string) (bool, error) {
	return false, nil
}
