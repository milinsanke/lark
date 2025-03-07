// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateMeetingRoomRoom 该接口用于更新会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTNwYjLzUDM24yM1AjN
func (r *MeetingRoomService) UpdateMeetingRoomRoom(ctx context.Context, request *UpdateMeetingRoomRoomReq, options ...MethodOptionFunc) (*UpdateMeetingRoomRoomResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomUpdateMeetingRoomRoom != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#UpdateMeetingRoomRoom mock enable")
		return r.cli.mock.mockMeetingRoomUpdateMeetingRoomRoom(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "UpdateMeetingRoomRoom",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateMeetingRoomRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomUpdateMeetingRoomRoom(f func(ctx context.Context, request *UpdateMeetingRoomRoomReq, options ...MethodOptionFunc) (*UpdateMeetingRoomRoomResp, *Response, error)) {
	r.mockMeetingRoomUpdateMeetingRoomRoom = f
}

func (r *Mock) UnMockMeetingRoomUpdateMeetingRoomRoom() {
	r.mockMeetingRoomUpdateMeetingRoomRoom = nil
}

type UpdateMeetingRoomRoomReq struct {
	RoomID       string  `json:"room_id,omitempty"`        // 要更新的会议室ID
	Name         *string `json:"name,omitempty"`           // 会议室名称
	Capacity     *int64  `json:"capacity,omitempty"`       // 容量
	IsDisabled   *bool   `json:"is_disabled,omitempty"`    // 是否禁用
	CustomRoomID *string `json:"custom_room_id,omitempty"` // 租户自定义会议室ID
}

type updateMeetingRoomRoomResp struct {
	Code int64                      `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *UpdateMeetingRoomRoomResp `json:"data,omitempty"`
}

type UpdateMeetingRoomRoomResp struct{}
