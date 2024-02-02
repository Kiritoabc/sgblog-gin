package vo

type AdminUserInfoVo struct {
	Permissions []string   `json:"permissions"`
	Roles       []string   `json:"roles"`
	UserInfoVo  UserInfoVo `json:"userInfoVo"`
}
