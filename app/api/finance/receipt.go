package finance

import (
	"gf-app/app/service/receipt"
	"gf-app/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

type Receipt struct{}

type GetReceiptAndReturnRequest struct {
	receipt.GetReceiptAndReturnInput
}

/**
获取收货退货列表
*/
func (rec *Receipt) GetMergeList(r *ghttp.Request) {
	var req *GetReceiptAndReturnRequest
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	//获取收货退货数据
	data, total, error := receipt.GetReceiptAndReturnData(req.GetReceiptAndReturnInput)
	if err := gerror.Cause(error); err != nil {
		response.JsonExit(r, 0, error.Error())
	}
	response.JsonExit(r, 0, "", response.ListData{List: data, Total: total})
}
