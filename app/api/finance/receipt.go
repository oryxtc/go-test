package receipt

import (
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Receipt struct{}

func (rec *Receipt) GetMergeList(r *ghttp.Request) {
	//校验请求数据
	type rValid struct {
		SettleId string `p:"settle_id" v:"required"`
		FinCode  string `p:"fin_code" v:"required"`
	}
	var req *rValid
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", req)
}
