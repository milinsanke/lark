// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetApplicationUserAdminScope
//
// 该接口用于获取应用管理员的管理范围，即该应用管理员能够管理哪些部门。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN
func (r *ApplicationService) GetApplicationUserAdminScope(ctx context.Context, request *GetApplicationUserAdminScopeReq, options ...MethodOptionFunc) (*GetApplicationUserAdminScopeResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUserAdminScope != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUserAdminScope mock enable")
		return r.cli.mock.mockApplicationGetApplicationUserAdminScope(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUserAdminScope",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v1/user/admin_scope/get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUserAdminScopeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApplicationGetApplicationUserAdminScope(f func(ctx context.Context, request *GetApplicationUserAdminScopeReq, options ...MethodOptionFunc) (*GetApplicationUserAdminScopeResp, *Response, error)) {
	r.mockApplicationGetApplicationUserAdminScope = f
}

func (r *Mock) UnMockApplicationGetApplicationUserAdminScope() {
	r.mockApplicationGetApplicationUserAdminScope = nil
}

type GetApplicationUserAdminScopeReq struct {
	EmployeeID string `query:"employee_id" json:"-"` // 支持通过 open_id 或者 employee_id 查询，不支持混合两种 ID 进行查询，其中 employee_id 同通讯录 v3 版本中的 user_id
	OpenID     string `query:"open_id" json:"-"`     // 支持通过 open_id 或者 employee_id 查询，不支持混合两种 ID 进行查询，其中 employee_id 同通讯录 v3 版本中的 user_id
}

type getApplicationUserAdminScopeResp struct {
	Code int64                             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApplicationUserAdminScopeResp `json:"data,omitempty"` // 返回业务数据
}

type GetApplicationUserAdminScopeResp struct {
	IsAll          bool     `json:"is_all,omitempty"`          // 是否管理所有部门
	DepartmentList []string `json:"department_list,omitempty"` // 管理的部门列表，当 is_all 为 true 时，不返回该字段
}
