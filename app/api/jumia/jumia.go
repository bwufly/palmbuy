package jumia

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"palmbuy/app/service/jumia"
)

// Download 下载jumia商品
func Download(r *ghttp.Request) {
	categoryName := r.GetString("categoryId")
	_ = jumia.JumiaService.SaveProductsToFileByCategoryName(categoryName)
	r.Response.WriteJson(g.Map{
		"code": 20000,
		"data": nil,
		"message": "下载成功",
	})
}

// List
func List(r *ghttp.Request)  {

	 total, items := jumia.JumiaService.GetFileList(r.GetInt("page"),r.GetInt("limit"))
	r.Response.WriteJson(g.Map{
		"code": 20000,
		"data": g.Map{
			"total":total,
			"items":items,
		},
		"message": "",
	})
}
