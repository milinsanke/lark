// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdatePublicMailboxPatch 更新公共邮箱部分字段，没有填写的字段不会被更新
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/patch
func (r *MailService) UpdatePublicMailboxPatch(ctx context.Context, request *UpdatePublicMailboxPatchReq, options ...MethodOptionFunc) (*UpdatePublicMailboxPatchResp, *Response, error) {
	if r.cli.mock.mockMailUpdatePublicMailboxPatch != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#UpdatePublicMailboxPatch mock enable")
		return r.cli.mock.mockMailUpdatePublicMailboxPatch(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "UpdatePublicMailboxPatch",
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updatePublicMailboxPatchResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMailUpdatePublicMailboxPatch(f func(ctx context.Context, request *UpdatePublicMailboxPatchReq, options ...MethodOptionFunc) (*UpdatePublicMailboxPatchResp, *Response, error)) {
	r.mockMailUpdatePublicMailboxPatch = f
}

func (r *Mock) UnMockMailUpdatePublicMailboxPatch() {
	r.mockMailUpdatePublicMailboxPatch = nil
}

type UpdatePublicMailboxPatchReq struct {
	PublicMailboxID string  `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	Name            *string `json:"name,omitempty"`             // 公共邮箱名称, 示例值："test public mailbox"
}

type updatePublicMailboxPatchResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *UpdatePublicMailboxPatchResp `json:"data,omitempty"`
}

type UpdatePublicMailboxPatchResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}
