package receipt

import (
	"gf-app/app/model"
	"gf-app/app/model/receipt_head"
	"strings"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
)

//校验请求数据
type GetReceiptAndReturnInput struct {
	SettleId string `c:"SettleId,omitempty"`
	FinCode  string `c:"FinCode,omitempty"`
	Vid      string `c:"Vid,omitempty"`
	Pn       int    `c:"Pn,omitempty"`
	Rn       int    `c:"Rn,omitempty"`
}

/**
获取收货和退货数据
*/
func GetReceiptAndReturnData(p GetReceiptAndReturnInput) (g.Array, int, error) {
	var whereSql string
	var where []string
	condition := gconv.Map(p)
	binds := g.Slice{model.VenderId, model.VenderId}
	//拼接sql
	if settleId, ok := condition["SettleId"]; ok {
		where = append(where, "settle_id=?")
		binds = append(binds, gconv.String(settleId))
	}
	if finCode, ok := condition["FinCode"]; ok {
		where = append(where, "fin_code=?")
		binds = append(binds, gconv.String(finCode))
	}
	if len(where) > 0 {
		whereSql += " WHERE " + strings.Join(where, " AND ")
	}
	if rn, ok := condition["Rn"]; ok {
		whereSql += " LIMIT ?"
		binds = append(binds, gconv.Int(rn))
	}
	if pn, ok := condition["Pn"]; ok {
		whereSql += " OFFSET ?"
		binds = append(binds, gconv.Int(pn))
	}
	data, total, error := receipt_head.GetReceiptAndReturnData(whereSql, binds)
	return g.Array{data}, total, error
}
