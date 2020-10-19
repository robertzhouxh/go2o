// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransferInRequest struct {
	PersonId             int64    `protobuf:"zigzag64,1,opt,name=personId,proto3" json:"personId,omitempty"`
	TransferWith         int32    `protobuf:"zigzag32,2,opt,name=transferWith,proto3" json:"transferWith,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferInRequest) Reset()         { *m = TransferInRequest{} }
func (m *TransferInRequest) String() string { return proto.CompactTextString(m) }
func (*TransferInRequest) ProtoMessage()    {}
func (*TransferInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{0}
}
func (m *TransferInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferInRequest.Unmarshal(m, b)
}
func (m *TransferInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferInRequest.Marshal(b, m, deterministic)
}
func (dst *TransferInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferInRequest.Merge(dst, src)
}
func (m *TransferInRequest) XXX_Size() int {
	return xxx_messageInfo_TransferInRequest.Size(m)
}
func (m *TransferInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferInRequest proto.InternalMessageInfo

func (m *TransferInRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *TransferInRequest) GetTransferWith() int32 {
	if m != nil {
		return m.TransferWith
	}
	return 0
}

func (m *TransferInRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type PersonId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonId) Reset()         { *m = PersonId{} }
func (m *PersonId) String() string { return proto.CompactTextString(m) }
func (*PersonId) ProtoMessage()    {}
func (*PersonId) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{1}
}
func (m *PersonId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonId.Unmarshal(m, b)
}
func (m *PersonId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonId.Marshal(b, m, deterministic)
}
func (dst *PersonId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonId.Merge(dst, src)
}
func (m *PersonId) XXX_Size() int {
	return xxx_messageInfo_PersonId.Size(m)
}
func (m *PersonId) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonId.DiscardUnknown(m)
}

var xxx_messageInfo_PersonId proto.InternalMessageInfo

func (m *PersonId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 收益总记录
type SRiseInfo struct {
	// 人员编号
	PersonId int64 `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`
	// 本金及收益的余额
	Balance float64 `protobuf:"fixed64,2,opt,name=Balance,proto3" json:"Balance,omitempty"`
	// 　结算金额
	SettlementAmount float64 `protobuf:"fixed64,3,opt,name=SettlementAmount,proto3" json:"SettlementAmount,omitempty"`
	// 当前的收益
	Rise float64 `protobuf:"fixed64,4,opt,name=Rise,proto3" json:"Rise,omitempty"`
	// 今日转入
	TransferIn float64 `protobuf:"fixed64,5,opt,name=TransferIn,proto3" json:"TransferIn,omitempty"`
	// 总金额
	TotalAmount float64 `protobuf:"fixed64,6,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	// 总收益
	TotalRise float64 `protobuf:"fixed64,7,opt,name=TotalRise,proto3" json:"TotalRise,omitempty"`
	// 结算日期,用于筛选需要结算的数据
	SettledDate int64 `protobuf:"varint,8,opt,name=SettledDate,proto3" json:"SettledDate,omitempty"`
	// 更新时间
	UpdateTime           int64    `protobuf:"varint,9,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SRiseInfo) Reset()         { *m = SRiseInfo{} }
func (m *SRiseInfo) String() string { return proto.CompactTextString(m) }
func (*SRiseInfo) ProtoMessage()    {}
func (*SRiseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{2}
}
func (m *SRiseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SRiseInfo.Unmarshal(m, b)
}
func (m *SRiseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SRiseInfo.Marshal(b, m, deterministic)
}
func (dst *SRiseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRiseInfo.Merge(dst, src)
}
func (m *SRiseInfo) XXX_Size() int {
	return xxx_messageInfo_SRiseInfo.Size(m)
}
func (m *SRiseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SRiseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SRiseInfo proto.InternalMessageInfo

func (m *SRiseInfo) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *SRiseInfo) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *SRiseInfo) GetSettlementAmount() float64 {
	if m != nil {
		return m.SettlementAmount
	}
	return 0
}

func (m *SRiseInfo) GetRise() float64 {
	if m != nil {
		return m.Rise
	}
	return 0
}

func (m *SRiseInfo) GetTransferIn() float64 {
	if m != nil {
		return m.TransferIn
	}
	return 0
}

func (m *SRiseInfo) GetTotalAmount() float64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *SRiseInfo) GetTotalRise() float64 {
	if m != nil {
		return m.TotalRise
	}
	return 0
}

func (m *SRiseInfo) GetSettledDate() int64 {
	if m != nil {
		return m.SettledDate
	}
	return 0
}

func (m *SRiseInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type RiseSettleRequest struct {
	PersonId             int64    `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`
	SettleDay            int64    `protobuf:"varint,2,opt,name=SettleDay,proto3" json:"SettleDay,omitempty"`
	Ratio                float64  `protobuf:"fixed64,3,opt,name=Ratio,proto3" json:"Ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiseSettleRequest) Reset()         { *m = RiseSettleRequest{} }
func (m *RiseSettleRequest) String() string { return proto.CompactTextString(m) }
func (*RiseSettleRequest) ProtoMessage()    {}
func (*RiseSettleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{3}
}
func (m *RiseSettleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiseSettleRequest.Unmarshal(m, b)
}
func (m *RiseSettleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiseSettleRequest.Marshal(b, m, deterministic)
}
func (dst *RiseSettleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiseSettleRequest.Merge(dst, src)
}
func (m *RiseSettleRequest) XXX_Size() int {
	return xxx_messageInfo_RiseSettleRequest.Size(m)
}
func (m *RiseSettleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RiseSettleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RiseSettleRequest proto.InternalMessageInfo

func (m *RiseSettleRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *RiseSettleRequest) GetSettleDay() int64 {
	if m != nil {
		return m.SettleDay
	}
	return 0
}

func (m *RiseSettleRequest) GetRatio() float64 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

type RiseTransferOutRequest struct {
	PersonId             int64    `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`
	TransferWith         int64    `protobuf:"varint,2,opt,name=TransferWith,proto3" json:"TransferWith,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiseTransferOutRequest) Reset()         { *m = RiseTransferOutRequest{} }
func (m *RiseTransferOutRequest) String() string { return proto.CompactTextString(m) }
func (*RiseTransferOutRequest) ProtoMessage()    {}
func (*RiseTransferOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{4}
}
func (m *RiseTransferOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiseTransferOutRequest.Unmarshal(m, b)
}
func (m *RiseTransferOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiseTransferOutRequest.Marshal(b, m, deterministic)
}
func (dst *RiseTransferOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiseTransferOutRequest.Merge(dst, src)
}
func (m *RiseTransferOutRequest) XXX_Size() int {
	return xxx_messageInfo_RiseTransferOutRequest.Size(m)
}
func (m *RiseTransferOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RiseTransferOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RiseTransferOutRequest proto.InternalMessageInfo

func (m *RiseTransferOutRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *RiseTransferOutRequest) GetTransferWith() int64 {
	if m != nil {
		return m.TransferWith
	}
	return 0
}

func (m *RiseTransferOutRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type CommitTransferRequest struct {
	PersonId             int64    `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId,omitempty"`
	LogId                int64    `protobuf:"varint,2,opt,name=LogId,proto3" json:"LogId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitTransferRequest) Reset()         { *m = CommitTransferRequest{} }
func (m *CommitTransferRequest) String() string { return proto.CompactTextString(m) }
func (*CommitTransferRequest) ProtoMessage()    {}
func (*CommitTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_e16057188c66f2fc, []int{5}
}
func (m *CommitTransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitTransferRequest.Unmarshal(m, b)
}
func (m *CommitTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitTransferRequest.Marshal(b, m, deterministic)
}
func (dst *CommitTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitTransferRequest.Merge(dst, src)
}
func (m *CommitTransferRequest) XXX_Size() int {
	return xxx_messageInfo_CommitTransferRequest.Size(m)
}
func (m *CommitTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitTransferRequest proto.InternalMessageInfo

func (m *CommitTransferRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *CommitTransferRequest) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferInRequest)(nil), "TransferInRequest")
	proto.RegisterType((*PersonId)(nil), "PersonId")
	proto.RegisterType((*SRiseInfo)(nil), "SRiseInfo")
	proto.RegisterType((*RiseSettleRequest)(nil), "RiseSettleRequest")
	proto.RegisterType((*RiseTransferOutRequest)(nil), "RiseTransferOutRequest")
	proto.RegisterType((*CommitTransferRequest)(nil), "CommitTransferRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinanceServiceClient is the client API for FinanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinanceServiceClient interface {
	// 获取用户的账户信息
	GetRiseInfo(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*SRiseInfo, error)
	// 转入
	RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error)
	// 转出
	RiseTransferOut(ctx context.Context, in *RiseTransferOutRequest, opts ...grpc.CallOption) (*Result, error)
	// 结算收益(按日期每天结息)
	RiseSettleByDay(ctx context.Context, in *RiseSettleRequest, opts ...grpc.CallOption) (*Result, error)
	// 提交转入/转出日志
	CommitTransfer(ctx context.Context, in *CommitTransferRequest, opts ...grpc.CallOption) (*Result, error)
	// 开通增利服务
	OpenRiseService(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Result, error)
}

type financeServiceClient struct {
	cc *grpc.ClientConn
}

func NewFinanceServiceClient(cc *grpc.ClientConn) FinanceServiceClient {
	return &financeServiceClient{cc}
}

func (c *financeServiceClient) GetRiseInfo(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*SRiseInfo, error) {
	out := new(SRiseInfo)
	err := c.cc.Invoke(ctx, "/FinanceService/GetRiseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseTransferIn(ctx context.Context, in *TransferInRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseTransferIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseTransferOut(ctx context.Context, in *RiseTransferOutRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseTransferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) RiseSettleByDay(ctx context.Context, in *RiseSettleRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/RiseSettleByDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) CommitTransfer(ctx context.Context, in *CommitTransferRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/CommitTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeServiceClient) OpenRiseService(ctx context.Context, in *PersonId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FinanceService/OpenRiseService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServiceServer is the server API for FinanceService service.
type FinanceServiceServer interface {
	// 获取用户的账户信息
	GetRiseInfo(context.Context, *PersonId) (*SRiseInfo, error)
	// 转入
	RiseTransferIn(context.Context, *TransferInRequest) (*Result, error)
	// 转出
	RiseTransferOut(context.Context, *RiseTransferOutRequest) (*Result, error)
	// 结算收益(按日期每天结息)
	RiseSettleByDay(context.Context, *RiseSettleRequest) (*Result, error)
	// 提交转入/转出日志
	CommitTransfer(context.Context, *CommitTransferRequest) (*Result, error)
	// 开通增利服务
	OpenRiseService(context.Context, *PersonId) (*Result, error)
}

func RegisterFinanceServiceServer(s *grpc.Server, srv FinanceServiceServer) {
	s.RegisterService(&_FinanceService_serviceDesc, srv)
}

func _FinanceService_GetRiseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetRiseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/GetRiseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetRiseInfo(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseTransferIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseTransferIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseTransferIn(ctx, req.(*TransferInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseTransferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiseTransferOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseTransferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseTransferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseTransferOut(ctx, req.(*RiseTransferOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_RiseSettleByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiseSettleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).RiseSettleByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/RiseSettleByDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).RiseSettleByDay(ctx, req.(*RiseSettleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_CommitTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).CommitTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/CommitTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).CommitTransfer(ctx, req.(*CommitTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceService_OpenRiseService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).OpenRiseService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceService/OpenRiseService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).OpenRiseService(ctx, req.(*PersonId))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FinanceService",
	HandlerType: (*FinanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRiseInfo",
			Handler:    _FinanceService_GetRiseInfo_Handler,
		},
		{
			MethodName: "RiseTransferIn",
			Handler:    _FinanceService_RiseTransferIn_Handler,
		},
		{
			MethodName: "RiseTransferOut",
			Handler:    _FinanceService_RiseTransferOut_Handler,
		},
		{
			MethodName: "RiseSettleByDay",
			Handler:    _FinanceService_RiseSettleByDay_Handler,
		},
		{
			MethodName: "CommitTransfer",
			Handler:    _FinanceService_CommitTransfer_Handler,
		},
		{
			MethodName: "OpenRiseService",
			Handler:    _FinanceService_OpenRiseService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance_service.proto",
}

func init() {
	proto.RegisterFile("finance_service.proto", fileDescriptor_finance_service_e16057188c66f2fc)
}

var fileDescriptor_finance_service_e16057188c66f2fc = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xda, 0xf5, 0xeb, 0xae, 0xea, 0xa8, 0xd9, 0x4a, 0x54, 0x4d, 0x53, 0xe5, 0x07, 0xa8,
	0x78, 0x30, 0x68, 0x13, 0x4f, 0x3c, 0x51, 0x26, 0x50, 0x24, 0xa4, 0x4d, 0x6e, 0x01, 0x89, 0x17,
	0xe4, 0xa6, 0x6e, 0x17, 0x91, 0xd8, 0x21, 0x71, 0x90, 0xf6, 0x7b, 0xf8, 0x03, 0xfc, 0x44, 0x64,
	0x3b, 0x69, 0xdc, 0x75, 0x93, 0xf6, 0xd4, 0x9e, 0xe3, 0x7b, 0xee, 0xc7, 0xb9, 0xb7, 0x85, 0xd3,
	0x75, 0x24, 0x98, 0x08, 0xf9, 0xcf, 0x9c, 0x67, 0x7f, 0xa2, 0x90, 0x93, 0x34, 0x93, 0x4a, 0x8e,
	0xfb, 0x9b, 0x58, 0x2e, 0x59, 0x6c, 0x11, 0xfe, 0x05, 0xc3, 0x45, 0xc6, 0x44, 0xbe, 0xe6, 0x59,
	0x20, 0x28, 0xff, 0x5d, 0xf0, 0x5c, 0xa1, 0x31, 0x74, 0x53, 0x9e, 0xe5, 0x52, 0x04, 0x2b, 0xdf,
	0x9b, 0x78, 0x53, 0x44, 0xb7, 0x18, 0x61, 0xe8, 0xab, 0x52, 0xf0, 0x3d, 0x52, 0xb7, 0x7e, 0x63,
	0xe2, 0x4d, 0x87, 0x74, 0x87, 0x43, 0x23, 0x68, 0xb3, 0x44, 0x16, 0x42, 0xf9, 0xcd, 0x89, 0x37,
	0xf5, 0x68, 0x89, 0xf0, 0x04, 0xba, 0x37, 0x55, 0x9e, 0x13, 0x68, 0x7d, 0x63, 0x71, 0xc1, 0x4d,
	0x81, 0x26, 0xb5, 0x00, 0xff, 0x6d, 0x40, 0x6f, 0x4e, 0xa3, 0x9c, 0x07, 0x62, 0x2d, 0x75, 0x1f,
	0x37, 0x6e, 0x1f, 0x4d, 0x5a, 0xeb, 0x7d, 0xe8, 0xcc, 0x58, 0xac, 0xe7, 0x33, 0x2d, 0x78, 0xb4,
	0x82, 0xe8, 0x35, 0x3c, 0x9b, 0x73, 0xa5, 0x62, 0x9e, 0x70, 0xa1, 0x3e, 0xb8, 0x7d, 0xec, 0xf1,
	0x08, 0xc1, 0xa1, 0xae, 0xe6, 0x1f, 0x9a, 0x77, 0xf3, 0x1d, 0x9d, 0x03, 0xd4, 0x96, 0xf8, 0x2d,
	0xf3, 0xe2, 0x30, 0x68, 0x02, 0x47, 0x0b, 0xa9, 0x58, 0x5c, 0xa6, 0x6e, 0x9b, 0x00, 0x97, 0x42,
	0x67, 0xd0, 0x33, 0xd0, 0xa4, 0xee, 0x98, 0xf7, 0x9a, 0xd0, 0x7a, 0xdb, 0xc7, 0xea, 0x8a, 0x29,
	0xee, 0x77, 0xcd, 0x60, 0x2e, 0xa5, 0x3b, 0xf8, 0x9a, 0xae, 0x98, 0xe2, 0x8b, 0x28, 0xe1, 0x7e,
	0xcf, 0x04, 0x38, 0x0c, 0x0e, 0x61, 0xa8, 0x33, 0x59, 0x89, 0xb3, 0xb4, 0x47, 0xcd, 0x3a, 0x83,
	0x9e, 0x0d, 0xbe, 0x62, 0x77, 0xc6, 0xae, 0x26, 0xad, 0x09, 0xbd, 0x0a, 0xca, 0x54, 0x24, 0x4b,
	0x97, 0x2c, 0xc0, 0x29, 0x8c, 0x74, 0x91, 0x6a, 0xf0, 0xeb, 0x42, 0x3d, 0xa5, 0x12, 0x86, 0xfe,
	0xe2, 0xfe, 0x79, 0x34, 0xe9, 0x0e, 0xa7, 0xcf, 0x63, 0x67, 0x2d, 0x25, 0xc2, 0x01, 0x9c, 0x7e,
	0x94, 0x49, 0x12, 0xa9, 0x2a, 0xfa, 0x29, 0x05, 0x4f, 0xa0, 0xf5, 0x45, 0x6e, 0x82, 0x55, 0x59,
	0xc9, 0x82, 0x8b, 0x7f, 0x0d, 0x18, 0x7c, 0xb2, 0xe7, 0x3f, 0xb7, 0xd7, 0x8f, 0x5e, 0xc2, 0xd1,
	0x67, 0xae, 0xb6, 0xb7, 0xd5, 0x23, 0x55, 0x8a, 0x31, 0x90, 0xed, 0xc9, 0xe1, 0x03, 0xf4, 0x06,
	0x06, 0xee, 0xdc, 0x81, 0x40, 0x88, 0xec, 0xfd, 0x44, 0xc6, 0x1d, 0x42, 0x79, 0x5e, 0xc4, 0x0a,
	0x1f, 0xa0, 0x77, 0x70, 0x7c, 0xcf, 0x28, 0xf4, 0x82, 0x3c, 0x6c, 0x9d, 0x2b, 0x7b, 0x6b, 0x65,
	0x76, 0x0d, 0xb3, 0x3b, 0xbd, 0x08, 0x44, 0xf6, 0xd6, 0xea, 0x2a, 0x2e, 0x61, 0xb0, 0xeb, 0x0f,
	0x1a, 0x91, 0x07, 0x0d, 0x73, 0x45, 0xaf, 0xe0, 0xf8, 0x3a, 0xe5, 0xc2, 0x26, 0xb6, 0x4e, 0x38,
	0xa3, 0xd7, 0x81, 0xb3, 0x73, 0x78, 0x1e, 0xca, 0x84, 0x6c, 0x22, 0x75, 0x5b, 0x2c, 0xc9, 0x46,
	0x5e, 0x48, 0x92, 0xa5, 0xe1, 0x8f, 0x0e, 0x79, 0x6f, 0xfe, 0x29, 0x96, 0x6d, 0xf3, 0x71, 0xf9,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0xad, 0x62, 0x2d, 0x62, 0x57, 0x04, 0x00, 0x00,
}
