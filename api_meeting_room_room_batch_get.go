// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetMeetingRoomRoom 该接口用于获取指定会议室的详细信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEDOyUjLxgjM14SM4ITN
func (r *MeetingRoomService) BatchGetMeetingRoomRoom(ctx context.Context, request *BatchGetMeetingRoomRoomReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomRoomResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomRoom != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomRoom mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomRoom(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomRoom",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomBatchGetMeetingRoomRoom(f func(ctx context.Context, request *BatchGetMeetingRoomRoomReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomRoomResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomRoom = f
}

func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomRoom() {
	r.mockMeetingRoomBatchGetMeetingRoomRoom = nil
}

type BatchGetMeetingRoomRoomReq struct {
	RoomIDs []string `query:"room_ids" json:"-"` // 用于查询指定会议室的 ID
	Fields  *string  `query:"fields" json:"-"`   // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,capacity,building_id,building_name,floor_name,is_disabled,display_id"，默认返回所有字段
}

type batchGetMeetingRoomRoomResp struct {
	Code int64                        `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetMeetingRoomRoomResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetMeetingRoomRoomResp struct {
	Rooms *BatchGetMeetingRoomRoomRespRoom `json:"rooms,omitempty"` // 会议室列表
}

type BatchGetMeetingRoomRoomRespRoom struct {
	RoomID       string `json:"room_id,omitempty"`       // 会议室 ID
	BuildingID   string `json:"building_id,omitempty"`   // 会议室所属建筑物 ID
	BuildingName string `json:"building_name,omitempty"` // 会议室所属建筑物名称
	Capacity     int64  `json:"capacity,omitempty"`      // 会议室能容纳的人数
	Description  string `json:"description,omitempty"`   // 会议室的相关描述
	DisplayID    string `json:"display_id,omitempty"`    // 会议室的展示 ID
	FloorName    string `json:"floor_name,omitempty"`    // 会议室所在楼层名称
	IsDisabled   bool   `json:"is_disabled,omitempty"`   // 会议室是否不可用，若会议室不可用，则该值为 true，否则为 false
	Name         string `json:"name,omitempty"`          // 会议室名称
}
