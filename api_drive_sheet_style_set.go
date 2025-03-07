// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetSheetStyle 该接口用于根据 spreadsheetToken 、range 和样式信息更新单元格样式；单次写入不超过5000行，100列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukjMzUjL5IzM14SOyMTN
func (r *DriveService) SetSheetStyle(ctx context.Context, request *SetSheetStyleReq, options ...MethodOptionFunc) (*SetSheetStyleResp, *Response, error) {
	if r.cli.mock.mockDriveSetSheetStyle != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#SetSheetStyle mock enable")
		return r.cli.mock.mockDriveSetSheetStyle(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "SetSheetStyle",
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/style",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(setSheetStyleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveSetSheetStyle(f func(ctx context.Context, request *SetSheetStyleReq, options ...MethodOptionFunc) (*SetSheetStyleResp, *Response, error)) {
	r.mockDriveSetSheetStyle = f
}

func (r *Mock) UnMockDriveSetSheetStyle() {
	r.mockDriveSetSheetStyle = nil
}

type SetSheetStyleReq struct {
	SpreadSheetToken string                       `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，详见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	AppendStyle      *SetSheetStyleReqAppendStyle `json:"appendStyle,omitempty"`     // 设置单元格样式
}

type SetSheetStyleReqAppendStyle struct {
	Range string                            `json:"range,omitempty"` // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Style *SetSheetStyleReqAppendStyleStyle `json:"style,omitempty"` // 需要更新的样式
}

type SetSheetStyleReqAppendStyleStyle struct {
	Font           *SetSheetStyleReqAppendStyleStyleFont `json:"font,omitempty"`           // 字体相关样式
	TextDecoration *int64                                `json:"textDecoration,omitempty"` // 文本装饰 ，0 默认，1 下划线，2 删除线 ，3 下划线和删除线
	Formatter      *string                               `json:"formatter,omitempty"`      // 数字格式，详见附录 [sheet支持数字格式类型](https://open.feishu.cn/document/ukTMukTMukTM/uMjM2UjLzIjN14yMyYTN)
	HAlign         *int64                                `json:"hAlign,omitempty"`         // 水平对齐，0 左对齐，1 中对齐，2 右对齐
	VAlign         *int64                                `json:"vAlign,omitempty"`         // 垂直对齐, 0 上对齐，1 中对齐, 2 下对齐
	ForeColor      *string                               `json:"foreColor,omitempty"`      // 字体颜色
	BackColor      *string                               `json:"backColor,omitempty"`      // 背景颜色
	BorderType     *string                               `json:"borderType,omitempty"`     // 边框类型，可选 "FULL_BORDER"，"OUTER_BORDER"，"INNER_BORDER"，"NO_BORDER"，"LEFT_BORDER"，"RIGHT_BORDER"，"TOP_BORDER"，"BOTTOM_BORDER"
	BorderColor    *string                               `json:"borderColor,omitempty"`    // 边框颜色
	Clean          *bool                                 `json:"clean,omitempty"`          // 是否清除所有格式,默认 false
}

type SetSheetStyleReqAppendStyleStyleFont struct {
	Bold     *bool   `json:"bold,omitempty"`     // 是否加粗
	Italic   *bool   `json:"italic,omitempty"`   // 是否斜体
	FontSize *string `json:"fontSize,omitempty"` // 字体大小 字号大小为9~36 行距固定为1.5，如:10pt/1.5
	Clean    *bool   `json:"clean,omitempty"`    // 清除 font 格式,默认 false
}

type setSheetStyleResp struct {
	Code int64              `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *SetSheetStyleResp `json:"data,omitempty"`
}

type SetSheetStyleResp struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 设置样式的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 设置样式的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 设置样式的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 设置样式的单元格总数
	Revision         int64  `json:"revision,omitempty"`         // sheet 的版本号
}
