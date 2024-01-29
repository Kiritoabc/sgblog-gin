package blog

import "time"

type SgLink struct {
	Id          int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	Name        string    `gorm:"column:name;type:VARCHAR(256);" json:"name"`
	Logo        string    `gorm:"column:logo;type:VARCHAR(256);" json:"logo"`
	Description string    `gorm:"column:description;type:VARCHAR(512);" json:"description"`
	Address     string    `gorm:"column:address;type:VARCHAR(128);" json:"address"`
	Status      string    `gorm:"column:status;type:CHAR(1);" json:"status"`
	CreateBy    int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;default:current_timestamp" json:"createTime"`
	UpdateBy    int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;default:current_timestamp on update current_timestamp" json:"updateTime"`
	DelFlag     int32     `gorm:"column:del_flag;type:INT(1);" json:"delFlag"`
}

func (SgLink) TableName() string {
	return "sg_link"
}
