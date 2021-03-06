// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub.proto

package sonm

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

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ListReply struct {
	Info map[string]*ListReply_ListValue `protobuf:"bytes,1,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListReply) GetInfo() map[string]*ListReply_ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListReply_ListValue struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListReply_ListValue) Reset()                    { *m = ListReply_ListValue{} }
func (m *ListReply_ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListReply_ListValue) ProtoMessage()               {}
func (*ListReply_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *ListReply_ListValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type HubInfoRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubInfoRequest) Reset()                    { *m = HubInfoRequest{} }
func (m *HubInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*HubInfoRequest) ProtoMessage()               {}
func (*HubInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *HubInfoRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type TaskRequirements struct {
	// How much resources consumes a task. This is used both for scheduling and for cgroups configuration.
	Resources *TaskResourceRequirements `protobuf:"bytes,1,opt,name=resources" json:"resources,omitempty"`
	// Optional miner ID restrictions (currently IP:port), that are allowed to start a task.
	Miners []string `protobuf:"bytes,2,rep,name=miners" json:"miners,omitempty"`
}

func (m *TaskRequirements) Reset()                    { *m = TaskRequirements{} }
func (m *TaskRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskRequirements) ProtoMessage()               {}
func (*TaskRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TaskRequirements) GetResources() *TaskResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskRequirements) GetMiners() []string {
	if m != nil {
		return m.Miners
	}
	return nil
}

type HubStartTaskRequest struct {
	Requirements  *TaskRequirements `protobuf:"bytes,1,opt,name=requirements" json:"requirements,omitempty"`
	Registry      string            `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string            `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	PublicKeyData string            `protobuf:"bytes,5,opt,name=PublicKeyData,json=publicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool              `protobuf:"varint,6,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string `protobuf:"bytes,7,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HubStartTaskRequest) Reset()                    { *m = HubStartTaskRequest{} }
func (m *HubStartTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskRequest) ProtoMessage()               {}
func (*HubStartTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *HubStartTaskRequest) GetRequirements() *TaskRequirements {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *HubStartTaskRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *HubStartTaskRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *HubStartTaskRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *HubStartTaskRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *HubStartTaskRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *HubStartTaskRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type HubStartTaskReply struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Endpoint []string `protobuf:"bytes,2,rep,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HubStartTaskReply) Reset()                    { *m = HubStartTaskReply{} }
func (m *HubStartTaskReply) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskReply) ProtoMessage()               {}
func (*HubStartTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *HubStartTaskReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HubStartTaskReply) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type HubStatusMapRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubStatusMapRequest) Reset()                    { *m = HubStatusMapRequest{} }
func (m *HubStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusMapRequest) ProtoMessage()               {}
func (*HubStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HubStatusMapRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type HubStatusRequest struct {
}

func (m *HubStatusRequest) Reset()                    { *m = HubStatusRequest{} }
func (m *HubStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusRequest) ProtoMessage()               {}
func (*HubStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type HubStatusReply struct {
	MinerCount uint64 `protobuf:"varint,1,opt,name=minerCount" json:"minerCount,omitempty"`
	Uptime     uint64 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Platform   string `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
	EthAddr    string `protobuf:"bytes,5,opt,name=ethAddr" json:"ethAddr,omitempty"`
}

func (m *HubStatusReply) Reset()                    { *m = HubStatusReply{} }
func (m *HubStatusReply) String() string            { return proto.CompactTextString(m) }
func (*HubStatusReply) ProtoMessage()               {}
func (*HubStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *HubStatusReply) GetMinerCount() uint64 {
	if m != nil {
		return m.MinerCount
	}
	return 0
}

func (m *HubStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *HubStatusReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HubStatusReply) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *HubStatusReply) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "sonm.ListRequest")
	proto.RegisterType((*ListReply)(nil), "sonm.ListReply")
	proto.RegisterType((*ListReply_ListValue)(nil), "sonm.ListReply.ListValue")
	proto.RegisterType((*HubInfoRequest)(nil), "sonm.HubInfoRequest")
	proto.RegisterType((*TaskRequirements)(nil), "sonm.TaskRequirements")
	proto.RegisterType((*HubStartTaskRequest)(nil), "sonm.HubStartTaskRequest")
	proto.RegisterType((*HubStartTaskReply)(nil), "sonm.HubStartTaskReply")
	proto.RegisterType((*HubStatusMapRequest)(nil), "sonm.HubStatusMapRequest")
	proto.RegisterType((*HubStatusRequest)(nil), "sonm.HubStatusRequest")
	proto.RegisterType((*HubStatusReply)(nil), "sonm.HubStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hub service

type HubClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error)
	StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error)
}

type hubClient struct {
	cc *grpc.ClientConn
}

func NewHubClient(cc *grpc.ClientConn) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error) {
	out := new(HubStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error) {
	out := new(HubStartTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StartTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StopTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/TaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error) {
	out := new(StatusMapReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/MinerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hub_serviceDesc.Streams[0], c.cc, "/sonm.Hub/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type hubTaskLogsClient struct {
	grpc.ClientStream
}

func (x *hubTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hub service

type HubServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Status(context.Context, *HubStatusRequest) (*HubStatusReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Info(context.Context, *HubInfoRequest) (*InfoReply, error)
	StartTask(context.Context, *HubStartTaskRequest) (*HubStartTaskReply, error)
	StopTask(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TaskStatus(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	MinerStatus(context.Context, *HubStatusMapRequest) (*StatusMapReply, error)
	TaskLogs(*TaskLogsRequest, Hub_TaskLogsServer) error
}

func RegisterHubServer(s *grpc.Server, srv HubServer) {
	s.RegisterService(&_Hub_serviceDesc, srv)
}

func _Hub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Status(ctx, req.(*HubStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Info(ctx, req.(*HubInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StartTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StartTask(ctx, req.(*HubStartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StopTask(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).TaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/TaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).TaskStatus(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_MinerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).MinerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/MinerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).MinerStatus(ctx, req.(*HubStatusMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).TaskLogs(m, &hubTaskLogsServer{stream})
}

type Hub_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type hubTaskLogsServer struct {
	grpc.ServerStream
}

func (x *hubTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Hub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hub_Ping_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Hub_Status_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Hub_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Hub_Info_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _Hub_StartTask_Handler,
		},
		{
			MethodName: "StopTask",
			Handler:    _Hub_StopTask_Handler,
		},
		{
			MethodName: "TaskStatus",
			Handler:    _Hub_TaskStatus_Handler,
		},
		{
			MethodName: "MinerStatus",
			Handler:    _Hub_MinerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskLogs",
			Handler:       _Hub_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hub.proto",
}

func init() { proto.RegisterFile("hub.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xed, 0x4e, 0x1b, 0x39,
	0x14, 0x65, 0x92, 0x21, 0x64, 0x2e, 0xdf, 0xe6, 0x6b, 0x98, 0x1f, 0x28, 0x9a, 0x5d, 0xad, 0x90,
	0x96, 0x0d, 0x2b, 0x5a, 0x55, 0x08, 0x55, 0x6d, 0x11, 0x45, 0xa2, 0x2a, 0xa8, 0x68, 0xa8, 0xfa,
	0x7f, 0x42, 0x4c, 0x62, 0x91, 0xb1, 0xa7, 0x1e, 0x3b, 0x52, 0x9e, 0xa0, 0xaf, 0xd0, 0xf7, 0xe8,
	0x83, 0xf4, 0x95, 0x2a, 0x7f, 0x4d, 0x1c, 0x48, 0xfb, 0x6b, 0x7c, 0x8e, 0xef, 0xb9, 0xd7, 0xc7,
	0xf7, 0x7a, 0x20, 0x1a, 0xca, 0x5e, 0xb7, 0xe4, 0x4c, 0x30, 0x14, 0x56, 0x8c, 0x16, 0xc9, 0x3a,
	0xa1, 0xea, 0x4b, 0x49, 0x6e, 0xe8, 0x74, 0x15, 0x96, 0xaf, 0x49, 0x25, 0x32, 0xfc, 0x55, 0xe2,
	0x4a, 0xa4, 0x3f, 0x02, 0x88, 0x0c, 0x2e, 0x47, 0x13, 0xf4, 0x1f, 0x84, 0x84, 0x3e, 0xb0, 0x38,
	0xe8, 0x34, 0x0f, 0x97, 0x4f, 0xf6, 0xbb, 0x4a, 0xda, 0xad, 0xb7, 0xbb, 0x1f, 0xe8, 0x03, 0xbb,
	0xa4, 0x82, 0x4f, 0x32, 0x1d, 0x96, 0xfc, 0x65, 0xb4, 0x5f, 0xf2, 0x91, 0xc4, 0x68, 0x17, 0x5a,
	0x63, 0xb5, 0xa8, 0xb4, 0x3a, 0xca, 0x2c, 0x4a, 0x32, 0x88, 0x6a, 0x1d, 0xda, 0x80, 0xe6, 0x23,
	0x9e, 0xc4, 0x41, 0x27, 0x38, 0x8c, 0x32, 0xb5, 0x44, 0xc7, 0xb0, 0xa8, 0x03, 0xe3, 0x46, 0x27,
	0x98, 0x57, 0xb3, 0x2e, 0x90, 0x99, 0xb8, 0xb3, 0xc6, 0x69, 0x90, 0xfe, 0x03, 0x6b, 0x57, 0xb2,
	0xa7, 0xd2, 0x5a, 0x1f, 0x68, 0x1b, 0x16, 0x0b, 0x42, 0x31, 0xb7, 0xa9, 0x0d, 0x48, 0x87, 0xb0,
	0xf1, 0x39, 0xaf, 0x1e, 0x55, 0x10, 0xe1, 0xb8, 0xc0, 0x54, 0x54, 0xe8, 0x35, 0x44, 0x1c, 0x57,
	0x4c, 0xf2, 0x7b, 0x7d, 0x54, 0x55, 0xf4, 0xc0, 0x14, 0x35, 0xa1, 0x66, 0xcb, 0x97, 0x64, 0x53,
	0x81, 0x72, 0xa9, 0x53, 0x57, 0x71, 0xc3, 0xb8, 0x34, 0x28, 0xfd, 0xd9, 0x80, 0xad, 0x2b, 0xd9,
	0xbb, 0x13, 0x39, 0x17, 0xae, 0xa4, 0x3a, 0xd7, 0x19, 0xac, 0x70, 0x2f, 0x95, 0x2d, 0xb8, 0xeb,
	0x17, 0xf4, 0x0a, 0xcd, 0xc4, 0xa2, 0x04, 0xda, 0x1c, 0x0f, 0x48, 0x25, 0xf8, 0x44, 0xdf, 0x4e,
	0x94, 0xd5, 0x58, 0xf9, 0x25, 0x45, 0x3e, 0xc0, 0x71, 0xd3, 0xf8, 0xd5, 0x00, 0x21, 0x08, 0x73,
	0x29, 0x86, 0x71, 0xa8, 0x49, 0xbd, 0x46, 0x7f, 0xc3, 0xea, 0xad, 0xec, 0x8d, 0xc8, 0xfd, 0x47,
	0x3c, 0x79, 0x9f, 0x8b, 0x3c, 0x5e, 0xd4, 0x9b, 0xab, 0xa5, 0x4f, 0xa2, 0x14, 0x56, 0xee, 0x59,
	0x51, 0x10, 0xf1, 0x89, 0xde, 0x09, 0x56, 0xc6, 0xad, 0x4e, 0x70, 0xd8, 0xce, 0x66, 0x38, 0xf4,
	0x12, 0x9a, 0x98, 0x8e, 0xe3, 0x25, 0x3d, 0x1c, 0xa9, 0xb1, 0x30, 0xc7, 0x73, 0xf7, 0x92, 0x8e,
	0xcd, 0x94, 0xa8, 0xf0, 0xe4, 0x15, 0xb4, 0x1d, 0x31, 0xa7, 0xfd, 0xdb, 0x7e, 0xfb, 0x23, 0xbf,
	0xc7, 0x6f, 0x61, 0x73, 0x36, 0xb9, 0x1a, 0xd0, 0x35, 0x68, 0x90, 0xbe, 0xd5, 0x37, 0x48, 0x5f,
	0x5d, 0x11, 0xa6, 0xfd, 0x92, 0x11, 0x2a, 0x6c, 0x43, 0x6a, 0x9c, 0xfe, 0xeb, 0x3a, 0x22, 0x64,
	0x75, 0x93, 0x97, 0x7f, 0x9e, 0x14, 0x04, 0x1b, 0x75, 0xb0, 0x7b, 0x1b, 0xdf, 0x03, 0x3d, 0x66,
	0x8e, 0x54, 0xf5, 0x0f, 0x00, 0x74, 0xfc, 0x05, 0x93, 0x54, 0xe8, 0x0c, 0x61, 0xe6, 0x31, 0x6a,
	0x3c, 0x64, 0x29, 0x48, 0x61, 0xfc, 0x84, 0x99, 0x45, 0x28, 0x86, 0xa5, 0x31, 0xe6, 0x15, 0x61,
	0xd4, 0x36, 0xcc, 0x41, 0xe5, 0xa0, 0x1c, 0xe5, 0xe2, 0x81, 0xf1, 0xc2, 0xb6, 0xad, 0xc6, 0x4a,
	0x85, 0xc5, 0xf0, 0xbc, 0xdf, 0xe7, 0xb6, 0x69, 0x0e, 0x9e, 0x7c, 0x0b, 0xa1, 0x79, 0x25, 0x7b,
	0xe8, 0x08, 0xc2, 0x5b, 0x42, 0x07, 0x68, 0xd3, 0x74, 0x43, 0xad, 0xed, 0xe9, 0x93, 0x75, 0x9f,
	0x2a, 0x47, 0x93, 0x74, 0x01, 0x9d, 0x42, 0xcb, 0x98, 0x41, 0xbb, 0x7e, 0xf7, 0xa6, 0x96, 0x93,
	0xed, 0x67, 0xbc, 0x51, 0x1e, 0x41, 0xa8, 0x1e, 0xa2, 0xab, 0xe3, 0xfd, 0x41, 0x5c, 0x9d, 0xfa,
	0xc5, 0xa6, 0x0b, 0xe8, 0x18, 0x42, 0xf5, 0x36, 0xd1, 0x34, 0x9b, 0xf7, 0x54, 0x9d, 0xc0, 0x50,
	0x46, 0x70, 0x0e, 0x51, 0xdd, 0x68, 0xb4, 0xff, 0xdb, 0xc9, 0x4a, 0xf6, 0xe6, 0x6d, 0x39, 0x6f,
	0x6d, 0x35, 0xa4, 0x3a, 0xc3, 0x8e, 0x09, 0x73, 0xd8, 0xa9, 0xb7, 0x9e, 0xd2, 0x46, 0xf9, 0x06,
	0x40, 0x41, 0x7b, 0x33, 0x7b, 0xd3, 0xa7, 0x39, 0x7b, 0x35, 0x3b, 0xcf, 0x37, 0x8c, 0xfe, 0x1d,
	0x2c, 0xdf, 0xa8, 0x09, 0xb0, 0x09, 0xf6, 0x9f, 0x5c, 0xe1, 0x74, 0xf4, 0xdc, 0xed, 0x7a, 0xbc,
	0xc9, 0x70, 0x06, 0x6d, 0x95, 0xf6, 0x9a, 0x0d, 0x2a, 0xe4, 0x95, 0x51, 0xf8, 0xc9, 0xd9, 0x1d,
	0x7d, 0x31, 0x94, 0xf4, 0x31, 0x5d, 0xf8, 0x3f, 0xe8, 0xb5, 0xf4, 0x6f, 0xfd, 0xc5, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0x93, 0xa6, 0x3a, 0xfa, 0x05, 0x00, 0x00,
}
