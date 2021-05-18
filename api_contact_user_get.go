// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetUser 该接口用于获取通讯录中单个用户的信息。只能访问有数据权限的字段，具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get
func (r *ContactService) GetUser(ctx context.Context, request *GetUserReq, options ...MethodOptionFunc) (*GetUserResp, *Response, error) {
	if r.cli.mock.mockContactGetUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetUser mock enable")
		return r.cli.mock.mockContactGetUser(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Contact#GetUser call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetUser request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/users/:user_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(getUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Contact#GetUser GET https://open.feishu.cn/open-apis/contact/v3/users/:user_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Contact#GetUser GET https://open.feishu.cn/open-apis/contact/v3/users/:user_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Contact", "GetUser", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetUser success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockContactGetUser(f func(ctx context.Context, request *GetUserReq, options ...MethodOptionFunc) (*GetUserResp, *Response, error)) {
	r.mockContactGetUser = f
}

func (r *Mock) UnMockContactGetUser() {
	r.mockContactGetUser = nil
}

type GetUserReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	UserID           string            `path:"user_id" json:"-"`             // 用户ID，需要与查询参数中的user_id_type类型保持一致。, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

type getUserResp struct {
	Code int64        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string       `json:"msg,omitempty"`  // 错误描述
	Data *GetUserResp `json:"data,omitempty"` //
}

type GetUserResp struct {
	User *GetUserRespUser `json:"user,omitempty"` // 用户信息
}

type GetUserRespUser struct {
	UnionID         string                       `json:"union_id,omitempty"`          // 用户的union_id
	UserID          string                       `json:"user_id,omitempty"`           // 租户内用户的唯一标识, 字段权限要求:  获取用户 userid
	OpenID          string                       `json:"open_id,omitempty"`           // 用户的open_id
	Name            string                       `json:"name,omitempty"`              // 用户名, 最小长度：`1` 字符
	EnName          string                       `json:"en_name,omitempty"`           // 英文名
	Email           string                       `json:"email,omitempty"`             // 邮箱, 字段权限要求:  获取用户邮箱
	Mobile          string                       `json:"mobile,omitempty"`            // 手机号, 字段权限要求:  获取用户手机号
	MobileVisible   bool                         `json:"mobile_visible,omitempty"`    // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码
	Gender          int64                        `json:"gender,omitempty"`            // 性别, 可选值有: `0`：保密, `1`：男, `2`：女
	Avatar          *GetUserRespUserAvatar       `json:"avatar,omitempty"`            // 用户头像信息
	Status          *GetUserRespUserStatus       `json:"status,omitempty"`            // 用户状态
	DepartmentIDs   []string                     `json:"department_ids,omitempty"`    // 用户所属部门的ID列表
	LeaderUserID    string                       `json:"leader_user_id,omitempty"`    // 用户的直接主管的用户ID
	City            string                       `json:"city,omitempty"`              // 城市
	Country         string                       `json:"country,omitempty"`           // 国家
	WorkStation     string                       `json:"work_station,omitempty"`      // 工位
	JoinTime        int64                        `json:"join_time,omitempty"`         // 入职时间
	IsTenantManager bool                         `json:"is_tenant_manager,omitempty"` // 是否是租户管理员
	EmployeeNo      string                       `json:"employee_no,omitempty"`       // 工号
	EmployeeType    int64                        `json:"employee_type,omitempty"`     // 员工类型
	Orders          []*GetUserRespUserOrder      `json:"orders,omitempty"`            // 用户排序信息
	CustomAttrs     []*GetUserRespUserCustomAttr `json:"custom_attrs,omitempty"`      // 自定义属性
	EnterpriseEmail string                       `json:"enterprise_email,omitempty"`  // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务
}

type GetUserRespUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type GetUserRespUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否冻结
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type GetUserRespUserOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前
}

type GetUserRespUserCustomAttr struct {
	Type  string                          `json:"type,omitempty"`  // 自定义属性类型
	ID    string                          `json:"id,omitempty"`    // 自定义属性ID
	Value *GetUserRespUserCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type GetUserRespUserCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	URL   string `json:"url,omitempty"`    // URL
	PcURL string `json:"pc_url,omitempty"` // PC上的URL
}
