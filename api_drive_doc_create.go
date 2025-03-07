// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateDriveDoc
//
// :::note
// 在使用此接口前，请仔细阅读[概述](https://open.feishu.cn/document/ukTMukTMukTM/ukjM5YjL5ITO24SOykjN)和[准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。
// 文档数据结构定义可参考：[文档数据结构概述](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)
// :::
// 该接口用于创建并初始化文档。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN
func (r *DriveService) CreateDriveDoc(ctx context.Context, request *CreateDriveDocReq, options ...MethodOptionFunc) (*CreateDriveDocResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveDoc != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveDoc mock enable")
		return r.cli.mock.mockDriveCreateDriveDoc(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveDoc",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/doc/v2/create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveDocResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCreateDriveDoc(f func(ctx context.Context, request *CreateDriveDocReq, options ...MethodOptionFunc) (*CreateDriveDocResp, *Response, error)) {
	r.mockDriveCreateDriveDoc = f
}

func (r *Mock) UnMockDriveCreateDriveDoc() {
	r.mockDriveCreateDriveDoc = nil
}

type CreateDriveDocReq struct {
	FolderToken *string `json:"FolderToken,omitempty"` // 文件夹 token，获取方式见[准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)；空表示根目录
	Content     *string `json:"Content,omitempty"`     // 传入符合[文档数据结构](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)的字符串，若为空表示创建空文档
}

type createDriveDocResp struct {
	Code int64               `json:"code,omitempty"`
	Msg  string              `json:"msg,omitempty"`
	Data *CreateDriveDocResp `json:"data,omitempty"`
}

type CreateDriveDocResp struct {
	ObjToken string `json:"objToken,omitempty"` // 新建文档的token
	URL      string `json:"url,omitempty"`      // 新建文档的访问链接
}
