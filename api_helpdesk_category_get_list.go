// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetCategoryList
//
// 该接口用于获取服务台知识库所有分类。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/list-categories
func (r *HelpdeskService) GetCategoryList(ctx context.Context, request *GetCategoryListReq, options ...MethodOptionFunc) (*GetCategoryListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetCategoryList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetCategoryList mock enable")
		return r.cli.mock.mockHelpdeskGetCategoryList(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Helpdesk#GetCategoryList call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetCategoryList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/categories",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedHelpdeskAuth: true,
	}
	resp := new(getCategoryListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#GetCategoryList GET https://open.feishu.cn/open-apis/helpdesk/v1/categories failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#GetCategoryList GET https://open.feishu.cn/open-apis/helpdesk/v1/categories failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "GetCategoryList", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetCategoryList success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskGetCategoryList(f func(ctx context.Context, request *GetCategoryListReq, options ...MethodOptionFunc) (*GetCategoryListResp, *Response, error)) {
	r.mockHelpdeskGetCategoryList = f
}

func (r *Mock) UnMockHelpdeskGetCategoryList() {
	r.mockHelpdeskGetCategoryList = nil
}

type GetCategoryListReq struct {
	Lang    *string `query:"lang" json:"-"`     // 知识库分类语言, 示例值："zh_cn"
	OrderBy *int64  `query:"order_by" json:"-"` // 排序键。1: 根据知识库分类更新时间排序, 示例值：1
	Asc     *bool   `query:"asc" json:"-"`      // 顺序。true: 正序；false：反序, 示例值：true
}

type getCategoryListResp struct {
	Code int64                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetCategoryListResp `json:"data,omitempty"` //
}

type GetCategoryListResp struct {
	Categories []*HelpdeskCategory `json:"categories,omitempty"` // 知识库分类列表
}
