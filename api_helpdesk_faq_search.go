// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SearchFAQ 该接口用于搜索服务台知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/search
func (r *HelpdeskService) SearchFAQ(ctx context.Context, request *SearchFAQReq, options ...MethodOptionFunc) (*SearchFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSearchFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SearchFAQ mock enable")
		return r.cli.mock.mockHelpdeskSearchFAQ(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Helpdesk#SearchFAQ call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SearchFAQ request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedHelpdeskAuth: true,
	}
	resp := new(searchFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#SearchFAQ GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs/search failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#SearchFAQ GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs/search failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "SearchFAQ", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SearchFAQ success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskSearchFAQ(f func(ctx context.Context, request *SearchFAQReq, options ...MethodOptionFunc) (*SearchFAQResp, *Response, error)) {
	r.mockHelpdeskSearchFAQ = f
}

func (r *Mock) UnMockHelpdeskSearchFAQ() {
	r.mockHelpdeskSearchFAQ = nil
}

type SearchFAQReq struct {
	Query     string  `query:"query" json:"-"`      // 搜索query, 示例值："点餐"
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："6936004780707807251"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`100`
}

type searchFAQResp struct {
	Code int64          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *SearchFAQResp `json:"data,omitempty"` //
}

type SearchFAQResp struct {
	HasMore   bool                 `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string               `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*SearchFAQRespItem `json:"items,omitempty"`      // 知识库列表
}

type SearchFAQRespItem struct {
	FaqID          string              `json:"faq_id,omitempty"`          // 知识库ID
	ID             string              `json:"id,omitempty"`              // 知识库旧版ID，请使用faq_id
	HelpdeskID     string              `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string              `json:"question,omitempty"`        // 问题
	Answer         string              `json:"answer,omitempty"`          // 答案
	AnswerRichtext string              `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64               `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64               `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory `json:"categories,omitempty"`      // 分类
	Tags           []string            `json:"tags,omitempty"`            // 关联词列表
	ExpireTime     int64               `json:"expire_time,omitempty"`     // 失效时间
}
