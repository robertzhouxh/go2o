// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance_service.proto

package proto // import "./"

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
	PersonId             int64    `protobuf:"zigzag64,1,opt,name=personId,proto3" json:"personId"`
	TransferWith         int32    `protobuf:"zigzag32,2,opt,name=transferWith,proto3" json:"transferWith"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferInRequest) Reset()         { *m = TransferInRequest{} }
func (m *TransferInRequest) String() string { return proto.CompactTextString(m) }
func (*TransferInRequest) ProtoMessage()    {}
func (*TransferInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{0}
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
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonId) Reset()         { *m = PersonId{} }
func (m *PersonId) String() string { return proto.CompactTextString(m) }
func (*PersonId) ProtoMessage()    {}
func (*PersonId) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{1}
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
	PersonId int64 `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId"`
	// 本金及收益的余额
	Balance float64 `protobuf:"fixed64,2,opt,name=Balance,proto3" json:"Balance"`
	// 结算金额
	SettlementAmount float64 `protobuf:"fixed64,3,opt,name=SettlementAmount,proto3" json:"SettlementAmount"`
	// 当前的收益
	Rise float64 `protobuf:"fixed64,4,opt,name=Rise,proto3" json:"Rise"`
	// 今日转入
	TransferIn float64 `protobuf:"fixed64,5,opt,name=TransferIn,proto3" json:"TransferIn"`
	// 总金额
	TotalAmount float64 `protobuf:"fixed64,6,opt,name=TotalAmount,proto3" json:"TotalAmount"`
	// 总收益
	TotalRise float64 `protobuf:"fixed64,7,opt,name=TotalRise,proto3" json:"TotalRise"`
	// 结算日期,用于筛选需要结算的数据
	SettledDate int64 `protobuf:"varint,8,opt,name=SettledDate,proto3" json:"SettledDate"`
	// 更新时间
	UpdateTime           int64    `protobuf:"varint,9,opt,name=UpdateTime,proto3" json:"UpdateTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SRiseInfo) Reset()         { *m = SRiseInfo{} }
func (m *SRiseInfo) String() string { return proto.CompactTextString(m) }
func (*SRiseInfo) ProtoMessage()    {}
func (*SRiseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{2}
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
	PersonId             int64    `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId"`
	SettleDay            int64    `protobuf:"varint,2,opt,name=SettleDay,proto3" json:"SettleDay"`
	Ratio                float64  `protobuf:"fixed64,3,opt,name=Ratio,proto3" json:"Ratio"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiseSettleRequest) Reset()         { *m = RiseSettleRequest{} }
func (m *RiseSettleRequest) String() string { return proto.CompactTextString(m) }
func (*RiseSettleRequest) ProtoMessage()    {}
func (*RiseSettleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{3}
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
	PersonId     int64   `protobuf:"varint,1,opt,name=PersonId,proto3" json:"PersonId"`
	TransferWith int64   `protobuf:"varint,2,opt,name=TransferWith,proto3" json:"TransferWith"`
	Amount       float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount"`
	// 提现银行账号
	BankAccountNo        string   `protobuf:"bytes,4,opt,name=BankAccountNo,proto3" json:"BankAccountNo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiseTransferOutRequest) Reset()         { *m = RiseTransferOutRequest{} }
func (m *RiseTransferOutRequest) String() string { return proto.CompactTextString(m) }
func (*RiseTransferOutRequest) ProtoMessage()    {}
func (*RiseTransferOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{4}
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

func (m *RiseTransferOutRequest) GetBankAccountNo() string {
	if m != nil {
		return m.BankAccountNo
	}
	return ""
}

type CommitTransferRequest struct {
	PersonId             int64    `protobuf:"varint,1,opt,name=personId,proto3" json:"personId"`
	LogId                int64    `protobuf:"varint,2,opt,name=logId,proto3" json:"logId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitTransferRequest) Reset()         { *m = CommitTransferRequest{} }
func (m *CommitTransferRequest) String() string { return proto.CompactTextString(m) }
func (*CommitTransferRequest) ProtoMessage()    {}
func (*CommitTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_finance_service_843503d3c54e798a, []int{5}
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
	proto.RegisterFile("finance_service.proto", fileDescriptor_finance_service_843503d3c54e798a)
}

var fileDescriptor_finance_service_843503d3c54e798a = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0x9a, 0xf5, 0x27, 0x67, 0xa5, 0xa3, 0x66, 0x2b, 0x51, 0x35, 0x41, 0x65, 0x21, 0xa8,
	0xb8, 0xf0, 0xd0, 0x26, 0xae, 0xb8, 0x5a, 0x99, 0x40, 0xb9, 0x61, 0x93, 0x5b, 0x40, 0xe2, 0x06,
	0xb9, 0xa9, 0xdb, 0x45, 0x4b, 0xec, 0x92, 0x38, 0x48, 0x7b, 0x11, 0x5e, 0x80, 0x17, 0xe0, 0x11,
	0x91, 0xed, 0xfe, 0xb8, 0xeb, 0x2a, 0xed, 0xaa, 0xfd, 0x3e, 0xfb, 0x3b, 0x3f, 0xdf, 0x39, 0x0e,
	0x1c, 0x4f, 0x13, 0xc1, 0x44, 0xcc, 0x7f, 0x16, 0x3c, 0xff, 0x9d, 0xc4, 0x9c, 0xcc, 0x73, 0xa9,
	0x64, 0xb7, 0x39, 0x4b, 0xe5, 0x98, 0xa5, 0x16, 0xe1, 0x5b, 0x68, 0x8f, 0x72, 0x26, 0x8a, 0x29,
	0xcf, 0x23, 0x41, 0xf9, 0xaf, 0x92, 0x17, 0x0a, 0x75, 0xa1, 0x31, 0xe7, 0x79, 0x21, 0x45, 0x34,
	0x09, 0xbd, 0x9e, 0xd7, 0x47, 0x74, 0x85, 0x11, 0x86, 0xa6, 0x5a, 0x08, 0xbe, 0x27, 0xea, 0x26,
	0xac, 0xf4, 0xbc, 0x7e, 0x9b, 0x6e, 0x70, 0xa8, 0x03, 0x35, 0x96, 0xc9, 0x52, 0xa8, 0xd0, 0xef,
	0x79, 0x7d, 0x8f, 0x2e, 0x10, 0xee, 0x41, 0xe3, 0x7a, 0x19, 0xe7, 0x08, 0xaa, 0xdf, 0x58, 0x5a,
	0x72, 0x93, 0xc0, 0xa7, 0x16, 0xe0, 0xbf, 0x15, 0x08, 0x86, 0x34, 0x29, 0x78, 0x24, 0xa6, 0x52,
	0xd7, 0x71, 0xed, 0xd6, 0xe1, 0xd3, 0xb5, 0x3e, 0x84, 0xfa, 0x80, 0xa5, 0xba, 0x3f, 0x53, 0x82,
	0x47, 0x97, 0x10, 0xbd, 0x85, 0xa7, 0x43, 0xae, 0x54, 0xca, 0x33, 0x2e, 0xd4, 0x85, 0x5b, 0xc7,
	0x16, 0x8f, 0x10, 0xec, 0xeb, 0x6c, 0xe1, 0xbe, 0x39, 0x37, 0xff, 0xd1, 0x0b, 0x80, 0xb5, 0x25,
	0x61, 0xd5, 0x9c, 0x38, 0x0c, 0xea, 0xc1, 0xc1, 0x48, 0x2a, 0x96, 0x2e, 0x42, 0xd7, 0xcc, 0x05,
	0x97, 0x42, 0x27, 0x10, 0x18, 0x68, 0x42, 0xd7, 0xcd, 0xf9, 0x9a, 0xd0, 0x7a, 0x5b, 0xc7, 0xe4,
	0x92, 0x29, 0x1e, 0x36, 0x4c, 0x63, 0x2e, 0xa5, 0x2b, 0xf8, 0x3a, 0x9f, 0x30, 0xc5, 0x47, 0x49,
	0xc6, 0xc3, 0xc0, 0x5c, 0x70, 0x18, 0x1c, 0x43, 0x5b, 0x47, 0xb2, 0x12, 0x67, 0x68, 0x3b, 0xcd,
	0x3a, 0x81, 0xc0, 0x5e, 0xbe, 0x64, 0x77, 0xc6, 0x2e, 0x9f, 0xae, 0x09, 0x3d, 0x0a, 0xca, 0x54,
	0x22, 0x17, 0x2e, 0x59, 0x80, 0xff, 0x78, 0xd0, 0xd1, 0x59, 0x96, 0x9d, 0x5f, 0x95, 0xea, 0x31,
	0xa9, 0x30, 0x34, 0x47, 0xf7, 0xf7, 0xc3, 0xa7, 0x1b, 0xdc, 0xae, 0xfd, 0x40, 0xaf, 0xe0, 0xc9,
	0x80, 0x89, 0xdb, 0x8b, 0x38, 0xd6, 0xf0, 0x8b, 0x34, 0x63, 0x09, 0xe8, 0x26, 0x89, 0x23, 0x38,
	0xfe, 0x28, 0xb3, 0x2c, 0x51, 0xcb, 0x98, 0xbb, 0xd6, 0xd6, 0x77, 0xd6, 0xf6, 0x08, 0xaa, 0xa9,
	0x9c, 0x45, 0x93, 0x45, 0x3d, 0x16, 0x9c, 0xfd, 0xab, 0x40, 0xeb, 0x93, 0x7d, 0x25, 0x43, 0xfb,
	0x48, 0xd0, 0x6b, 0x38, 0xf8, 0xcc, 0xd5, 0x6a, 0x05, 0x03, 0xb2, 0xec, 0xac, 0x0b, 0x64, 0xb5,
	0x99, 0x78, 0x0f, 0x9d, 0x42, 0xcb, 0x75, 0x27, 0x12, 0x08, 0x91, 0xad, 0x97, 0xd4, 0xad, 0x13,
	0xca, 0x8b, 0x32, 0x55, 0x78, 0x0f, 0xbd, 0x87, 0xc3, 0x7b, 0x76, 0xa2, 0xe7, 0xe4, 0x61, 0x83,
	0x5d, 0xd9, 0x3b, 0x2b, 0xb3, 0xd3, 0x1a, 0xdc, 0xe9, 0x79, 0x21, 0xb2, 0x35, 0x7d, 0x57, 0x71,
	0x0e, 0xad, 0x4d, 0x7f, 0x50, 0x87, 0x3c, 0x68, 0x98, 0x2b, 0x7a, 0x03, 0x87, 0x57, 0x73, 0x2e,
	0x6c, 0x60, 0xeb, 0x84, 0xd3, 0xfa, 0xfa, 0xe2, 0xe0, 0x25, 0x3c, 0x8b, 0x65, 0x46, 0x66, 0x89,
	0xba, 0x29, 0xc7, 0x64, 0x26, 0xcf, 0x24, 0xc9, 0xe7, 0xf1, 0x8f, 0x06, 0x39, 0xfd, 0x60, 0xbe,
	0x28, 0xe3, 0x9a, 0xf9, 0x39, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xcc, 0xe5, 0x87, 0x7f,
	0x04, 0x00, 0x00,
}
