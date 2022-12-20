package jwt

import (
	"cloud-api-go/pkg/app"
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	// 秘钥，用于加密 JWT，使用 app.key 读取配置信息
	SignKey []byte
	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

type CustomClaims struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	ExpireAtTime int64  `json:"expire_at_time"`
	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtPkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件使用
func (jwt *JWT) ParserToken(c *gin.Context) (*CustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFormHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	token, err := jwt.parseTokenString(tokenString)
	// 解析出错时
	if err != nil {
		// 类型断言：err 是否为 jwt 的 ValidationError 类型
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			// 如果从请求头中的 Authorization 中获取不到 Token
			if validationErr.Errors == jwtPkg.ValidationErrorMalformed {
				return nil, ErrHeaderMalformed
				// 如果 Token 已过期
			} else if validationErr.Errors == jwtPkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	// 将 Token 中的载荷信息解析出来和 CustomClaims 进行比对
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新 Token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 从请求头中获取 Token
	tokenString, parseErr := jwt.getTokenFormHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 解析 Token
	token, err := jwt.parseTokenString(tokenString)
	// 解析出错时，如果未报错，证明是合法的 Token，（包括未到期）
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 满足刷新条件时：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 解析 CustomClaims 的数据
	claims := token.Claims.(*CustomClaims)
	// 检查是否过了最大允许刷新的时间
	x := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成 Token，此方法供外部使用
func (jwt *JWT) IssueToken(userID string, userName string) string {
	expireAtTime := jwt.expireAtTime()
	// 构造用户 claims 信息(负荷)
	claims := CustomClaims{
		UserID:       userID,
		Username:     userName,
		ExpireAtTime: expireAtTime,
		StandardClaims: jwtPkg.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间，刷新 Token 时不会更新
			ExpiresAt: expireAtTime,                   // 前面过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
		},
	}
	// 根据 claims 生成 token 对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// createToken 创建 Token，此方法内部使用
func (jwt *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// parseTokenString 使用 jwtPkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	timeNow := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute

	return timeNow.Add(expire).Unix()
}

// getTokenFormHeader 从请求头中取出 Token 格式为：Authorization:Bearer token
func (jwt *JWT) getTokenFormHeader(c *gin.Context) (string, error) {
	// 从请求头中获取 Authorization
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		return "", ErrHeaderEmpty
	}
	// 请求头中的 Token 格式为： Bearer xxx 所以需要按空格分割
	// 分割后结果为：["Bearer", "token"]，要取数组中的第二个元素
	parts := strings.SplitN(header, "", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}
