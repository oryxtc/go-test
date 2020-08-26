package router

import (
	"gf-app/app/api/finance"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/finance", func(group *ghttp.RouterGroup) {
		group.ALL("/{.struct}/{.method}", new(receipt.Receipt))
	})
}
