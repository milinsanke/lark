// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateAttendanceUserStatisticsSettings
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/update-user-stats-settings
func (r *AttendanceService) UpdateAttendanceUserStatisticsSettings(ctx context.Context, request *UpdateAttendanceUserStatisticsSettingsReq, options ...MethodOptionFunc) (*UpdateAttendanceUserStatisticsSettingsResp, *Response, error) {
	if r.cli.mock.mockAttendanceUpdateAttendanceUserStatisticsSettings != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#UpdateAttendanceUserStatisticsSettings mock enable")
		return r.cli.mock.mockAttendanceUpdateAttendanceUserStatisticsSettings(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "UpdateAttendanceUserStatisticsSettings",
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateAttendanceUserStatisticsSettingsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceUpdateAttendanceUserStatisticsSettings(f func(ctx context.Context, request *UpdateAttendanceUserStatisticsSettingsReq, options ...MethodOptionFunc) (*UpdateAttendanceUserStatisticsSettingsResp, *Response, error)) {
	r.mockAttendanceUpdateAttendanceUserStatisticsSettings = f
}

func (r *Mock) UnMockAttendanceUpdateAttendanceUserStatisticsSettings() {
	r.mockAttendanceUpdateAttendanceUserStatisticsSettings = nil
}

type UpdateAttendanceUserStatisticsSettingsReq struct {
	EmployeeType    EmployeeType                                   `query:"employee_type" json:"-"`     // 用户 ID 类型, 可选值有: `employee_id`：用户员工 ID, `employee_no`：用户员工工号
	UserStatsViewID string                                         `path:"user_stats_view_id" json:"-"` // 用户视图 ID, 示例值："TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09"
	View            *UpdateAttendanceUserStatisticsSettingsReqView `json:"view,omitempty"`              // 统计视图
}

type UpdateAttendanceUserStatisticsSettingsReqView struct {
	ViewID    string                                               `json:"view_id,omitempty"`    // 视图 ID, 示例值："TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
	StatsType string                                               `json:"stats_type,omitempty"` // 统计类型, 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                               `json:"user_id,omitempty"`    // 用户 ID
	Items     []*UpdateAttendanceUserStatisticsSettingsReqViewItem `json:"items,omitempty"`      // 一级标题
}

type UpdateAttendanceUserStatisticsSettingsReqViewItem struct {
	Code       string                                                        `json:"code,omitempty"`        // 编号, 示例值："501"
	Title      *string                                                       `json:"title,omitempty"`       // 标题名称, 示例值："基本信息"
	ChildItems []*UpdateAttendanceUserStatisticsSettingsReqViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

type UpdateAttendanceUserStatisticsSettingsReqViewItemChildItem struct {
	Code  string `json:"code,omitempty"`  // 标题编号, 示例值："50101"
	Value string `json:"value,omitempty"` // 开关字段,      , 可选值有: `0`：关闭, `1`：开启,非开关字段场景,  code = 51501  **可选值为1～6**
}

type updateAttendanceUserStatisticsSettingsResp struct {
	Code int64                                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAttendanceUserStatisticsSettingsResp `json:"data,omitempty"`
}

type UpdateAttendanceUserStatisticsSettingsResp struct {
	View *UpdateAttendanceUserStatisticsSettingsRespView `json:"view,omitempty"` // 用户视图
}

type UpdateAttendanceUserStatisticsSettingsRespView struct {
	ViewID    string                                                `json:"view_id,omitempty"`    // 统计视图 ID, 示例值："TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
	StatsType string                                                `json:"stats_type,omitempty"` // 统计类型, 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                                `json:"user_id,omitempty"`    // 用户 ID
	Items     []*UpdateAttendanceUserStatisticsSettingsRespViewItem `json:"items,omitempty"`      // 一级标题
}

type UpdateAttendanceUserStatisticsSettingsRespViewItem struct {
	Code       string                                                         `json:"code,omitempty"`        // 标题编码
	Title      string                                                         `json:"title,omitempty"`       // 标题名称
	ChildItems []*UpdateAttendanceUserStatisticsSettingsRespViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

type UpdateAttendanceUserStatisticsSettingsRespViewItemChildItem struct {
	Code  string `json:"code,omitempty"`  // 标题编号
	Value string `json:"value,omitempty"` // 是否开启,      , 可选值有: `0`：关闭, `1`：开启
}
