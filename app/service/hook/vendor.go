package hook

import (
	"gf-app/app/model"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Vendor(r *ghttp.Request) {
	//获取venderId
	venderId := r.Header.Get("Venderid")
	if len(venderId) == 0 {
		response.Json(r, -1, "缺少venderId")
		r.ExitAll()
	}
	//赋值venderId
	model.VenderId = gconv.Int(venderId)
}
