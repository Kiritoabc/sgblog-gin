package blog

import (
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
)

type LoginService struct{}

func (s LoginService) Login(user blog.SysUser) (*vo.BlogUserLoginVo, error) {
	return nil, nil
}
