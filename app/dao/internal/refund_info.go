// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// RefundInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type RefundInfoDao struct {
	gmvc.M                    // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB            // DB is the raw underlying database management object.
	Table   string            // Table is the table name of the DAO.
	Columns refundInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// RefundInfoColumns defines and stores column names for table refund_info.
type refundInfoColumns struct {
	Id        string // 售后退款表
	Number    string // 售后订单号
	OrderId   string // 订单id
	GoodsId   string // 要售后的商品id
	Reason    string // 退款原因
	Status    string // 状态 1待处理 2同意退款 3拒绝退款
	UserId    string // 用户id
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

func NewRefundInfoDao() *RefundInfoDao {
	return &RefundInfoDao{
		M:     g.DB("default").Model("refund_info").Safe(),
		DB:    g.DB("default"),
		Table: "refund_info",
		Columns: refundInfoColumns{
			Id:        "id",
			Number:    "number",
			OrderId:   "order_id",
			GoodsId:   "goods_id",
			Reason:    "reason",
			Status:    "status",
			UserId:    "user_id",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
			DeletedAt: "deleted_at",
		},
	}
}
