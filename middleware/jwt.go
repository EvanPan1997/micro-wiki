package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"micro-wiki/util"
	"micro-wiki/util/msg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(util.JwtSecret)
var code int

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	claim    jwt.StandardClaims
}

func (c Claims) Valid() error {
	panic("implement me")
}

// 生成token
func GenerateToken(username string, password string) (string, int) {
	expireTime := time.Now().AddDate(0, 0, 7)
	SetClaims := Claims{
		Username: username,
		Password: password,
		claim: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "micro-wiki",
			IssuedAt:  time.Now().Unix(),
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", msg.ERR_MSG
	}
	return token, msg.SUCCESS
}

// 认证token
func checkToken(token string) (*Claims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if key, code := setToken.Claims.(*Claims); code && setToken.Valid {
		return key, msg.SUCCESS
	} else {
		return nil, msg.ERROR
	}
}

// Jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code = msg.SUCCESS
		if tokenHeader == "" {
			// 若为空, 则203重定向
			code = msg.REDIRECT
			c.Abort()
		}
		// 解析tokenHeader
		checkTokens := strings.SplitN(tokenHeader, " ", 2)
		if len(checkTokens) != 2 && checkTokens[0] != "Bearer" {
			code = msg.ERROR_TOKEN_TYPE_WRONG
			c.Abort()
		}
		key, checkCode := checkToken(checkTokens[1])
		if checkCode == msg.ERROR {
			code = msg.ERROR_TOKEN_WRONG
			c.Abort()
		}
		if time.Now().Unix() > key.claim.ExpiresAt {
			code = msg.ERROR_TOKEN_OVERDUE
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
