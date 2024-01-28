package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
)

type LoginRouter struct{}

func (s *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	loginApi := v1.ApiGroupApp.BlogApiGroup.BlogLoginApi
	{
		Router.POST("/login", loginApi.Login)
	}
}
