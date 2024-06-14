package authentication

import (
	"WlFrame-gin/conf"
	"WlFrame-gin/utils/global"
	"WlFrame-gin/utils/jwt"
	"fmt"
	casbin "github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Enforcer *casbin.Enforcer

// 初始化casbin
func CasbinSetup() {
	global.DBConfig = conf.GetDatabaseConfig()

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		global.DBConfig.Username,
		global.DBConfig.Password,
		global.DBConfig.Host,
		global.DBConfig.Port,
		global.DBConfig.DbName,
	)
	a, _ := xormadapter.NewAdapter("mysql", dataSourceName, true)
	e, _ := casbin.NewEnforcer("conf/rbac_models.conf", a)

	Enforcer = e
}

// 拦截器
func Rbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		//校验token
		claims, ok := jwt.AnalysisToken(token)

		if !ok {
			fmt.Println("token校验失败")
			// 重定向到首页
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else if token == "---000---" {
			//跳过鉴权
			c.Next()
		} else {
			var e *casbin.Enforcer
			e = Enforcer

			//从mysql中加载策略
			err := e.LoadPolicy()
			if err != nil {
				log.Println("从mysql中加载策略失败", err)
			}

			//获取请求的URI
			obj := c.Request.URL.RequestURI()
			//获取请求方法
			act := c.Request.Method
			//获取用户的角色列表
			//value, ok := c.Get("role")

			if !ok {
				//角色信息为空
				fmt.Println("很遗憾,权限验证没有通过")
				c.Abort()
			} else if len(claims.Role) == 0 {
				//角色信息为空
				fmt.Println("很遗憾,权限验证没有通过")
				c.Abort()
			} else {
				c.Set("roles", claims.Role)
				c.Set("sysPermissionIDs", claims.SysPermissionIDs)
				if obj == "/api/v1/system/permission/tree" {
					c.Next()
					return
				}
				//默认权限校验不通过
				isTrue := false
				for _, roleName := range claims.Role {
					sub := roleName
					//判断策略中是否存在
					ok, _ := e.Enforce(sub, obj, "GET")
					if ok {
						log.Println("权限校验通过：", obj, act, sub)
						//只要有一个角色的权限校验通过，那么就该用户权限校验通过
						isTrue = true
						break
					}
				}
				if !isTrue {
					log.Println("很遗憾,权限验证没有通过:", act, obj)
					c.Abort()
				} else {
					fmt.Println("权限验证通过")
					c.Next()
				}
			}
		}
	}
}
