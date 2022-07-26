package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-wiki/global"
	"micro-wiki/model/common/response"
	"micro-wiki/model/system"
	"micro-wiki/utils"
	"net/http"
)

type UserApi struct {
}

// 注册用户
func (api *UserApi) Register(c *gin.Context) {
	// bind json
	var registerReq system.RegisterReq
	if err := c.BindJSON(&registerReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// 创建结构体实例, 并加入初始数据
	var wikiUser *system.WikiUser = new(system.WikiUser)
	wikiUser.RegisterFormToWikiUser(&registerReq)
	// 插入记录到wiki_user表中
	db := global.MW_DB
	// 查询是否有重复用户
	result := db.Table("wiki_user").Take(&system.WikiUser{Username: wikiUser.Username})
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("数据库查询错误", c)
		return
	}
	if result.RowsAffected > 0 {
		log.Printf("用户:%s, 已存在", wikiUser.Username)
		response.FailWithMessage("该用户已存在", c)
		// 返回不再继续, 否则会继续执行插入
		return
	}

	// 查询操作员, 确认是否有权限进行操作
	if registerReq.Operator != "system" {
		var roleName string
		db.Raw("SELECT r.role_name FROM wiki_role r,wiki_user u WHERE u.user_id = ? AND u.role_id = r.role_id", registerReq.Operator).Scan(&roleName)
		if roleName == "Admin" {
			response.FailWithMessage("无权限进行该操作", c)
			return
		}
	}

	// 插入数据库
	db.Table("wiki_user").Create(&wikiUser)

	response.OkWithMessage("Register Success", c)
}

// 查询用户详细信息
func (api *UserApi) QueryDetail(c *gin.Context) {
	userId := c.Query("user_id")

	// 查询用户信息
	db := global.MW_DB
	var wikiUser system.WikiUser
	result := db.Table("wiki_user").Find(&wikiUser, &system.WikiUser{UserID: userId})
	// 判断是否报错
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("用户查询失败", c)
		return
	}
	// 判断是否查到用户
	if result.RowsAffected == 0 {
		log.Printf("UserID:%s, 不存在", userId)
		response.FailWithMessage("该用户不存在", c)
		return
	}

	response.OkWithMessage("Query Success", c)
}

// 用户修改密码
func (api *UserApi) ChangePassword(c *gin.Context) {
	// bind json
	var changePasswordReq system.ChangePasswordReq
	if err := c.BindJSON(&changePasswordReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if changePasswordReq.UserID != changePasswordReq.Operator {
		response.FailWithMessage("不能修改他人的密码", c)
		return
	}

	var wikiUser system.WikiUser
	db := global.MW_DB
	result := db.Table("wiki_user").Find(&wikiUser, changePasswordReq.UserID)
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("数据库查询错误", c)
	}

	if result.RowsAffected == 0 {
		response.FailWithMessage("该用户不存在", c)
		return
	} else if result.RowsAffected == 1 {
		wikiUser.Password = utils.BcryptHash(changePasswordReq.Password)
		db.Table("wiki_user").Save(&wikiUser)
		response.OkWithMessage("密码修改成功", c)
	} else {
		log.Printf("[平台异常]-该用户: %s, 查询结果不为0也不为1")
		response.FailWithMessage("数据异常", c)
	}
}

// 管理员重置用户密码
func (api *UserApi) ResetPassword(c *gin.Context) {
	// bind json
	var resetPasswordReq system.ResetPasswordReq
	if err := c.BindJSON(&resetPasswordReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	// 确认操作者权限权限
	if resetPasswordReq.Operator != "admin" {
		response.FailWithMessage("权限不足", c)
	}

	var wikiUser system.WikiUser
	db := global.MW_DB
	result := db.Table("wiki_user").Find(&wikiUser, resetPasswordReq.UserID)
	if result.Error != nil {
		log.Println(result.Error)
		response.FailWithMessage("数据库查询错误", c)
		return
	}

	// ResetPassword
	if result.RowsAffected == 0 {
		response.FailWithMessage("该用户不存在", c)
	} else if result.RowsAffected == 1 {
		wikiUser.ResetPassword()
		db.Table("wiki_user").Save(&wikiUser)
		response.OkWithMessage("密码重置成功", c)
	} else {
		log.Printf("[平台异常]-该用户: %s, 查询结果不为0也不为1")
		response.FailWithMessage("数据异常", c)
	}
}
