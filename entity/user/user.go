package user

type WikiUser struct {
	DataId       string `gorm:"column:data_id"`
	DataDate     string `gorm:"column:data_date"`
	DataDateTime string `gorm:"column:data_date_time"`

	CrtUser string `gorm:"column:crt_user"`
	CrtDate string `gorm:"column:crt_date"`
	CrtTime string `gorm:"column:crt_time"`
	ChgUser string `gorm:"column:chg_user"`
	ChgDate string `gorm:"column:chg_date"`
	ChgTime string `gorm:"column:chg_time"`
	ApvUser string `gorm:"column:apv_user"`
	ApvDate string `gorm:"column:apv_date"`
	ApvTime string `gorm:"column:apv_time"`
	Remarks string `gorm:"column:remarks"`

	UserId          string `gorm:"column:user_id"`
	UserName        string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	RoleId          string `gorm:"column:role_id"`
	LoginIp         string `gorm:"column:login_ip"`
	Status          string `gorm:"column:status"`
	IsLock          string `gorm:"column:is_lock"`
	LastAccessTime  string `gorm:"column:last_access_time"`
	LastLogoutTime  string `gorm:"column:last_logout_time"`
	LastFailTime    string `gorm:"column:last_fail_time"`
	PswdErrCnt      int8   `gorm:"column:pswd_err_cnt"`
	TotPswdErr      int8   `gorm:"column:tot_pswd_err"`
	LastChgPswdTime string `gorm:"column:last_chg_pswd_time"`

	Rsv1          string `gorm:"column:rsv1"`
	Rsv2          string `gorm:"column:rsv2"`
	Rsv3          string `gorm:"column:rsv3"`
	Rsv4          string `gorm:"column:rsv4"`
	Rsv5          string `gorm:"column:rsv5"`
	AccessTokenId string `gorm:"column:access_token_id"`
}
