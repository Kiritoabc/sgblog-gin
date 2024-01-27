package model

import "time"

type SysUser struct {
	Id          int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	UserName    string    `gorm:"column:user_name;type:VARCHAR(64);NOT NULL"`
	NickName    string    `gorm:"column:nick_name;type:VARCHAR(64);NOT NULL"`
	Password    string    `gorm:"column:password;type:VARCHAR(64);NOT NULL"`
	Type        string    `gorm:"column:type;type:CHAR(1);"`
	Status      string    `gorm:"column:status;type:CHAR(1);"`
	Email       string    `gorm:"column:email;type:VARCHAR(64);"`
	Phonenumber string    `gorm:"column:phonenumber;type:VARCHAR(32);"`
	Sex         string    `gorm:"column:sex;type:CHAR(1);"`
	Avatar      string    `gorm:"column:avatar;type:VARCHAR(128);"`
	CreateBy    int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy    int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag     int32     `gorm:"column:del_flag;type:INT(11);"`
}
