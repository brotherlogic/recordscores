// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: recordscores.proto

package proto

import (
	proto1 "github.com/brotherlogic/recordcollection/proto"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Scores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *Scores) Reset() {
	*x = Scores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scores) ProtoMessage() {}

func (x *Scores) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scores.ProtoReflect.Descriptor instead.
func (*Scores) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{0}
}

func (x *Scores) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32                           `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Rating     int32                           `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Category   proto1.ReleaseMetadata_Category `protobuf:"varint,3,opt,name=category,proto3,enum=recordcollection.ReleaseMetadata_Category" json:"category,omitempty"`
	ScoreTime  int64                           `protobuf:"varint,4,opt,name=score_time,json=scoreTime,proto3" json:"score_time,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{1}
}

func (x *Score) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *Score) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Score) GetCategory() proto1.ReleaseMetadata_Category {
	if x != nil {
		return x.Category
	}
	return proto1.ReleaseMetadata_UNKNOWN
}

func (x *Score) GetScoreTime() int64 {
	if x != nil {
		return x.ScoreTime
	}
	return 0
}

type GetScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32                           `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Category   proto1.ReleaseMetadata_Category `protobuf:"varint,2,opt,name=category,proto3,enum=recordcollection.ReleaseMetadata_Category" json:"category,omitempty"`
}

func (x *GetScoreRequest) Reset() {
	*x = GetScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreRequest) ProtoMessage() {}

func (x *GetScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoreRequest.ProtoReflect.Descriptor instead.
func (*GetScoreRequest) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{2}
}

func (x *GetScoreRequest) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *GetScoreRequest) GetCategory() proto1.ReleaseMetadata_Category {
	if x != nil {
		return x.Category
	}
	return proto1.ReleaseMetadata_UNKNOWN
}

type GetScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *GetScoreResponse) Reset() {
	*x = GetScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreResponse) ProtoMessage() {}

func (x *GetScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoreResponse.ProtoReflect.Descriptor instead.
func (*GetScoreResponse) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{3}
}

func (x *GetScoreResponse) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_recordscores_proto protoreflect.FileDescriptor

var file_recordscores_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x06, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x22, 0xa7, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x46,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x32, 0x61, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recordscores_proto_rawDescOnce sync.Once
	file_recordscores_proto_rawDescData = file_recordscores_proto_rawDesc
)

func file_recordscores_proto_rawDescGZIP() []byte {
	file_recordscores_proto_rawDescOnce.Do(func() {
		file_recordscores_proto_rawDescData = protoimpl.X.CompressGZIP(file_recordscores_proto_rawDescData)
	})
	return file_recordscores_proto_rawDescData
}

var file_recordscores_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_recordscores_proto_goTypes = []interface{}{
	(*Scores)(nil),                       // 0: recordscores.Scores
	(*Score)(nil),                        // 1: recordscores.Score
	(*GetScoreRequest)(nil),              // 2: recordscores.GetScoreRequest
	(*GetScoreResponse)(nil),             // 3: recordscores.GetScoreResponse
	(proto1.ReleaseMetadata_Category)(0), // 4: recordcollection.ReleaseMetadata.Category
}
var file_recordscores_proto_depIdxs = []int32{
	1, // 0: recordscores.Scores.scores:type_name -> recordscores.Score
	4, // 1: recordscores.Score.category:type_name -> recordcollection.ReleaseMetadata.Category
	4, // 2: recordscores.GetScoreRequest.category:type_name -> recordcollection.ReleaseMetadata.Category
	1, // 3: recordscores.GetScoreResponse.scores:type_name -> recordscores.Score
	2, // 4: recordscores.RecordScoreService.GetScore:input_type -> recordscores.GetScoreRequest
	3, // 5: recordscores.RecordScoreService.GetScore:output_type -> recordscores.GetScoreResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_recordscores_proto_init() }
func file_recordscores_proto_init() {
	if File_recordscores_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recordscores_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scores); i {
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
		file_recordscores_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_recordscores_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoreRequest); i {
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
		file_recordscores_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoreResponse); i {
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
			RawDescriptor: file_recordscores_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recordscores_proto_goTypes,
		DependencyIndexes: file_recordscores_proto_depIdxs,
		MessageInfos:      file_recordscores_proto_msgTypes,
	}.Build()
	File_recordscores_proto = out.File
	file_recordscores_proto_rawDesc = nil
	file_recordscores_proto_goTypes = nil
	file_recordscores_proto_depIdxs = nil
}
