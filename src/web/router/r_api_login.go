package router

import (
	"github.com/LyricTian/gin-admin/src/web/ctl"
	"github.com/gin-gonic/gin"
)

// APILoginRouter 注册登录相关路由
func APILoginRouter(g *gin.RouterGroup, c *ctl.Login) {
	POST(g, "/login", c.Login, "用户登录")
	POST(g, "/logout", c.Logout, "用户登出")
	GET(g, "/current/user", c.GetCurrentUserInfo, "获取当前用户信息")
	GET(g, "/current/menus", c.QueryCurrentUserMenus, "查询当前用户菜单")
}
