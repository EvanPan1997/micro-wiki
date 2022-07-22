package system

import (
	"github.com/gin-gonic/gin"
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
	db.Table("wiki_user").Create(&wikiUser)

	response.OkWithMessage("Register Success", c)
}
