// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBuildingList 该接口用于获取本企业下的建筑物（办公大楼）。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugzNyUjL4cjM14CO3ITN
func (r *MeetingRoomService) GetBuildingList(ctx context.Context, request *GetBuildingListReq, options ...MethodOptionFunc) (*GetBuildingListResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomGetBuildingList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#GetBuildingList mock enable")
		return r.cli.mock.mockMeetingRoomGetBuildingList(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] MeetingRoom#GetBuildingList call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#GetBuildingList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=*",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getBuildingListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] MeetingRoom#GetBuildingList GET https://open.feishu.cn/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=* failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] MeetingRoom#GetBuildingList GET https://open.feishu.cn/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=* failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "GetBuildingList", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#GetBuildingList success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomGetBuildingList(f func(ctx context.Context, request *GetBuildingListReq, options ...MethodOptionFunc) (*GetBuildingListResp, *Response, error)) {
	r.mockMeetingRoomGetBuildingList = f
}

func (r *Mock) UnMockMeetingRoomGetBuildingList() {
	r.mockMeetingRoomGetBuildingList = nil
}

type GetBuildingListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 请求期望返回的建筑物数量，不足则返回全部，该值默认为 10，最大为 100
	PageToken *string `query:"page_token" json:"-"` // 用于标记当前请求的分页标记，将返回以当前分页标记开始，往后 page_size 个元素
	OrderBy   *string `query:"order_by" json:"-"`   // 提供用于对名称进行升序/降序排序的方式查询，可选项有："name-asc,name-desc"，传入其他字符串不做处理，默认无序
	Fields    *string `query:"fields" json:"-"`     // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,floors"，默认返回所有字段
}

type getBuildingListResp struct {
	Code int64                `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetBuildingListResp `json:"data,omitempty"` // 返回业务信息
}

type GetBuildingListResp struct {
	PageToken string                        `json:"page_token,omitempty"` // 分页标记，存在下一页时返回
	HasMore   bool                          `json:"has_more,omitempty"`   // 存在下一页时，该值为 true，否则为 false
	Buildings *GetBuildingListRespBuildings `json:"buildings,omitempty"`  // 建筑列表
}

type GetBuildingListRespBuildings struct {
	BuildingID  string   `json:"building_id,omitempty"` // 建筑物 ID
	Description string   `json:"description,omitempty"` // 建筑物的相关描述
	Floors      []string `json:"floors,omitempty"`      // 属于当前建筑物的所有楼层列表
	Name        string   `json:"name,omitempty"`        // 建筑物名称
}
