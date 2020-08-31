package receipt

import (
	"gf-app/app/model/receipt"

	"github.com/gogf/gf/frame/g"
)

/**
获取收货和退货数据
*/
func GetReceiptAndReturnData(condition g.Map) (g.Array, int, error) {
	data, total, error := receipt.GetReceiptAndReturnData(condition)
	return g.Array{data}, total, error
}
