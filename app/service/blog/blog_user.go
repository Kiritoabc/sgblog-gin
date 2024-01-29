package blog

import (
	"github.com/jinzhu/copier"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
)

type UserService struct{}

func (s *UserService) UserInfo(userId int64) (*vo.UserInfoVo, error) {
	var userInfoVo vo.UserInfoVo
	var user blog.SysUser
	if err := global.SG_BLOG_DB.
		Model(&blog.SysUser{}).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		return nil, err
	}
	err := copier.Copy(&userInfoVo, &user)
	if err != nil {
		return nil, err
	}
	return &userInfoVo, nil
}
