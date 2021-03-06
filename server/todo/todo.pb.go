// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TodoObject struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoObject) Reset()         { *m = TodoObject{} }
func (m *TodoObject) String() string { return proto.CompactTextString(m) }
func (*TodoObject) ProtoMessage()    {}
func (*TodoObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *TodoObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoObject.Unmarshal(m, b)
}
func (m *TodoObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoObject.Marshal(b, m, deterministic)
}
func (m *TodoObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoObject.Merge(m, src)
}
func (m *TodoObject) XXX_Size() int {
	return xxx_messageInfo_TodoObject.Size(m)
}
func (m *TodoObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoObject.DiscardUnknown(m)
}

var xxx_messageInfo_TodoObject proto.InternalMessageInfo

func (m *TodoObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TodoObject) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

// list
type ListTodoParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoParams) Reset()         { *m = ListTodoParams{} }
func (m *ListTodoParams) String() string { return proto.CompactTextString(m) }
func (*ListTodoParams) ProtoMessage()    {}
func (*ListTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *ListTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoParams.Unmarshal(m, b)
}
func (m *ListTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoParams.Marshal(b, m, deterministic)
}
func (m *ListTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoParams.Merge(m, src)
}
func (m *ListTodoParams) XXX_Size() int {
	return xxx_messageInfo_ListTodoParams.Size(m)
}
func (m *ListTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoParams proto.InternalMessageInfo

type ListTodoResponse struct {
	Todos                []*TodoObject `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (m *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(m, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetTodos() []*TodoObject {
	if m != nil {
		return m.Todos
	}
	return nil
}

// get
type GetTodoParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoParams) Reset()         { *m = GetTodoParams{} }
func (m *GetTodoParams) String() string { return proto.CompactTextString(m) }
func (*GetTodoParams) ProtoMessage()    {}
func (*GetTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *GetTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoParams.Unmarshal(m, b)
}
func (m *GetTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoParams.Marshal(b, m, deterministic)
}
func (m *GetTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoParams.Merge(m, src)
}
func (m *GetTodoParams) XXX_Size() int {
	return xxx_messageInfo_GetTodoParams.Size(m)
}
func (m *GetTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoParams proto.InternalMessageInfo

func (m *GetTodoParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoResponse struct {
	Todo                 *TodoObject `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetTodo() *TodoObject {
	if m != nil {
		return m.Todo
	}
	return nil
}

// add
type AddTodoParams struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoParams) Reset()         { *m = AddTodoParams{} }
func (m *AddTodoParams) String() string { return proto.CompactTextString(m) }
func (*AddTodoParams) ProtoMessage()    {}
func (*AddTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *AddTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoParams.Unmarshal(m, b)
}
func (m *AddTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoParams.Marshal(b, m, deterministic)
}
func (m *AddTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoParams.Merge(m, src)
}
func (m *AddTodoParams) XXX_Size() int {
	return xxx_messageInfo_AddTodoParams.Size(m)
}
func (m *AddTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoParams proto.InternalMessageInfo

func (m *AddTodoParams) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type AddTodoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoResponse) Reset()         { *m = AddTodoResponse{} }
func (m *AddTodoResponse) String() string { return proto.CompactTextString(m) }
func (*AddTodoResponse) ProtoMessage()    {}
func (*AddTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *AddTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoResponse.Unmarshal(m, b)
}
func (m *AddTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoResponse.Marshal(b, m, deterministic)
}
func (m *AddTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoResponse.Merge(m, src)
}
func (m *AddTodoResponse) XXX_Size() int {
	return xxx_messageInfo_AddTodoResponse.Size(m)
}
func (m *AddTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoResponse proto.InternalMessageInfo

func (m *AddTodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// update
type UpdateTodoParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoParams) Reset()         { *m = UpdateTodoParams{} }
func (m *UpdateTodoParams) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoParams) ProtoMessage()    {}
func (*UpdateTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *UpdateTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoParams.Unmarshal(m, b)
}
func (m *UpdateTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoParams.Marshal(b, m, deterministic)
}
func (m *UpdateTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoParams.Merge(m, src)
}
func (m *UpdateTodoParams) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoParams.Size(m)
}
func (m *UpdateTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoParams proto.InternalMessageInfo

func (m *UpdateTodoParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTodoParams) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type UpdateTodoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoResponse) Reset()         { *m = UpdateTodoResponse{} }
func (m *UpdateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoResponse) ProtoMessage()    {}
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{8}
}

func (m *UpdateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoResponse.Unmarshal(m, b)
}
func (m *UpdateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoResponse.Merge(m, src)
}
func (m *UpdateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoResponse.Size(m)
}
func (m *UpdateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoResponse proto.InternalMessageInfo

func (m *UpdateTodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// delete
type DeleteTodoParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoParams) Reset()         { *m = DeleteTodoParams{} }
func (m *DeleteTodoParams) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoParams) ProtoMessage()    {}
func (*DeleteTodoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{9}
}

func (m *DeleteTodoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoParams.Unmarshal(m, b)
}
func (m *DeleteTodoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoParams.Marshal(b, m, deterministic)
}
func (m *DeleteTodoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoParams.Merge(m, src)
}
func (m *DeleteTodoParams) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoParams.Size(m)
}
func (m *DeleteTodoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoParams proto.InternalMessageInfo

func (m *DeleteTodoParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{10}
}

func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(m, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

func (m *DeleteTodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TodoObject)(nil), "todo.todoObject")
	proto.RegisterType((*ListTodoParams)(nil), "todo.listTodoParams")
	proto.RegisterType((*ListTodoResponse)(nil), "todo.ListTodoResponse")
	proto.RegisterType((*GetTodoParams)(nil), "todo.getTodoParams")
	proto.RegisterType((*GetTodoResponse)(nil), "todo.getTodoResponse")
	proto.RegisterType((*AddTodoParams)(nil), "todo.addTodoParams")
	proto.RegisterType((*AddTodoResponse)(nil), "todo.addTodoResponse")
	proto.RegisterType((*UpdateTodoParams)(nil), "todo.updateTodoParams")
	proto.RegisterType((*UpdateTodoResponse)(nil), "todo.updateTodoResponse")
	proto.RegisterType((*DeleteTodoParams)(nil), "todo.deleteTodoParams")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.deleteTodoResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x4a, 0xe3, 0x40,
	0x14, 0xc6, 0x49, 0xda, 0x6e, 0x76, 0x4f, 0xb7, 0x6d, 0x38, 0xbb, 0x5d, 0x42, 0x58, 0xd8, 0x32,
	0x2b, 0x52, 0x2a, 0x34, 0x52, 0x41, 0xa1, 0x2f, 0x50, 0x04, 0x41, 0x89, 0xfa, 0x00, 0xd3, 0xce,
	0x10, 0xa2, 0x6d, 0x26, 0x74, 0x46, 0x6f, 0xc4, 0x1b, 0x5f, 0x41, 0x7c, 0x32, 0x5f, 0xc1, 0x07,
	0x91, 0x4c, 0x92, 0xe6, 0x8f, 0xf5, 0xcf, 0xdd, 0xe4, 0xcc, 0x37, 0xbf, 0x99, 0xf3, 0x7d, 0x27,
	0x00, 0x4a, 0x30, 0x31, 0x8e, 0xd7, 0x42, 0x09, 0x6c, 0x26, 0x6b, 0xf7, 0x6f, 0x20, 0x44, 0xb0,
	0xe4, 0x1e, 0x8d, 0x43, 0x8f, 0x46, 0x91, 0x50, 0x54, 0x85, 0x22, 0x92, 0xa9, 0x86, 0xec, 0xa7,
	0x27, 0x4e, 0xe7, 0x57, 0x7c, 0xa1, 0xb0, 0x0b, 0x66, 0xc8, 0x1c, 0x63, 0x60, 0x0c, 0x7f, 0xf8,
	0x66, 0xc8, 0x10, 0xa1, 0xa9, 0xa8, 0xbc, 0x76, 0x4c, 0x5d, 0xd1, 0x6b, 0x62, 0x43, 0x77, 0x19,
	0x4a, 0x75, 0x21, 0x98, 0x38, 0xa3, 0x6b, 0xba, 0x92, 0x64, 0x0a, 0xf6, 0x49, 0x56, 0xf1, 0xb9,
	0x8c, 0x45, 0x24, 0x39, 0xee, 0x42, 0x2b, 0xe1, 0x4a, 0xc7, 0x18, 0x34, 0x86, 0xed, 0x89, 0x3d,
	0xd6, 0xef, 0x2a, 0xae, 0xf2, 0xd3, 0x6d, 0xf2, 0x0f, 0x3a, 0x01, 0x2f, 0xc1, 0xea, 0x4f, 0x20,
	0x47, 0xd0, 0xcb, 0x04, 0x1b, 0xf6, 0x0e, 0xe8, 0xce, 0xb4, 0x68, 0x1b, 0x5a, 0xef, 0x92, 0xff,
	0xd0, 0xa1, 0x8c, 0x95, 0xc8, 0x79, 0x33, 0x46, 0xa9, 0x99, 0x3d, 0xe8, 0x65, 0xa2, 0x0d, 0xdd,
	0x01, 0x6b, 0xc5, 0xa5, 0xa4, 0x01, 0xcf, 0x94, 0xf9, 0x27, 0x39, 0x04, 0xfb, 0x26, 0x66, 0x54,
	0xf1, 0xf7, 0x9f, 0xbb, 0xd5, 0xb1, 0x31, 0x60, 0x71, 0xee, 0x0b, 0xf7, 0x10, 0xb0, 0x19, 0x5f,
	0xf2, 0x8f, 0xee, 0x49, 0x98, 0x85, 0xe6, 0x73, 0xe6, 0xe4, 0xa9, 0x01, 0xed, 0xc4, 0x96, 0x73,
	0xbe, 0xbe, 0x0d, 0x17, 0x1c, 0x67, 0xf0, 0x3d, 0x4f, 0x11, 0x7f, 0xa7, 0x0e, 0x56, 0x53, 0x75,
	0xff, 0xa4, 0xd5, 0x7a, 0xb2, 0xa4, 0xf3, 0xf0, 0xfc, 0xf2, 0x68, 0x5a, 0xd8, 0xf2, 0x92, 0x7d,
	0x9c, 0x81, 0x95, 0x39, 0x88, 0xbf, 0xd2, 0x13, 0x15, 0xd7, 0xdd, 0x7e, 0xa5, 0xb8, 0xa1, 0xd8,
	0x9a, 0x02, 0x24, 0xa5, 0x4c, 0x8d, 0x11, 0x1e, 0x83, 0x95, 0x05, 0x9d, 0x83, 0x2a, 0x83, 0x91,
	0x83, 0x6a, 0xc3, 0x40, 0x50, 0x83, 0x7e, 0x22, 0x68, 0x90, 0x77, 0x17, 0xb2, 0x7b, 0xf4, 0x01,
	0x0a, 0x73, 0x30, 0x6b, 0xa4, 0x6e, 0xa9, 0xeb, 0xd4, 0xeb, 0x75, 0xe6, 0xa8, 0xcc, 0xbc, 0x04,
	0x28, 0x42, 0xcc, 0x99, 0xf5, 0x71, 0xc8, 0x99, 0x6f, 0xe3, 0x26, 0x7d, 0xcd, 0xec, 0xb9, 0x25,
	0xe6, 0xd4, 0x18, 0xcd, 0xbf, 0xe9, 0xdf, 0xf0, 0xe0, 0x35, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x13,
	0x77, 0x62, 0xb8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	ListTodo(ctx context.Context, in *ListTodoParams, opts ...grpc.CallOption) (*ListTodoResponse, error)
	AddTodo(ctx context.Context, in *AddTodoParams, opts ...grpc.CallOption) (*AddTodoResponse, error)
	GetTodo(ctx context.Context, in *GetTodoParams, opts ...grpc.CallOption) (*GetTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoParams, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoParams, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoParams, opts ...grpc.CallOption) (*ListTodoResponse, error) {
	out := new(ListTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/listTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) AddTodo(ctx context.Context, in *AddTodoParams, opts ...grpc.CallOption) (*AddTodoResponse, error) {
	out := new(AddTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/addTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *GetTodoParams, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/getTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoParams, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/deleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoParams, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.todoService/updateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	ListTodo(context.Context, *ListTodoParams) (*ListTodoResponse, error)
	AddTodo(context.Context, *AddTodoParams) (*AddTodoResponse, error)
	GetTodo(context.Context, *GetTodoParams) (*GetTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoParams) (*DeleteTodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoParams) (*UpdateTodoResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodo(ctx, req.(*AddTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*GetTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.todoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
		{
			MethodName: "addTodo",
			Handler:    _TodoService_AddTodo_Handler,
		},
		{
			MethodName: "getTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
		{
			MethodName: "deleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "updateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
