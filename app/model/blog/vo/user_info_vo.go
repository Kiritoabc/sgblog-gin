package vo

type UserInfoVo struct {
	// 主键
	ID int64 `json:"id"`
	// 昵称
	NickName string `json:"nick_name"`
	// 头像
	Avatar string `json:"avatar"`
	// 性别
	Sex string `json:"sex"`
	// 邮箱
	Email string `json:"email"`
}
