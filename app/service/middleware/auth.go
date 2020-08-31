package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

func Vendor(r *ghttp.Request) {
	r.Middleware.Next()
}
