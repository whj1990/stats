package repo

import (
	"context"
	"time"

	"github.com/whj1990/go-core/handler"
	"github.com/whj1990/go-core/repository"
)

type OrderData struct {
	Id             int64     `form:"id" json:"id"` // 唯一标识
	OrderNo        string    //订单号
	ChannelOrderNo string    //通道返回订单号
	TenantId       string    //租户ID
	UserId         int64     //买家用户id
	UserCode       string    //买家用户编号
	Count          int64     //商品数量
	CollectionId   int64     //系列ID
	TotalPrice     float64   //买家价格
	TotalPoundage  float64   //手续费
	TenantProfit   float64   //一级平台利润
	Status         int8      //订单状态 0-等待提交到通道，1:待付款 2:待发货 3:待收货(已发货) 5:成功 6:失败
	PayTime        time.Time //付款时间
	WalletCode     string    //钱包支付类型
	MarketType     int8      //1-首发，2-转卖
	IsPayed        int8      //是否已支付，1.已支付0.未支付
	CloseType      int8      //订单关闭原因 1-超时未支付 4-买家取消 15-已通过货到付款交易
	DeleteStatus   int8      //用户订单删除状态，0：没有删除， 1：回收站， 2：永久删除
	CreateTime     time.Time
	UpdateTime     time.Time
}

func (*OrderData) TableName() string {
	return "t_order"
}

type OrderRepo interface {
	TradeStats(ctx context.Context, tenantId string, startTime, endTime time.Time) (amount float64, count int64, err error)
}

type orderRepo struct {
	repository.BaseRepo
}

func (r *orderRepo) TradeStats(ctx context.Context, tenantId string, startTime, endTime time.Time) (amount float64, count int64, err error) {
	whereData := OrderData{
		Status: 5,
	}
	query := r.Db.WithContext(ctx).Model(&OrderData{}).Where("delete_status = 0")
	if tenantId != "" {
		whereData.TenantId = tenantId
	}
	query.Where(whereData).Where("pay_time between ? and ?", startTime.Format(time.DateTime), endTime.Format(time.DateTime))
	if result := query.Pluck("SUM(total_price) as total_amount", &amount); result.Error != nil {
		err = handler.HandleError(err)
	}
	if result := query.Pluck("SUM(count) as total_count", &count); result.Error != nil {
		err = handler.HandleError(err)
	}
	return
}
