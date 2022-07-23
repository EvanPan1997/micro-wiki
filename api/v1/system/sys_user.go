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
		response.Fail(c)
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
}
