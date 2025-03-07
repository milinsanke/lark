// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CopyDriveFile 该接口用于根据文件 token 复制 Doc 或 Sheet  到目标文件夹中。
//
// 若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。复制的文档将会在「我的空间」的「归我所有」列表里。
// **注意**：该接口不支持并发调用，且调用频率上限为 5QPS 且 10000 次每天
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTNzUjL2UzM14iN1MTN
func (r *DriveService) CopyDriveFile(ctx context.Context, request *CopyDriveFileReq, options ...MethodOptionFunc) (*CopyDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveCopyDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CopyDriveFile mock enable")
		return r.cli.mock.mockDriveCopyDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "CopyDriveFile",
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/drive/explorer/v2/file/copy/files/:fileToken",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(copyDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCopyDriveFile(f func(ctx context.Context, request *CopyDriveFileReq, options ...MethodOptionFunc) (*CopyDriveFileResp, *Response, error)) {
	r.mockDriveCopyDriveFile = f
}

func (r *Mock) UnMockDriveCopyDriveFile() {
	r.mockDriveCopyDriveFile = nil
}

type CopyDriveFileReq struct {
	FileToken      string `path:"fileToken" json:"-"`       // 需要复制的源文件的 token, 获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction)
	Type           string `json:"type,omitempty"`           // 需要复制的文档类型，"doc" or "sheet"
	DstFolderToken string `json:"dstFolderToken,omitempty"` // 目标文件夹的 token, 获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction)
	DstName        string `json:"dstName,omitempty"`        // 复制的副本文件的新名称
	CommentNeeded  *bool  `json:"commentNeeded,omitempty"`  // 是否复制评论
}

type copyDriveFileResp struct {
	Code int64              `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *CopyDriveFileResp `json:"data,omitempty"`
}

type CopyDriveFileResp struct {
	FolderToken string `json:"folderToken,omitempty"` // 目标文件夹的 token
	Revision    int64  `json:"revision,omitempty"`    // 新创建文档的版本号
	Token       string `json:"token,omitempty"`       // 新创建文档的 token
	Type        string `json:"type,omitempty"`        // 新建文档的类型，"doc" or "sheet"
	URL         string `json:"url,omitempty"`         // 新创建文档的 url
}
