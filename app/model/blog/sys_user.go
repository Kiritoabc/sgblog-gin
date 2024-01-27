package blog

import "time"

type SysUser struct {
	Id          int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	UserName    string    `gorm:"column:user_name;type:VARCHAR(64);NOT NULL" json:"userName"`
	NickName    string    `gorm:"column:nick_name;type:VARCHAR(64);NOT NULL" json:"nickName"`
	Password    string    `gorm:"column:password;type:VARCHAR(64);NOT NULL" json:"password"`
	Type        string    `gorm:"column:type;type:CHAR(1);" json:"type"`
	Status      string    `gorm:"column:status;type:CHAR(1);" json:"status"`
	Email       string    `gorm:"column:email;type:VARCHAR(64);" json:"email"`
	Phonenumber string    `gorm:"column:phonenumber;type:VARCHAR(32);" json:"phonenumber"`
	Sex         string    `gorm:"column:sex;type:CHAR(1);" json:"sex"`
	Avatar      string    `gorm:"column:avatar;type:VARCHAR(128);" json:"avatar"`
	CreateBy    int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;" json:"createTime"`
	UpdateBy    int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;" json:"updateTime"`
	DelFlag     int32     `gorm:"column:del_flag;type:INT(11);" json:"delFlag"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
