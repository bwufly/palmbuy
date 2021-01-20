package user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Login(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"code": 20000,
		"data": "token",
	})
}

func Info(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"code": 20000,
		"data": g.Map{
			"roles":        []string{"admin"},
			"introduction": "I am a super administrator",
			"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"name":         "Super Admin",
		},
	})
}
