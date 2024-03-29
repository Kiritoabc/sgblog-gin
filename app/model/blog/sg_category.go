package blog

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type SgCategory struct {
	Id          int64                 `gorm:"column:id;type:BIGINT(200);AUTO_INCREMENT;NOT NULL" json:"id"`
	Name        string                `gorm:"column:name;type:VARCHAR(128);" json:"name"`
	Pid         int64                 `gorm:"column:pid;type:BIGINT(200);" json:"pid"`
	Description string                `gorm:"column:description;type:VARCHAR(512);" json:"description"`
	Status      string                `gorm:"column:status;type:CHAR(1);" json:"status"`
	CreateBy    int64                 `gorm:"column:create_by;type:BIGINT(200);" json:"createBy"`
	CreateTime  time.Time             `gorm:"column:create_time;type:DATETIME;default:current_timestamp" json:"createTime"`
	UpdateBy    int64                 `gorm:"column:update_by;type:BIGINT(200);" json:"updateBy"`
	UpdateTime  time.Time             `gorm:"column:update_time;type:DATETIME;default:current_timestamp on update current_timestamp" json:"updateTime"`
	DelFlag     soft_delete.DeletedAt `gorm:"column:del_flag;type:INT(1);softDelete:flag" json:"delFlag"`
}

func (SgCategory) TableName() string {
	return "sg_category"
}
