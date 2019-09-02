package jwt

import (
	"echo_shop/app/cache"
	"echo_shop/app/models"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/errno"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	cacheTokenKeyName   = "jwt_token_"
	jwtTokenRefreshTime = 20160 * time.Minute // 允许刷新 token 的时间 (14 天); 期间内允许使用之前颁发的 token (即使过期)来获取新token
	jwtTokenExpiredTime = 60 * time.Minute    // token 60 分钟过期
	jwtTokenRemainTime  = 2 * time.Minute     // token 刷新后，旧的 token 还有 2 分钟的使用时间 (前提是旧 token 没过过期时间)
)

func getCacheKey(tokenString string) string {
	return cacheTokenKeyName + tokenString
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*CustomClaims, *errno.Errno) {
	token, err := parse(tokenString)
	if err != nil {
		// token 过期
		if isExpired(err) {
			if claims, ok := token.Claims.(*CustomClaims); ok {
				return claims, errno.JWTTokenExpireErr
			}
		}
		if e, ok := err.(*errno.Errno); ok {
			return nil, e
		}
		return nil, errno.JWTTokenErr.SetMessage(err.Error())
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errno.JWTTokenErr.SetMessage("jwt token parse error")
}

// 创建 token
func create(u *models.User) (string, CustomClaims, error) {
	claims := CustomClaims{}
	claims.SetUser(u)
	claims.SetExpiredAt()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(config.String("APP.KEY")))
	if err != nil {
		return "", claims, err
	}
	return s, claims, nil
}

// 刷新 token
func refresh(tokenString string) (string, CustomClaims, *errno.Errno) {
	token, err := parse(tokenString)
	if err != nil {
		// 非过期错误
		if !isExpired(err) {
			return "", CustomClaims{}, errno.JWTTokenErr.SetMessage(err.Error())
		}
		// 判断是否过了可刷新时间
		if claims, ok := token.Claims.(*CustomClaims); ok {
			now := time.Now().Unix()
			if now > claims.RefreshTime {
				return "", CustomClaims{}, errno.JWTTokenExpireErr
			}
		}
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", CustomClaims{}, errno.JWTTokenErr
	}

	forget(tokenString, jwtTokenRemainTime) // 将之前的 token 加入黑名单使之失效
	u := &models.User{}
	u.ID = claims.UserID
	newToken, newClaims, err := create(u)
	if err != nil {
		return "", CustomClaims{}, errno.JWTTokenErr.SetMessage(err.Error())
	}

	return newToken, newClaims, nil
}

// 使 token 失效
func forget(tokenString string, remainTime time.Duration) {
	now := time.Now()
	cache.PutInt64(getCacheKey(tokenString), now.Add(remainTime).Unix(), jwtTokenExpiredTime) // val 为 token 还可以使用的过渡时间
}

func parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.String("APP.KEY")), nil
	})

	// 在黑名单中
	if t, ok := cache.GetInt64(getCacheKey(tokenString)); ok {
		now := time.Now().Unix()
		// 过了留存时间了
		if now > t {
			return nil, errno.JWTTokenInBlackListErr
		}
	}

	return token, err
}

// 判断是否过期
func isExpired(err error) bool {
	switch err.(type) {
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			// token 过期了
			return true
		default:
			return false
		}

	default:
		return false
	}
}
