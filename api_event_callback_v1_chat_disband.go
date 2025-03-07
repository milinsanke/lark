// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV1ChatDisband
//
// ##  解散群
// 群聊被解散后触发此事件。
// * 特殊说明：只有开启机器人能力并且机器人所在的群被解散时才能触发此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukjNxYjL5YTM24SO2EjN
func (r *EventCallbackService) HandlerEventV1ChatDisband(f eventV1ChatDisbandHandler) {
	r.cli.eventHandler.eventV1ChatDisbandHandler = f
}

type eventV1ChatDisbandHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ChatDisband) (string, error)

type EventV1ChatDisband struct {
	AppID     string                      `json:"app_id,omitempty"`  // 如: cli_9c8609450f78d102
	ChatID    string                      `json:"chat_id,omitempty"` // 如: oc_9f2df3c095c9395334bb6e70ced0fa83
	Operator  *EventV1ChatDisbandOperator `json:"operator,omitempty"`
	TenantKey string                      `json:"tenant_key,omitempty"` // 如: 736588c9260f175d
	Type      string                      `json:"type,omitempty"`       // 如: chat_disband
}

type EventV1ChatDisbandOperator struct {
	OpenID string `json:"open_id,omitempty"` // 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	UserID string `json:"user_id,omitempty"` // 如: ca51d83b
}
