// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// AddChatMember 将用户或机器人拉入群聊。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 如需拉用户进群，需要机器人对用户有可见性
// - 在开启[仅群主可添加群成员]的设置时，仅有群主 或 创建群组且具备[更新应用所创建群的群信息]权限的机器人，可以拉用户或者机器人进群
// - 在未开启[仅群主可添加群成员]的设置时，所有群成员都可以拉用户或机器人进群
// - 每次请求，最多拉50个用户或者5个机器人，并且群组最多容纳15个机器人
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create
func (r *ChatService) AddChatMember(ctx context.Context, request *AddChatMemberReq, options ...MethodOptionFunc) (*AddChatMemberResp, *Response, error) {
	if r.cli.mock.mockChatAddChatMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#AddChatMember mock enable")
		return r.cli.mock.mockChatAddChatMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "AddChatMember",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(addChatMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatAddChatMember(f func(ctx context.Context, request *AddChatMemberReq, options ...MethodOptionFunc) (*AddChatMemberResp, *Response, error)) {
	r.mockChatAddChatMember = f
}

func (r *Mock) UnMockChatAddChatMember() {
	r.mockChatAddChatMember = nil
}

type AddChatMemberReq struct {
	MemberIDType *IDType  `query:"member_id_type" json:"-"` // 进群成员 id 类型 open_id/user_id/union_id/app_id, 示例值："user_id", 可选值有: `user_id`：以 user_id 来识别成员, `union_id`：以 union_id 来识别成员, `open_id`：以 open_id 来识别成员, `app_id`：以 app_id 来识别成员
	ChatID       string   `path:"chat_id" json:"-"`         // 群 ID, 示例值："oc_a0553eda9014c201e6969b478895c230"
	IDList       []string `json:"id_list,omitempty"`        // 成员列表
}

type addChatMemberResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *AddChatMemberResp `json:"data,omitempty"`
}

type AddChatMemberResp struct {
	InvalidIDList []string `json:"invalid_id_list,omitempty"` // 无效成员列表
}
