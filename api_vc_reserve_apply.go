// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// ApplyVCReserve 创建一个会议预约。
//
// 支持预约最近30天内的会议（到期时间距离当前时间不超过30天），预约到期后会议号将被释放，如需继续使用可通过"更新预约"接口进行续期；预约会议时可配置参会人在会中的权限，以达到控制会议的目的
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/apply
func (r *VCService) ApplyVCReserve(ctx context.Context, request *ApplyVCReserveReq, options ...MethodOptionFunc) (*ApplyVCReserveResp, *Response, error) {
	if r.cli.mock.mockVCApplyVCReserve != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#ApplyVCReserve mock enable")
		return r.cli.mock.mockVCApplyVCReserve(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "ApplyVCReserve",
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/reserves/apply",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(applyVCReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockVCApplyVCReserve(f func(ctx context.Context, request *ApplyVCReserveReq, options ...MethodOptionFunc) (*ApplyVCReserveResp, *Response, error)) {
	r.mockVCApplyVCReserve = f
}

func (r *Mock) UnMockVCApplyVCReserve() {
	r.mockVCApplyVCReserve = nil
}

type ApplyVCReserveReq struct {
	UserIDType      *IDType                           `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	EndTime         *string                           `json:"end_time,omitempty"`         // 预约到期时间（unix时间，单位sec）, 示例值："1608888867"
	MeetingSettings *ApplyVCReserveReqMeetingSettings `json:"meeting_settings,omitempty"` // 会议设置
}

type ApplyVCReserveReqMeetingSettings struct {
	Topic              *string                                             `json:"topic,omitempty"`                // 会议主题, 示例值："my meeting"
	ActionPermissions  []*ApplyVCReserveReqMeetingSettingsActionPermission `json:"action_permissions,omitempty"`   // 会议权限配置列表，如果存在相同的权限配置项则它们之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
	MeetingInitialType *int64                                              `json:"meeting_initial_type,omitempty"` // 会议初始类型, 示例值：1, 可选值有: `1`：多人会议, `2`：1v1呼叫
	CallSetting        *ApplyVCReserveReqMeetingSettingsCallSetting        `json:"call_setting,omitempty"`         // 1v1呼叫相关参数
}

type ApplyVCReserveReqMeetingSettingsActionPermission struct {
	Permission         int64                                                                `json:"permission,omitempty"`          // 权限项, 示例值：1, 可选值有: `1`：是否能成为主持人, `2`：是否能邀请参会人, `3`：是否能加入会议
	PermissionCheckers []*ApplyVCReserveReqMeetingSettingsActionPermissionPermissionChecker `json:"permission_checkers,omitempty"` // 权限检查器列表，权限检查器之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
}

type ApplyVCReserveReqMeetingSettingsActionPermissionPermissionChecker struct {
	CheckField int64    `json:"check_field,omitempty"` // 检查字段类型, 示例值：1, 可选值有: `1`：用户ID, `2`：用户类型, `3`：租户ID
	CheckMode  int64    `json:"check_mode,omitempty"`  // 检查方式, 示例值：1, 可选值有: `1`：在check_list中为有权限（白名单）, `2`：不在check_list中为有权限（黑名单）
	CheckList  []string `json:"check_list,omitempty"`  // 检查字段列表
}

type ApplyVCReserveReqMeetingSettingsCallSetting struct {
	Callee *ApplyVCReserveReqMeetingSettingsCallSettingCallee `json:"callee,omitempty"` // 被呼叫的用户
}

type ApplyVCReserveReqMeetingSettingsCallSettingCallee struct {
	ID          *string                                                       `json:"id,omitempty"`            // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType    int64                                                         `json:"user_type,omitempty"`     // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
	PstnSipInfo *ApplyVCReserveReqMeetingSettingsCallSettingCalleePstnSipInfo `json:"pstn_sip_info,omitempty"` // pstn/sip信息
}

type ApplyVCReserveReqMeetingSettingsCallSettingCalleePstnSipInfo struct {
	Nickname    *string `json:"nickname,omitempty"`     // 给pstn/sip用户设置的临时昵称, 示例值："dodo"
	MainAddress string  `json:"main_address,omitempty"` // pstn/sip主机号, 示例值："1234"
}

type applyVCReserveResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *ApplyVCReserveResp `json:"data,omitempty"`
}

type ApplyVCReserveResp struct {
	Reserve *ApplyVCReserveRespReserve `json:"reserve,omitempty"` // 预约数据
}

type ApplyVCReserveRespReserve struct {
	ID        string `json:"id,omitempty"`         // 预约ID
	MeetingNo string `json:"meeting_no,omitempty"` // 9位会议号
	URL       string `json:"url,omitempty"`        // 会议链接
	AppLink   string `json:"app_link,omitempty"`   // APPLink用于唤起飞书APP入会。"{?}"为占位符，用于配置入会参数，使用时需替换具体值：0表示关闭，1表示打开。preview为入会前的设置页，mic为麦克风，speaker为扬声器，camera为摄像头
	EndTime   string `json:"end_time,omitempty"`   // 预约到期时间（unix时间，单位sec）
}
