package blog

import "time"

type SgTag struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	Name       string    `gorm:"column:name;type:VARCHAR(128);" json:"name"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;" json:"createTime"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;" json:"updateTime"`
	DelFlag    int32     `gorm:"column:del_flag;type:INT(1);" json:"delFlag"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);" json:"remark"`
}

func (SgTag) TableName() string {
	return "sg_tag"
}
