// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SubscribeCalendarEvent 该接口用于以用户身份订阅指定日历下的日程变更事件。
//
// 用户必须对日历有访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/subscription
func (r *CalendarService) SubscribeCalendarEvent(ctx context.Context, request *SubscribeCalendarEventReq, options ...MethodOptionFunc) (*SubscribeCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarSubscribeCalendarEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SubscribeCalendarEvent mock enable")
		return r.cli.mock.mockCalendarSubscribeCalendarEvent(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Calendar#SubscribeCalendarEvent call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SubscribeCalendarEvent request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/subscription",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(subscribeCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#SubscribeCalendarEvent POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/subscription failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#SubscribeCalendarEvent POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/subscription failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Calendar", "SubscribeCalendarEvent", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SubscribeCalendarEvent success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockCalendarSubscribeCalendarEvent(f func(ctx context.Context, request *SubscribeCalendarEventReq, options ...MethodOptionFunc) (*SubscribeCalendarEventResp, *Response, error)) {
	r.mockCalendarSubscribeCalendarEvent = f
}

func (r *Mock) UnMockCalendarSubscribeCalendarEvent() {
	r.mockCalendarSubscribeCalendarEvent = nil
}

type SubscribeCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type subscribeCalendarEventResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeCalendarEventResp `json:"data,omitempty"`
}

type SubscribeCalendarEventResp struct{}
