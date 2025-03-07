// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateApprovalInstance
//
// 创建一个审批实例，调用方需对审批定义的表单有详细了解，将按照定义的表单结构，将表单 Value 通过接口传入。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDNyUjLyQjM14iM0ITN
func (r *ApprovalService) CreateApprovalInstance(ctx context.Context, request *CreateApprovalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalCreateApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#CreateApprovalInstance mock enable")
		return r.cli.mock.mockApprovalCreateApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CreateApprovalInstance",
		Method:                "POST",
		URL:                   "https://www.feishu.cn/approval/openapi/v2/instance/create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApprovalCreateApprovalInstance(f func(ctx context.Context, request *CreateApprovalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalInstanceResp, *Response, error)) {
	r.mockApprovalCreateApprovalInstance = f
}

func (r *Mock) UnMockApprovalCreateApprovalInstance() {
	r.mockApprovalCreateApprovalInstance = nil
}

type CreateApprovalInstanceReq struct {
	ApprovalCode           string              `json:"approval_code,omitempty"`              // 审批定义 code
	UserID                 *string             `json:"user_id,omitempty"`                    // 发起审批用户
	TenantID               *string             `json:"tenant_id,omitempty"`                  // 平台租户ID
	OpenID                 string              `json:"open_id,omitempty"`                    // 发起审批用户 open id, 如果传了 user_id 则优先使用 user_id
	DepartmentID           *string             `json:"department_id,omitempty"`              // 发起审批用户部门，如果用户只属于一个部门，可以不填，如果属于多个部门，必须填其中一个部门
	Form                   ApprovalWidgetList  `json:"form,omitempty"`                       // json 数组，**控件值**
	NodeApproverUserIDList map[string][]string `json:"node_approver_user_id_list,omitempty"` // 如果有发起人自选节点，则需要填写对应节点的审批人<br>key:  node id 或 custom node id , 通过 [查看审批定义](https://open.feishu.cn/document/ukTMukTMukTM/uADNyUjLwQjM14CM0ITN) 获取<br> value: 审批人列表
	NodeApproverOpenIDList map[string][]string `json:"node_approver_open_id_list,omitempty"` // 发起人自选 open id
	UUID                   *string             `json:"uuid,omitempty"`                       // 审批实例 uuid，用于幂等操作，同一个 uuid 只能用于创建一个审批实例，如果冲突，返回错误码 60012 ，格式必须为 XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX，不区分大小写
}

type createApprovalInstanceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非0表示失败
	Msg  string                      `json:"msg,omitempty"`  // 返回码的描述
	Data *CreateApprovalInstanceResp `json:"data,omitempty"` // 返回业务信息
}

type CreateApprovalInstanceResp struct {
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例 Code
}
