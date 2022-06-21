package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"micro-wiki/entity/user"
	"micro-wiki/util"
	"micro-wiki/util/db"
	"micro-wiki/util/msg"
	"net/http"
	"strings"
	"time"
)

const (
	alg = "HS256"
	typ = "JWT"

	author = "Evan.Pan"
	iss    = "micro-wiki platform"

	//secret = "20220618@Evan.Pan"
)

type TokenHeader struct {
	alg string
	typ string
}

type TokenPayload struct {
	author   string
	iss      string
	exp      int64
	userId   string
	password string
}

// 仅登录时, 下发token使用
func GenerateToken(userId string, password string) (string, error) {
	header := TokenHeader{alg, typ}
	payload := TokenPayload{author: author, iss: iss, exp: time.Now().AddDate(0, 0, 7).Unix(), userId: userId, password: password}
	headerJson, err := json.Marshal(header)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	headerBase64 := util.EncodeToBase64(headerJson)
	payloadBase64 := util.EncodeToBase64(payloadJson)
	/* 后续增加盐值加密 */
	signature := util.EncodeMD5(util.BytesToString(headerJson) + "." + util.BytesToString(payloadJson))
	return headerBase64 + "." + payloadBase64 + "." + signature, nil
}

// Jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		code := msg.SUCCESS
		if token == "" {
			// 若为空, 则203重定向
			code = msg.REDIRECT
			c.Abort()
		}
		/*
			解析token:
			1. 解析
			2. 是否超时
			3. 是否密码未被修改
		*/
		// 解析
		tokens := strings.Split(token, ".")
		if len(tokens) != 3 {
			code = msg.ERROR
			// token错误
			c.Abort()
		}
		header := tokens[0]
		payload := tokens[1]
		signature := tokens[2]
		// 检查是否被篡改
		if !(util.EncodeMD5(header)+"."+util.EncodeMD5(payload) == signature) {
			code = msg.ERROR
			// token被篡改
			c.Abort()
		}
		// 是否超时
		var tokenPayload TokenPayload
		err := json.Unmarshal(util.StringToBytes(payload), &tokenPayload)
		if err != nil {
			code = msg.ERROR
			// json解码错误
			c.Abort()
		}
		exp := tokenPayload.exp
		if exp < time.Now().Unix() {
			code = msg.REDIRECT
			// 超时重定向回首页或登录页
			c.Abort()
		}
		/*
			是否密码未被修改: select password from wiki_user where data_id = :"userId" and status = '1' and is_lock = '0'
			若存在且一致, 则通过
			若存在但不一致, 密码被修改, 重定向
			若不存在, 则账号被注销, 重定向
		*/
		userId := tokenPayload.userId
		dbConn, err := db.GetDbConnection()
		if err != nil {
			log.Fatalln(err)
		}
		var wikiUser user.WikiUser
		wikiUser.UserId = userId
		dbConn.Table("wiki_user").Select("password").Find(&wikiUser)
		if tokenPayload.password != wikiUser.Password {
			code = msg.REDIRECT
			//密码已更改
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
		c.Next()
	}
}
