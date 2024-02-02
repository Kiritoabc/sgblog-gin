package admin

import (
	"errors"
	"github.com/jinzhu/copier"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/utils"
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

func (s *UserService) UpdateUserInfo(user *blog.SysUser) error {

	return global.SG_BLOG_DB.Model(&blog.SysUser{}).Where("id = ?", user.Id).Updates(user).Error
}

func (s UserService) Register(user blog.SysUser) error {
	existUser, err := isExistUser(user)
	if err != nil {
		return err
	}
	if existUser.UserName == user.UserName {
		return errors.New("用户名重复")
	}
	// 密码加密
	user.Password = utils.MD5V([]byte(user.Password))
	// 创建用户
	if err := global.SG_BLOG_DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
