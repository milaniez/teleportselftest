// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protojob.proto

package protojob

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

type Status_RunStatus int32

const (
	Status_RUN_STATUS_PENDING  Status_RunStatus = 0
	Status_RUN_STATUS_RUNNING  Status_RunStatus = 1
	Status_RUN_STATUS_FINISHED Status_RunStatus = 2
)

// Enum value maps for Status_RunStatus.
var (
	Status_RunStatus_name = map[int32]string{
		0: "RUN_STATUS_PENDING",
		1: "RUN_STATUS_RUNNING",
		2: "RUN_STATUS_FINISHED",
	}
	Status_RunStatus_value = map[string]int32{
		"RUN_STATUS_PENDING":  0,
		"RUN_STATUS_RUNNING":  1,
		"RUN_STATUS_FINISHED": 2,
	}
)

func (x Status_RunStatus) Enum() *Status_RunStatus {
	p := new(Status_RunStatus)
	*p = x
	return p
}

func (x Status_RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protojob_proto_enumTypes[0].Descriptor()
}

func (Status_RunStatus) Type() protoreflect.EnumType {
	return &file_protojob_proto_enumTypes[0]
}

func (x Status_RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_RunStatus.Descriptor instead.
func (Status_RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{4, 0}
}

type Status_ExitStatus int32

const (
	Status_EXIT_STATUS_OK            Status_ExitStatus = 0
	Status_EXIT_STATUS_ABORTED       Status_ExitStatus = 1
	Status_EXIT_STATUS_FAILURE       Status_ExitStatus = 2
	Status_EXIT_STATUS_TIME_EXCEEDED Status_ExitStatus = 3
	Status_EXIT_STATUS_CPU_EXCEEDED  Status_ExitStatus = 4
	Status_EXIT_STATUS_MEM_EXCEEDED  Status_ExitStatus = 5
)

// Enum value maps for Status_ExitStatus.
var (
	Status_ExitStatus_name = map[int32]string{
		0: "EXIT_STATUS_OK",
		1: "EXIT_STATUS_ABORTED",
		2: "EXIT_STATUS_FAILURE",
		3: "EXIT_STATUS_TIME_EXCEEDED",
		4: "EXIT_STATUS_CPU_EXCEEDED",
		5: "EXIT_STATUS_MEM_EXCEEDED",
	}
	Status_ExitStatus_value = map[string]int32{
		"EXIT_STATUS_OK":            0,
		"EXIT_STATUS_ABORTED":       1,
		"EXIT_STATUS_FAILURE":       2,
		"EXIT_STATUS_TIME_EXCEEDED": 3,
		"EXIT_STATUS_CPU_EXCEEDED":  4,
		"EXIT_STATUS_MEM_EXCEEDED":  5,
	}
)

func (x Status_ExitStatus) Enum() *Status_ExitStatus {
	p := new(Status_ExitStatus)
	*p = x
	return p
}

func (x Status_ExitStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_ExitStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protojob_proto_enumTypes[1].Descriptor()
}

func (Status_ExitStatus) Type() protoreflect.EnumType {
	return &file_protojob_proto_enumTypes[1]
}

func (x Status_ExitStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_ExitStatus.Descriptor instead.
func (Status_ExitStatus) EnumDescriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{4, 1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{0}
}

type JobID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobID) Reset() {
	*x = JobID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobID) ProtoMessage() {}

func (x *JobID) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobID.ProtoReflect.Descriptor instead.
func (*JobID) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{1}
}

func (x *JobID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type JobIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *JobIDList) Reset() {
	*x = JobIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobIDList) ProtoMessage() {}

func (x *JobIDList) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobIDList.ProtoReflect.Descriptor instead.
func (*JobIDList) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{2}
}

func (x *JobIDList) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{3}
}

func (x *Cmd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cmd) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunStatus  Status_RunStatus  `protobuf:"varint,1,opt,name=runStatus,proto3,enum=protojob.Status_RunStatus" json:"runStatus,omitempty"`
	ExitStatus Status_ExitStatus `protobuf:"varint,2,opt,name=exitStatus,proto3,enum=protojob.Status_ExitStatus" json:"exitStatus,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetRunStatus() Status_RunStatus {
	if x != nil {
		return x.RunStatus
	}
	return Status_RUN_STATUS_PENDING
}

func (x *Status) GetExitStatus() Status_ExitStatus {
	if x != nil {
		return x.ExitStatus
	}
	return Status_EXIT_STATUS_OK
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Finished bool   `protobuf:"varint,1,opt,name=finished,proto3" json:"finished,omitempty"`
	Output   string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *Result) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protojob_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_protojob_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_protojob_proto_rawDescGZIP(), []int{6}
}

func (x *Output) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_protojob_proto protoreflect.FileDescriptor

var file_protojob_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x17, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09,
	0x4a, 0x6f, 0x62, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x03, 0x43, 0x6d, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x85, 0x03, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a,
	0x0a, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a,
	0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x09, 0x52, 0x75,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x55, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02,
	0x22, 0xad, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x0e, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x50, 0x55, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4d, 0x45, 0x4d, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x05,
	0x22, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x20,
	0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x83, 0x02, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e,
	0x43, 0x6d, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x49, 0x44, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x49, 0x44, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f,
	0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6c, 0x61, 0x6e, 0x69, 0x65, 0x7a, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x65, 0x6c, 0x66, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protojob_proto_rawDescOnce sync.Once
	file_protojob_proto_rawDescData = file_protojob_proto_rawDesc
)

func file_protojob_proto_rawDescGZIP() []byte {
	file_protojob_proto_rawDescOnce.Do(func() {
		file_protojob_proto_rawDescData = protoimpl.X.CompressGZIP(file_protojob_proto_rawDescData)
	})
	return file_protojob_proto_rawDescData
}

var file_protojob_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protojob_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protojob_proto_goTypes = []interface{}{
	(Status_RunStatus)(0),  // 0: protojob.Status.RunStatus
	(Status_ExitStatus)(0), // 1: protojob.Status.ExitStatus
	(*Empty)(nil),          // 2: protojob.Empty
	(*JobID)(nil),          // 3: protojob.JobID
	(*JobIDList)(nil),      // 4: protojob.JobIDList
	(*Cmd)(nil),            // 5: protojob.Cmd
	(*Status)(nil),         // 6: protojob.Status
	(*Result)(nil),         // 7: protojob.Result
	(*Output)(nil),         // 8: protojob.Output
}
var file_protojob_proto_depIdxs = []int32{
	0, // 0: protojob.Status.runStatus:type_name -> protojob.Status.RunStatus
	1, // 1: protojob.Status.exitStatus:type_name -> protojob.Status.ExitStatus
	5, // 2: protojob.Worker.Start:input_type -> protojob.Cmd
	3, // 3: protojob.Worker.GetStatus:input_type -> protojob.JobID
	2, // 4: protojob.Worker.GetJobIDs:input_type -> protojob.Empty
	3, // 5: protojob.Worker.GetResult:input_type -> protojob.JobID
	3, // 6: protojob.Worker.StreamOutput:input_type -> protojob.JobID
	3, // 7: protojob.Worker.Start:output_type -> protojob.JobID
	6, // 8: protojob.Worker.GetStatus:output_type -> protojob.Status
	4, // 9: protojob.Worker.GetJobIDs:output_type -> protojob.JobIDList
	7, // 10: protojob.Worker.GetResult:output_type -> protojob.Result
	8, // 11: protojob.Worker.StreamOutput:output_type -> protojob.Output
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protojob_proto_init() }
func file_protojob_proto_init() {
	if File_protojob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protojob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protojob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobID); i {
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
		file_protojob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobIDList); i {
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
		file_protojob_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
		file_protojob_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_protojob_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_protojob_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_protojob_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protojob_proto_goTypes,
		DependencyIndexes: file_protojob_proto_depIdxs,
		EnumInfos:         file_protojob_proto_enumTypes,
		MessageInfos:      file_protojob_proto_msgTypes,
	}.Build()
	File_protojob_proto = out.File
	file_protojob_proto_rawDesc = nil
	file_protojob_proto_goTypes = nil
	file_protojob_proto_depIdxs = nil
}
