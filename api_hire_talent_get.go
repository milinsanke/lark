// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHireTalent 根据人才 ID 获取人才信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/talent/get
func (r *HireService) GetHireTalent(ctx context.Context, request *GetHireTalentReq, options ...MethodOptionFunc) (*GetHireTalentResp, *Response, error) {
	if r.cli.mock.mockHireGetHireTalent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireTalent mock enable")
		return r.cli.mock.mockHireGetHireTalent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireTalent",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/talents/:talent_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireTalentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireGetHireTalent(f func(ctx context.Context, request *GetHireTalentReq, options ...MethodOptionFunc) (*GetHireTalentResp, *Response, error)) {
	r.mockHireGetHireTalent = f
}

func (r *Mock) UnMockHireGetHireTalent() {
	r.mockHireGetHireTalent = nil
}

type GetHireTalentReq struct {
	TalentID string `path:"talent_id" json:"-"` // 人才ID, 示例值："6891560630172518670"
}

type getHireTalentResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetHireTalentResp `json:"data,omitempty"`
}

type GetHireTalentResp struct {
	Talent *GetHireTalentRespTalent `json:"talent,omitempty"` // 人才信息
}

type GetHireTalentRespTalent struct {
	ID                        string                                          `json:"id,omitempty"`                          // ID
	BasicInfo                 *GetHireTalentRespTalentBasicInfo               `json:"basic_info,omitempty"`                  // 基础信息
	EducationList             []*GetHireTalentRespTalentEducation             `json:"education_list,omitempty"`              // 教育经历
	CareerList                []*GetHireTalentRespTalentCareer                `json:"career_list,omitempty"`                 // 工作经历
	ProjectList               []*GetHireTalentRespTalentProject               `json:"project_list,omitempty"`                // 项目经历
	WorksList                 []*GetHireTalentRespTalentWorks                 `json:"works_list,omitempty"`                  // 作品集
	AwardList                 []*GetHireTalentRespTalentAward                 `json:"award_list,omitempty"`                  // 获奖列表
	CompetitionList           []*GetHireTalentRespTalentCompetition           `json:"competition_list,omitempty"`            // 竞赛列表
	CertificateList           []*GetHireTalentRespTalentCertificate           `json:"certificate_list,omitempty"`            // 证书列表
	LanguageList              []*GetHireTalentRespTalentLanguage              `json:"language_list,omitempty"`               // 语言列表
	SnsList                   []*GetHireTalentRespTalentSns                   `json:"sns_list,omitempty"`                    // SNS列表
	ResumeSourceList          []*GetHireTalentRespTalentResumeSource          `json:"resume_source_list,omitempty"`          // 简历来源
	InterviewRegistrationList []*GetHireTalentRespTalentInterviewRegistration `json:"interview_registration_list,omitempty"` // 面试登记表
	ResumeAttachmentIDList    []string                                        `json:"resume_attachment_id_list,omitempty"`   // 简历附件id列表（按照简历创建时间降序）
}

type GetHireTalentRespTalentBasicInfo struct {
	Name                 string                                           `json:"name,omitempty"`                  // 名字
	Mobile               string                                           `json:"mobile,omitempty"`                // 手机
	MobileCountryCode    string                                           `json:"mobile_country_code,omitempty"`   // 手机国家代码
	Email                string                                           `json:"email,omitempty"`                 // 邮箱
	ExperienceYears      int64                                            `json:"experience_years,omitempty"`      // 工作年限
	Age                  int64                                            `json:"age,omitempty"`                   // 年龄
	Nationality          *GetHireTalentRespTalentBasicInfoNationality     `json:"nationality,omitempty"`           // 国籍
	Gender               int64                                            `json:"gender,omitempty"`                // 性别, 可选值有: `1`：男, `2`：女, `3`：其他
	CurrentCity          *GetHireTalentRespTalentBasicInfoCurrentCity     `json:"current_city,omitempty"`          // 当前所在城市信息
	HometownCity         *GetHireTalentRespTalentBasicInfoHometownCity    `json:"hometown_city,omitempty"`         // 家乡
	PreferredCityList    []*GetHireTalentRespTalentBasicInfoPreferredCity `json:"preferred_city_list,omitempty"`   // 偏好城市
	IdentificationType   int64                                            `json:"identification_type,omitempty"`   // 证件类型, 可选值有: `1`：中国 - 居民身份证, `2`：护照, `3`：中国 - 港澳居民居住证, `4`：中国 - 台湾居民来往大陆通行证, `5`：其他, `6`：中国 - 港澳居民来往内地通行证, `9`：中国 - 台湾居民居住证
	IdentificationNumber string                                           `json:"identification_number,omitempty"` // 证件号
	Birthday             int64                                            `json:"birthday,omitempty"`              // 生日
}

type GetHireTalentRespTalentBasicInfoNationality struct {
	NationalityCode string `json:"nationality_code,omitempty"` // 国家编码
	ZhName          string `json:"zh_name,omitempty"`          // 名字
	EnName          string `json:"en_name,omitempty"`          // 英文名
}

type GetHireTalentRespTalentBasicInfoCurrentCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 名字
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

type GetHireTalentRespTalentBasicInfoHometownCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 名字
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

type GetHireTalentRespTalentBasicInfoPreferredCity struct {
	CityCode string `json:"city_code,omitempty"` // 城市码
	ZhName   string `json:"zh_name,omitempty"`   // 名字
	EnName   string `json:"en_name,omitempty"`   // 英文名
}

type GetHireTalentRespTalentEducation struct {
	ID              string `json:"id,omitempty"`               // ID
	Degree          int64  `json:"degree,omitempty"`           // 学位, 可选值有: `1`：小学, `2`：初中, `3`：专职, `4`：高中, `5`：大专, `6`：本科, `7`：硕士, `8`：博士, `9`：其他
	School          string `json:"school,omitempty"`           // 学校
	FieldOfStudy    string `json:"field_of_study,omitempty"`   // 专业
	StartTime       int64  `json:"start_time,omitempty"`       // 开始时间
	EndTime         int64  `json:"end_time,omitempty"`         // 结束时间
	EducationType   int64  `json:"education_type,omitempty"`   // 学历类型, 可选值有: `1`：非中国大陆, `2`：统招全日制, `3`：非全日制, `4`：自考, `5`：其他
	AcademicRanking int64  `json:"academic_ranking,omitempty"` // 成绩排名, 可选值有: `5`：前 5 %, `10`：前 10 %, `20`：前 20 %, `30`：前 30 %, `50`：前 50 %, `-1`：其他
	TagList         string `json:"tag_list,omitempty"`         // 标记, 可选值有: `1`：985学校, `2`：211学校, `3`：一本, `4`：国外院校QS200, `5`：百度 阿里 腾讯, `6`：头条, 美团, 滴滴, `7`：其它大厂, `8`：猎头渠道, `9`：内推渠道, `10`：互联网大厂（包含 BAT/TMD）, `11`：熟人内推, `100`：email, `101`：mobile, `102`：猎头保护中, `103`：已入职, `104`：已离职
}

type GetHireTalentRespTalentCareer struct {
	ID        string `json:"id,omitempty"`         // ID
	Company   string `json:"company,omitempty"`    // 公司
	Title     string `json:"title,omitempty"`      // 职位
	Desc      string `json:"desc,omitempty"`       // 描述
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
	TagList   int64  `json:"tag_list,omitempty"`   // 标签, 可选值有: `1`：985学校, `2`：211学校, `3`：一本, `4`：国外院校QS200, `5`：百度 阿里 腾讯, `6`：头条, 美团, 滴滴, `7`：其它大厂, `8`：猎头渠道, `9`：内推渠道, `10`：互联网大厂（包含 BAT/TMD）, `11`：熟人内推, `100`：email, `101`：mobile, `102`：猎头保护中, `103`：已入职, `104`：已离职
}

type GetHireTalentRespTalentProject struct {
	ID        string `json:"id,omitempty"`         // ID
	Name      string `json:"name,omitempty"`       // 项目名称
	Role      string `json:"role,omitempty"`       // 项目角色
	Link      string `json:"link,omitempty"`       // 项目链接
	Desc      string `json:"desc,omitempty"`       // 描述
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

type GetHireTalentRespTalentWorks struct {
	ID   string `json:"id,omitempty"`   // ID
	Link string `json:"link,omitempty"` // 链接
	Desc string `json:"desc,omitempty"` // 描述
	Name string `json:"name,omitempty"` // 名字
}

type GetHireTalentRespTalentAward struct {
	ID        string `json:"id,omitempty"`         // ID
	Title     string `json:"title,omitempty"`      // 名称
	AwardTime string `json:"award_time,omitempty"` // 获奖时间
	Desc      string `json:"desc,omitempty"`       // 描述
}

type GetHireTalentRespTalentCompetition struct {
	ID   string `json:"id,omitempty"`   // ID
	Name string `json:"name,omitempty"` // 竞赛名称
	Desc string `json:"desc,omitempty"` // 竞赛描述
}

type GetHireTalentRespTalentCertificate struct {
	ID   string `json:"id,omitempty"`   // ID
	Name string `json:"name,omitempty"` // 证件名称
	Desc string `json:"desc,omitempty"` // 证件描述
}

type GetHireTalentRespTalentLanguage struct {
	ID          string `json:"id,omitempty"`          // ID
	Language    int64  `json:"language,omitempty"`    // 语言, 可选值有: `1`：英语, `2`：法语, `3`：日语, `4`：韩语, `5`：德语, `6`：俄语, `7`：西班牙语, `8`：葡萄牙语, `9`：阿拉伯语, `10`：印地语, `11`：印度斯坦语, `12`：孟加拉语, `13`：豪萨语, `14`：旁遮普语, `15`：波斯语, `16`：斯瓦西里语, `17`：泰卢固语, `18`：土耳其语, `19`：意大利语, `20`：爪哇语, `21`：泰米尔语, `22`：马拉地语, `23`：越南语, `24`：普通话, `25`：粤语
	Proficiency int64  `json:"proficiency,omitempty"` // 熟练程度, 可选值有: `1`：入门, `2`：日常会话, `3`：商务会话, `4`：无障碍沟通, `5`：母语
}

type GetHireTalentRespTalentSns struct {
	ID      string `json:"id,omitempty"`       // ID
	SnsType int64  `json:"sns_type,omitempty"` // SNS类型, 可选值有: `1`：领英, `2`：脉脉, `3`：微信, `4`：微博, `5`：Github, `6`：知乎, `7`：脸书, `8`：推特, `9`：Whatsapp, `10`：个人网站, `11`：QQ
	Link    string `json:"link,omitempty"`     // SNS链接
}

type GetHireTalentRespTalentResumeSource struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名
	EnName string `json:"en_name,omitempty"` // 英文名
}

type GetHireTalentRespTalentInterviewRegistration struct {
	ID               string `json:"id,omitempty"`                // ID
	RegistrationTime int64  `json:"registration_time,omitempty"` // 创建时间
}
