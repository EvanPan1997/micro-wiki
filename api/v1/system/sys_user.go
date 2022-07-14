package system

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/model/common/response"
)

type UserApi struct {
}

func (s *UserApi) Register(c *gin.Context) {
	// 接收请求体

	// 插入记录到wiki_user表中

	response.OkWithMessage("Register Success", c)
}
