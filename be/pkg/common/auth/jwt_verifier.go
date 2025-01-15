package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrJwtBadScheme         = errors.New("JWT authorization scheme is invalid")
	ErrJwtBadAlg            = errors.New("JWT Inconsistent Algorithm")
	ErrJwtInvalid           = errors.New("JWT is invalid")
	ErrJwtNotFoundRealm     = errors.New("JWT tenant not set")
	ErrJwtInvalidSubject    = errors.New("JWT subject is invalid")
	ErrJwtNotFoundMapClaims = errors.New("JWT MapClaims not found")
)

type (
	Verifier interface {
		Verify(jwtPayload string) (AuthInfo, error)
	}

	verifier struct {
		signMethod string
		verifyKey  *rsa.PublicKey
	}
)

const (
	realmKey string = "realm"
	idKey    string = "jti"
	subKey   string = "sub"
	rolesKey string = "roles"
	algKey   string = "alg"
)

var _ (Verifier) = (verifier)(verifier{})

// TODO testing
func (c verifier) Verify(jwtPayload string) (AuthInfo, error) {

	jwtTokenBytes, err := base64.StdEncoding.DecodeString(jwtPayload)
	if err != nil {
		return AuthInfo{}, err
	}

	jwtToken, err := jwt.Parse(string(jwtTokenBytes), func(token *jwt.Token) (interface{}, error) {
		return c.verifyKey, nil
	})
	if err != nil {
		return AuthInfo{}, err
	}

	jwtAlg := jwtToken.Header[algKey]
	if c.signMethod != jwtAlg {
		return AuthInfo{}, ErrJwtBadAlg
	}

	if !jwtToken.Valid {
		return AuthInfo{}, ErrJwtInvalid
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return AuthInfo{}, ErrJwtNotFoundMapClaims
	}

	subject, ok := (claims[subKey]).(string)
	if !ok {
		return AuthInfo{}, ErrJwtInvalidSubject
	}

	return AuthInfo{
		User: subject,
	}, nil
}

func NewVerifier(crtFile string) (*verifier, error) {
	crtBytes, err := os.ReadFile(crtFile)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(crtBytes)
	if err != nil {
		return nil, err
	}

	return &verifier{
		signMethod: "RS256",
		verifyKey:  verifyKey,
	}, nil
}
