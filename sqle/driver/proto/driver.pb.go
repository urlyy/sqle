// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	driver.proto

It has these top-level messages:
	InitRequest
	Empty
	ExecRequest
	ExecResponse
	TxRequest
	TxResponse
	DatabasesResponse
	ParseRequest
	Node
	ParseResponse
	Rule
	AuditRequest
	AuditResult
	AuditResponse
	GenRollbackSQLRequest
	GenRollbackSQLResponse
	MetasResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type InitRequest struct {
	InstanceHost string `protobuf:"bytes,1,opt,name=instanceHost" json:"instanceHost,omitempty"`
	InstancePort string `protobuf:"bytes,2,opt,name=instancePort" json:"instancePort,omitempty"`
	InstanceUser string `protobuf:"bytes,3,opt,name=instanceUser" json:"instanceUser,omitempty"`
	InstancePass string `protobuf:"bytes,4,opt,name=instancePass" json:"instancePass,omitempty"`
	DatabaseOpen string `protobuf:"bytes,5,opt,name=databaseOpen" json:"databaseOpen,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto1.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InitRequest) GetInstanceHost() string {
	if m != nil {
		return m.InstanceHost
	}
	return ""
}

func (m *InitRequest) GetInstancePort() string {
	if m != nil {
		return m.InstancePort
	}
	return ""
}

func (m *InitRequest) GetInstanceUser() string {
	if m != nil {
		return m.InstanceUser
	}
	return ""
}

func (m *InitRequest) GetInstancePass() string {
	if m != nil {
		return m.InstancePass
	}
	return ""
}

func (m *InitRequest) GetDatabaseOpen() string {
	if m != nil {
		return m.DatabaseOpen
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ExecRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExecRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type ExecResponse struct {
	LastInsertId      int64  `protobuf:"varint,1,opt,name=lastInsertId" json:"lastInsertId,omitempty"`
	LastInsertIdError string `protobuf:"bytes,2,opt,name=lastInsertIdError" json:"lastInsertIdError,omitempty"`
	RowsAffected      int64  `protobuf:"varint,3,opt,name=rowsAffected" json:"rowsAffected,omitempty"`
	RowsAffectedError string `protobuf:"bytes,4,opt,name=rowsAffectedError" json:"rowsAffectedError,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto1.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExecResponse) GetLastInsertId() int64 {
	if m != nil {
		return m.LastInsertId
	}
	return 0
}

func (m *ExecResponse) GetLastInsertIdError() string {
	if m != nil {
		return m.LastInsertIdError
	}
	return ""
}

func (m *ExecResponse) GetRowsAffected() int64 {
	if m != nil {
		return m.RowsAffected
	}
	return 0
}

func (m *ExecResponse) GetRowsAffectedError() string {
	if m != nil {
		return m.RowsAffectedError
	}
	return ""
}

type TxRequest struct {
	Queries []string `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *TxRequest) Reset()                    { *m = TxRequest{} }
func (m *TxRequest) String() string            { return proto1.CompactTextString(m) }
func (*TxRequest) ProtoMessage()               {}
func (*TxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TxRequest) GetQueries() []string {
	if m != nil {
		return m.Queries
	}
	return nil
}

type TxResponse struct {
	Resluts []*ExecResponse `protobuf:"bytes,1,rep,name=resluts" json:"resluts,omitempty"`
}

func (m *TxResponse) Reset()                    { *m = TxResponse{} }
func (m *TxResponse) String() string            { return proto1.CompactTextString(m) }
func (*TxResponse) ProtoMessage()               {}
func (*TxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TxResponse) GetResluts() []*ExecResponse {
	if m != nil {
		return m.Resluts
	}
	return nil
}

type DatabasesResponse struct {
	Databases []string `protobuf:"bytes,1,rep,name=databases" json:"databases,omitempty"`
}

func (m *DatabasesResponse) Reset()                    { *m = DatabasesResponse{} }
func (m *DatabasesResponse) String() string            { return proto1.CompactTextString(m) }
func (*DatabasesResponse) ProtoMessage()               {}
func (*DatabasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DatabasesResponse) GetDatabases() []string {
	if m != nil {
		return m.Databases
	}
	return nil
}

type ParseRequest struct {
	SqlText string `protobuf:"bytes,1,opt,name=sqlText" json:"sqlText,omitempty"`
}

func (m *ParseRequest) Reset()                    { *m = ParseRequest{} }
func (m *ParseRequest) String() string            { return proto1.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()               {}
func (*ParseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ParseRequest) GetSqlText() string {
	if m != nil {
		return m.SqlText
	}
	return ""
}

type Node struct {
	Text        string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Fingerprint string `protobuf:"bytes,3,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto1.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Node) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Node) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

type ParseResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ParseResponse) Reset()                    { *m = ParseResponse{} }
func (m *ParseResponse) String() string            { return proto1.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()               {}
func (*ParseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ParseResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Rule struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Desc      string `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Level     string `protobuf:"bytes,4,opt,name=level" json:"level,omitempty"`
	Typ       string `protobuf:"bytes,5,opt,name=typ" json:"typ,omitempty"`
	IsDefault bool   `protobuf:"varint,6,opt,name=isDefault" json:"isDefault,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto1.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Rule) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Rule) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *Rule) GetTyp() string {
	if m != nil {
		return m.Typ
	}
	return ""
}

func (m *Rule) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

type AuditRequest struct {
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
	Sql   string  `protobuf:"bytes,2,opt,name=sql" json:"sql,omitempty"`
}

func (m *AuditRequest) Reset()                    { *m = AuditRequest{} }
func (m *AuditRequest) String() string            { return proto1.CompactTextString(m) }
func (*AuditRequest) ProtoMessage()               {}
func (*AuditRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AuditRequest) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *AuditRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

type AuditResult struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Level   string `protobuf:"bytes,2,opt,name=level" json:"level,omitempty"`
}

func (m *AuditResult) Reset()                    { *m = AuditResult{} }
func (m *AuditResult) String() string            { return proto1.CompactTextString(m) }
func (*AuditResult) ProtoMessage()               {}
func (*AuditResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AuditResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AuditResult) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type AuditResponse struct {
	Results []*AuditResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *AuditResponse) Reset()                    { *m = AuditResponse{} }
func (m *AuditResponse) String() string            { return proto1.CompactTextString(m) }
func (*AuditResponse) ProtoMessage()               {}
func (*AuditResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AuditResponse) GetResults() []*AuditResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type GenRollbackSQLRequest struct {
	Sql string `protobuf:"bytes,1,opt,name=sql" json:"sql,omitempty"`
}

func (m *GenRollbackSQLRequest) Reset()                    { *m = GenRollbackSQLRequest{} }
func (m *GenRollbackSQLRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenRollbackSQLRequest) ProtoMessage()               {}
func (*GenRollbackSQLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GenRollbackSQLRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

type GenRollbackSQLResponse struct {
	Sql    string `protobuf:"bytes,1,opt,name=sql" json:"sql,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *GenRollbackSQLResponse) Reset()                    { *m = GenRollbackSQLResponse{} }
func (m *GenRollbackSQLResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenRollbackSQLResponse) ProtoMessage()               {}
func (*GenRollbackSQLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GenRollbackSQLResponse) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *GenRollbackSQLResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type MetasResponse struct {
	Name  string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Rules []*Rule `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *MetasResponse) Reset()                    { *m = MetasResponse{} }
func (m *MetasResponse) String() string            { return proto1.CompactTextString(m) }
func (*MetasResponse) ProtoMessage()               {}
func (*MetasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *MetasResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetasResponse) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto1.RegisterType((*InitRequest)(nil), "proto.InitRequest")
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
	proto1.RegisterType((*ExecRequest)(nil), "proto.ExecRequest")
	proto1.RegisterType((*ExecResponse)(nil), "proto.ExecResponse")
	proto1.RegisterType((*TxRequest)(nil), "proto.TxRequest")
	proto1.RegisterType((*TxResponse)(nil), "proto.TxResponse")
	proto1.RegisterType((*DatabasesResponse)(nil), "proto.DatabasesResponse")
	proto1.RegisterType((*ParseRequest)(nil), "proto.ParseRequest")
	proto1.RegisterType((*Node)(nil), "proto.Node")
	proto1.RegisterType((*ParseResponse)(nil), "proto.ParseResponse")
	proto1.RegisterType((*Rule)(nil), "proto.Rule")
	proto1.RegisterType((*AuditRequest)(nil), "proto.AuditRequest")
	proto1.RegisterType((*AuditResult)(nil), "proto.AuditResult")
	proto1.RegisterType((*AuditResponse)(nil), "proto.AuditResponse")
	proto1.RegisterType((*GenRollbackSQLRequest)(nil), "proto.GenRollbackSQLRequest")
	proto1.RegisterType((*GenRollbackSQLResponse)(nil), "proto.GenRollbackSQLResponse")
	proto1.RegisterType((*MetasResponse)(nil), "proto.MetasResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Driver service

type DriverClient interface {
	// Metas returns some base info from plugin server.
	Metas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetasResponse, error)
	// Init will should be called at first before calling following methods.
	// It will pass some necessary info to plugin server. In the begginning,
	// we consider that put this info to the executable binary environment.
	// We put all communication on gRPC for unification in the end.
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	Tx(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error)
	Databases(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DatabasesResponse, error)
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error)
	GenRollbackSQL(ctx context.Context, in *GenRollbackSQLRequest, opts ...grpc.CallOption) (*GenRollbackSQLResponse, error)
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) Metas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetasResponse, error) {
	out := new(MetasResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Metas", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.Driver/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.Driver/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.Driver/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Tx(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error) {
	out := new(TxResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Tx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Databases(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DatabasesResponse, error) {
	out := new(DatabasesResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Databases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Parse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error) {
	out := new(AuditResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/Audit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GenRollbackSQL(ctx context.Context, in *GenRollbackSQLRequest, opts ...grpc.CallOption) (*GenRollbackSQLResponse, error) {
	out := new(GenRollbackSQLResponse)
	err := grpc.Invoke(ctx, "/proto.Driver/GenRollbackSQL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Driver service

type DriverServer interface {
	// Metas returns some base info from plugin server.
	Metas(context.Context, *Empty) (*MetasResponse, error)
	// Init will should be called at first before calling following methods.
	// It will pass some necessary info to plugin server. In the begginning,
	// we consider that put this info to the executable binary environment.
	// We put all communication on gRPC for unification in the end.
	Init(context.Context, *InitRequest) (*Empty, error)
	Close(context.Context, *Empty) (*Empty, error)
	Ping(context.Context, *Empty) (*Empty, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	Tx(context.Context, *TxRequest) (*TxResponse, error)
	Databases(context.Context, *Empty) (*DatabasesResponse, error)
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	Audit(context.Context, *AuditRequest) (*AuditResponse, error)
	GenRollbackSQL(context.Context, *GenRollbackSQLRequest) (*GenRollbackSQLResponse, error)
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_Metas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Metas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Metas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Metas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Close(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Tx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Tx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Tx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Tx(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Databases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Databases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Databases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Databases(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Audit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Audit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/Audit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Audit(ctx, req.(*AuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GenRollbackSQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenRollbackSQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GenRollbackSQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Driver/GenRollbackSQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GenRollbackSQL(ctx, req.(*GenRollbackSQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Metas",
			Handler:    _Driver_Metas_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Driver_Init_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Driver_Close_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Driver_Ping_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Driver_Exec_Handler,
		},
		{
			MethodName: "Tx",
			Handler:    _Driver_Tx_Handler,
		},
		{
			MethodName: "Databases",
			Handler:    _Driver_Databases_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _Driver_Parse_Handler,
		},
		{
			MethodName: "Audit",
			Handler:    _Driver_Audit_Handler,
		},
		{
			MethodName: "GenRollbackSQL",
			Handler:    _Driver_GenRollbackSQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}

func init() { proto1.RegisterFile("driver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0x96, 0x93, 0x38, 0x90, 0x49, 0x38, 0x82, 0x3d, 0x1c, 0x64, 0x45, 0x1c, 0x29, 0x2c, 0x3a,
	0x52, 0xd0, 0xe1, 0x80, 0x4e, 0xb8, 0x44, 0x5c, 0xf0, 0xd7, 0x16, 0xa9, 0xb4, 0xa9, 0x4b, 0x6f,
	0x7a, 0x67, 0xe2, 0x09, 0xb2, 0xba, 0xd8, 0xc9, 0xee, 0x9a, 0x26, 0xaf, 0xd0, 0xc7, 0xe9, 0x1b,
	0xf4, 0x8d, 0xfa, 0x08, 0xd5, 0xae, 0x77, 0x9d, 0x75, 0x02, 0xbd, 0xf2, 0xce, 0xb7, 0x9f, 0x67,
	0xbe, 0x99, 0x9d, 0x19, 0xe8, 0xc4, 0x3c, 0x79, 0x42, 0x7e, 0x34, 0xe1, 0x99, 0xcc, 0x88, 0xaf,
	0x3f, 0xf4, 0x87, 0x07, 0xed, 0x9b, 0x34, 0x91, 0x21, 0x4e, 0x73, 0x14, 0x92, 0x50, 0xe8, 0x24,
	0xa9, 0x90, 0x51, 0x3a, 0xc2, 0x37, 0x99, 0x90, 0x81, 0xd7, 0xf3, 0xfa, 0xad, 0xb0, 0x82, 0xb9,
	0x9c, 0x61, 0xc6, 0x65, 0x50, 0xab, 0x72, 0x14, 0xe6, 0x72, 0x3e, 0x09, 0xe4, 0x41, 0xbd, 0xca,
	0x51, 0x58, 0xc5, 0x4f, 0x24, 0x44, 0xd0, 0x58, 0xf2, 0x13, 0x09, 0xa1, 0x38, 0x71, 0x24, 0xa3,
	0xfb, 0x48, 0xe0, 0xfb, 0x09, 0xa6, 0x81, 0x5f, 0x70, 0x5c, 0x8c, 0xae, 0x81, 0x7f, 0xfd, 0x38,
	0x91, 0x73, 0xba, 0x0f, 0xed, 0xeb, 0x19, 0x8e, 0x6c, 0x2e, 0xdb, 0xe0, 0x4f, 0x73, 0xe4, 0x73,
	0x93, 0x44, 0x61, 0xd0, 0xef, 0x1e, 0x74, 0x0a, 0x96, 0x98, 0x64, 0xa9, 0x40, 0x15, 0x82, 0x45,
	0x42, 0xde, 0xa4, 0x02, 0xb9, 0xbc, 0x89, 0x35, 0xbb, 0x1e, 0x56, 0x30, 0x72, 0x08, 0x5b, 0xae,
	0x7d, 0xcd, 0x79, 0xc6, 0x4d, 0xde, 0xab, 0x17, 0xca, 0x23, 0xcf, 0xbe, 0x8a, 0xf3, 0xf1, 0x18,
	0x47, 0x12, 0x63, 0x9d, 0x7c, 0x3d, 0xac, 0x60, 0xca, 0xa3, 0x6b, 0x17, 0x1e, 0x8b, 0x0a, 0xac,
	0x5e, 0xd0, 0x7f, 0xa0, 0x75, 0x37, 0xb3, 0x79, 0x05, 0xb0, 0xa6, 0x52, 0x49, 0x50, 0x04, 0x5e,
	0xaf, 0xde, 0x6f, 0x85, 0xd6, 0xa4, 0xa7, 0x00, 0x8a, 0x66, 0x12, 0xfb, 0x0f, 0xd6, 0x38, 0x0a,
	0x96, 0xcb, 0x82, 0xd7, 0x1e, 0xfc, 0x59, 0xbc, 0xfd, 0x91, 0x9b, 0x7e, 0x68, 0x39, 0xf4, 0x7f,
	0xd8, 0xba, 0x32, 0x65, 0x15, 0xa5, 0x8f, 0x5d, 0x68, 0xd9, 0x5a, 0xdb, 0x68, 0x0b, 0x80, 0xf6,
	0xa1, 0x33, 0x8c, 0xb8, 0x40, 0x47, 0x99, 0x98, 0xb2, 0x3b, 0x9c, 0xd9, 0xc6, 0xb1, 0x26, 0x1d,
	0x42, 0xe3, 0x5d, 0x16, 0x23, 0x21, 0xd0, 0x90, 0x8b, 0x6b, 0x7d, 0xd6, 0xd8, 0x7c, 0x82, 0xa6,
	0x9e, 0xfa, 0x4c, 0x7a, 0xd0, 0x1e, 0x27, 0xe9, 0x03, 0xf2, 0x09, 0x4f, 0x52, 0x69, 0xda, 0xc7,
	0x85, 0xe8, 0x00, 0x36, 0x4c, 0x6c, 0x23, 0x75, 0x0f, 0xfc, 0x34, 0x8b, 0xd1, 0x26, 0xdb, 0x36,
	0xc9, 0xaa, 0xb0, 0x61, 0x71, 0x43, 0xbf, 0x79, 0xd0, 0x08, 0x73, 0xa6, 0x65, 0xa4, 0xd1, 0x23,
	0x5a, 0x19, 0xea, 0xac, 0xb0, 0x18, 0xc5, 0xc8, 0xca, 0x50, 0x67, 0xd5, 0x42, 0x4f, 0x11, 0xcb,
	0xd1, 0x08, 0x28, 0x0c, 0x85, 0x32, 0x7c, 0x42, 0x66, 0xde, 0xab, 0x30, 0xc8, 0x26, 0xd4, 0xe5,
	0x7c, 0x62, 0x3a, 0x54, 0x1d, 0x55, 0xf1, 0x12, 0x71, 0x85, 0xe3, 0x28, 0x67, 0x32, 0x68, 0xf6,
	0xbc, 0xfe, 0x7a, 0xb8, 0x00, 0xe8, 0x25, 0x74, 0xce, 0xf3, 0x78, 0x31, 0x7a, 0x7b, 0xe0, 0xf3,
	0x9c, 0xad, 0xe8, 0x57, 0x7a, 0xc3, 0xe2, 0x46, 0x85, 0x10, 0x53, 0x66, 0x14, 0xaa, 0x23, 0x3d,
	0x83, 0xb6, 0x71, 0x22, 0x72, 0xa6, 0x1f, 0xe0, 0x11, 0x85, 0x88, 0x1e, 0x6c, 0x6a, 0xd6, 0x5c,
	0x68, 0xae, 0x39, 0x9a, 0xe9, 0x19, 0x6c, 0xd8, 0xdf, 0x8b, 0x22, 0x1e, 0xea, 0x9e, 0xc9, 0x59,
	0xd9, 0x33, 0xc4, 0xc8, 0x70, 0xa2, 0x84, 0x96, 0x42, 0x0f, 0xe0, 0xaf, 0xd7, 0x98, 0x86, 0x19,
	0x63, 0xf7, 0xd1, 0xe8, 0xcb, 0xc7, 0x0f, 0x6f, 0x6d, 0x2e, 0x46, 0xa8, 0xb7, 0x10, 0x7a, 0x01,
	0x3b, 0xcb, 0x54, 0x13, 0x72, 0x85, 0x4b, 0x76, 0xa0, 0xc9, 0x31, 0x12, 0x59, 0x6a, 0xc4, 0x1a,
	0x8b, 0xbe, 0x82, 0x8d, 0x5b, 0x94, 0xd1, 0xa2, 0x3b, 0x9f, 0x7b, 0xc6, 0xb2, 0x8c, 0xb5, 0x97,
	0xca, 0x38, 0xf8, 0x59, 0x87, 0xe6, 0x95, 0x5e, 0x86, 0xe4, 0x5f, 0xf0, 0xb5, 0x4b, 0xd2, 0xb1,
	0xb3, 0xa1, 0x36, 0x49, 0x77, 0xdb, 0x58, 0xd5, 0x70, 0x7d, 0x68, 0xa8, 0x5d, 0x49, 0x6c, 0x4d,
	0x9c, 0xc5, 0xd9, 0xad, 0xfc, 0x4f, 0xf6, 0xc1, 0xbf, 0x64, 0x99, 0xc0, 0x25, 0xb7, 0x55, 0x12,
	0x85, 0xc6, 0x30, 0x49, 0x1f, 0x7e, 0xcb, 0x39, 0x86, 0x86, 0x9a, 0xd6, 0x32, 0xa4, 0xb3, 0xdf,
	0xba, 0xcf, 0x8d, 0x33, 0x39, 0x80, 0xda, 0xdd, 0x8c, 0x6c, 0x9a, 0xab, 0x72, 0x69, 0x74, 0xb7,
	0x1c, 0xc4, 0x50, 0x4f, 0xa0, 0x55, 0x0e, 0xfc, 0x92, 0x88, 0xc0, 0x58, 0xab, 0x0b, 0x61, 0x00,
	0xbe, 0x1e, 0x3b, 0x62, 0xa3, 0xbb, 0x0b, 0xa0, 0xac, 0x5b, 0x75, 0x32, 0x07, 0xe0, 0xeb, 0xf6,
	0x29, 0xff, 0x71, 0xfb, 0xbe, 0xfc, 0xa7, 0xda, 0x88, 0xb7, 0xf0, 0x47, 0xb5, 0x5f, 0xc8, 0xae,
	0xe1, 0x3d, 0xdb, 0x71, 0xdd, 0xbf, 0x5f, 0xb8, 0x2d, 0xdc, 0x5d, 0xc0, 0xe7, 0xf5, 0xa3, 0xe3,
	0x53, 0x4d, 0xb9, 0x6f, 0xea, 0xcf, 0xc9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xae, 0x86,
	0xf6, 0x11, 0x07, 0x00, 0x00,
}
