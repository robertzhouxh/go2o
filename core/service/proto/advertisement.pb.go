// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/advertisement.proto

package proto // import "."

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

// 广告分组
type SAdGroup struct {
	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 是否开放给外部
	Opened               bool     `protobuf:"varint,3,opt,name=Opened,proto3" json:"Opened,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdGroup) Reset()         { *m = SAdGroup{} }
func (m *SAdGroup) String() string { return proto.CompactTextString(m) }
func (*SAdGroup) ProtoMessage()    {}
func (*SAdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{0}
}
func (m *SAdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdGroup.Unmarshal(m, b)
}
func (m *SAdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdGroup.Marshal(b, m, deterministic)
}
func (dst *SAdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdGroup.Merge(dst, src)
}
func (m *SAdGroup) XXX_Size() int {
	return xxx_messageInfo_SAdGroup.Size(m)
}
func (m *SAdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SAdGroup proto.InternalMessageInfo

func (m *SAdGroup) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAdGroup) GetOpened() bool {
	if m != nil {
		return m.Opened
	}
	return false
}

func (m *SAdGroup) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// 广告位编号
type AdPositionId struct {
	GroupId              int64    `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	PositionId           int64    `protobuf:"varint,2,opt,name=PositionId,proto3" json:"PositionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdPositionId) Reset()         { *m = AdPositionId{} }
func (m *AdPositionId) String() string { return proto.CompactTextString(m) }
func (*AdPositionId) ProtoMessage()    {}
func (*AdPositionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{1}
}
func (m *AdPositionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdPositionId.Unmarshal(m, b)
}
func (m *AdPositionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdPositionId.Marshal(b, m, deterministic)
}
func (dst *AdPositionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdPositionId.Merge(dst, src)
}
func (m *AdPositionId) XXX_Size() int {
	return xxx_messageInfo_AdPositionId.Size(m)
}
func (m *AdPositionId) XXX_DiscardUnknown() {
	xxx_messageInfo_AdPositionId.DiscardUnknown(m)
}

var xxx_messageInfo_AdPositionId proto.InternalMessageInfo

func (m *AdPositionId) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *AdPositionId) GetPositionId() int64 {
	if m != nil {
		return m.PositionId
	}
	return 0
}

type AdGroupListResponse struct {
	Value                []*SAdGroup `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AdGroupListResponse) Reset()         { *m = AdGroupListResponse{} }
func (m *AdGroupListResponse) String() string { return proto.CompactTextString(m) }
func (*AdGroupListResponse) ProtoMessage()    {}
func (*AdGroupListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{2}
}
func (m *AdGroupListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupListResponse.Unmarshal(m, b)
}
func (m *AdGroupListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupListResponse.Marshal(b, m, deterministic)
}
func (dst *AdGroupListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupListResponse.Merge(dst, src)
}
func (m *AdGroupListResponse) XXX_Size() int {
	return xxx_messageInfo_AdGroupListResponse.Size(m)
}
func (m *AdGroupListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupListResponse proto.InternalMessageInfo

func (m *AdGroupListResponse) GetValue() []*SAdGroup {
	if m != nil {
		return m.Value
	}
	return nil
}

// 广告位
type SAdPosition struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 分组编号
	GroupId int64 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	// 引用键
	Key string `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	// 名称
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	// todo:广告位类型限制
	// 广告类型限制,0为无限制
	TypeLimit int32 `protobuf:"varint,5,opt,name=TypeLimit,proto3" json:"TypeLimit,omitempty"`
	// 是否开放给外部
	Opened bool `protobuf:"varint,6,opt,name=Opened,proto3" json:"Opened,omitempty"`
	// 是否启用
	Enabled bool `protobuf:"varint,7,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	// 默认广告编号
	DefaultId            int64    `protobuf:"varint,8,opt,name=DefaultId,proto3" json:"DefaultId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdPosition) Reset()         { *m = SAdPosition{} }
func (m *SAdPosition) String() string { return proto.CompactTextString(m) }
func (*SAdPosition) ProtoMessage()    {}
func (*SAdPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{3}
}
func (m *SAdPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdPosition.Unmarshal(m, b)
}
func (m *SAdPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdPosition.Marshal(b, m, deterministic)
}
func (dst *SAdPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdPosition.Merge(dst, src)
}
func (m *SAdPosition) XXX_Size() int {
	return xxx_messageInfo_SAdPosition.Size(m)
}
func (m *SAdPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdPosition.DiscardUnknown(m)
}

var xxx_messageInfo_SAdPosition proto.InternalMessageInfo

func (m *SAdPosition) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdPosition) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *SAdPosition) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SAdPosition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAdPosition) GetTypeLimit() int32 {
	if m != nil {
		return m.TypeLimit
	}
	return 0
}

func (m *SAdPosition) GetOpened() bool {
	if m != nil {
		return m.Opened
	}
	return false
}

func (m *SAdPosition) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *SAdPosition) GetDefaultId() int64 {
	if m != nil {
		return m.DefaultId
	}
	return 0
}

// 广告用户设置
type SAdUserSet struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 广告位编号
	PosId int64 `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId,omitempty"`
	// 广告用户编号
	AdUserId int64 `protobuf:"varint,3,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	// 广告编号
	AdId                 int64    `protobuf:"varint,4,opt,name=AdId,proto3" json:"AdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdUserSet) Reset()         { *m = SAdUserSet{} }
func (m *SAdUserSet) String() string { return proto.CompactTextString(m) }
func (*SAdUserSet) ProtoMessage()    {}
func (*SAdUserSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{4}
}
func (m *SAdUserSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdUserSet.Unmarshal(m, b)
}
func (m *SAdUserSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdUserSet.Marshal(b, m, deterministic)
}
func (dst *SAdUserSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdUserSet.Merge(dst, src)
}
func (m *SAdUserSet) XXX_Size() int {
	return xxx_messageInfo_SAdUserSet.Size(m)
}
func (m *SAdUserSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdUserSet.DiscardUnknown(m)
}

var xxx_messageInfo_SAdUserSet proto.InternalMessageInfo

func (m *SAdUserSet) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdUserSet) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SAdUserSet) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SAdUserSet) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

// 广告
type SAd struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 广告用户编号
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// 广告类型
	Type int32 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	// 展现次数
	ShowTimes int32 `protobuf:"varint,5,opt,name=ShowTimes,proto3" json:"ShowTimes,omitempty"`
	// 点击次数
	ClickTimes int32 `protobuf:"varint,6,opt,name=ClickTimes,proto3" json:"ClickTimes,omitempty"`
	// 展现天数
	ShowDays int32 `protobuf:"varint,7,opt,name=ShowDays,proto3" json:"ShowDays,omitempty"`
	// 修改时间
	UpdateTime           int64    `protobuf:"varint,8,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAd) Reset()         { *m = SAd{} }
func (m *SAd) String() string { return proto.CompactTextString(m) }
func (*SAd) ProtoMessage()    {}
func (*SAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{5}
}
func (m *SAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAd.Unmarshal(m, b)
}
func (m *SAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAd.Marshal(b, m, deterministic)
}
func (dst *SAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAd.Merge(dst, src)
}
func (m *SAd) XXX_Size() int {
	return xxx_messageInfo_SAd.Size(m)
}
func (m *SAd) XXX_DiscardUnknown() {
	xxx_messageInfo_SAd.DiscardUnknown(m)
}

var xxx_messageInfo_SAd proto.InternalMessageInfo

func (m *SAd) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAd) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SAd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SAd) GetShowTimes() int32 {
	if m != nil {
		return m.ShowTimes
	}
	return 0
}

func (m *SAd) GetClickTimes() int32 {
	if m != nil {
		return m.ClickTimes
	}
	return 0
}

func (m *SAd) GetShowDays() int32 {
	if m != nil {
		return m.ShowDays
	}
	return 0
}

func (m *SAd) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

// 广告数据传输对象
type SAdDto struct {
	Id                   int64             `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AdType               int32             `protobuf:"varint,2,opt,name=AdType,proto3" json:"AdType,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SAdDto) Reset()         { *m = SAdDto{} }
func (m *SAdDto) String() string { return proto.CompactTextString(m) }
func (*SAdDto) ProtoMessage()    {}
func (*SAdDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{6}
}
func (m *SAdDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdDto.Unmarshal(m, b)
}
func (m *SAdDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdDto.Marshal(b, m, deterministic)
}
func (dst *SAdDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdDto.Merge(dst, src)
}
func (m *SAdDto) XXX_Size() int {
	return xxx_messageInfo_SAdDto.Size(m)
}
func (m *SAdDto) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdDto.DiscardUnknown(m)
}

var xxx_messageInfo_SAdDto proto.InternalMessageInfo

func (m *SAdDto) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdDto) GetAdType() int32 {
	if m != nil {
		return m.AdType
	}
	return 0
}

func (m *SAdDto) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

//  超链接
type SHyperLink struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	LinkUrl              string   `protobuf:"bytes,4,opt,name=LinkUrl,proto3" json:"LinkUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SHyperLink) Reset()         { *m = SHyperLink{} }
func (m *SHyperLink) String() string { return proto.CompactTextString(m) }
func (*SHyperLink) ProtoMessage()    {}
func (*SHyperLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{7}
}
func (m *SHyperLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SHyperLink.Unmarshal(m, b)
}
func (m *SHyperLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SHyperLink.Marshal(b, m, deterministic)
}
func (dst *SHyperLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SHyperLink.Merge(dst, src)
}
func (m *SHyperLink) XXX_Size() int {
	return xxx_messageInfo_SHyperLink.Size(m)
}
func (m *SHyperLink) XXX_DiscardUnknown() {
	xxx_messageInfo_SHyperLink.DiscardUnknown(m)
}

var xxx_messageInfo_SHyperLink proto.InternalMessageInfo

func (m *SHyperLink) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SHyperLink) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SHyperLink) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SHyperLink) GetLinkUrl() string {
	if m != nil {
		return m.LinkUrl
	}
	return ""
}

// 广告图片
type SImage struct {
	// 图片编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 广告编号
	AdId int64 `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	// 图片标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	// 链接
	LinkUrl string `protobuf:"bytes,4,opt,name=LinkUrl,proto3" json:"LinkUrl,omitempty"`
	// 图片地址
	ImageUrl string `protobuf:"bytes,5,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	// 排列序号
	SortNum int32 `protobuf:"varint,6,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,7,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SImage) Reset()         { *m = SImage{} }
func (m *SImage) String() string { return proto.CompactTextString(m) }
func (*SImage) ProtoMessage()    {}
func (*SImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{8}
}
func (m *SImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SImage.Unmarshal(m, b)
}
func (m *SImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SImage.Marshal(b, m, deterministic)
}
func (dst *SImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SImage.Merge(dst, src)
}
func (m *SImage) XXX_Size() int {
	return xxx_messageInfo_SImage.Size(m)
}
func (m *SImage) XXX_DiscardUnknown() {
	xxx_messageInfo_SImage.DiscardUnknown(m)
}

var xxx_messageInfo_SImage proto.InternalMessageInfo

func (m *SImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SImage) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SImage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SImage) GetLinkUrl() string {
	if m != nil {
		return m.LinkUrl
	}
	return ""
}

func (m *SImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *SImage) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SImage) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SetDefaultAdRequest struct {
	GroupId              int64    `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	PosId                int64    `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId,omitempty"`
	AdId                 int64    `protobuf:"varint,3,opt,name=AdId,proto3" json:"AdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDefaultAdRequest) Reset()         { *m = SetDefaultAdRequest{} }
func (m *SetDefaultAdRequest) String() string { return proto.CompactTextString(m) }
func (*SetDefaultAdRequest) ProtoMessage()    {}
func (*SetDefaultAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{9}
}
func (m *SetDefaultAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefaultAdRequest.Unmarshal(m, b)
}
func (m *SetDefaultAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefaultAdRequest.Marshal(b, m, deterministic)
}
func (dst *SetDefaultAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefaultAdRequest.Merge(dst, src)
}
func (m *SetDefaultAdRequest) XXX_Size() int {
	return xxx_messageInfo_SetDefaultAdRequest.Size(m)
}
func (m *SetDefaultAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefaultAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefaultAdRequest proto.InternalMessageInfo

func (m *SetDefaultAdRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *SetDefaultAdRequest) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SetDefaultAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type SetUserAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	PosId                int64    `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId,omitempty"`
	AdId                 int64    `protobuf:"varint,3,opt,name=AdId,proto3" json:"AdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserAdRequest) Reset()         { *m = SetUserAdRequest{} }
func (m *SetUserAdRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserAdRequest) ProtoMessage()    {}
func (*SetUserAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{10}
}
func (m *SetUserAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserAdRequest.Unmarshal(m, b)
}
func (m *SetUserAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserAdRequest.Marshal(b, m, deterministic)
}
func (dst *SetUserAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserAdRequest.Merge(dst, src)
}
func (m *SetUserAdRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserAdRequest.Size(m)
}
func (m *SetUserAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserAdRequest proto.InternalMessageInfo

func (m *SetUserAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SetUserAdRequest) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SetUserAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type AdIdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdIdRequest) Reset()         { *m = AdIdRequest{} }
func (m *AdIdRequest) String() string { return proto.CompactTextString(m) }
func (*AdIdRequest) ProtoMessage()    {}
func (*AdIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{11}
}
func (m *AdIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdIdRequest.Unmarshal(m, b)
}
func (m *AdIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdIdRequest.Marshal(b, m, deterministic)
}
func (dst *AdIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdIdRequest.Merge(dst, src)
}
func (m *AdIdRequest) XXX_Size() int {
	return xxx_messageInfo_AdIdRequest.Size(m)
}
func (m *AdIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdIdRequest proto.InternalMessageInfo

func (m *AdIdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *AdIdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type AdKeyRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	AdPosKey             string   `protobuf:"bytes,2,opt,name=AdPosKey,proto3" json:"AdPosKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdKeyRequest) Reset()         { *m = AdKeyRequest{} }
func (m *AdKeyRequest) String() string { return proto.CompactTextString(m) }
func (*AdKeyRequest) ProtoMessage()    {}
func (*AdKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{12}
}
func (m *AdKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdKeyRequest.Unmarshal(m, b)
}
func (m *AdKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdKeyRequest.Marshal(b, m, deterministic)
}
func (dst *AdKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdKeyRequest.Merge(dst, src)
}
func (m *AdKeyRequest) XXX_Size() int {
	return xxx_messageInfo_AdKeyRequest.Size(m)
}
func (m *AdKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdKeyRequest proto.InternalMessageInfo

func (m *AdKeyRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *AdKeyRequest) GetAdPosKey() string {
	if m != nil {
		return m.AdPosKey
	}
	return ""
}

type SaveAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	Value                *SAd     `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveAdRequest) Reset()         { *m = SaveAdRequest{} }
func (m *SaveAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveAdRequest) ProtoMessage()    {}
func (*SaveAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{13}
}
func (m *SaveAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAdRequest.Unmarshal(m, b)
}
func (m *SaveAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAdRequest.Merge(dst, src)
}
func (m *SaveAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveAdRequest.Size(m)
}
func (m *SaveAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAdRequest proto.InternalMessageInfo

func (m *SaveAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveAdRequest) GetValue() *SAd {
	if m != nil {
		return m.Value
	}
	return nil
}

type SaveLinkAdRequest struct {
	AdUserId             int64       `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	AdId                 int64       `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	Value                *SHyperLink `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SaveLinkAdRequest) Reset()         { *m = SaveLinkAdRequest{} }
func (m *SaveLinkAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveLinkAdRequest) ProtoMessage()    {}
func (*SaveLinkAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{14}
}
func (m *SaveLinkAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveLinkAdRequest.Unmarshal(m, b)
}
func (m *SaveLinkAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveLinkAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveLinkAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveLinkAdRequest.Merge(dst, src)
}
func (m *SaveLinkAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveLinkAdRequest.Size(m)
}
func (m *SaveLinkAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveLinkAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveLinkAdRequest proto.InternalMessageInfo

func (m *SaveLinkAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveLinkAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SaveLinkAdRequest) GetValue() *SHyperLink {
	if m != nil {
		return m.Value
	}
	return nil
}

type SaveImageAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	Value                *SImage  `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveImageAdRequest) Reset()         { *m = SaveImageAdRequest{} }
func (m *SaveImageAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveImageAdRequest) ProtoMessage()    {}
func (*SaveImageAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{15}
}
func (m *SaveImageAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveImageAdRequest.Unmarshal(m, b)
}
func (m *SaveImageAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveImageAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveImageAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveImageAdRequest.Merge(dst, src)
}
func (m *SaveImageAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveImageAdRequest.Size(m)
}
func (m *SaveImageAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveImageAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveImageAdRequest proto.InternalMessageInfo

func (m *SaveImageAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveImageAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SaveImageAdRequest) GetValue() *SImage {
	if m != nil {
		return m.Value
	}
	return nil
}

type ImageAdIdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId,omitempty"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId,omitempty"`
	ImageId              int64    `protobuf:"varint,3,opt,name=ImageId,proto3" json:"ImageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageAdIdRequest) Reset()         { *m = ImageAdIdRequest{} }
func (m *ImageAdIdRequest) String() string { return proto.CompactTextString(m) }
func (*ImageAdIdRequest) ProtoMessage()    {}
func (*ImageAdIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_d0385e73d21f1d07, []int{16}
}
func (m *ImageAdIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageAdIdRequest.Unmarshal(m, b)
}
func (m *ImageAdIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageAdIdRequest.Marshal(b, m, deterministic)
}
func (dst *ImageAdIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageAdIdRequest.Merge(dst, src)
}
func (m *ImageAdIdRequest) XXX_Size() int {
	return xxx_messageInfo_ImageAdIdRequest.Size(m)
}
func (m *ImageAdIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageAdIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageAdIdRequest proto.InternalMessageInfo

func (m *ImageAdIdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *ImageAdIdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *ImageAdIdRequest) GetImageId() int64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func init() {
	proto.RegisterType((*SAdGroup)(nil), "SAdGroup")
	proto.RegisterType((*AdPositionId)(nil), "AdPositionId")
	proto.RegisterType((*AdGroupListResponse)(nil), "AdGroupListResponse")
	proto.RegisterType((*SAdPosition)(nil), "SAdPosition")
	proto.RegisterType((*SAdUserSet)(nil), "SAdUserSet")
	proto.RegisterType((*SAd)(nil), "SAd")
	proto.RegisterType((*SAdDto)(nil), "SAdDto")
	proto.RegisterMapType((map[string]string)(nil), "SAdDto.DataEntry")
	proto.RegisterType((*SHyperLink)(nil), "SHyperLink")
	proto.RegisterType((*SImage)(nil), "SImage")
	proto.RegisterType((*SetDefaultAdRequest)(nil), "SetDefaultAdRequest")
	proto.RegisterType((*SetUserAdRequest)(nil), "SetUserAdRequest")
	proto.RegisterType((*AdIdRequest)(nil), "AdIdRequest")
	proto.RegisterType((*AdKeyRequest)(nil), "AdKeyRequest")
	proto.RegisterType((*SaveAdRequest)(nil), "SaveAdRequest")
	proto.RegisterType((*SaveLinkAdRequest)(nil), "SaveLinkAdRequest")
	proto.RegisterType((*SaveImageAdRequest)(nil), "SaveImageAdRequest")
	proto.RegisterType((*ImageAdIdRequest)(nil), "ImageAdIdRequest")
}

func init() {
	proto.RegisterFile("message/advertisement.proto", fileDescriptor_advertisement_d0385e73d21f1d07)
}

var fileDescriptor_advertisement_d0385e73d21f1d07 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x96, 0xe3, 0xfc, 0x4e, 0xce, 0x39, 0x6a, 0xdd, 0xa3, 0xca, 0xea, 0x39, 0x94, 0x60, 0x09,
	0x29, 0x57, 0x46, 0x2a, 0x12, 0x20, 0x10, 0x17, 0x86, 0x94, 0x36, 0x6a, 0x55, 0xaa, 0x75, 0x8b,
	0x00, 0x71, 0x51, 0x27, 0x3b, 0x4d, 0xad, 0xc6, 0xb1, 0xb1, 0x37, 0x45, 0x79, 0x0b, 0x1e, 0x84,
	0x37, 0xe1, 0x82, 0x57, 0x42, 0x33, 0xfe, 0x2d, 0x4d, 0x51, 0x54, 0x89, 0xab, 0xec, 0x37, 0xe3,
	0xfd, 0x66, 0xe7, 0x9b, 0x6f, 0x37, 0xf0, 0x5f, 0x80, 0x49, 0xe2, 0x4d, 0xf0, 0x91, 0x27, 0xaf,
	0x30, 0x56, 0x7e, 0x82, 0x01, 0xce, 0x94, 0x1d, 0xc5, 0xa1, 0x0a, 0xad, 0x33, 0x68, 0xbb, 0x8e,
	0xdc, 0x8b, 0xc3, 0x79, 0x64, 0xfc, 0x03, 0xb5, 0xa1, 0x34, 0xb5, 0x9e, 0xd6, 0xd7, 0x45, 0x6d,
	0x28, 0x0d, 0x03, 0xea, 0x47, 0x5e, 0x80, 0x66, 0xad, 0xa7, 0xf5, 0x3b, 0x82, 0xd7, 0xc6, 0x26,
	0x34, 0xdf, 0x46, 0x38, 0x43, 0x69, 0xea, 0x3d, 0xad, 0xdf, 0x16, 0x19, 0x32, 0x4c, 0x68, 0xed,
	0xce, 0xbc, 0xd1, 0x14, 0xa5, 0x59, 0xe7, 0x44, 0x0e, 0xad, 0x7d, 0xf8, 0xcb, 0x91, 0xc7, 0x61,
	0xe2, 0x2b, 0x3f, 0x9c, 0x0d, 0xf9, 0x4b, 0x2e, 0x57, 0x94, 0xca, 0xa1, 0xb1, 0x0d, 0x50, 0x7e,
	0xc7, 0x55, 0x75, 0x51, 0x89, 0x58, 0x4f, 0x60, 0x23, 0x3b, 0xea, 0xa1, 0x9f, 0x28, 0x81, 0x49,
	0x14, 0xce, 0x12, 0x34, 0xee, 0x43, 0xe3, 0x9d, 0x37, 0x9d, 0xa3, 0xa9, 0xf5, 0xf4, 0x7e, 0x77,
	0xa7, 0x63, 0xe7, 0x0d, 0x89, 0x34, 0x6e, 0x7d, 0xd7, 0xa0, 0xeb, 0x96, 0x67, 0xb8, 0xd1, 0x67,
	0xe5, 0x44, 0xb5, 0xeb, 0x27, 0x5a, 0x03, 0xfd, 0x00, 0x17, 0xdc, 0x6a, 0x47, 0xd0, 0xb2, 0xd0,
	0xa4, 0x5e, 0xd1, 0xe4, 0x7f, 0xe8, 0x9c, 0x2c, 0x22, 0x3c, 0xf4, 0x03, 0x5f, 0x99, 0x8d, 0x9e,
	0xd6, 0x6f, 0x88, 0x32, 0x50, 0x51, 0xac, 0x79, 0x9b, 0x62, 0xad, 0x6b, 0x8a, 0x11, 0xdf, 0x00,
	0xcf, 0xbd, 0xf9, 0x54, 0x0d, 0xa5, 0xd9, 0xe6, 0x13, 0x95, 0x01, 0x6b, 0x04, 0xe0, 0x3a, 0xf2,
	0x34, 0xc1, 0xd8, 0x45, 0x75, 0xa3, 0x97, 0x7f, 0xa1, 0x71, 0x1c, 0x26, 0x45, 0x27, 0x29, 0x30,
	0xb6, 0xa0, 0x9d, 0x6e, 0x19, 0xa6, 0x73, 0xd3, 0x45, 0x81, 0xa9, 0x23, 0x47, 0x0e, 0xd3, 0xb1,
	0xe9, 0x82, 0xd7, 0xd6, 0x0f, 0x0d, 0x74, 0xd7, 0x91, 0x37, 0xd8, 0x37, 0xa1, 0x99, 0xb1, 0xa4,
	0xf4, 0xcd, 0x92, 0x83, 0x55, 0xd1, 0x2b, 0xaa, 0x18, 0x50, 0x27, 0x11, 0x98, 0xb7, 0x21, 0x78,
	0x4d, 0x9d, 0xb9, 0x17, 0xe1, 0x97, 0x13, 0x3f, 0xc0, 0x24, 0x57, 0xaa, 0x08, 0xd0, 0xfc, 0x5f,
	0x4f, 0xfd, 0xf1, 0x65, 0x9a, 0x6e, 0x72, 0xba, 0x12, 0xa1, 0x2e, 0xe8, 0xe3, 0x81, 0xb7, 0x48,
	0x58, 0xb2, 0x86, 0x28, 0x30, 0xed, 0x3d, 0x8d, 0xa4, 0xa7, 0x90, 0x3e, 0xcd, 0x44, 0xab, 0x44,
	0xac, 0xaf, 0x1a, 0x34, 0x5d, 0x47, 0x0e, 0x54, 0xb8, 0xac, 0x29, 0x47, 0xf2, 0x51, 0x6b, 0x4c,
	0x9a, 0x21, 0xe3, 0x21, 0xd4, 0x07, 0x9e, 0xf2, 0x4c, 0x9d, 0x6d, 0xb5, 0x6e, 0xa7, 0xdb, 0x6d,
	0x8a, 0xed, 0xce, 0x54, 0xbc, 0x10, 0x9c, 0xde, 0x7a, 0x0a, 0x9d, 0x22, 0x44, 0x86, 0xb9, 0xc4,
	0x05, 0x93, 0x77, 0x04, 0x2d, 0x69, 0x20, 0x57, 0xec, 0xce, 0xf4, 0x16, 0xa5, 0xe0, 0x79, 0xed,
	0x99, 0x66, 0x9d, 0x01, 0xb8, 0xfb, 0x8b, 0x08, 0xe3, 0x43, 0x7f, 0x76, 0xb9, 0xec, 0xf2, 0xf1,
	0x58, 0x6a, 0xe5, 0x58, 0x88, 0xeb, 0xc4, 0x57, 0xd3, 0x5c, 0xe7, 0x14, 0x90, 0x91, 0x88, 0xe1,
	0x34, 0x9e, 0x66, 0xae, 0xcc, 0xa1, 0xf5, 0x8d, 0x9a, 0x1e, 0x06, 0xde, 0x04, 0xff, 0x04, 0x3d,
	0xcd, 0x83, 0xc9, 0x29, 0xd5, 0xe0, 0x54, 0x81, 0x69, 0x97, 0x1b, 0xc6, 0xea, 0x68, 0x1e, 0x64,
	0x83, 0xcc, 0xe1, 0xed, 0xbe, 0xb7, 0x3e, 0xc0, 0x86, 0x8b, 0x2a, 0x73, 0xba, 0x23, 0x05, 0x7e,
	0x9e, 0x63, 0xa2, 0x7e, 0xf3, 0x60, 0x2c, 0x37, 0x7b, 0xde, 0x9a, 0x5e, 0x31, 0xf4, 0x7b, 0x58,
	0x73, 0x51, 0x91, 0x5b, 0x4b, 0xde, 0xea, 0xa5, 0xd0, 0x7e, 0xb9, 0x14, 0xab, 0x33, 0xbf, 0x84,
	0x2e, 0xfd, 0xae, 0x42, 0xba, 0x44, 0x73, 0xeb, 0x0d, 0xbd, 0x8e, 0x07, 0xb8, 0x58, 0x65, 0x3f,
	0xe7, 0x8e, 0xc3, 0x84, 0x9e, 0xa4, 0xd4, 0x4d, 0x05, 0xb6, 0xf6, 0xe0, 0x6f, 0xd7, 0xbb, 0xc2,
	0xd5, 0xba, 0xdb, 0xca, 0x5f, 0x4c, 0x62, 0xe9, 0xee, 0xd4, 0xc9, 0xda, 0xf9, 0x63, 0x79, 0x0e,
	0xeb, 0x44, 0x44, 0x33, 0x76, 0xee, 0xda, 0x95, 0xf1, 0x20, 0x2f, 0xa0, 0x73, 0x81, 0xae, 0x5d,
	0x1a, 0x3d, 0xaf, 0x33, 0x06, 0x83, 0xea, 0xb0, 0x61, 0xee, 0x5e, 0xe8, 0xde, 0xf5, 0x42, 0x2d,
	0x3b, 0xb5, 0x7b, 0x5e, 0xe4, 0x13, 0xac, 0x65, 0x05, 0xee, 0x3c, 0x21, 0xb2, 0x1f, 0x73, 0x14,
	0x73, 0xcf, 0xe1, 0xab, 0x6d, 0xd8, 0x18, 0x87, 0x81, 0x3d, 0xf1, 0xd5, 0xc5, 0x7c, 0x64, 0x4f,
	0xc2, 0x9d, 0xd0, 0x8e, 0xa3, 0xf1, 0xc7, 0x96, 0xfd, 0x82, 0xff, 0x5b, 0x47, 0x4d, 0xfe, 0x79,
	0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xed, 0x5b, 0x80, 0xfa, 0x81, 0x07, 0x00, 0x00,
}
