package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"palmbuy/app/api/jumia"
	"palmbuy/app/api/user"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func init() {
	s := g.Server()

	// palmbuy
	s.Group("/palmbuy", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.Group("/user", func(userGroup *ghttp.RouterGroup) {
			userGroup.POST("/login", user.Login)
			userGroup.GET("/info", user.Info)
		})

		// jumia
		group.Group("/jumia", func(jumiaGroup *ghttp.RouterGroup) {
			jumiaGroup.ALL("/download", jumia.Download)
			jumiaGroup.GET("/list", jumia.List)
		})

	})
}
