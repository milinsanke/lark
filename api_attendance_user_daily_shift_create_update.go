// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateUpdateAttendanceUserDailyShift
//
// 班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//CreateandEditShifts
func (r *AttendanceService) CreateUpdateAttendanceUserDailyShift(ctx context.Context, request *CreateUpdateAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*CreateUpdateAttendanceUserDailyShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateUpdateAttendanceUserDailyShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateUpdateAttendanceUserDailyShift mock enable")
		return r.cli.mock.mockAttendanceCreateUpdateAttendanceUserDailyShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateUpdateAttendanceUserDailyShift",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUpdateAttendanceUserDailyShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceCreateUpdateAttendanceUserDailyShift(f func(ctx context.Context, request *CreateUpdateAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*CreateUpdateAttendanceUserDailyShiftResp, *Response, error)) {
	r.mockAttendanceCreateUpdateAttendanceUserDailyShift = f
}

func (r *Mock) UnMockAttendanceCreateUpdateAttendanceUserDailyShift() {
	r.mockAttendanceCreateUpdateAttendanceUserDailyShift = nil
}

type CreateUpdateAttendanceUserDailyShiftReq struct {
	EmployeeType    EmployeeType                                             `query:"employee_type" json:"-"`     // 请求体中的 user_id 的员工工号类型可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserDailyShifts []*CreateUpdateAttendanceUserDailyShiftReqUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

type CreateUpdateAttendanceUserDailyShiftReqUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID，休息为 0
	Month   int64  `json:"month,omitempty"`    // 月份
	UserID  string `json:"user_id,omitempty"`  // 用户
	DayNo   int64  `json:"day_no,omitempty"`   // 日期
}

type createUpdateAttendanceUserDailyShiftResp struct {
	Code int64                                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                    `json:"msg,omitempty"`  // 错误描述
	Data *CreateUpdateAttendanceUserDailyShiftResp `json:"data,omitempty"` // -
}

type CreateUpdateAttendanceUserDailyShiftResp struct {
	UserDailyShifts []*CreateUpdateAttendanceUserDailyShiftRespUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

type CreateUpdateAttendanceUserDailyShiftRespUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID
	Month   int64  `json:"month,omitempty"`    // 月份
	UserID  string `json:"user_id,omitempty"`  // 用户
	DayNo   int64  `json:"day_no,omitempty"`   // 日期
}
