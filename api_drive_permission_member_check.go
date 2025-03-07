// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CheckDriveMemberPermission 该接口用于根据 filetoken 判断当前登录用户是否具有某权限。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzN3UjL2czN14iN3cTN
func (r *DriveService) CheckDriveMemberPermission(ctx context.Context, request *CheckDriveMemberPermissionReq, options ...MethodOptionFunc) (*CheckDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveCheckDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CheckDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveCheckDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CheckDriveMemberPermission",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/permitted",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(checkDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCheckDriveMemberPermission(f func(ctx context.Context, request *CheckDriveMemberPermissionReq, options ...MethodOptionFunc) (*CheckDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveCheckDriveMemberPermission = f
}

func (r *Mock) UnMockDriveCheckDriveMemberPermission() {
	r.mockDriveCheckDriveMemberPermission = nil
}

type CheckDriveMemberPermissionReq struct {
	Token string `json:"token,omitempty"` // 文件的 token，获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type  string `json:"type,omitempty"`  // 文档类型  "doc"  or  "sheet" or "file"
	Perm  string `json:"perm,omitempty"`  // 权限，"view" or "edit" or "share"
}

type checkDriveMemberPermissionResp struct {
	Code int64                           `json:"code,omitempty"`
	Msg  string                          `json:"msg,omitempty"`
	Data *CheckDriveMemberPermissionResp `json:"data,omitempty"`
}

type CheckDriveMemberPermissionResp struct {
	IsPermitted bool `json:"is_permitted,omitempty"` // 是否具有指定权限
}
