package auth

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
	"fmt"

	"github.com/labstack/echo/v4"
)

const (
	tokenParamsKeyName = "token"
	tokenHeaderKeyName = "Authorization"
)

// GetTokenFromRequest 从请求中获取 token
func GetTokenFromRequest(c echo.Context) (string, *errno.Errno) {
	if token, ok := getTokenFromHeader(c); ok {
		return token, nil
	}
	if token, ok := getTokenFromParams(c); ok {
		return token, nil
	}

	return "", errno.JWTTokenMissingErr
}

// ParseAndGetUser 解析 token 获取 user
func ParseAndGetUser(c echo.Context, token string) (*models.User, *errno.Errno) {
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	if err := models.DB().Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		return nil, errno.DatabaseErr.SetMessage(err.Error())
	}

	c.Set(tokenHeaderKeyName+"User", user)
	c.Set(tokenHeaderKeyName+"Token", token)
	return user, nil
}

// GetTokenUserFromContext -
func GetTokenUserFromContext(c echo.Context) (string, *models.User, bool) {
	user := c.Get(tokenHeaderKeyName + "User")
	if user == nil {
		return "", nil, false
	}
	t := c.Get(tokenHeaderKeyName + "Token")
	if t == nil {
		return "", nil, false
	}

	u, ok := user.(*models.User)
	s, ok := t.(string)
	if !ok {
		return "", nil, false
	}

	return s, u, true
}

// ---------------- private
func getTokenFromHeader(c echo.Context) (string, bool) {
	header := c.Request().Header.Get(tokenHeaderKeyName)
	if header == "" {
		return "", false
	}

	var token string
	fmt.Sscanf(header, jwt.TokenInHeaderIdentification+" %s", &token)
	if token == "" {
		return "", false
	}
	return token, true
}

func getTokenFromParams(c echo.Context) (string, bool) {
	token := c.QueryParam(tokenParamsKeyName)
	if token != "" {
		return token, true
	}
	token = c.FormValue(tokenParamsKeyName)
	if token != "" {
		return token, true
	}

	return "", false
}
