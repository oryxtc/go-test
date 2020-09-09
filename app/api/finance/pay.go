package finance

import (
	"gf-app/app/service/pay"
	"gf-app/library/response"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/net/ghttp"
)

/**
支付单
*/
type Pay struct{}

func (p *Pay) GetPayDetail(r *ghttp.Request) {
	payId := r.GetString("pay_id")
	if e := gvalid.Check(payId, "required", nil); e != nil {
		response.JsonExit(r, 1, e.String())
	}
	//获取详情
	data, error := pay.GetDetail(payId)
	if err := gerror.Cause(error); err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	response.JsonExit(r, 0, "", data)
}
