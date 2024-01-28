package blog

type UserLogin struct {
	User        *SysUser `json:"user"`
	Permissions []string `json:"permissions"`
}
