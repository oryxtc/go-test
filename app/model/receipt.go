package receipt

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

/**
获取收货和退货数据
*/
func GetReceiptAndReturnData(condition g.Map) (gdb.Result, int, error) {
	db := g.DB()
	sql := "SELECT SQL_CALC_FOUND_ROWS * FROM (" +
		"SELECT '1' order_type,b.vid,b.fin_code,b.`name` AS vender_name,b.type AS vender_type,b.settle_type,a.venderId,a.zone_id,a.receipt_id AS rr_num,a.allocate_id,a.order_date,a.order_id,a.order_name,a.reference,a.type,a.enter_date AS rr_date,a.created_at,a.it_amount,a.nt_amount,a.`status`,a.settle_id,a.is_write_off,a.warehouse_code FROM receipt_head AS a LEFT JOIN vender_main AS b ON a.vid=b.vid AND a.venderId=b.venderId WHERE a.venderId=? " +
		"UNION ALL " +
		"SELECT '2' order_type,b.vid,b.fin_code, b.name AS vender_name,b.type AS vender_type,b.settle_type,a.venderId,a.zone_id,a.return_id AS order_num,a.allocate_id,a.order_date,'3' order_id,'' order_name,'' reference,'0' type,a.date AS rr_date,a.created_at,a.it_amount,a.nt_amount,a.`status`,a.settle_id,a.is_write_off,a.warehouse_code FROM return_head AS a LEFT JOIN vender_main AS b ON a.vid=b.vid AND a.venderId=b.venderId WHERE a.venderId=?" +
		") a"
	//拼接sql
	if rn, ok := condition["rn"]; ok {
		sql += " LIMIT " + gconv.String(rn)
	}
	if pn, ok := condition["pn"]; ok {
		sql += " OFFSET " + gconv.String(pn)
	}
	data, error := db.GetAll(sql, g.Slice{10099, 10099})
	if error != nil {
		return nil, 0, gerror.New("获取数据失败!")
	}
	total, _ := db.GetCount("SELECT FOUND_ROWS()")
	return data, total, error
}
