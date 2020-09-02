package finance

import (
	"gf-app/app/service/receipt"
	"gf-app/library/response"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/net/ghttp"
)

type Pay struct{}

func (p *Pay) GetPayDetail(r *ghttp.Request) {
	//校验数据
	type v struct {
		PayId string
	}
	var req *v
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	//获取详情
	data, error := receipt.GetDetail(req.PayId)
	if err := gerror.Cause(error); err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	response.JsonExit(r, 0, "", data)

}
