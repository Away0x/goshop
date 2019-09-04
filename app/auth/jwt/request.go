package jwt

import (
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetToken 获取 token
func GetToken(c echo.Context) (string, *errno.Errno) {
	if token, ok := getTokenFromHeader(c); ok {
		c.Set(tokenHeaderKeyName+"Token", token)
		return token, nil
	}
	if token, ok := getTokenFromParams(c); ok {
		c.Set(tokenHeaderKeyName+"Token", token)
		return token, nil
	}

	return "", errno.JWTTokenMissingErr
}

// GetUser 根据 token 获取用户数据 (并且将用户数据和 token 存储到 context 中)
func GetUser(c echo.Context, token string) (*models.User, *errno.Errno) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	user, e := models.GetUser(claims.UserID)
	if e != nil {
		return nil, errno.DatabaseErr.SetErrorContent(e)
	}

	c.Set(tokenHeaderKeyName+"User", user)
	c.Set(tokenHeaderKeyName+"Token", token)
	return user, nil
}

// GetTokenAndUser 获取 token 和 user (并且将用户数据和 token 存储到 context 中)
func GetTokenAndUser(c echo.Context) (tokenStr string, user *models.User, err *errno.Errno) {
	// 如果 context 中存在
	if tokenStr, user, ok := GetTokenAndUserFromContext(c); ok {
		return tokenStr, user, nil
	}

	tokenStr, err = GetToken(c)
	if err != nil {
		return "", nil, err
	}

	user, err = GetUser(c, tokenStr)
	if err != nil {
		return "", nil, err
	}

	c.Set(tokenHeaderKeyName+"User", user)
	c.Set(tokenHeaderKeyName+"Token", tokenStr)
	return tokenStr, user, nil
}

// GetTokenAndUserFromContext 从 context 中获取 user 和 token
func GetTokenAndUserFromContext(c echo.Context) (string, *models.User, bool) {
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
	if !ok || u == nil || s == "" {
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
	fmt.Sscanf(header, tokenInHeaderIdentification+" %s", &token)
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
