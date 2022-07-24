package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-wiki/global"
	"micro-wiki/model/common/response"
	"micro-wiki/model/system"
)

type UserApi struct {
}

func (s *UserApi) Register(c *gin.Context) {
	// 插入记录到wiki_user表中
	var wikiUser *system.WikiUser = new(system.WikiUser)
	wikiUser.RegisterFormToWikiUser(c)

	db := global.MW_DB
	// 查询是否有重复用户
	result := db.Table("wiki_user").Take(&system.WikiUser{Username: wikiUser.Username})
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("用户注册失败", c)
	}
	if result.RowsAffected > 0 {
		log.Printf("用户:%s, 已存在", wikiUser.Username)
		response.FailWithMessage("该用户已存在", c)
		// 返回不再继续, 否则会继续执行插入
		return
	}

	// 插入数据库
	db.Table("wiki_user").Create(&wikiUser)

	response.OkWithMessage("Register Success", c)
	return
}

func (s *UserApi) QueryUserDetail(c *gin.Context) {
	userId := c.Query("user_id")

	// 确认是否有权限查询

	// 查询用户信息
	db := global.MW_DB
	var wikiUser system.WikiUser
	result := db.Table("wiki_user").Find(&wikiUser, &system.WikiUser{UserID: userId})
	// 判断是否报错
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("用户查询失败", c)
	}
	// 判断是否查到用户
	if result.RowsAffected == 0 {
		log.Printf("UserID:%s, 不存在", userId)
		response.FailWithMessage("该用户不存在", c)
		return
	}

	response.OkWithMessage("Register Success", c)
	return
}
