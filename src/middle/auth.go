package middle

import (
	"down-date-server/src/global"
	"down-date-server/src/model"
	"down-date-server/src/model/request"
	"down-date-server/src/utils/response"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SignKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage(401, "未授权的访问", c)
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseJWT(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithMessage(401, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithMessage(401, "认证失败", c)
			c.Abort()
			return
		}
		// if err, _ = service.FindUserByUuid(claims.UUID.String()); err != nil {
		// 	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		// 	c.Abort()
		// }
		var u model.User
		if err = global.DB.Where("`uuid` = ?", claims.UUID.String()).First(&u).Error; err != nil {
			response.FailWithMessage(401, "认证失败", c)
			c.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateJWT(*claims)
			newClaims, _ := j.ParseJWT(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			// if global.GVA_CONFIG.System.UseMultipoint {
			// 	err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
			// 	if err != nil {
			// 		global.GVA_LOG.Error("get redis jwt failed", zap.Any("err", err))
			// 	} else { // 当之前的取成功时才进行拉黑操作
			// 		_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
			// 	}
			// 	// 无论如何都要记录当前的活跃状态
			// 	_ = service.SetRedisJWT(newToken, newClaims.Username)
			// }
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SignKey),
	}
}

// 创建一个Token
func (j *JWT) CreateJWT(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

// 解析Token
func (j *JWT) ParseJWT(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
