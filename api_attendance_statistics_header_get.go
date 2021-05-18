// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetStatisticsHeader
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-header
func (r *AttendanceService) GetStatisticsHeader(ctx context.Context, request *GetStatisticsHeaderReq, options ...MethodOptionFunc) (*GetStatisticsHeaderResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetStatisticsHeader != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetStatisticsHeader mock enable")
		return r.cli.mock.mockAttendanceGetStatisticsHeader(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Attendance#GetStatisticsHeader call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetStatisticsHeader request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_fields/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getStatisticsHeaderResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#GetStatisticsHeader POST https://open.feishu.cn/open-apis/attendance/v1/user_stats_fields/query failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#GetStatisticsHeader POST https://open.feishu.cn/open-apis/attendance/v1/user_stats_fields/query failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "GetStatisticsHeader", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetStatisticsHeader success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceGetStatisticsHeader(f func(ctx context.Context, request *GetStatisticsHeaderReq, options ...MethodOptionFunc) (*GetStatisticsHeaderResp, *Response, error)) {
	r.mockAttendanceGetStatisticsHeader = f
}

func (r *Mock) UnMockAttendanceGetStatisticsHeader() {
	r.mockAttendanceGetStatisticsHeader = nil
}

type GetStatisticsHeaderReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 用户 ID 类型, 可选值有: `employee_id`, `employee_no`
	Locale       string       `json:"locale,omitempty"`        // 语言类型, 可选值有: `en`：英文, `ja`：日文, `zh`：中文
	StatsType    string       `json:"stats_type,omitempty"`    // 统计类型,      , **可选值有**：     , `daily`：日度统计, `month`：月度统计
	StartDate    int64        `json:"start_date,omitempty"`    // 开始时间, 示例值：20210316,      ,      （时间间隔不超过 40 天）
	EndDate      int64        `json:"end_date,omitempty"`      // 结束时间, 示例值：20210323
}

type getStatisticsHeaderResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetStatisticsHeaderResp `json:"data,omitempty"` //
}

type GetStatisticsHeaderResp struct {
	UserStatsField *GetStatisticsHeaderRespUserStatsField `json:"user_stats_field,omitempty"` // 统计数据表头
}

type GetStatisticsHeaderRespUserStatsField struct {
	StatsType string                                        `json:"stats_type,omitempty"` // 统计类型,    , 可选值有: `daily`：日度统计, `month`：月度统计
	UserID    string                                        `json:"user_id,omitempty"`    // 用户 ID
	Fields    []*GetStatisticsHeaderRespUserStatsFieldField `json:"fields,omitempty"`     // 字段列表
}

type GetStatisticsHeaderRespUserStatsFieldField struct {
	Code        string                                                  `json:"code,omitempty"`         // 字段编号
	Title       string                                                  `json:"title,omitempty"`        // 字段标题
	ChildFields []*GetStatisticsHeaderRespUserStatsFieldFieldChildField `json:"child_fields,omitempty"` // 子字段列表
}

type GetStatisticsHeaderRespUserStatsFieldFieldChildField struct {
	Code     string `json:"code,omitempty"`      // 字段编号
	Title    string `json:"title,omitempty"`     // 字段名称
	TimeUnit string `json:"time_unit,omitempty"` // 时间类型
}
