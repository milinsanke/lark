// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// JoinChat 用户或机器人主动加入群聊。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 目前仅支持加入公开群
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/me_join
func (r *ChatService) JoinChat(ctx context.Context, request *JoinChatReq, options ...MethodOptionFunc) (*JoinChatResp, *Response, error) {
	if r.cli.mock.mockChatJoinChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#JoinChat mock enable")
		return r.cli.mock.mockChatJoinChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "JoinChat",
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/me_join",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(joinChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatJoinChat(f func(ctx context.Context, request *JoinChatReq, options ...MethodOptionFunc) (*JoinChatResp, *Response, error)) {
	r.mockChatJoinChat = f
}

func (r *Mock) UnMockChatJoinChat() {
	r.mockChatJoinChat = nil
}

type JoinChatReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群 ID, 示例值："oc_a0553eda9014c201e6969b478895c230"
}

type joinChatResp struct {
	Code int64         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string        `json:"msg,omitempty"`  // 错误描述
	Data *JoinChatResp `json:"data,omitempty"`
}

type JoinChatResp struct{}
