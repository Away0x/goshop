package jwt

import (
	"echo_shop/app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims -
type CustomClaims struct {
	jwt.StandardClaims
	UserID      uint  `json:"userid"`
	RefreshTime int64 `json:"refresh_time,omitempty"`
}

// SetUser 设置 token 有效期
func (c *CustomClaims) SetUser(u *models.User) {
	c.UserID = u.ID
}

// SetExpiredAt 设置 user data
func (c *CustomClaims) SetExpiredAt() {
	now := time.Now()
	c.IssuedAt = now.Unix()
	c.ExpiresAt = now.Add(jwtTokenExpiredTime).Unix()
	c.RefreshTime = now.Add(jwtTokenRefreshTime).Unix()
}
