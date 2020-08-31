package router

import (
	receipt "gf-app/app/api/finance"
	"gf-app/app/service/hook"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		//绑定全局路由回调事件
		group.Hook("/*", ghttp.HOOK_BEFORE_SERVE, hook.Vendor)
		//分组匹配路由
		s.Group("/finance", func(group *ghttp.RouterGroup) {
			group.ALL("/{.struct}/{.method}", new(receipt.Receipt))
		})
	})

}
