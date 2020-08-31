package receipt

import (
	base "gf-app/app/api"
	"gf-app/app/service/receipt"

	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Receipt struct{}

/**
获取收货退货列表
*/
func (rec *Receipt) GetMergeList(r *ghttp.Request) {
	//校验请求数据
	type rValid struct {
		SettleId string `c:"SettleId,omitempty"`
		FinCode  string `c:"FinCode,omitempty"`
		Vid      string `c:"Vid,omitempty"`
		Pn       int
		Rn       int
	}
	var req *rValid
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	//获取收货退货数据
	data, total, error := receipt.GetReceiptAndReturnData(gconv.Map(req))
	if error := gerror.Cause(error); error != nil {
		response.JsonExit(r, 0, error.Error())
	}
	response.JsonExit(r, 0, "", base.ResponseListTem{List: data, Total: total})
}
