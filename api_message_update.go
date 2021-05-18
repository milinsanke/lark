// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateMessage 更新应用已发送的消息卡片内容。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 当前仅更新支持卡片消息
// - 只支持对所有人都更新的「共享卡片」。如果你只想更新特定人的消息卡片，必须要用户在卡片操作交互后触发，开发文档参考[「独享卡片」](https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN#49904b71)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch
func (r *MessageService) UpdateMessage(ctx context.Context, request *UpdateMessageReq, options ...MethodOptionFunc) (*UpdateMessageResp, *Response, error) {
	if r.cli.mock.mockMessageUpdateMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#UpdateMessage mock enable")
		return r.cli.mock.mockMessageUpdateMessage(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Message#UpdateMessage call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Message#UpdateMessage request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(updateMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Message#UpdateMessage PATCH https://open.feishu.cn/open-apis/im/v1/messages/:message_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Message#UpdateMessage PATCH https://open.feishu.cn/open-apis/im/v1/messages/:message_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Message", "UpdateMessage", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Message#UpdateMessage success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMessageUpdateMessage(f func(ctx context.Context, request *UpdateMessageReq, options ...MethodOptionFunc) (*UpdateMessageResp, *Response, error)) {
	r.mockMessageUpdateMessage = f
}

func (r *Mock) UnMockMessageUpdateMessage() {
	r.mockMessageUpdateMessage = nil
}

type UpdateMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待更新的消息的ID, 示例值："om_dc13264520392913993dd051dba21dcf"
	Content   string `json:"content,omitempty"`   // 消息内容 json 格式，[发送消息 content 说明](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json)，参考文档中的卡片格式|, 示例值："参考链接"
}

type updateMessageResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateMessageResp `json:"data,omitempty"`
}

type UpdateMessageResp struct{}
