// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateUpdateAttendanceGroup
//
// 考勤组，是对部门或者员工在某个特定场所及特定时间段内的出勤情况（包括上下班、迟到、早退、病假、婚假、丧假、公休、工作时间、加班情况等）的一种规则设定。
// 通过设置考勤组，可以从部门、员工两个维度，来设定考勤方式、考勤时间、考勤地点等考勤规则。
// 出于安全考虑，目前通过该接口只允许修改自己创建的考勤组。
// ### 考勤组负责人
// 考勤组负责人可修改该考勤组的排班，并查看该考勤组的考勤数据。
// 如果考勤组负责人同时被企业管理员赋予了考勤管理员的角色，则该考勤组负责人还拥有考勤管理员的权限，可以编辑及删除考勤规则。
// ### 考勤组人员
// 可按部门、员工两个维度，设置需要参加考勤或无需参加考勤的人员。
// - 若是按部门维度添加的考勤人员，当有新员工加入该部门时，其会自动加入该考勤组。
// - 若是按员工维度添加的考勤人员，当其上级部门被添加到其他考勤组时，该员工不会更换考勤组。
// ### 考勤组类型
// 提供 3 种不同的考勤类型：固定班制、排班制、自由班制。
// - 固定班制：指考勤组内每位人员的上下班时间一致，适用于上下班时间固定或无需安排多个班次的考勤组。
// - 排班制：指考勤组人员的上下班时间不完全一致，可自定义安排每位人员的上下班时间，适用于存在多个班次如早晚班的考勤组。
// - 自由班制：指没有具体的班次，考勤组人员可以在打卡时段内自由打卡，按照打卡时段统计上班时长。
// ### 考勤班次
// - 固定班制下，需设置周一到周日每天安排哪个班次，以及可针对特殊日期进行打卡设置。
// - 排班制下，需对考勤组内每一位人员的每一天进行班次指定。
// - 自由班制下，需设置一天中最早打卡时间和最晚打卡时间，以及一周中哪几天需要打卡。
// ### 考勤方式
// 支持 3 种考勤方式：GPS 打卡、Wi-Fi 打卡、考勤机打卡。
// - GPS 打卡：需设置经纬度信息及考勤地点名称。
// - Wi-Fi 打卡：需设置 Wi-Fi 名称及 Wi-Fi 的 MAC 地址。
// - 考勤机打卡：需设置考勤机名称及考勤机序号。
// ### 考勤其他设置
// - 规则设置：支持设置是否允许外勤打卡，是否允许补卡以及一个月补卡的次数，是否允许 PC 端打卡。
// - 安全设置：支持设置是否开启人脸识别打卡，以及什么情况下开启人脸识别。
// - 统计设置：支持设置考勤组人员是否可以查看到某些维度的统计数据。
// - 加班设置：支持配置加班时间的计算规则。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_create_update
func (r *AttendanceService) CreateUpdateAttendanceGroup(ctx context.Context, request *CreateUpdateAttendanceGroupReq, options ...MethodOptionFunc) (*CreateUpdateAttendanceGroupResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateUpdateAttendanceGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateUpdateAttendanceGroup mock enable")
		return r.cli.mock.mockAttendanceCreateUpdateAttendanceGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateUpdateAttendanceGroup",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/groups",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUpdateAttendanceGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceCreateUpdateAttendanceGroup(f func(ctx context.Context, request *CreateUpdateAttendanceGroupReq, options ...MethodOptionFunc) (*CreateUpdateAttendanceGroupResp, *Response, error)) {
	r.mockAttendanceCreateUpdateAttendanceGroup = f
}

func (r *Mock) UnMockAttendanceCreateUpdateAttendanceGroup() {
	r.mockAttendanceCreateUpdateAttendanceGroup = nil
}

type CreateUpdateAttendanceGroupReq struct {
	EmployeeType EmployeeType                         `query:"employee_type" json:"-"` // 用户 ID 的类型，必选字段，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】
	DeptType     string                               `query:"dept_type" json:"-"`     // 部门 ID 的类型，必选字段，可用值：【open_id（暂时只支持部门的 openid）】，示例值：“od-fcb45c28a45311afd441b8869541ece8”
	Group        *CreateUpdateAttendanceGroupReqGroup `json:"group,omitempty"`         // 考勤组
}

type CreateUpdateAttendanceGroupReqGroup struct {
	GroupID                *string                                                     `json:"group_id,omitempty"`                   // 考勤组的 ID, 需要从获取用户打卡结果的接口中获取 groupId
	GroupName              string                                                      `json:"group_name,omitempty"`                 // 考勤组名称
	TimeZone               string                                                      `json:"time_zone,omitempty"`                  // 时区，可参考时区列表 https://www.zeitverschiebung.net/cn/all-time-zones.html
	BindDeptIDs            []string                                                    `json:"bind_dept_ids,omitempty"`              // 绑定的部门 ID
	ExceptDeptIDs          []string                                                    `json:"except_dept_ids,omitempty"`            // 排除的部门 ID
	BindUserIDs            []string                                                    `json:"bind_user_ids,omitempty"`              // 绑定的用户 ID
	ExceptUserIDs          []string                                                    `json:"except_user_ids,omitempty"`            // 排除的用户 ID
	GroupLeaderIDs         []string                                                    `json:"group_leader_ids,omitempty"`           // 考勤负责人 ID 列表，需至少存在一名考勤负责人
	AllowOutPunch          *bool                                                       `json:"allow_out_punch,omitempty"`            // 是否允许外勤打卡
	AllowPcPunch           *bool                                                       `json:"allow_pc_punch,omitempty"`             // 是否允许 PC 端打卡
	AllowRemedy            *bool                                                       `json:"allow_remedy,omitempty"`               // 是否允许补卡
	RemedyLimit            *bool                                                       `json:"remedy_limit,omitempty"`               // 是否限制补卡次数
	RemedyLimitCount       *int64                                                      `json:"remedy_limit_count,omitempty"`         // 补卡次数
	RemedyDateLimit        *bool                                                       `json:"remedy_date_limit,omitempty"`          // 是否限制补卡时间
	RemedyDateNum          *int64                                                      `json:"remedy_date_num,omitempty"`            // 补卡时间
	ShowCumulativeTime     *bool                                                       `json:"show_cumulative_time,omitempty"`       // 是否展示上班累计时长
	ShowOverTime           *bool                                                       `json:"show_over_time,omitempty"`             // 是否展示加班累计时长
	HideStaffPunchTime     *bool                                                       `json:"hide_staff_punch_time,omitempty"`      // 是否隐藏员工打卡具体时间
	FacePunch              *bool                                                       `json:"face_punch,omitempty"`                 // 是否开启人脸识别打卡
	FacePunchCfg           *int64                                                      `json:"face_punch_cfg,omitempty"`             // 人脸识别打卡规则，1：每次打卡均需人脸识别，2：疑似作弊打卡时需要人脸识别
	FaceDowngrade          *bool                                                       `json:"face_downgrade,omitempty"`             // 人脸识别失败时是否允许普通拍照打卡
	ReplaceBasicPic        *bool                                                       `json:"replace_basic_pic,omitempty"`          // 人脸识别失败时是否允许替换基准图片
	Machines               []*CreateUpdateAttendanceGroupReqGroupMachine               `json:"machines,omitempty"`                   // 考勤机列表
	GpsRange               int64                                                       `json:"gps_range,omitempty"`                  // GPS 打卡的有效范围，必选字段
	Locations              []*CreateUpdateAttendanceGroupReqGroupLocation              `json:"locations,omitempty"`                  // 地址列表
	GroupType              int64                                                       `json:"group_type,omitempty"`                 // 考勤类型，0：固定班制，2：排班制，3：自由班制
	PunchDayShiftIDs       []string                                                    `json:"punch_day_shift_ids,omitempty"`        // 固定班制必须填，长度必须等于7
	FreePunchCfg           *CreateUpdateAttendanceGroupReqGroupFreePunchCfg            `json:"free_punch_cfg,omitempty"`             // 配置自由班制
	CalendarID             int64                                                       `json:"calendar_id,omitempty"`                // 国家法定节假日历 ID，0：不根据国家法定节假日历排休，1：中国，2：美国，3：日本，4：印度，5：新加坡，默认为 1，必选字段
	NeedPunchSpecialDays   []*CreateUpdateAttendanceGroupReqGroupNeedPunchSpecialDay   `json:"need_punch_special_days,omitempty"`    // 必须打卡的特殊日期
	NoNeedPunchSpecialDays []*CreateUpdateAttendanceGroupReqGroupNoNeedPunchSpecialDay `json:"no_need_punch_special_days,omitempty"` // 无需打卡的特殊日期
	EffectNow              *bool                                                       `json:"effect_now,omitempty"`                 // 是否立即生效，默认为 false
}

type CreateUpdateAttendanceGroupReqGroupMachine struct {
	MachineSn   string `json:"machine_sn,omitempty"`   // 考勤机序列号
	MachineName string `json:"machine_name,omitempty"` // 考勤机名称
}

type CreateUpdateAttendanceGroupReqGroupLocation struct {
	LocationID   *string  `json:"location_id,omitempty"`   // 地址 ID
	LocationName string   `json:"location_name,omitempty"` // 地址名称，必选字段
	LocationType *int64   `json:"location_type,omitempty"` // 地址类型，1：GPS，2：Wifi，8：IP
	Latitude     *float64 `json:"latitude,omitempty"`      // 地址纬度
	Longitude    *float64 `json:"longitude,omitempty"`     // 地址经度
	Ssid         *string  `json:"ssid,omitempty"`          // Wi-Fi 名称
	Bssid        *string  `json:"bssid,omitempty"`         // Wi-Fi 的 MAC 地址
	MapType      *int64   `json:"map_type,omitempty"`      // 地图类型，1：高德, 2：谷歌
	Address      *string  `json:"address,omitempty"`       // 地址名称
	Ip           *string  `json:"ip,omitempty"`            // IP 地址
	Feature      *string  `json:"feature,omitempty"`       // 额外信息，例如运营商信息
}

type CreateUpdateAttendanceGroupReqGroupFreePunchCfg struct {
	FreeStartTime        string `json:"free_start_time,omitempty"`           // 自由班制的打卡开始时间
	FreeEndTime          string `json:"free_end_time,omitempty"`             // 自由班制的打卡结束时间
	PunchDay             int64  `json:"punch_day,omitempty"`                 // 打卡时间，格式 1111100
	WorkDayNoPunchAsLack *bool  `json:"work_day_no_punch_as_lack,omitempty"` // 工作日不打卡是否记为缺卡
}

type CreateUpdateAttendanceGroupReqGroupNeedPunchSpecialDay struct {
	PunchDay *int64 `json:"punch_day,omitempty"` // 打卡日期，格式 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

type CreateUpdateAttendanceGroupReqGroupNoNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期，格式 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

type createUpdateAttendanceGroupResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateUpdateAttendanceGroupResp `json:"data,omitempty"` // -
}

type CreateUpdateAttendanceGroupResp struct {
	Group *CreateUpdateAttendanceGroupRespGroup `json:"group,omitempty"` // 考勤组
}

type CreateUpdateAttendanceGroupRespGroup struct {
	GroupID                string                                                       `json:"group_id,omitempty"`                   // 考勤组的 ID, 需要从获取用户打卡结果的接口中获取 groupId
	GroupName              string                                                       `json:"group_name,omitempty"`                 // 考勤组名称
	TimeZone               string                                                       `json:"time_zone,omitempty"`                  // 时区
	BindDeptIDs            []string                                                     `json:"bind_dept_ids,omitempty"`              // 绑定的部门 ID
	ExceptDeptIDs          []string                                                     `json:"except_dept_ids,omitempty"`            // 排除的部门 ID
	BindUserIDs            []string                                                     `json:"bind_user_ids,omitempty"`              // 绑定的用户 ID
	ExceptUserIDs          []string                                                     `json:"except_user_ids,omitempty"`            // 排除的用户 ID
	GroupLeaderIDs         []string                                                     `json:"group_leader_ids,omitempty"`           // 考勤负责人 ID 列表，必选字段
	AllowOutPunch          bool                                                         `json:"allow_out_punch,omitempty"`            // 是否允许外勤打卡
	AllowPcPunch           bool                                                         `json:"allow_pc_punch,omitempty"`             // 是否允许 PC 端打卡
	AllowRemedy            bool                                                         `json:"allow_remedy,omitempty"`               // 是否允许补卡
	RemedyLimit            bool                                                         `json:"remedy_limit,omitempty"`               // 是否限制补卡次数
	RemedyLimitCount       int64                                                        `json:"remedy_limit_count,omitempty"`         // 补卡次数
	RemedyDateLimit        bool                                                         `json:"remedy_date_limit,omitempty"`          // 是否限制补卡时间
	RemedyDateNum          int64                                                        `json:"remedy_date_num,omitempty"`            // 补卡时间
	ShowCumulativeTime     bool                                                         `json:"show_cumulative_time,omitempty"`       // 是否展示上班累计时长
	ShowOverTime           bool                                                         `json:"show_over_time,omitempty"`             // 是否展示加班累计时长
	HideStaffPunchTime     bool                                                         `json:"hide_staff_punch_time,omitempty"`      // 是否隐藏员工打卡具体时间
	FacePunch              bool                                                         `json:"face_punch,omitempty"`                 // 是否开启人脸识别打卡
	FacePunchCfg           int64                                                        `json:"face_punch_cfg,omitempty"`             // 人脸识别打卡规则，1：每次打卡均需人脸识别，2：疑似作弊打卡时需要人脸识别
	FaceDowngrade          bool                                                         `json:"face_downgrade,omitempty"`             // 人脸识别失败时是否允许普通拍照打卡
	ReplaceBasicPic        bool                                                         `json:"replace_basic_pic,omitempty"`          // 人脸识别失败时是否允许替换基准图片
	Machines               []*CreateUpdateAttendanceGroupRespGroupMachine               `json:"machines,omitempty"`                   // 考勤机列表
	GpsRange               int64                                                        `json:"gps_range,omitempty"`                  // GPS 打卡的有效范围
	Locations              []*CreateUpdateAttendanceGroupRespGroupLocation              `json:"locations,omitempty"`                  // 地址列表
	GroupType              int64                                                        `json:"group_type,omitempty"`                 // 考勤类型，0：固定班制，2：排班制，3：自由班制
	PunchDayShiftIDs       []string                                                     `json:"punch_day_shift_ids,omitempty"`        // 固定班制必须填
	FreePunchCfg           *CreateUpdateAttendanceGroupRespGroupFreePunchCfg            `json:"free_punch_cfg,omitempty"`             // 配置自由班制
	CalendarID             int64                                                        `json:"calendar_id,omitempty"`                // 国家法定节假日历 ID，0：不根据国家法定节假日历排休，1：中国，2：美国，3：日本，4：印度，5：新加坡，默认为 1
	NeedPunchSpecialDays   []*CreateUpdateAttendanceGroupRespGroupNeedPunchSpecialDay   `json:"need_punch_special_days,omitempty"`    // 必须打卡的特殊日期
	NoNeedPunchSpecialDays []*CreateUpdateAttendanceGroupRespGroupNoNeedPunchSpecialDay `json:"no_need_punch_special_days,omitempty"` // 无需打卡的特殊日期
	WorkDayNoPunchAsLack   bool                                                         `json:"work_day_no_punch_as_lack,omitempty"`  // 自由班制下工作日不打卡是否记为缺卡
}

type CreateUpdateAttendanceGroupRespGroupMachine struct {
	MachineSn   string `json:"machine_sn,omitempty"`   // 考勤机序列号
	MachineName string `json:"machine_name,omitempty"` // 考勤机名称
}

type CreateUpdateAttendanceGroupRespGroupLocation struct {
	LocationID   string  `json:"location_id,omitempty"`   // 地址 ID
	LocationName string  `json:"location_name,omitempty"` // 地址名称
	LocationType int64   `json:"location_type,omitempty"` // 地址类型，1：GPS，2：Wifi，8：IP
	Latitude     float64 `json:"latitude,omitempty"`      // 地址纬度
	Longitude    float64 `json:"longitude,omitempty"`     // 地址经度
	Ssid         string  `json:"ssid,omitempty"`          // Wi-Fi 名称
	Bssid        string  `json:"bssid,omitempty"`         // Wi-Fi 的 MAC 地址
	MapType      int64   `json:"map_type,omitempty"`      // 地图类型，1：高德，2：谷歌
	Address      string  `json:"address,omitempty"`       // 地址名称
	Ip           string  `json:"ip,omitempty"`            // IP 地址
	Feature      string  `json:"feature,omitempty"`       // 额外信息，例如运营商信息
}

type CreateUpdateAttendanceGroupRespGroupFreePunchCfg struct {
	FreeStartTime        string `json:"free_start_time,omitempty"`           // 自由班制的打卡开始时间
	FreeEndTime          string `json:"free_end_time,omitempty"`             // 自由班制的打卡结束时间
	PunchDay             int64  `json:"punch_day,omitempty"`                 // 打卡时间，格式 1111100
	WorkDayNoPunchAsLack bool   `json:"work_day_no_punch_as_lack,omitempty"` // 工作日不打卡是否记为缺卡
}

type CreateUpdateAttendanceGroupRespGroupNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期，格式 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

type CreateUpdateAttendanceGroupRespGroupNoNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期，格式 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}
