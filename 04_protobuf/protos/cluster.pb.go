// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: examples/demo/protos/rpc.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RPCMsg message to be sent using rpc
type RPCMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *RPCMsg) Reset() {
	*x = RPCMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMsg) ProtoMessage() {}

func (x *RPCMsg) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMsg.ProtoReflect.Descriptor instead.
func (*RPCMsg) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *RPCMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// RPCRes is the rpc response
type RPCRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *RPCRes) Reset() {
	*x = RPCRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRes) ProtoMessage() {}

func (x *RPCRes) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRes.ProtoReflect.Descriptor instead.
func (*RPCRes) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *RPCRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// UserMessage represents a message that user sent
type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *UserMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Stats exports the room status
type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutboundBytes int64 `protobuf:"varint,1,opt,name=outbound_bytes,json=outboundBytes,proto3" json:"outbound_bytes,omitempty"`
	InboundBytes  int64 `protobuf:"varint,2,opt,name=inbound_bytes,json=inboundBytes,proto3" json:"inbound_bytes,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Stats) GetOutboundBytes() int64 {
	if x != nil {
		return x.OutboundBytes
	}
	return 0
}

func (x *Stats) GetInboundBytes() int64 {
	if x != nil {
		return x.InboundBytes
	}
	return 0
}

// SendRPCMsg represents a rpc message
type SendRPCMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Route    string `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Msg      string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendRPCMsg) Reset() {
	*x = SendRPCMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRPCMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRPCMsg) ProtoMessage() {}

func (x *SendRPCMsg) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRPCMsg.ProtoReflect.Descriptor instead.
func (*SendRPCMsg) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *SendRPCMsg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *SendRPCMsg) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *SendRPCMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// NewUser message will be received when new user join room
type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *NewUser) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// AllMembers contains all members uid
type AllMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []string `protobuf:"bytes,1,rep,name=Members,proto3" json:"Members,omitempty"`
}

func (x *AllMembers) Reset() {
	*x = AllMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllMembers) ProtoMessage() {}

func (x *AllMembers) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllMembers.ProtoReflect.Descriptor instead.
func (*AllMembers) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *AllMembers) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

// JoinResponse represents the result of joining room
type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{7}
}

func (x *JoinResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *JoinResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// Response struct
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_demo_protos_cluster_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_examples_demo_protos_cluster_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_examples_demo_protos_cluster_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_examples_demo_protos_cluster_proto protoreflect.FileDescriptor

var file_examples_demo_protos_cluster_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x52, 0x50, 0x43, 0x4d, 0x73, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67,
	0x22, 0x1a, 0x0a, 0x06, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x3b, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x51,
	0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x50, 0x43, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x23, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x3a,
	0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x16, 0x5a, 0x14,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_demo_protos_cluster_proto_rawDescOnce sync.Once
	file_examples_demo_protos_cluster_proto_rawDescData = file_examples_demo_protos_cluster_proto_rawDesc
)

func file_examples_demo_protos_cluster_proto_rawDescGZIP() []byte {
	file_examples_demo_protos_cluster_proto_rawDescOnce.Do(func() {
		file_examples_demo_protos_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_demo_protos_cluster_proto_rawDescData)
	})
	return file_examples_demo_protos_cluster_proto_rawDescData
}

var file_examples_demo_protos_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_examples_demo_protos_cluster_proto_goTypes = []interface{}{
	(*RPCMsg)(nil),       // 0: cluster_protos.RPCMsg
	(*RPCRes)(nil),       // 1: cluster_protos.RPCRes
	(*UserMessage)(nil),  // 2: cluster_protos.UserMessage
	(*Stats)(nil),        // 3: cluster_protos.Stats
	(*SendRPCMsg)(nil),   // 4: cluster_protos.SendRPCMsg
	(*NewUser)(nil),      // 5: cluster_protos.NewUser
	(*AllMembers)(nil),   // 6: cluster_protos.AllMembers
	(*JoinResponse)(nil), // 7: cluster_protos.JoinResponse
	(*Response)(nil),     // 8: cluster_protos.Response
}
var file_examples_demo_protos_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_demo_protos_cluster_proto_init() }
func file_examples_demo_protos_cluster_proto_init() {
	if File_examples_demo_protos_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_demo_protos_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRPCMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllMembers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_demo_protos_cluster_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_demo_protos_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_demo_protos_cluster_proto_goTypes,
		DependencyIndexes: file_examples_demo_protos_cluster_proto_depIdxs,
		MessageInfos:      file_examples_demo_protos_cluster_proto_msgTypes,
	}.Build()
	File_examples_demo_protos_cluster_proto = out.File
	file_examples_demo_protos_cluster_proto_rawDesc = nil
	file_examples_demo_protos_cluster_proto_goTypes = nil
	file_examples_demo_protos_cluster_proto_depIdxs = nil
}
