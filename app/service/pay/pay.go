package pay

import (
	"gf-app/app/model/pay_head"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

func GetDetail(payId string) (*pay_head.Entity, error) {
	detail, error := pay_head.Model.FindOne(g.Map{"pay_id": payId})
	if error != nil {
		return nil, gerror.New("获取数据失败!")
	}
	return detail, error
}
