// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EndMeeting 结束一个进行中的会议
//
// 会议正在进行中，且操作者须具有相应的权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/end
func (r *VCService) EndMeeting(ctx context.Context, request *EndMeetingReq, options ...MethodOptionFunc) (*EndMeetingResp, *Response, error) {
	if r.cli.mock.mockVCEndMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#EndMeeting mock enable")
		return r.cli.mock.mockVCEndMeeting(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] VC#EndMeeting call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] VC#EndMeeting request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "PATCH",
		URL:          "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/end",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(endMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] VC#EndMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/end failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] VC#EndMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/end failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("VC", "EndMeeting", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] VC#EndMeeting success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCEndMeeting(f func(ctx context.Context, request *EndMeetingReq, options ...MethodOptionFunc) (*EndMeetingResp, *Response, error)) {
	r.mockVCEndMeeting = f
}

func (r *Mock) UnMockVCEndMeeting() {
	r.mockVCEndMeeting = nil
}

type EndMeetingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID, 示例值："6911188411932033028"
}

type endMeetingResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *EndMeetingResp `json:"data,omitempty"`
}

type EndMeetingResp struct{}
