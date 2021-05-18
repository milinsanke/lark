// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// InviteMeeting 邀请参会人进入会议
//
// 发起邀请的操作者必须具有相应的权限（如果操作者为用户，则必须在会中），如果会议被锁定、或参会人数如果达到上限，则会邀请失败
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/invite
func (r *VCService) InviteMeeting(ctx context.Context, request *InviteMeetingReq, options ...MethodOptionFunc) (*InviteMeetingResp, *Response, error) {
	if r.cli.mock.mockVCInviteMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#InviteMeeting mock enable")
		return r.cli.mock.mockVCInviteMeeting(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] VC#InviteMeeting call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] VC#InviteMeeting request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "PATCH",
		URL:          "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/invite",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(inviteMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] VC#InviteMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/invite failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] VC#InviteMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/invite failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("VC", "InviteMeeting", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] VC#InviteMeeting success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCInviteMeeting(f func(ctx context.Context, request *InviteMeetingReq, options ...MethodOptionFunc) (*InviteMeetingResp, *Response, error)) {
	r.mockVCInviteMeeting = f
}

func (r *Mock) UnMockVCInviteMeeting() {
	r.mockVCInviteMeeting = nil
}

type InviteMeetingReq struct {
	UserIDType *IDType                    `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	MeetingID  string                     `path:"meeting_id" json:"-"`    // 会议ID, 示例值："6911188411932033028"
	Invitees   []*InviteMeetingReqInvitee `json:"invitees,omitempty"`     // 被邀请的用户列表
}

type InviteMeetingReqInvitee struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int64  `json:"user_type,omitempty"` // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type inviteMeetingResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *InviteMeetingResp `json:"data,omitempty"` //
}

type InviteMeetingResp struct {
	InviteResults []*InviteMeetingRespInviteResult `json:"invite_results,omitempty"` // 邀请结果
}

type InviteMeetingRespInviteResult struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
	Status   int64  `json:"status,omitempty"`    // 邀请结果, 可选值有: `1`：邀请成功, `2`：邀请失败
}
