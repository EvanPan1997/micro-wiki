package system

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/model/common/response"
)

type UserApi struct {
}

func (s *UserApi) Register(c *gin.Context) {
	response.OkWithMessage("Register Success", c)
}
