package vo

import "time"

type UserVo struct {
	// 主键
	Id uint64 `json:"id"`
	// 用户名
	UserName string `json:"userName"`
	// 昵称
	NickName string `json:"nickName"`
	// 账号状态（0正常 1停用）
	Status string `json:"status"`
	// 邮箱
	Email string `json:"email"`
	// 手机号
	Phonenumber string `json:"phonenumber"`
	// 用户性别（0男，1女，2未知）
	Sex string `json:"sex"`
	// 头像
	Avatar string `json:"avatar"`

	// 创建人的用户id
	CreateBy uint64 `json:"createBy"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	UpdateBy uint64 `json:"updateBy"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}
