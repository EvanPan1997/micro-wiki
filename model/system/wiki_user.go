package system

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/utils"
	"time"
)

var (
	InitPassword = "123456"
)

type WikiUser struct {
	// 数据ID
	DataID string
	// 用户信息
	UserID            string
	Username          string
	Password          string
	RoleID            string
	Status            string
	IsLocked          string
	LastLoginTime     time.Time
	LastLogoutTime    time.Time
	LoginIP           string
	PwdErrorCount     int
	TotalPwdErr       int
	LastChangePwdTime time.Time
	AccessToken       string
	// 系统维护信息
	CreateUser string
	CreateAt   time.Time
	ChangeUser string
	ChangeAt   time.Time
	DeleteUser string
	DeleteAt   time.Time
	// 补充字段
	remarks  string
	Reserve1 string
	Reserve2 string
	Reserve3 string
	Reserve4 string
	Reserve5 string
}

/* 将用户注册表单, 转换为WikiUser实例 */
//func (wikiUser *WikiUser) RegisterFormToWikiUser(form map[string]string) {
func (wikiUser *WikiUser) RegisterFormToWikiUser(c *gin.Context) {

	ids := utils.GenUuidV4()
	username := c.PostForm("username")
	operator := c.PostForm("operator")

	wikiUser.DataID = ids
	wikiUser.UserID = ids
	//wikiUser.Username 		= form["username"]
	wikiUser.Username = username

	wikiUser.Password = utils.BcryptHash(InitPassword)
	wikiUser.RoleID = "User"
	wikiUser.Status = "1"
	wikiUser.IsLocked = "N"
	wikiUser.PwdErrorCount = 0
	wikiUser.TotalPwdErr = 0

	//wikiUser.CreateUser 	= form["operator"]
	wikiUser.CreateUser = operator
	wikiUser.CreateAt = time.Now()
	//wikiUser.ChangeUser 	= form["operator"]
	wikiUser.ChangeUser = operator
	wikiUser.ChangeAt = time.Now()
}
