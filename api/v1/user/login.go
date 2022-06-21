package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-wiki/entity/user"
	"micro-wiki/util/db"
	"time"
)

/*
登录接口服务
*/
func LoginService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nowStr := time.Now().Format("20060102150405")
		// 验证用户名和密码
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		var wikiUser user.WikiUser
		wikiUser.UserName = username

		// 验证登录
		dbConn, err := db.GetDbConnection()
		if err != nil {
			log.Println(err)
		}
		dbConn.Table("wiki_user").Select("password").Find(&wikiUser)
		// 检查状态和是否锁定
		if wikiUser.Status == "0" || wikiUser.IsLock == "1" {
			wikiUser.LastFailTime = nowStr

			dbConn.Table("wiki_user").Save(&wikiUser)
			// 状态禁用或锁定, 日志打印
			log.Println("该用户状态不可用")
		}

		// 密码是否正确
		if wikiUser.Password != password {
			/* 连续失败次数+1 失败总次数+1 最近失败时间 */
			wikiUser.PswdErrCnt++
			wikiUser.TotPswdErr++
			wikiUser.LastFailTime = nowStr
			wikiUser.LastFailTime = nowStr

			dbConn.Table("wiki_user").Save(&wikiUser)
			// 密码错误, 日志打印
			log.Println("密码错误")
		}
		// 登录成功更新表数据
		/* 时间类 登录IP 连续失败次数清零 */
		wikiUser.LoginIp = ctx.ClientIP()
		wikiUser.LastAccessTime = nowStr
		dbConn.Table("wiki_user").Save(&wikiUser)
	}
}
