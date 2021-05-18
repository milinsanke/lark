// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateUserStatisticsSettings
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/update-user-stats-settings
func (r *AttendanceService) UpdateUserStatisticsSettings(ctx context.Context, request *UpdateUserStatisticsSettingsReq, options ...MethodOptionFunc) (*UpdateUserStatisticsSettingsResp, *Response, error) {
	if r.cli.mock.mockAttendanceUpdateUserStatisticsSettings != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#UpdateUserStatisticsSettings mock enable")
		return r.cli.mock.mockAttendanceUpdateUserStatisticsSettings(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Attendance#UpdateUserStatisticsSettings call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#UpdateUserStatisticsSettings request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateUserStatisticsSettingsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#UpdateUserStatisticsSettings PUT https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#UpdateUserStatisticsSettings PUT https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "UpdateUserStatisticsSettings", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#UpdateUserStatisticsSettings success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceUpdateUserStatisticsSettings(f func(ctx context.Context, request *UpdateUserStatisticsSettingsReq, options ...MethodOptionFunc) (*UpdateUserStatisticsSettingsResp, *Response, error)) {
	r.mockAttendanceUpdateUserStatisticsSettings = f
}

func (r *Mock) UnMockAttendanceUpdateUserStatisticsSettings() {
	r.mockAttendanceUpdateUserStatisticsSettings = nil
}

type UpdateUserStatisticsSettingsReq struct {
	EmployeeType    EmployeeType                         `query:"employee_type" json:"-"`     // 用户 ID 类型, 可选值有: `employee_id`：用户员工 ID, `employee_no`：用户员工工号
	UserStatsViewID string                               `path:"user_stats_view_id" json:"-"` // 用户视图 ID, 示例值："TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09"
	View            *UpdateUserStatisticsSettingsReqView `json:"view,omitempty"`              // 统计视图
}

type UpdateUserStatisticsSettingsReqView struct {
	ViewID    string                                     `json:"view_id,omitempty"`    // 视图 ID, 示例值："TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
	StatsType string                                     `json:"stats_type,omitempty"` // 统计类型, 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                     `json:"user_id,omitempty"`    // 用户 ID
	Items     []*UpdateUserStatisticsSettingsReqViewItem `json:"items,omitempty"`      // 一级标题
}

type UpdateUserStatisticsSettingsReqViewItem struct {
	Code       string                                              `json:"code,omitempty"`        // 编号, 示例值："501"
	Title      *string                                             `json:"title,omitempty"`       // 标题名称, 示例值："基本信息"
	ChildItems []*UpdateUserStatisticsSettingsReqViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

type UpdateUserStatisticsSettingsReqViewItemChildItem struct {
	Code  string `json:"code,omitempty"`  // 标题编号, 示例值："50101"
	Value string `json:"value,omitempty"` // 开关字段,      , 可选值有: `0`：关闭, `1`：开启,非开关字段场景,  code = 51501  **可选值为1～6**
}

type updateUserStatisticsSettingsResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *UpdateUserStatisticsSettingsResp `json:"data,omitempty"` //
}

type UpdateUserStatisticsSettingsResp struct {
	View *UpdateUserStatisticsSettingsRespView `json:"view,omitempty"` // 用户视图
}

type UpdateUserStatisticsSettingsRespView struct {
	ViewID    string                                      `json:"view_id,omitempty"`    // 统计视图 ID, 示例值："TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
	StatsType string                                      `json:"stats_type,omitempty"` // 统计类型, 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                      `json:"user_id,omitempty"`    // 用户 ID
	Items     []*UpdateUserStatisticsSettingsRespViewItem `json:"items,omitempty"`      // 一级标题
}

type UpdateUserStatisticsSettingsRespViewItem struct {
	Code       string                                               `json:"code,omitempty"`        // 标题编码
	Title      string                                               `json:"title,omitempty"`       // 标题名称
	ChildItems []*UpdateUserStatisticsSettingsRespViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

type UpdateUserStatisticsSettingsRespViewItemChildItem struct {
	Code  string `json:"code,omitempty"`  // 标题编号
	Value string `json:"value,omitempty"` // 是否开启,      , 可选值有: `0`：关闭, `1`：开启
}
