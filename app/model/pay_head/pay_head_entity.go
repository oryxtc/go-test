// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package pay_head

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table pay_head.
type Entity struct {
    Id               uint64  `orm:"id,primary"        json:"id"`                //                                              
    Vid              uint64  `orm:"vid"               json:"vid"`               // id                                           
    PayId            uint64  `orm:"pay_id,unique"     json:"pay_id"`            // id                                           
    ItAmount         float64 `orm:"it_amount"         json:"it_amount"`         // 结算单金额-预付款金额                        
    Status           int     `orm:"status"            json:"status"`            // 2-待确认, 3-已确认,待审核,4-待付款,5-已付款  
    DueDate          uint    `orm:"due_date"          json:"due_date"`          // 账期到期日 = 结算单生成日期+账期天数         
    ApproveUid       uint64  `orm:"approve_uid"       json:"approve_uid"`       // id                                           
    ApproveUsername  string  `orm:"approve_username"  json:"approve_username"`  // 审核人                                       
    ApproveDate      uint    `orm:"approve_date"      json:"approve_date"`      // 审核日期                                     
    IssueUid         uint64  `orm:"issue_uid"         json:"issue_uid"`         // id                                           
    IssueUsername    string  `orm:"issue_username"    json:"issue_username"`    //                                              
    IssueDate        uint    `orm:"issue_date"        json:"issue_date"`        //                                              
    ModifiedUid      uint64  `orm:"modified_uid"      json:"modified_uid"`      // 修改人id                                     
    ModifiedUsername string  `orm:"modified_username" json:"modified_username"` // 修改人                                       
    ModifiedDate     uint    `orm:"modified_date"     json:"modified_date"`     // 修改时间                                     
    CompanyId        string  `orm:"company_id"        json:"company_id"`        // 支付单生成时供商的结算公司                   
    Type             uint    `orm:"type"              json:"type"`              // 支付单类型：1-正常；2-预付款支付单           
    Comment          string  `orm:"comment"           json:"comment"`           // 备注                                         
    IsDeleted        int     `orm:"is_deleted"        json:"is_deleted"`        // 是否已删除，0-未删除，1-已删除               
    CreatedAt        uint    `orm:"created_at"        json:"created_at"`        //                                              
    UpdatedAt        uint    `orm:"updated_at"        json:"updated_at"`        //                                              
    VenderId         int64   `orm:"venderId"          json:"vender_id"`         // 租户id/商家id                                
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}