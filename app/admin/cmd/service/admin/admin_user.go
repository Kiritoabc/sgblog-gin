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

func (s *UserService) Register(user blog.SysUser) error {
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

func (s *UserService) SelectUserPage(user blog.SysUser, pageNum int, pageSize int) ([]*vo.UserVo, int64, error) {
	var userList []*vo.UserVo

	query := global.SG_BLOG_DB.Model(&blog.SysUser{})

	// build query
	if user.UserName != "" {
		query = query.Where("user_name like ?", "%"+user.UserName+"%")
	}

	if user.Status != "" {
		query = query.Where("status = ?", user.Status)
	}

	if user.Phonenumber != "" {
		query = query.Where("phonenumber =  ?", user.Phonenumber)
	}

	// page query
	offset := (pageNum - 1) * pageSize
	limit := pageSize

	var total int64
	if err := query.Count(&total).Offset(offset).Limit(limit).Find(&userList).Error; err != nil {
		return nil, 0, err
	}

	var userVoList []*vo.UserVo
	err := copier.Copy(&userVoList, &userList)
	if err != nil {
		return nil, 0, err
	}
	return userVoList, total, nil
}

func (s *UserService) CheckUserNameUnique(user blog.SysUser) bool {
	var count int64
	global.SG_BLOG_DB.Model(&blog.SysUser{}).Where("user_name = ?", user.UserName).Count(&count)
	return count == 0
}

func (s *UserService) CheckPhoneUnique(user blog.SysUser) bool {
	var count int64
	global.SG_BLOG_DB.Model(&blog.SysUser{}).Where("phonenumber = ?", user.Phonenumber).Count(&count)
	return count == 0
}

func (s *UserService) CheckEmailUnique(user blog.SysUser) bool {
	var count int64
	global.SG_BLOG_DB.Model(&blog.SysUser{}).Where("email = ?", user.Email).Count(&count)
	return count == 0
}

func (s *UserService) GetById(userId int64) (*blog.SysUser, error) {

	var user *blog.SysUser
	err := global.SG_BLOG_DB.Model(&blog.SysUser{}).Where("id = ?", userId).First(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user blog.SysUser) error {
	db := global.SG_BLOG_DB

	// 开启事务
	tx := db.Begin()

	// 删除用户与角色关联
	if err := tx.Where("user_id = ?", user.Id).Delete(&blog.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 新增用户与角色关联
	userRoles := make([]blog.SysUserRole, 0)
	for _, roleId := range user.RoleIds { // 假设RoleIds是从User结构体中提取的角色ID列表
		userRole := blog.SysUserRole{UserId: user.Id, RoleId: roleId}
		userRoles = append(userRoles, userRole)
	}

	if err := tx.Create(&userRoles).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新用户信息
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
