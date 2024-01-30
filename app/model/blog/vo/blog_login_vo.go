package vo

type BlogUserLoginVo struct {
	// 登录令牌
	Token string `json:"token"`
	// 用户信息
	UserInfo *UserInfoVo `json:"userInfo"`
}
