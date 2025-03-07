// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHireNote 根据备注 ID 获取备注信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/get
func (r *HireService) GetHireNote(ctx context.Context, request *GetHireNoteReq, options ...MethodOptionFunc) (*GetHireNoteResp, *Response, error) {
	if r.cli.mock.mockHireGetHireNote != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireNote mock enable")
		return r.cli.mock.mockHireGetHireNote(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireNote",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/notes/:note_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireNoteResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireGetHireNote(f func(ctx context.Context, request *GetHireNoteReq, options ...MethodOptionFunc) (*GetHireNoteResp, *Response, error)) {
	r.mockHireGetHireNote = f
}

func (r *Mock) UnMockHireGetHireNote() {
	r.mockHireGetHireNote = nil
}

type GetHireNoteReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	NoteID     string  `path:"note_id" json:"-"`       // Note ID, 示例值："6950620009265891614"
}

type getHireNoteResp struct {
	Code int64            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetHireNoteResp `json:"data,omitempty"`
}

type GetHireNoteResp struct {
	Note *GetHireNoteRespNote `json:"note,omitempty"` // 备注数据
}

type GetHireNoteRespNote struct {
	ID            string `json:"id,omitempty"`             // ID备注
	TalentID      string `json:"talent_id,omitempty"`      // 人才ID
	ApplicationID string `json:"application_id,omitempty"` // 投递ID
	IsPrivate     bool   `json:"is_private,omitempty"`     // 是否私密
	CreateTime    int64  `json:"create_time,omitempty"`    // 创建时间
	ModifyTime    int64  `json:"modify_time,omitempty"`    // 更新时间
	CreatorID     string `json:"creator_id,omitempty"`     // 创建人ID
	Content       string `json:"content,omitempty"`        // 内容
}
