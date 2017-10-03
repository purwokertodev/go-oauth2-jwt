package token

import (
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Issuer   string
	Audience string
	Subject  string
	Expired  time.Duration
	Lock     *sync.RWMutex
}

func NewClaim(issuer string, audience string, subject string, expired time.Duration) *Claim {
	return &Claim{
		Issuer:   issuer,
		Audience: audience,
		Subject:  subject,
		Expired:  expired,
		Lock:     new(sync.RWMutex),
	}
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type AccessTokenResponse struct {
	Error       error
	AccessToken *AccessToken
}

func (cl *Claim) GenerateToken(signKey *rsa.PrivateKey) <-chan AccessTokenResponse {
	result := make(chan AccessTokenResponse)
	go func() {
		cl.Lock.Lock()
		defer close(result)
		defer cl.Lock.Unlock()
		token := jwt.New(jwt.SigningMethodRS256)
		claims := make(jwt.MapClaims)
		claims["iss"] = cl.Issuer
		claims["aud"] = cl.Audience
		claims["exp"] = time.Now().Add(cl.Expired).Unix()
		claims["iat"] = time.Now().Unix()
		claims["sub"] = cl.Subject
		token.Claims = claims

		tokenString, err := token.SignedString(signKey)
		if err != nil {
			result <- AccessTokenResponse{Error: err, AccessToken: nil}
			return
		}
		result <- AccessTokenResponse{Error: nil, AccessToken: &AccessToken{fmt.Sprintf("Bearer %v", tokenString)}}
	}()
	return result
}
