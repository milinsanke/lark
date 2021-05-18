// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetPublicMailbox 获取公共邮箱信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/get
func (r *MailService) GetPublicMailbox(ctx context.Context, request *GetPublicMailboxReq, options ...MethodOptionFunc) (*GetPublicMailboxResp, *Response, error) {
	if r.cli.mock.mockMailGetPublicMailbox != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetPublicMailbox mock enable")
		return r.cli.mock.mockMailGetPublicMailbox(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Mail#GetPublicMailbox call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetPublicMailbox request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Mail#GetPublicMailbox GET https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Mail#GetPublicMailbox GET https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Mail", "GetPublicMailbox", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetPublicMailbox success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMailGetPublicMailbox(f func(ctx context.Context, request *GetPublicMailboxReq, options ...MethodOptionFunc) (*GetPublicMailboxResp, *Response, error)) {
	r.mockMailGetPublicMailbox = f
}

func (r *Mock) UnMockMailGetPublicMailbox() {
	r.mockMailGetPublicMailbox = nil
}

type GetPublicMailboxReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
}

type getPublicMailboxResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetPublicMailboxResp `json:"data,omitempty"` //
}

type GetPublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}
