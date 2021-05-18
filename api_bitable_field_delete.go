// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteBitableField 该接口用于在数据表中删除一个字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete
func (r *BitableService) DeleteBitableField(ctx context.Context, request *DeleteBitableFieldReq, options ...MethodOptionFunc) (*DeleteBitableFieldResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableField mock enable")
		return r.cli.mock.mockBitableDeleteBitableField(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Bitable#DeleteBitableField call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableField request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "DELETE",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(deleteBitableFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#DeleteBitableField DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#DeleteBitableField DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "DeleteBitableField", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableField success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableDeleteBitableField(f func(ctx context.Context, request *DeleteBitableFieldReq, options ...MethodOptionFunc) (*DeleteBitableFieldResp, *Response, error)) {
	r.mockBitableDeleteBitableField = f
}

func (r *Mock) UnMockBitableDeleteBitableField() {
	r.mockBitableDeleteBitableField = nil
}

type DeleteBitableFieldReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值："tblsRc9GRRXKqhvW"
	FieldID  string `path:"field_id" json:"-"`  // field id, 示例值："fldPTb0U2y"
}

type deleteBitableFieldResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteBitableFieldResp `json:"data,omitempty"` //
}

type DeleteBitableFieldResp struct {
	FieldID string `json:"field_id,omitempty"` // field id
	Deleted bool   `json:"deleted,omitempty"`  // 删除标记
}
