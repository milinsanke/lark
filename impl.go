// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"net/http"
	"time"
)

type Lark struct {
	appID               string
	appSecret           string
	encryptKey          string
	verificationToken   string
	helpdeskID          string
	helpdeskToken       string
	timeout             time.Duration
	isISV               bool
	tenantKey           string
	customBotWebHookURL string
	customBotSecret     string

	httpClient   *http.Client
	logger       Logger
	logLevel     LogLevel
	store        Store
	mock         *Mock
	eventHandler *eventHandler

	// service
	Auth          *AuthService
	Contact       *ContactService
	Message       *MessageService
	Chat          *ChatService
	Bot           *BotService
	Calendar      *CalendarService
	Drive         *DriveService
	Bitable       *BitableService
	MeetingRoom   *MeetingRoomService
	VC            *VCService
	Application   *ApplicationService
	Mail          *MailService
	Approval      *ApprovalService
	Helpdesk      *HelpdeskService
	Admin         *AdminService
	HumanAuth     *HumanAuthService
	AI            *AIService
	Attendance    *AttendanceService
	File          *FileService
	EventCallback *EventCallbackService
	OKR           *OKRService
	EHR           *EHRService
	Tenant        *TenantService
	Search        *SearchService
	Hire          *HireService
}

func (r *Lark) initService() {
	r.Auth = &AuthService{cli: r}
	r.Contact = &ContactService{cli: r}
	r.Message = &MessageService{cli: r}
	r.Chat = &ChatService{cli: r}
	r.Bot = &BotService{cli: r}
	r.Calendar = &CalendarService{cli: r}
	r.Drive = &DriveService{cli: r}
	r.Bitable = &BitableService{cli: r}
	r.MeetingRoom = &MeetingRoomService{cli: r}
	r.VC = &VCService{cli: r}
	r.Application = &ApplicationService{cli: r}
	r.Mail = &MailService{cli: r}
	r.Approval = &ApprovalService{cli: r}
	r.Helpdesk = &HelpdeskService{cli: r}
	r.Admin = &AdminService{cli: r}
	r.HumanAuth = &HumanAuthService{cli: r}
	r.AI = &AIService{cli: r}
	r.Attendance = &AttendanceService{cli: r}
	r.File = &FileService{cli: r}
	r.EventCallback = &EventCallbackService{cli: r}
	r.OKR = &OKRService{cli: r}
	r.EHR = &EHRService{cli: r}
	r.Tenant = &TenantService{cli: r}
	r.Search = &SearchService{cli: r}
	r.Hire = &HireService{cli: r}
}

type (
	AuthService          struct{ cli *Lark }
	ContactService       struct{ cli *Lark }
	MessageService       struct{ cli *Lark }
	ChatService          struct{ cli *Lark }
	BotService           struct{ cli *Lark }
	CalendarService      struct{ cli *Lark }
	DriveService         struct{ cli *Lark }
	BitableService       struct{ cli *Lark }
	MeetingRoomService   struct{ cli *Lark }
	VCService            struct{ cli *Lark }
	ApplicationService   struct{ cli *Lark }
	MailService          struct{ cli *Lark }
	ApprovalService      struct{ cli *Lark }
	HelpdeskService      struct{ cli *Lark }
	AdminService         struct{ cli *Lark }
	HumanAuthService     struct{ cli *Lark }
	AIService            struct{ cli *Lark }
	AttendanceService    struct{ cli *Lark }
	FileService          struct{ cli *Lark }
	EventCallbackService struct{ cli *Lark }
	OKRService           struct{ cli *Lark }
	EHRService           struct{ cli *Lark }
	TenantService        struct{ cli *Lark }
	SearchService        struct{ cli *Lark }
	HireService          struct{ cli *Lark }
)
