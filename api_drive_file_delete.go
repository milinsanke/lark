// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteDriveFile
//
// **本文档包含两个接口，分别用于删除 Doc 和 Sheet，对应的文档类型请调用对应的接口**<br>
// **文档只能被文档所有者删除，文档被删除后将会放到回收站里**
// 该接口用于根据 docToken 删除对应的 Docs 文档。
// **注意**：该接口不支持并发调用，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATM2UjLwEjN14CMxYTN
func (r *DriveService) DeleteDriveFile(ctx context.Context, request *DeleteDriveFileReq, options ...MethodOptionFunc) (*DeleteDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveFile mock enable")
		return r.cli.mock.mockDriveDeleteDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "DeleteDriveFile",
		Method:              "DELETE",
		URL:                 "https://open.feishu.cn/open-apis/drive/explorer/v2/file/docs/:docToken",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(deleteDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveDeleteDriveFile(f func(ctx context.Context, request *DeleteDriveFileReq, options ...MethodOptionFunc) (*DeleteDriveFileResp, *Response, error)) {
	r.mockDriveDeleteDriveFile = f
}

func (r *Mock) UnMockDriveDeleteDriveFile() {
	r.mockDriveDeleteDriveFile = nil
}

type DeleteDriveFileReq struct {
	SpreadSheetToken string `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction)
}

type deleteDriveFileResp struct {
	Code int64                `json:"code,omitempty"`
	Msg  string               `json:"msg,omitempty"`
	Data *DeleteDriveFileResp `json:"data,omitempty"`
}

type DeleteDriveFileResp struct {
	ID     string `json:"id,omitempty"`     // sheet 的 id 「字符串类型」
	Result bool   `json:"result,omitempty"` // 删除结果
}
