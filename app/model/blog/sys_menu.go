package blog

import "time"

type SysMenu struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	MenuName   string    `gorm:"column:menu_name;type:VARCHAR(50);NOT NULL" json:"menuName"`
	ParentId   int64     `gorm:"column:parent_id;type:BIGINT(20);" json:"parentId"`
	OrderNum   int32     `gorm:"column:order_num;type:INT(4);" json:"orderNum"`
	Path       string    `gorm:"column:path;type:VARCHAR(200);" json:"path"`
	Component  string    `gorm:"column:component;type:VARCHAR(255);" json:"component"`
	IsFrame    int32     `gorm:"column:is_frame;type:INT(1);" json:"isFrame"`
	MenuType   string    `gorm:"column:menu_type;type:CHAR(1);" json:"menuType"`
	Visible    string    `gorm:"column:visible;type:CHAR(1);" json:"visible"`
	Status     string    `gorm:"column:status;type:CHAR(1);" json:"status"`
	Perms      string    `gorm:"column:perms;type:VARCHAR(100);" json:"perms"`
	Icon       string    `gorm:"column:icon;type:VARCHAR(100);" json:"icon"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;" json:"createTime"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;" json:"updateTime"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);" json:"remark"`
	DelFlag    string    `gorm:"column:del_flag;type:CHAR(1);" json:"delFlag"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}