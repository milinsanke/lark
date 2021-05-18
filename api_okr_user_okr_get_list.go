// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetUserOKRList 根据用户的id获取OKR列表
//
// 使用<md-tag mode="inline" type="token-tenant">tenant_access_token</md-tag>需要额外申请权限以应用身份访问OKR信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/user-okr/list
func (r *OKRService) GetUserOKRList(ctx context.Context, request *GetUserOKRListReq, options ...MethodOptionFunc) (*GetUserOKRListResp, *Response, error) {
	if r.cli.mock.mockOKRGetUserOKRList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetUserOKRList mock enable")
		return r.cli.mock.mockOKRGetUserOKRList(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] OKR#GetUserOKRList call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetUserOKRList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/okr/v1/users/:user_id/okrs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(getUserOKRListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] OKR#GetUserOKRList GET https://open.feishu.cn/open-apis/okr/v1/users/:user_id/okrs failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] OKR#GetUserOKRList GET https://open.feishu.cn/open-apis/okr/v1/users/:user_id/okrs failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("OKR", "GetUserOKRList", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetUserOKRList success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockOKRGetUserOKRList(f func(ctx context.Context, request *GetUserOKRListReq, options ...MethodOptionFunc) (*GetUserOKRListResp, *Response, error)) {
	r.mockOKRGetUserOKRList = f
}

func (r *Mock) UnMockOKRGetUserOKRList() {
	r.mockOKRGetUserOKRList = nil
}

type GetUserOKRListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	Offset     string  `query:"offset" json:"-"`       // 请求列表的偏移，offset>=0，请求Query中, 示例值："0"
	Limit      string  `query:"limit" json:"-"`        // 请求列表的长度，0<limit<=10，请求Query中, 示例值："0"
	Lang       *string `query:"lang" json:"-"`         // 请求OKR的语言版本（比如@的人名），lang=en_us/zh_cn，请求 Query中, 示例值："zh_cn", 默认值: `zh_cn`
	UserID     string  `path:"user_id" json:"-"`       // 目标用户id, 示例值："ou-asdasdasdasdasd"
}

type getUserOKRListResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetUserOKRListResp `json:"data,omitempty"` //
}

type GetUserOKRListResp struct {
	Total   int64                    `json:"total,omitempty"`    // OKR周期总数
	OkrList []*GetUserOKRListRespOkr `json:"okr_list,omitempty"` // OKR 列表
}

type GetUserOKRListRespOkr struct {
	ID            string                            `json:"id,omitempty"`             // id
	Permission    int64                             `json:"permission,omitempty"`     // OKR的访问权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	Name          string                            `json:"name,omitempty"`           // 名称
	ObjectiveList []*GetUserOKRListRespOkrObjective `json:"objective_list,omitempty"` // Objective列表
}

type GetUserOKRListRespOkrObjective struct {
	ID                    string                                             `json:"id,omitempty"`                      // Objective ID
	Permission            int64                                              `json:"permission,omitempty"`              // 权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	Content               string                                             `json:"content,omitempty"`                 // Objective 内容
	ProgressReport        string                                             `json:"progress_report,omitempty"`         // Objective 进度记录内容
	Score                 int64                                              `json:"score,omitempty"`                   // Objective 分数（0 - 100）
	Weight                float64                                            `json:"weight,omitempty"`                  // Objective的权重（0 - 100）
	ProgressRate          *GetUserOKRListRespOkrObjectiveProgressRate        `json:"progress_rate,omitempty"`           // Objective进度
	KrList                []*GetUserOKRListRespOkrObjectiveKr                `json:"kr_list,omitempty"`                 // Objective KeyResult 列表
	AlignedObjectiveList  []*GetUserOKRListRespOkrObjectiveAlignedObjective  `json:"aligned_objective_list,omitempty"`  // 对齐到该Objective的Objective列表
	AligningObjectiveList []*GetUserOKRListRespOkrObjectiveAligningObjective `json:"aligning_objective_list,omitempty"` // 该Objective对齐到的Objective列表
}

type GetUserOKRListRespOkrObjectiveProgressRate struct {
	Percent int64  `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

type GetUserOKRListRespOkrObjectiveKr struct {
	ID           string                                        `json:"id,omitempty"`            // Key Result ID
	Content      string                                        `json:"content,omitempty"`       // KeyResult 内容
	Score        int64                                         `json:"score,omitempty"`         // KeyResult打分（0 - 100）
	Weight       int64                                         `json:"weight,omitempty"`        // KeyResult权重（0 - 100）（废弃）
	KrWeight     float64                                       `json:"kr_weight,omitempty"`     // KeyResult的权重（0 - 100）
	ProgressRate *GetUserOKRListRespOkrObjectiveKrProgressRate `json:"progress_rate,omitempty"` // KR进度
}

type GetUserOKRListRespOkrObjectiveKrProgressRate struct {
	Percent int64  `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

type GetUserOKRListRespOkrObjectiveAlignedObjective struct {
	ID    string                                               `json:"id,omitempty"`     // Objective的ID
	OkrID string                                               `json:"okr_id,omitempty"` // OKR的ID
	Owner *GetUserOKRListRespOkrObjectiveAlignedObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

type GetUserOKRListRespOkrObjectiveAlignedObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}

type GetUserOKRListRespOkrObjectiveAligningObjective struct {
	ID    string                                                `json:"id,omitempty"`     // Objective的ID
	OkrID string                                                `json:"okr_id,omitempty"` // OKR的ID
	Owner *GetUserOKRListRespOkrObjectiveAligningObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

type GetUserOKRListRespOkrObjectiveAligningObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}
