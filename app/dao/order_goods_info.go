// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/app/dao/internal"
)

// orderGoodsInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type orderGoodsInfoDao struct {
	*internal.OrderGoodsInfoDao
}

var (
	// OrderGoodsInfo is globally public accessible object for table order_goods_info operations.
	OrderGoodsInfo orderGoodsInfoDao
)

func init() {
	OrderGoodsInfo = orderGoodsInfoDao{
		internal.NewOrderGoodsInfoDao(),
	}
}

// Fill with you ideas below.
