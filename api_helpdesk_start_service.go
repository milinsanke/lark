// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// StartHelpdeskService 该接口用于创建服务台对话。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/start_service
func (r *HelpdeskService) StartHelpdeskService(ctx context.Context, request *StartHelpdeskServiceReq, options ...MethodOptionFunc) (*StartHelpdeskServiceResp, *Response, error) {
	if r.cli.mock.mockHelpdeskStartHelpdeskService != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#StartHelpdeskService mock enable")
		return r.cli.mock.mockHelpdeskStartHelpdeskService(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "StartHelpdeskService",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/start_service",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(startHelpdeskServiceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskStartHelpdeskService(f func(ctx context.Context, request *StartHelpdeskServiceReq, options ...MethodOptionFunc) (*StartHelpdeskServiceResp, *Response, error)) {
	r.mockHelpdeskStartHelpdeskService = f
}

func (r *Mock) UnMockHelpdeskStartHelpdeskService() {
	r.mockHelpdeskStartHelpdeskService = nil
}

type StartHelpdeskServiceReq struct {
	HumanService    *bool    `json:"human_service,omitempty"`    // 是否直接进入人工, 示例值：false
	AppointedAgents []string `json:"appointed_agents,omitempty"` // 客服 open ids (获取方式参考[获取单个用户信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get))，human_service需要为true
	OpenID          string   `json:"open_id,omitempty"`          // 用户 open id,(获取方式参考[获取单个用户信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)), 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	CustomizedInfo  *string  `json:"customized_info,omitempty"`  // 工单来源自定义信息，长度限制1024字符，如设置，[获取工单详情](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get)会返回此信息, 示例值："test customized info"
}

type startHelpdeskServiceResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *StartHelpdeskServiceResp `json:"data,omitempty"`
}

type StartHelpdeskServiceResp struct {
	ChatID   string `json:"chat_id,omitempty"`   // 客服群open ID
	TicketID string `json:"ticket_id,omitempty"` // 工单ID
}
