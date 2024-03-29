package system

import (
	"micro-wiki/utils"
	"time"
)

var (
	InitPassword = "123456"
)

type WikiUser struct {
	// 数据ID
	DataID string `gorm:"primaryKey" column:"data_id"`
	// 用户信息
	UserID            string `column:"user_id"`
	Username          string `column:"username"`
	Password          string `column:"password"`
	RoleID            string `column:"role_id"`
	Status            string `column:"status"`
	IsLocked          string `column:"is_locked"`
	LastLoginTime     string `column:"last_login_time"`
	LastLogoutTime    string `column:"last_logout_time"`
	LoginIP           string `column:"login_ip"`
	PwdErrorCount     int    `column:"pwd_error_count"`
	TotalPwdError     int    `column:"total_pwd_error"`
	LastChangePwdTime string `column:"last_change_pwd_time"`
	AccessToken       string `column:"access_token"`
	// 系统维护信息
	CreateUser string `column:"create_user"`
	CreateAt   string `column:"create_at"`
	ChangeUser string `column:"change_user"`
	ChangeAt   string `column:"change_at"`
	DeleteUser string `column:"delete_user"`
	DeleteAt   string `column:"delete_at"`
	// 补充字段
	remarks  string `column:"remarks"`
	Reserve1 string `column:"reserve1"`
	Reserve2 string `column:"reserve2"`
	Reserve3 string `column:"reserve3"`
	Reserve4 string `column:"reserve4"`
	Reserve5 string `column:"reserve5"`
}

// 注册表单结构体
type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Operator string `form:"operator" json:"operator" binding:"required"`
}

/* 将用户注册表单, 转换为WikiUser实例 */
func (wikiUser *WikiUser) RegisterFormToWikiUser(registerReq *RegisterReq) {
	// 提供hash值作为本结构体的id
	ids := utils.GenUuidV4()
	nowStr := utils.TimeToString(time.Now())

	// 基本信息
	wikiUser.DataID = ids
	wikiUser.UserID = ids
	wikiUser.Username = registerReq.Username
	wikiUser.Password = utils.BcryptHash(InitPassword)
	wikiUser.RoleID = "User"
	wikiUser.Status = "1"
	wikiUser.IsLocked = "N"
	wikiUser.PwdErrorCount = 0
	wikiUser.TotalPwdError = 0

	// 创建信息
	wikiUser.CreateUser = registerReq.Operator
	wikiUser.CreateAt = nowStr
	// 修改信息
	wikiUser.ChangeUser = registerReq.Operator
	wikiUser.ChangeAt = nowStr
}

// 修改密码表单结构体
type ChangePasswordReq struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Operator string `form:"operator" json:"operator" binding:"required"`
}

// 重置密码表单结构体
type ResetPasswordReq struct {
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Operator string `form:"operator" json:"operator" binding:"required"`
}

// 重置密码函数
func (wikiUser *WikiUser) ResetPassword() {
	wikiUser.Password = utils.BcryptHash(InitPassword)
}
