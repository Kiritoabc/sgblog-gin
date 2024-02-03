package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	blog_utils "sgblog-go/app/admin/cmd/utils"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := blog_utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, blog_utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		//c.Set("claims", claims)
		key := fmt.Sprintf("login:%d", claims.BaseClaims.ID)
		jsonResult, err := global.SG_BLOG_REDIS.Get(context.Background(), key).Result()
		if err != nil {
			response.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		var loginResponse *blog.UserLogin
		err = json.Unmarshal([]byte(jsonResult), &loginResponse)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("loginUser", loginResponse)
		c.Set("userId", loginResponse.User.Id)
		c.Next()
	}
}
