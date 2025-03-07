// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetApplicationOrder
//
// 该接口用于查询某个订单的具体信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN
func (r *ApplicationService) GetApplicationOrder(ctx context.Context, request *GetApplicationOrderReq, options ...MethodOptionFunc) (*GetApplicationOrderResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationOrder != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationOrder mock enable")
		return r.cli.mock.mockApplicationGetApplicationOrder(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationOrder",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/pay/v1/order/get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationOrderResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApplicationGetApplicationOrder(f func(ctx context.Context, request *GetApplicationOrderReq, options ...MethodOptionFunc) (*GetApplicationOrderResp, *Response, error)) {
	r.mockApplicationGetApplicationOrder = f
}

func (r *Mock) UnMockApplicationGetApplicationOrder() {
	r.mockApplicationGetApplicationOrder = nil
}

type GetApplicationOrderReq struct {
	OrderID string `query:"order_id" json:"-"` // 订单ID
}

type getApplicationOrderResp struct {
	Code int64                    `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApplicationOrderResp `json:"data,omitempty"` // 返回的业务信息
}

type GetApplicationOrderResp struct {
	Order *GetApplicationOrderRespOrder `json:"order,omitempty"` // 订单信息
}

type GetApplicationOrderRespOrder struct {
	OrderID       string `json:"order_id,omitempty"`        // 订单ID，唯一标识
	PricePlanID   string `json:"price_plan_id,omitempty"`   // 价格方案ID，唯一标识
	PricePlanType string `json:"price_plan_type,omitempty"` // 价格方案类型 "trial" -试用；"permanent"-一次性付费；"per_year"-企业年付费；"per_month"-企业月付费；"per_seat_per_year"-按人按年付费；"per_seat_per_month"-按人按月付费；"permanent_count"-按次付费；
	Seats         int64  `json:"seats,omitempty"`           // 实际购买人数 仅对price_plan_type为per_seat_per_year和per_seat_per_month 有效
	BuyCount      int64  `json:"buy_count,omitempty"`       // 购买数量 总是为1
	CreateTime    string `json:"create_time,omitempty"`     // 订单创建时间戳
	PayTime       string `json:"pay_time,omitempty"`        // 订单支付时间戳
	Status        string `json:"status,omitempty"`          // 订单当前状态，"normal" -正常；"refund"-已退款；
	BuyType       string `json:"buy_type,omitempty"`        // 购买类型，"buy" - 普通购买;"upgrade"-为升级购买(仅price_plan_type 为per_year，per_month，per_seat_per_year，per_seat_per_month时可升级购买);"renew" - 续费购买；
	SrcOrderID    string `json:"src_order_id,omitempty"`    // 源订单ID，当前订单为升级购买时，即buy_type为upgrade时，此字段记录源订单等ID
	DstOrderID    string `json:"dst_order_id,omitempty"`    // 升级后的新订单ID，当前订单如果做过升级购买，此字段记录升级购买后生成的新订单ID，当前订单仍然有效
	OrderPayPrice int64  `json:"order_pay_price,omitempty"` // 订单实际支付金额, 单位分
	TenantKey     string `json:"tenant_key,omitempty"`      // 租户唯一标识
}
