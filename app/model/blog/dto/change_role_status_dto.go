package dto

type ChangeRoleStatusDto struct {
	RoleId int64  `json:"roleId"`
	Status string `json:"status"`
}
