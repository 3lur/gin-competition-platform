package jwt

import (
	"competition-backend/pkg/app"
	"competition-backend/pkg/config"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

type CustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`
	jwtPkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// ParseToken 解析 Token，中间件使用
func (j *JWT) ParseToken(c *gin.Context) (*CustomClaims, error) {
	// 1. 从请求头获取 Token 字符
	tokenString, parseErr := j.getTokenFromRequestHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	// 2. 解析 Token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtPkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtPkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	// 3. 将 token 中的 claims 信息解析出来和 CustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新 Token
func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 1. 从请求头获取 Token
	tokenString, parseErr := j.getTokenFromRequestHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 解析 Token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		vErr, ok := err.(*jwtPkg.ValidationError)
		if !ok || vErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", nil
		}
	}

	// 3. 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)
	x := app.TimeNowInTimezone().Add(-j.MaxRefresh).Unix()

	// 5. 检查是否过了最大允许刷新的时间
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 颁发 Token 在登录成功后调用
func (j *JWT) IssueToken(userID, userName string) string {
	eat := j.expireAtTime()
	// 1. 构造 claims 信息
	c := CustomClaims{
		UserID:       userID,
		UserName:     userName,
		ExpireAtTime: eat,
		StandardClaims: jwtPkg.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间，refresh token 不会更新
			ExpiresAt: eat,                            // 签名过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
		},
	}

	// 2. 根据 claims 信息生成 token
	t, err := j.createToken(c)
	if err != nil {
		return ""
	}
	return t
}

// createToken 创建 Token
func (j *JWT) createToken(c CustomClaims) (string, error) {
	t := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, c)
	return t.SignedString(j.SignKey)
}

// parseTokenString 解析从请求头获取到的 Token 字符串
func (j *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

// getTokenFromRequestHeader 从请求头获取到 Token
func (j *JWT) getTokenFromRequestHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割，格式为：Bearer Token
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

func (j *JWT) expireAtTime() int64 {
	tn := app.TimeNowInTimezone()
	var eTime int64
	if config.GetBool("app.debug") {
		eTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		eTime = config.GetInt64("jwt.expire_time")
	}
	e := time.Duration(eTime) * time.Minute
	return tn.Add(e).Unix()
}
