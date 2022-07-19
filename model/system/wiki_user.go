package system

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/utils"
	"net/http"
	"time"
)

var (
	InitPassword = "123456"
)

type WikiUser struct {
	// 数据ID
	DataID string `gorm:"primaryKey" column:"data_id"`
	// 用户信息
	UserID            string    `column:"user_id"`
	Username          string    `column:"username"`
	Password          string    `column:"password"`
	RoleID            string    `column:"role_id"`
	Status            string    `column:"status"`
	IsLocked          string    `column:"is_locked"`
	LastLoginTime     time.Time `column:"last_login_time"`
	LastLogoutTime    time.Time `column:"last_logout_time"`
	LoginIP           string    `column:"login_ip"`
	PwdErrorCount     int       `column:"pwd_error_count"`
	TotalPwdErr       int       `column:"total_pwd_err"`
	LastChangePwdTime time.Time `column:"last_change_pwd_time"`
	AccessToken       string    `column:"access_token"`
	// 系统维护信息
	CreateUser string    `column:"create_user"`
	CreateAt   time.Time `column:"create_at"`
	ChangeUser string    `column:"change_user"`
	ChangeAt   time.Time `column:"change_at"`
	DeleteUser string    `column:"delete_user"`
	DeleteAt   time.Time `column:"delete_at"`
	// 补充字段
	remarks  string `column:"remarks"`
	Reserve1 string `column:"reserve1"`
	Reserve2 string `column:"reserve2"`
	Reserve3 string `column:"reserve3"`
	Reserve4 string `column:"reserve4"`
	Reserve5 string `column:"reserve5"`
}

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Operator string `form:"operator" json:"operator" binding:"required"`
}

/* 将用户注册表单, 转换为WikiUser实例 */
func (wikiUser *WikiUser) RegisterFormToWikiUser(c *gin.Context) {
	// 提供hash值作为本结构体的id
	ids := utils.GenUuidV4()

	// bind json
	var registerReq RegisterReq
	if err := c.BindJSON(&registerReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
	}

	// 基本信息
	wikiUser.DataID = ids
	wikiUser.UserID = ids
	wikiUser.Username = registerReq.Username
	wikiUser.Password = utils.BcryptHash(InitPassword)
	wikiUser.RoleID = "User"
	wikiUser.Status = "1"
	wikiUser.IsLocked = "N"
	wikiUser.PwdErrorCount = 0
	wikiUser.TotalPwdErr = 0

	// 创建信息
	wikiUser.CreateUser = registerReq.Operator
	wikiUser.CreateAt = time.Now()
	// 修改信息
	wikiUser.ChangeUser = registerReq.Operator
	wikiUser.ChangeAt = time.Now()
}
