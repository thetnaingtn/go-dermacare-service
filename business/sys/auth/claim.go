package auth

import "github.com/golang-jwt/jwt/v4"

type Claim struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func (c *Claim) Authorize(roles ...string) bool {
	for _, has := range c.Roles {
		for _, want := range roles {
			if has == want {
				return true
			}
		}
	}

	return false
}
