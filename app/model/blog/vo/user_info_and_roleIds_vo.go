package vo

import "sgblog-go/app/model/blog"

type UserInfoAndRoleIdsVo struct {
	User    *blog.SysUser   `json:"user"`
	Roles   []*blog.SysRole `json:"roles"`
	RoleIds []int64         `json:"roleIds"`
}
