// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: ituserver/ituserver.proto

package ituserver

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

type CourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CourseRequest) Reset() {
	*x = CourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ituserver_ituserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseRequest) ProtoMessage() {}

func (x *CourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ituserver_ituserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseRequest.ProtoReflect.Descriptor instead.
func (*CourseRequest) Descriptor() ([]byte, []int) {
	return file_ituserver_ituserver_proto_rawDescGZIP(), []int{0}
}

func (x *CourseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Teacher         string  `protobuf:"bytes,3,opt,name=teacher,proto3" json:"teacher,omitempty"`
	CourseManager   string  `protobuf:"bytes,4,opt,name=courseManager,proto3" json:"courseManager,omitempty"`
	NumParticipants int32   `protobuf:"varint,5,opt,name=numParticipants,proto3" json:"numParticipants,omitempty"`
	Semester        string  `protobuf:"bytes,6,opt,name=semester,proto3" json:"semester,omitempty"`
	ECTS            float32 `protobuf:"fixed32,7,opt,name=ECTS,proto3" json:"ECTS,omitempty"`
	ProgrammingLang string  `protobuf:"bytes,8,opt,name=programmingLang,proto3" json:"programmingLang,omitempty"`
}

func (x *CourseReply) Reset() {
	*x = CourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ituserver_ituserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseReply) ProtoMessage() {}

func (x *CourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_ituserver_ituserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseReply.ProtoReflect.Descriptor instead.
func (*CourseReply) Descriptor() ([]byte, []int) {
	return file_ituserver_ituserver_proto_rawDescGZIP(), []int{1}
}

func (x *CourseReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CourseReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CourseReply) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *CourseReply) GetCourseManager() string {
	if x != nil {
		return x.CourseManager
	}
	return ""
}

func (x *CourseReply) GetNumParticipants() int32 {
	if x != nil {
		return x.NumParticipants
	}
	return 0
}

func (x *CourseReply) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *CourseReply) GetECTS() float32 {
	if x != nil {
		return x.ECTS
	}
	return 0
}

func (x *CourseReply) GetProgrammingLang() string {
	if x != nil {
		return x.ProgrammingLang
	}
	return ""
}

var File_ituserver_ituserver_proto protoreflect.FileDescriptor

var file_ituserver_ituserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x74, 0x75, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x74, 0x75, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x74, 0x75,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x6e,
	0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x45, 0x43, 0x54, 0x53, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x45, 0x43, 0x54, 0x53, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x32,
	0x4b, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x69, 0x74, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x74, 0x75, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x4c, 0x46, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x44, 0x53, 0x59, 0x53,
	0x45, 0x58, 0x31, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x74,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x74, 0x75, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ituserver_ituserver_proto_rawDescOnce sync.Once
	file_ituserver_ituserver_proto_rawDescData = file_ituserver_ituserver_proto_rawDesc
)

func file_ituserver_ituserver_proto_rawDescGZIP() []byte {
	file_ituserver_ituserver_proto_rawDescOnce.Do(func() {
		file_ituserver_ituserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_ituserver_ituserver_proto_rawDescData)
	})
	return file_ituserver_ituserver_proto_rawDescData
}

var file_ituserver_ituserver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ituserver_ituserver_proto_goTypes = []interface{}{
	(*CourseRequest)(nil), // 0: ituserver.CourseRequest
	(*CourseReply)(nil),   // 1: ituserver.CourseReply
}
var file_ituserver_ituserver_proto_depIdxs = []int32{
	0, // 0: ituserver.Cources.GetCourses:input_type -> ituserver.CourseRequest
	1, // 1: ituserver.Cources.GetCourses:output_type -> ituserver.CourseReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ituserver_ituserver_proto_init() }
func file_ituserver_ituserver_proto_init() {
	if File_ituserver_ituserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ituserver_ituserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseRequest); i {
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
		file_ituserver_ituserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseReply); i {
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
			RawDescriptor: file_ituserver_ituserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ituserver_ituserver_proto_goTypes,
		DependencyIndexes: file_ituserver_ituserver_proto_depIdxs,
		MessageInfos:      file_ituserver_ituserver_proto_msgTypes,
	}.Build()
	File_ituserver_ituserver_proto = out.File
	file_ituserver_ituserver_proto_rawDesc = nil
	file_ituserver_ituserver_proto_goTypes = nil
	file_ituserver_ituserver_proto_depIdxs = nil
}
