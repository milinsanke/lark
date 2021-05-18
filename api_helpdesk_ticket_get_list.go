// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetTicketList 该接口用于获取全部工单详情。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list
func (r *HelpdeskService) GetTicketList(ctx context.Context, request *GetTicketListReq, options ...MethodOptionFunc) (*GetTicketListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetTicketList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetTicketList mock enable")
		return r.cli.mock.mockHelpdeskGetTicketList(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Helpdesk#GetTicketList call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetTicketList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedHelpdeskAuth: true,
	}
	resp := new(getTicketListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#GetTicketList GET https://open.feishu.cn/open-apis/helpdesk/v1/tickets failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#GetTicketList GET https://open.feishu.cn/open-apis/helpdesk/v1/tickets failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "GetTicketList", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetTicketList success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskGetTicketList(f func(ctx context.Context, request *GetTicketListReq, options ...MethodOptionFunc) (*GetTicketListResp, *Response, error)) {
	r.mockHelpdeskGetTicketList = f
}

func (r *Mock) UnMockHelpdeskGetTicketList() {
	r.mockHelpdeskGetTicketList = nil
}

type GetTicketListReq struct {
	TicketID         *string                    `query:"ticket_id" json:"-"`         // 搜索条件：工单ID, 示例值："123456"
	AgentID          *string                    `query:"agent_id" json:"-"`          // 搜索条件: 客服id, 示例值："ou_b5de90429xxx"
	ClosedByID       *string                    `query:"closed_by_id" json:"-"`      // 搜索条件: 关单客服id, 示例值："ou_b5de90429xxx"
	Type             *int64                     `query:"type" json:"-"`              // 搜索条件: 工单类型 1:bot 2:人工, 示例值：1
	Channel          *int64                     `query:"channel" json:"-"`           // 搜索条件: 工单渠道, 示例值：0
	Solved           *int64                     `query:"solved" json:"-"`            // 搜索条件: 工单是否解决 1:没解决 2:已解决	, 示例值：1
	Score            *int64                     `query:"score" json:"-"`             // 搜索条件: 工单评分, 示例值：1
	StatusList       []int64                    `query:"status_list" json:"-"`       // 搜索条件: 工单状态列表
	GuestName        *string                    `query:"guest_name" json:"-"`        // 搜索条件: 用户名称, 示例值："abc"
	GuestID          *string                    `query:"guest_id" json:"-"`          // 搜索条件: 用户id, 示例值："ou_b5de90429xxx"
	CustomizedFields []*HelpdeskCustomizedField `query:"customized_fields" json:"-"` // 搜索条件: 自定义字段列表
	Tags             []string                   `query:"tags" json:"-"`              // 搜索条件: 用户标签列表
	Page             *int64                     `query:"page" json:"-"`              // 页数, 从1开始, 默认为1, 示例值：1
	PageSize         *int64                     `query:"page_size" json:"-"`         // 当前页大小，最大为200, 默认为20, 示例值：20
	CreateTimeStart  *int64                     `query:"create_time_start" json:"-"` // 搜索条件: 工单创建起始时间 ms, 示例值：1616920429000
	CreateTimeEnd    *int64                     `query:"create_time_end" json:"-"`   // 搜索条件: 工单创建结束时间 ms, 示例值：1616920429000
	UpdateTimeStart  *int64                     `query:"update_time_start" json:"-"` // 搜索条件: 工单修改起始时间 ms, 示例值：1616920429000
	UpdateTimeEnd    *int64                     `query:"update_time_end" json:"-"`   // 搜索条件: 工单修改结束时间 ms, 示例值：1616920429000
}

type getTicketListResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketListResp `json:"data,omitempty"` //
}

type GetTicketListResp struct {
	Total   int64                      `json:"total,omitempty"`   // 工单总数
	Tickets []*GetTicketListRespTicket `json:"tickets,omitempty"` // 工单
}

type GetTicketListRespTicket struct {
	TicketID         string                                    `json:"ticket_id,omitempty"`         // 工单ID
	HelpdeskID       string                                    `json:"helpdesk_id,omitempty"`       // 服务台ID
	Guest            *GetTicketListRespTicketGuest             `json:"guest,omitempty"`             // 工单创建用户
	Stage            int64                                     `json:"stage,omitempty"`             // 工单阶段，1：bot，2：人工
	Status           int64                                     `json:"status,omitempty"`            // 工单状态，1：已创建 2: 处理中 3: 排队中 4：待定 5：待用户响应 50: 被机器人关闭 51: 被人工关闭
	Score            int64                                     `json:"score,omitempty"`             // 工单评分，1：不满意，2:一般，3:满意
	CreatedAt        int64                                     `json:"created_at,omitempty"`        // 工单创建时间
	UpdatedAt        int64                                     `json:"updated_at,omitempty"`        // 工单更新时间，没有值时为-1
	ClosedAt         int64                                     `json:"closed_at,omitempty"`         // 工单结束时间
	Agents           []*GetTicketListRespTicketAgent           `json:"agents,omitempty"`            // 工单客服
	Channel          int64                                     `json:"channel,omitempty"`           // 工单渠道
	Solve            int64                                     `json:"solve,omitempty"`             // 工单是否解决 1:没解决 2:已解决
	ClosedBy         *GetTicketListRespTicketClosedBy          `json:"closed_by,omitempty"`         // 关单用户ID
	Collaborators    []*GetTicketListRespTicketCollaborator    `json:"collaborators,omitempty"`     // 工单协作者
	CustomizedFields []*GetTicketListRespTicketCustomizedField `json:"customized_fields,omitempty"` // 自定义字段列表，没有值时不设置
}

type GetTicketListRespTicketGuest struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketAgent struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketClosedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketCollaborator struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketCustomizedField struct {
	ID          string `json:"id,omitempty"`           // 自定义字段ID
	Value       string `json:"value,omitempty"`        // 自定义字段值
	KeyName     string `json:"key_name,omitempty"`     // 键名
	DisplayName string `json:"display_name,omitempty"` // 展示名称
	Position    int64  `json:"position,omitempty"`     // 展示位置
	Required    bool   `json:"required,omitempty"`     // 是否必填
	Editable    bool   `json:"editable,omitempty"`     // 是否可修改
}
