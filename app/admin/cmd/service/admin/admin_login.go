package admin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	uuid2 "github.com/gofrs/uuid/v5"
	"github.com/google/uuid"
	"sgblog-go/app/admin/cmd/global"
	blog_utils "sgblog-go/app/admin/cmd/utils"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/request"
	"sgblog-go/common/constants"
	"sgblog-go/common/utils"
	"time"
)

type LoginService struct{}

func (s *LoginService) Login(user blog.SysUser) (string, error) {
	var user1 blog.SysUser
	// 1.判断用户是否存在
	if user1, _ = isExistUser(user); user1.UserName == "" {
		return "", errors.New("用户名不存在")
	}
	// 2.判断密码是否正确
	if ok := utils.BcryptCheck(user.Password, user1.Password); !ok {
		return "", errors.New("密码错误")
	}
	// 3.获取userid生成token
	loginUser, err := LoadUserByUserName(user.UserName)
	j := &blog_utils.JWT{SigningKey: []byte(global.SG_BLOG_COFIG.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID: uuid2.UUID(uuid.New()),
		ID:   uint(loginUser.User.Id),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}

	//把用户信息存入redis
	var ctx = context.Background()
	key := fmt.Sprintf("login:%d", loginUser.User.Id)
	jsonUser1, err := json.Marshal(&loginUser)
	redis := global.SG_BLOG_REDIS
	err = redis.Set(ctx, key, jsonUser1, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func isExistUser(user blog.SysUser) (blog.SysUser, error) {
	var user1 blog.SysUser
	db := global.SG_BLOG_DB.Model(&blog.SysUser{})
	err := db.Where("user_name = ?", user.UserName).First(&user1).Error
	if err != nil {
		return user1, err
	}
	return user1, nil
}

// TODO:转移
func LoadUserByUserName(name string) (*blog.UserLogin, error) {
	db := global.SG_BLOG_DB.Model(&blog.SysUser{})
	var user blog.SysUser
	err := db.Where("user_name = ? ", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Type == constants.Admin {
		var permission []string
		err := global.SG_BLOG_DB.Model(&blog.SysUserRole{}).
			Joins("LEFT JOIN sys_role_menu ON sys_user_role.role_id = sys_role_menu.role_id").
			Joins("LEFT JOIN sys_menu ON sys_menu.id = sys_role_menu.menu_id").
			Where("sys_user_role.user_id = ?", user.Id).
			Where("sys_menu.menu_type IN (?) AND sys_menu.status = ? AND sys_menu.del_flag = ?", []string{"C", "F"}, 0, 0).
			Pluck("DISTINCT sys_menu.perms", &permission).Error
		if err != nil {
			return nil, err
		}
		return &blog.UserLogin{
			User:        &user,
			Permissions: permission,
		}, nil
	}

	return &blog.UserLogin{
		User:        &user,
		Permissions: nil,
	}, nil
}
