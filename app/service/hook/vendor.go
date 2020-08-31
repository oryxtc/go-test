package hook

import (
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Vendor(r *ghttp.Request) {
	//获取venderId
	venderId := r.Header["Venderid"]
	if len(venderId) == 0 {
		response.Json(r, -1, "缺少venderId")
		r.ExitAll()
	}
	r.Middleware.Next()
}
