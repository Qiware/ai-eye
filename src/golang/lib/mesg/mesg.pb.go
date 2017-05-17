// Code generated by protoc-gen-go.
// source: mesg.proto
// DO NOT EDIT!

/*
Package mesg is a generated protocol buffer package.

It is generated from these files:
	mesg.proto

It has these top-level messages:
	MesgOnline
	MesgOnlineAck
	MesgKick
	BrowserPluginInfo
	BrowserUseragentInfo
	BrowserScreenInfo
	MesgBrowserEnv
	MesgBrowserEnvAck
	MesgEventStatistic
	MesgEventStatisticAck
	MesgLsndInfo
	MesgFrwdInfo
*/
package mesg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 控件定义
type CtrlType int32

const (
	CtrlType_CTRL_USR_NAME_INPUT_BOX CtrlType = 0
	CtrlType_CTRL_PASSWD_INPUT_BOX   CtrlType = 1
	CtrlType_CTRL_IMG_CODE_INPUT_BOX CtrlType = 2
	CtrlType_CTRL_IMG_CODE_BTN       CtrlType = 3
	CtrlType_CTRL_TEL_INPUT_BOX      CtrlType = 4
	CtrlType_CTRL_SMS_INPUT_BOX      CtrlType = 5
	CtrlType_CTRL_SMS_BTN            CtrlType = 6
	CtrlType_CTRL_LOGIN_BTN          CtrlType = 7
)

var CtrlType_name = map[int32]string{
	0: "CTRL_USR_NAME_INPUT_BOX",
	1: "CTRL_PASSWD_INPUT_BOX",
	2: "CTRL_IMG_CODE_INPUT_BOX",
	3: "CTRL_IMG_CODE_BTN",
	4: "CTRL_TEL_INPUT_BOX",
	5: "CTRL_SMS_INPUT_BOX",
	6: "CTRL_SMS_BTN",
	7: "CTRL_LOGIN_BTN",
}
var CtrlType_value = map[string]int32{
	"CTRL_USR_NAME_INPUT_BOX": 0,
	"CTRL_PASSWD_INPUT_BOX":   1,
	"CTRL_IMG_CODE_INPUT_BOX": 2,
	"CTRL_IMG_CODE_BTN":       3,
	"CTRL_TEL_INPUT_BOX":      4,
	"CTRL_SMS_INPUT_BOX":      5,
	"CTRL_SMS_BTN":            6,
	"CTRL_LOGIN_BTN":          7,
}

func (x CtrlType) Enum() *CtrlType {
	p := new(CtrlType)
	*p = x
	return p
}
func (x CtrlType) String() string {
	return proto.EnumName(CtrlType_name, int32(x))
}
func (x *CtrlType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CtrlType_value, data, "CtrlType")
	if err != nil {
		return err
	}
	*x = CtrlType(value)
	return nil
}
func (CtrlType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 事件定义
type EventType int32

const (
	EventType_EV_CHANGE      EventType = 0
	EventType_EV_CLICK       EventType = 1
	EventType_EV_DBLCLICK    EventType = 2
	EventType_EV_FOCUS       EventType = 3
	EventType_EV_KEYDOWN     EventType = 4
	EventType_EV_KEYPRESS    EventType = 5
	EventType_EV_KEYUP       EventType = 6
	EventType_EV_MOUSEDOWN   EventType = 7
	EventType_EV_MOUSEMOVE   EventType = 8
	EventType_EV_MOUSEOUT    EventType = 9
	EventType_EV_MOUSEOVER   EventType = 10
	EventType_EV_MOUSEUP     EventType = 11
	EventType_EV_TOUCHSTART  EventType = 12
	EventType_EV_TOUCHMOVE   EventType = 13
	EventType_EV_TOUCHEND    EventType = 14
	EventType_EV_TOUCHCANCEL EventType = 15
)

var EventType_name = map[int32]string{
	0:  "EV_CHANGE",
	1:  "EV_CLICK",
	2:  "EV_DBLCLICK",
	3:  "EV_FOCUS",
	4:  "EV_KEYDOWN",
	5:  "EV_KEYPRESS",
	6:  "EV_KEYUP",
	7:  "EV_MOUSEDOWN",
	8:  "EV_MOUSEMOVE",
	9:  "EV_MOUSEOUT",
	10: "EV_MOUSEOVER",
	11: "EV_MOUSEUP",
	12: "EV_TOUCHSTART",
	13: "EV_TOUCHMOVE",
	14: "EV_TOUCHEND",
	15: "EV_TOUCHCANCEL",
}
var EventType_value = map[string]int32{
	"EV_CHANGE":      0,
	"EV_CLICK":       1,
	"EV_DBLCLICK":    2,
	"EV_FOCUS":       3,
	"EV_KEYDOWN":     4,
	"EV_KEYPRESS":    5,
	"EV_KEYUP":       6,
	"EV_MOUSEDOWN":   7,
	"EV_MOUSEMOVE":   8,
	"EV_MOUSEOUT":    9,
	"EV_MOUSEOVER":   10,
	"EV_MOUSEUP":     11,
	"EV_TOUCHSTART":  12,
	"EV_TOUCHMOVE":   13,
	"EV_TOUCHEND":    14,
	"EV_TOUCHCANCEL": 15,
}

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}
func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (x *EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EventType_value, data, "EventType")
	if err != nil {
		return err
	}
	*x = EventType(value)
	return nil
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

//
// 命令ID: 0x0101
// 命令描述: 上线请求(ONLINE)
// 协议格式:
type MesgOnline struct {
	Sid              *uint64 `protobuf:"varint,1,req,name=sid" json:"sid,omitempty"`
	Token            *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	App              *string `protobuf:"bytes,3,req,name=app" json:"app,omitempty"`
	Version          *string `protobuf:"bytes,4,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgOnline) Reset()                    { *m = MesgOnline{} }
func (m *MesgOnline) String() string            { return proto.CompactTextString(m) }
func (*MesgOnline) ProtoMessage()               {}
func (*MesgOnline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MesgOnline) GetSid() uint64 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *MesgOnline) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *MesgOnline) GetApp() string {
	if m != nil && m.App != nil {
		return *m.App
	}
	return ""
}

func (m *MesgOnline) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

//
// 命令ID: 0x0102
// 命令描述: 上线请求应答(ONLINE-ACK)
// 协议格式:
type MesgOnlineAck struct {
	Sid              *uint64 `protobuf:"varint,1,req,name=sid" json:"sid,omitempty"`
	App              *string `protobuf:"bytes,2,req,name=app" json:"app,omitempty"`
	Version          *string `protobuf:"bytes,3,req,name=version" json:"version,omitempty"`
	Code             *uint32 `protobuf:"varint,4,req,name=code" json:"code,omitempty"`
	Errmsg           *string `protobuf:"bytes,5,req,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgOnlineAck) Reset()                    { *m = MesgOnlineAck{} }
func (m *MesgOnlineAck) String() string            { return proto.CompactTextString(m) }
func (*MesgOnlineAck) ProtoMessage()               {}
func (*MesgOnlineAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MesgOnlineAck) GetSid() uint64 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *MesgOnlineAck) GetApp() string {
	if m != nil && m.App != nil {
		return *m.App
	}
	return ""
}

func (m *MesgOnlineAck) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *MesgOnlineAck) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *MesgOnlineAck) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

//
// 命令ID: 0x0107
// 命令描述: 踢连接下线(KICK)
// 协议格式:
type MesgKick struct {
	Code             *uint32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Errmsg           *string `protobuf:"bytes,2,req,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgKick) Reset()                    { *m = MesgKick{} }
func (m *MesgKick) String() string            { return proto.CompactTextString(m) }
func (*MesgKick) ProtoMessage()               {}
func (*MesgKick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MesgKick) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *MesgKick) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

// 插件信息
type BrowserPluginInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Desc             *string `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BrowserPluginInfo) Reset()                    { *m = BrowserPluginInfo{} }
func (m *BrowserPluginInfo) String() string            { return proto.CompactTextString(m) }
func (*BrowserPluginInfo) ProtoMessage()               {}
func (*BrowserPluginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BrowserPluginInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *BrowserPluginInfo) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

// 用户代理信息
type BrowserUseragentInfo struct {
	Ua               *string `protobuf:"bytes,1,opt,name=ua" json:"ua,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BrowserUseragentInfo) Reset()                    { *m = BrowserUseragentInfo{} }
func (m *BrowserUseragentInfo) String() string            { return proto.CompactTextString(m) }
func (*BrowserUseragentInfo) ProtoMessage()               {}
func (*BrowserUseragentInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BrowserUseragentInfo) GetUa() string {
	if m != nil && m.Ua != nil {
		return *m.Ua
	}
	return ""
}

// 屏幕信息
type BrowserScreenInfo struct {
	Width            *uint32 `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height           *uint32 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	AvailWidth       *uint32 `protobuf:"varint,3,opt,name=avail_width" json:"avail_width,omitempty"`
	AvailHeight      *uint32 `protobuf:"varint,4,opt,name=avail_height" json:"avail_height,omitempty"`
	AvailLeft        *uint32 `protobuf:"varint,5,opt,name=avail_left" json:"avail_left,omitempty"`
	AvailTop         *uint32 `protobuf:"varint,6,opt,name=avail_top" json:"avail_top,omitempty"`
	OuterWidth       *uint32 `protobuf:"varint,7,opt,name=outer_width" json:"outer_width,omitempty"`
	OuterHeight      *uint32 `protobuf:"varint,8,opt,name=outer_height" json:"outer_height,omitempty"`
	InnerWidth       *uint32 `protobuf:"varint,9,opt,name=inner_width" json:"inner_width,omitempty"`
	InnerHeight      *uint32 `protobuf:"varint,10,opt,name=inner_height" json:"inner_height,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BrowserScreenInfo) Reset()                    { *m = BrowserScreenInfo{} }
func (m *BrowserScreenInfo) String() string            { return proto.CompactTextString(m) }
func (*BrowserScreenInfo) ProtoMessage()               {}
func (*BrowserScreenInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BrowserScreenInfo) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *BrowserScreenInfo) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *BrowserScreenInfo) GetAvailWidth() uint32 {
	if m != nil && m.AvailWidth != nil {
		return *m.AvailWidth
	}
	return 0
}

func (m *BrowserScreenInfo) GetAvailHeight() uint32 {
	if m != nil && m.AvailHeight != nil {
		return *m.AvailHeight
	}
	return 0
}

func (m *BrowserScreenInfo) GetAvailLeft() uint32 {
	if m != nil && m.AvailLeft != nil {
		return *m.AvailLeft
	}
	return 0
}

func (m *BrowserScreenInfo) GetAvailTop() uint32 {
	if m != nil && m.AvailTop != nil {
		return *m.AvailTop
	}
	return 0
}

func (m *BrowserScreenInfo) GetOuterWidth() uint32 {
	if m != nil && m.OuterWidth != nil {
		return *m.OuterWidth
	}
	return 0
}

func (m *BrowserScreenInfo) GetOuterHeight() uint32 {
	if m != nil && m.OuterHeight != nil {
		return *m.OuterHeight
	}
	return 0
}

func (m *BrowserScreenInfo) GetInnerWidth() uint32 {
	if m != nil && m.InnerWidth != nil {
		return *m.InnerWidth
	}
	return 0
}

func (m *BrowserScreenInfo) GetInnerHeight() uint32 {
	if m != nil && m.InnerHeight != nil {
		return *m.InnerHeight
	}
	return 0
}

//
// 命令ID: 0x0201
// 命令描述: 上报浏览器环境信息(BROWSER-ENV)
// 协议格式:
type MesgBrowserEnv struct {
	Orig             *uint64 `protobuf:"varint,1,req,name=orig" json:"orig,omitempty"`
	Dest             *uint64 `protobuf:"varint,2,req,name=dest" json:"dest,omitempty"`
	Level            *uint32 `protobuf:"varint,3,req,name=level" json:"level,omitempty"`
	Time             *uint64 `protobuf:"varint,4,req,name=time" json:"time,omitempty"`
	Text             *string `protobuf:"bytes,5,req,name=text" json:"text,omitempty"`
	Data             []byte  `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgBrowserEnv) Reset()                    { *m = MesgBrowserEnv{} }
func (m *MesgBrowserEnv) String() string            { return proto.CompactTextString(m) }
func (*MesgBrowserEnv) ProtoMessage()               {}
func (*MesgBrowserEnv) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MesgBrowserEnv) GetOrig() uint64 {
	if m != nil && m.Orig != nil {
		return *m.Orig
	}
	return 0
}

func (m *MesgBrowserEnv) GetDest() uint64 {
	if m != nil && m.Dest != nil {
		return *m.Dest
	}
	return 0
}

func (m *MesgBrowserEnv) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *MesgBrowserEnv) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *MesgBrowserEnv) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *MesgBrowserEnv) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//
// 命令ID: 0x0202
// 命令描述: 上报浏览器环境信息应答(BROWSER-ENV-ACK)
// 协议格式:
type MesgBrowserEnvAck struct {
	Code             *uint32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Errmsg           *string `protobuf:"bytes,2,req,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgBrowserEnvAck) Reset()                    { *m = MesgBrowserEnvAck{} }
func (m *MesgBrowserEnvAck) String() string            { return proto.CompactTextString(m) }
func (*MesgBrowserEnvAck) ProtoMessage()               {}
func (*MesgBrowserEnvAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MesgBrowserEnvAck) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *MesgBrowserEnvAck) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

//
// 命令ID: 0x0203
// 命令描述: 上报事件统计信息(EVENT-STATISTIC)
// 协议格式:
type MesgEventStatistic struct {
	Orig             *uint64 `protobuf:"varint,1,req,name=orig" json:"orig,omitempty"`
	Dest             *uint64 `protobuf:"varint,2,req,name=dest" json:"dest,omitempty"`
	Mark             *string `protobuf:"bytes,3,req,name=mark" json:"mark,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgEventStatistic) Reset()                    { *m = MesgEventStatistic{} }
func (m *MesgEventStatistic) String() string            { return proto.CompactTextString(m) }
func (*MesgEventStatistic) ProtoMessage()               {}
func (*MesgEventStatistic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MesgEventStatistic) GetOrig() uint64 {
	if m != nil && m.Orig != nil {
		return *m.Orig
	}
	return 0
}

func (m *MesgEventStatistic) GetDest() uint64 {
	if m != nil && m.Dest != nil {
		return *m.Dest
	}
	return 0
}

func (m *MesgEventStatistic) GetMark() string {
	if m != nil && m.Mark != nil {
		return *m.Mark
	}
	return ""
}

//
// 命令ID: 0x0204
// 命令描述: 上报事件统计信息应答(EVENT-STATISTIC-ACK)
// 协议格式: NONE
type MesgEventStatisticAck struct {
	Code             *uint32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Errmsg           *string `protobuf:"bytes,2,req,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgEventStatisticAck) Reset()                    { *m = MesgEventStatisticAck{} }
func (m *MesgEventStatisticAck) String() string            { return proto.CompactTextString(m) }
func (*MesgEventStatisticAck) ProtoMessage()               {}
func (*MesgEventStatisticAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MesgEventStatisticAck) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *MesgEventStatisticAck) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

//
// 命令ID: 0x0301
// 命令描述: 帧听层信息上报(LSND-INFO)
// 协议格式:
type MesgLsndInfo struct {
	Type             *uint32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Nid              *uint32 `protobuf:"varint,2,req,name=nid" json:"nid,omitempty"`
	Opid             *uint32 `protobuf:"varint,3,req,name=opid" json:"opid,omitempty"`
	Nation           *string `protobuf:"bytes,4,req,name=nation" json:"nation,omitempty"`
	Ip               *string `protobuf:"bytes,5,req,name=ip" json:"ip,omitempty"`
	Port             *uint32 `protobuf:"varint,6,req,name=port" json:"port,omitempty"`
	Connections      *uint32 `protobuf:"varint,7,req,name=connections" json:"connections,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgLsndInfo) Reset()                    { *m = MesgLsndInfo{} }
func (m *MesgLsndInfo) String() string            { return proto.CompactTextString(m) }
func (*MesgLsndInfo) ProtoMessage()               {}
func (*MesgLsndInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MesgLsndInfo) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *MesgLsndInfo) GetNid() uint32 {
	if m != nil && m.Nid != nil {
		return *m.Nid
	}
	return 0
}

func (m *MesgLsndInfo) GetOpid() uint32 {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return 0
}

func (m *MesgLsndInfo) GetNation() string {
	if m != nil && m.Nation != nil {
		return *m.Nation
	}
	return ""
}

func (m *MesgLsndInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *MesgLsndInfo) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *MesgLsndInfo) GetConnections() uint32 {
	if m != nil && m.Connections != nil {
		return *m.Connections
	}
	return 0
}

//
// 命令ID: 0x0303
// 命令描述: 转发层信息上报 (FRWD-INFO)
// 协议格式:
type MesgFrwdInfo struct {
	Nid              *uint32 `protobuf:"varint,1,req,name=nid" json:"nid,omitempty"`
	Ip               *string `protobuf:"bytes,2,req,name=ip" json:"ip,omitempty"`
	ForwardPort      *uint32 `protobuf:"varint,3,req,name=forward_port" json:"forward_port,omitempty"`
	BackendPort      *uint32 `protobuf:"varint,4,req,name=backend_port" json:"backend_port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgFrwdInfo) Reset()                    { *m = MesgFrwdInfo{} }
func (m *MesgFrwdInfo) String() string            { return proto.CompactTextString(m) }
func (*MesgFrwdInfo) ProtoMessage()               {}
func (*MesgFrwdInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MesgFrwdInfo) GetNid() uint32 {
	if m != nil && m.Nid != nil {
		return *m.Nid
	}
	return 0
}

func (m *MesgFrwdInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *MesgFrwdInfo) GetForwardPort() uint32 {
	if m != nil && m.ForwardPort != nil {
		return *m.ForwardPort
	}
	return 0
}

func (m *MesgFrwdInfo) GetBackendPort() uint32 {
	if m != nil && m.BackendPort != nil {
		return *m.BackendPort
	}
	return 0
}

func init() {
	proto.RegisterType((*MesgOnline)(nil), "mesg_online")
	proto.RegisterType((*MesgOnlineAck)(nil), "mesg_online_ack")
	proto.RegisterType((*MesgKick)(nil), "mesg_kick")
	proto.RegisterType((*BrowserPluginInfo)(nil), "browser_plugin_info")
	proto.RegisterType((*BrowserUseragentInfo)(nil), "browser_useragent_info")
	proto.RegisterType((*BrowserScreenInfo)(nil), "browser_screen_info")
	proto.RegisterType((*MesgBrowserEnv)(nil), "mesg_browser_env")
	proto.RegisterType((*MesgBrowserEnvAck)(nil), "mesg_browser_env_ack")
	proto.RegisterType((*MesgEventStatistic)(nil), "mesg_event_statistic")
	proto.RegisterType((*MesgEventStatisticAck)(nil), "mesg_event_statistic_ack")
	proto.RegisterType((*MesgLsndInfo)(nil), "mesg_lsnd_info")
	proto.RegisterType((*MesgFrwdInfo)(nil), "mesg_frwd_info")
	proto.RegisterEnum("CtrlType", CtrlType_name, CtrlType_value)
	proto.RegisterEnum("EventType", EventType_name, EventType_value)
}

func init() { proto.RegisterFile("mesg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0xc6, 0x6b, 0xf9, 0x5f, 0x7d, 0x2c, 0x3b, 0x0c, 0x93, 0x76, 0x1a, 0x76, 0x53, 0x08, 0xbb,
	0xe8, 0x7a, 0x31, 0x60, 0xc0, 0x2e, 0x76, 0x39, 0xc7, 0xe6, 0xd2, 0x20, 0xb6, 0x65, 0x58, 0x96,
	0xbb, 0xec, 0x86, 0x50, 0x65, 0xc6, 0x21, 0x6c, 0x53, 0x82, 0xc4, 0x38, 0xdd, 0x9b, 0xec, 0xa5,
	0xf6, 0x06, 0x7b, 0x98, 0x81, 0x87, 0xb4, 0xeb, 0x06, 0x43, 0xb1, 0x4b, 0xfe, 0xce, 0xf7, 0x7d,
	0x87, 0x3a, 0x87, 0x02, 0xd8, 0x89, 0x6a, 0xfd, 0x63, 0x51, 0xe6, 0x3a, 0x0f, 0xaf, 0xa1, 0x6b,
	0x4e, 0x3c, 0x57, 0x5b, 0xa9, 0x04, 0xed, 0x42, 0xbd, 0x92, 0xab, 0xa0, 0xf6, 0xc6, 0x7b, 0xdb,
	0xa0, 0x3d, 0x68, 0xea, 0x7c, 0x23, 0x54, 0xe0, 0xbd, 0xf1, 0xde, 0x76, 0x4c, 0x2d, 0x2d, 0x8a,
	0xa0, 0x8e, 0x87, 0x33, 0x68, 0xef, 0x45, 0x59, 0xc9, 0x5c, 0x05, 0x0d, 0x03, 0xc2, 0x3f, 0xe0,
	0xec, 0x24, 0x88, 0xa7, 0xd9, 0xe6, 0xcb, 0x30, 0xe7, 0xf6, 0x9e, 0xbb, 0x6d, 0x9c, 0x0f, 0x8d,
	0x2c, 0x5f, 0x09, 0xcc, 0xea, 0xd1, 0x3e, 0xb4, 0x44, 0x59, 0xee, 0xaa, 0x75, 0xd0, 0xc4, 0xec,
	0x1f, 0xa0, 0x83, 0xd9, 0x1b, 0x99, 0x6d, 0x8e, 0xd2, 0xda, 0x33, 0x29, 0x26, 0x87, 0x3f, 0xc1,
	0xc5, 0xc7, 0x32, 0x7f, 0xaa, 0x44, 0xc9, 0x8b, 0xed, 0xe3, 0x5a, 0x2a, 0x2e, 0xd5, 0x7d, 0x6e,
	0x4c, 0x2a, 0xdd, 0x19, 0x53, 0xcd, 0x76, 0x5b, 0x89, 0x2a, 0x0b, 0x3c, 0x73, 0x0a, 0xbf, 0x87,
	0xd7, 0x07, 0xcb, 0x63, 0x25, 0xca, 0x74, 0x2d, 0x94, 0xb6, 0x2e, 0x00, 0xef, 0x31, 0xb5, 0x9e,
	0xf0, 0x9f, 0xda, 0xe7, 0xe4, 0x2a, 0x2b, 0x85, 0x70, 0xc9, 0x3d, 0x68, 0x3e, 0xc9, 0x95, 0x7e,
	0x40, 0x19, 0xde, 0xe7, 0x41, 0xc8, 0xf5, 0x83, 0xc6, 0xf0, 0x1e, 0xbd, 0x80, 0x6e, 0xba, 0x4f,
	0xe5, 0x96, 0x5b, 0x51, 0x1d, 0xe1, 0x25, 0xf8, 0x16, 0x3a, 0x69, 0x03, 0x29, 0x05, 0xb0, 0x74,
	0x2b, 0xee, 0x75, 0xd0, 0x44, 0x76, 0x0e, 0x1d, 0xcb, 0x74, 0x5e, 0x04, 0xad, 0x43, 0x62, 0xfe,
	0xa8, 0x45, 0xe9, 0x12, 0xdb, 0x87, 0x44, 0x0b, 0x5d, 0xe2, 0xcb, 0x83, 0x54, 0x2a, 0x75, 0x94,
	0x76, 0x0e, 0x52, 0x0b, 0x9d, 0x14, 0x0c, 0x0d, 0xd7, 0x40, 0x70, 0xc4, 0x87, 0x4f, 0x14, 0x6a,
	0x6f, 0xc6, 0x94, 0x97, 0x72, 0xed, 0x16, 0x68, 0x87, 0xa6, 0x71, 0xce, 0xf8, 0x36, 0xb6, 0x62,
	0x2f, 0xb6, 0xb8, 0xbf, 0x9e, 0x29, 0x6a, 0xb9, 0xb3, 0xfb, 0x43, 0xa9, 0x16, 0x9f, 0xb4, 0xdd,
	0x1e, 0x1a, 0x53, 0x9d, 0xe2, 0xf5, 0xfd, 0xf0, 0x67, 0xb8, 0x7c, 0xde, 0x08, 0x1f, 0xcb, 0xd7,
	0xd7, 0xfa, 0xab, 0x73, 0x89, 0xbd, 0x59, 0x4e, 0xa5, 0x53, 0x2d, 0x2b, 0x2d, 0xb3, 0xaf, 0x5e,
	0xd1, 0x87, 0xc6, 0x2e, 0x2d, 0x37, 0xf6, 0x85, 0x85, 0xbf, 0x40, 0xf0, 0x5f, 0x09, 0xff, 0xa3,
	0xf7, 0x27, 0xe8, 0xa3, 0x73, 0x5b, 0xa9, 0xd5, 0xf1, 0x35, 0xe9, 0x3f, 0x8b, 0x83, 0xbe, 0x0b,
	0x75, 0x25, 0x57, 0x28, 0xc6, 0x41, 0xe4, 0x85, 0x5c, 0xb9, 0xb1, 0xf4, 0xa1, 0xa5, 0x52, 0x7d,
	0xfc, 0x49, 0xcc, 0x83, 0x92, 0xc5, 0xe7, 0xb1, 0x14, 0x79, 0xa9, 0x83, 0x16, 0x2a, 0x2f, 0xa0,
	0x9b, 0xe5, 0x4a, 0x89, 0xcc, 0xa8, 0xab, 0xa0, 0x6d, 0x60, 0x98, 0xb8, 0xce, 0xf7, 0xe5, 0x93,
	0xeb, 0xec, 0x7a, 0xd9, 0xc6, 0x36, 0xcd, 0xfe, 0x51, 0x97, 0xe0, 0xdf, 0xe7, 0xe5, 0x53, 0x5a,
	0xae, 0x38, 0xa6, 0xda, 0xfe, 0x97, 0xe0, 0x7f, 0x4c, 0xb3, 0x8d, 0x50, 0x8e, 0xe2, 0xef, 0xf5,
	0xee, 0xef, 0x1a, 0x74, 0x32, 0x5d, 0x6e, 0xb9, 0xf9, 0x08, 0xfa, 0x1d, 0x7c, 0x33, 0x5c, 0xcc,
	0xc7, 0x3c, 0x89, 0xe7, 0x7c, 0x3a, 0x98, 0x30, 0x7e, 0x33, 0x9d, 0x25, 0x0b, 0x7e, 0x15, 0xfd,
	0x4e, 0x5e, 0xd0, 0x6f, 0xe1, 0x15, 0x16, 0x67, 0x83, 0x38, 0xfe, 0x30, 0x3a, 0x29, 0xd5, 0x8e,
	0xbe, 0x9b, 0xc9, 0x35, 0x1f, 0x46, 0xa3, 0x53, 0x9f, 0x47, 0x5f, 0xc1, 0xf9, 0x97, 0xc5, 0xab,
	0xc5, 0x94, 0xd4, 0xe9, 0x6b, 0xa0, 0x88, 0x17, 0x6c, 0x7c, 0x22, 0x6f, 0x1c, 0x79, 0x3c, 0x89,
	0x4f, 0x78, 0x93, 0x12, 0xf0, 0x8f, 0xdc, 0x24, 0xb4, 0x28, 0x85, 0x3e, 0x92, 0x71, 0x74, 0x7d,
	0x33, 0x45, 0xd6, 0x7e, 0xf7, 0x97, 0x07, 0x60, 0xd7, 0x8a, 0x1f, 0xd4, 0x83, 0x0e, 0x5b, 0xf2,
	0xe1, 0xfb, 0xc1, 0xf4, 0x9a, 0x91, 0x17, 0xd4, 0x87, 0x97, 0xe6, 0x38, 0xbe, 0x19, 0xde, 0x92,
	0x1a, 0x3d, 0x83, 0x2e, 0x5b, 0xf2, 0xd1, 0xd5, 0xd8, 0x02, 0xcf, 0x95, 0x7f, 0x8b, 0x86, 0x49,
	0x4c, 0xea, 0xb4, 0x0f, 0xc0, 0x96, 0xfc, 0x96, 0xdd, 0x8d, 0xa2, 0x0f, 0x53, 0xd2, 0x70, 0xf2,
	0x5b, 0x76, 0x37, 0x9b, 0xb3, 0x38, 0x26, 0x4d, 0x27, 0xbf, 0x65, 0x77, 0xc9, 0x8c, 0xb4, 0xcc,
	0xfd, 0xd8, 0x92, 0x4f, 0xa2, 0x24, 0x66, 0x68, 0x68, 0x9f, 0x92, 0x49, 0xb4, 0x64, 0xe4, 0xa5,
	0x8b, 0x40, 0x12, 0x25, 0x0b, 0xd2, 0x39, 0x95, 0x44, 0x4b, 0x36, 0x27, 0xe0, 0xba, 0x22, 0x49,
	0x66, 0xa4, 0x4b, 0xcf, 0xa1, 0xc7, 0x96, 0x7c, 0x11, 0x25, 0xc3, 0xf7, 0xf1, 0x62, 0x30, 0x5f,
	0x10, 0xdf, 0x99, 0x10, 0x61, 0x6e, 0xcf, 0xe5, 0x22, 0x61, 0xd3, 0x11, 0xe9, 0x9b, 0xd1, 0x1c,
	0xc0, 0x70, 0x30, 0x1d, 0xb2, 0x31, 0x39, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x4c, 0xf5,
	0x35, 0xeb, 0x05, 0x00, 0x00,
}
