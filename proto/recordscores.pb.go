// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: recordscores.proto

package proto

import (
	proto "github.com/brotherlogic/recordcollection/proto"
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

type ScoreAdjustment_AdjustmentType int32

const (
	ScoreAdjustment_IGNORED_ADJUSTMENT          ScoreAdjustment_AdjustmentType = 0
	ScoreAdjustment_OTHER_VERSIONS_ADJUSTMENT   ScoreAdjustment_AdjustmentType = 1
	ScoreAdjustment_DIGITAL_VERSIONS_ADJUSTMENT ScoreAdjustment_AdjustmentType = 2
	ScoreAdjustment_OWN_OTHER_ADJUSTMENT        ScoreAdjustment_AdjustmentType = 3
	ScoreAdjustment_PREVIOUSY_SOLD_ADJUSTMENT   ScoreAdjustment_AdjustmentType = 4
	ScoreAdjustment_KEEP_ADJUSTMENT             ScoreAdjustment_AdjustmentType = 5
	ScoreAdjustment_UNKNOWN_KEEP_ADJUSTMENT     ScoreAdjustment_AdjustmentType = 6
)

// Enum value maps for ScoreAdjustment_AdjustmentType.
var (
	ScoreAdjustment_AdjustmentType_name = map[int32]string{
		0: "IGNORED_ADJUSTMENT",
		1: "OTHER_VERSIONS_ADJUSTMENT",
		2: "DIGITAL_VERSIONS_ADJUSTMENT",
		3: "OWN_OTHER_ADJUSTMENT",
		4: "PREVIOUSY_SOLD_ADJUSTMENT",
		5: "KEEP_ADJUSTMENT",
		6: "UNKNOWN_KEEP_ADJUSTMENT",
	}
	ScoreAdjustment_AdjustmentType_value = map[string]int32{
		"IGNORED_ADJUSTMENT":          0,
		"OTHER_VERSIONS_ADJUSTMENT":   1,
		"DIGITAL_VERSIONS_ADJUSTMENT": 2,
		"OWN_OTHER_ADJUSTMENT":        3,
		"PREVIOUSY_SOLD_ADJUSTMENT":   4,
		"KEEP_ADJUSTMENT":             5,
		"UNKNOWN_KEEP_ADJUSTMENT":     6,
	}
)

func (x ScoreAdjustment_AdjustmentType) Enum() *ScoreAdjustment_AdjustmentType {
	p := new(ScoreAdjustment_AdjustmentType)
	*p = x
	return p
}

func (x ScoreAdjustment_AdjustmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScoreAdjustment_AdjustmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_recordscores_proto_enumTypes[0].Descriptor()
}

func (ScoreAdjustment_AdjustmentType) Type() protoreflect.EnumType {
	return &file_recordscores_proto_enumTypes[0]
}

func (x ScoreAdjustment_AdjustmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScoreAdjustment_AdjustmentType.Descriptor instead.
func (ScoreAdjustment_AdjustmentType) EnumDescriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{3, 0}
}

type Scores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores    []*Score                 `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	LastScore map[int32]*ComputedScore `protobuf:"bytes,2,rep,name=last_score,json=lastScore,proto3" json:"last_score,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (x *Scores) GetLastScore() map[int32]*ComputedScore {
	if x != nil {
		return x.LastScore
	}
	return nil
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32                          `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Rating     int32                          `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Category   proto.ReleaseMetadata_Category `protobuf:"varint,3,opt,name=category,proto3,enum=recordcollection.ReleaseMetadata_Category" json:"category,omitempty"`
	ScoreTime  int64                          `protobuf:"varint,4,opt,name=score_time,json=scoreTime,proto3" json:"score_time,omitempty"`
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

func (x *Score) GetCategory() proto.ReleaseMetadata_Category {
	if x != nil {
		return x.Category
	}
	return proto.ReleaseMetadata_UNKNOWN
}

func (x *Score) GetScoreTime() int64 {
	if x != nil {
		return x.ScoreTime
	}
	return 0
}

type ComputedScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRating  int32                  `protobuf:"varint,1,opt,name=base_rating,json=baseRating,proto3" json:"base_rating,omitempty"`
	Adjustments []*ScoreAdjustment     `protobuf:"bytes,2,rep,name=adjustments,proto3" json:"adjustments,omitempty"`
	Overall     float32                `protobuf:"fixed32,3,opt,name=overall,proto3" json:"overall,omitempty"`
	CurrFolder  int32                  `protobuf:"varint,4,opt,name=curr_folder,json=currFolder,proto3" json:"curr_folder,omitempty"`
	Location    proto.PurchaseLocation `protobuf:"varint,5,opt,name=location,proto3,enum=recordcollection.PurchaseLocation" json:"location,omitempty"`
}

func (x *ComputedScore) Reset() {
	*x = ComputedScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedScore) ProtoMessage() {}

func (x *ComputedScore) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ComputedScore.ProtoReflect.Descriptor instead.
func (*ComputedScore) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{2}
}

func (x *ComputedScore) GetBaseRating() int32 {
	if x != nil {
		return x.BaseRating
	}
	return 0
}

func (x *ComputedScore) GetAdjustments() []*ScoreAdjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

func (x *ComputedScore) GetOverall() float32 {
	if x != nil {
		return x.Overall
	}
	return 0
}

func (x *ComputedScore) GetCurrFolder() int32 {
	if x != nil {
		return x.CurrFolder
	}
	return 0
}

func (x *ComputedScore) GetLocation() proto.PurchaseLocation {
	if x != nil {
		return x.Location
	}
	return proto.PurchaseLocation_LOCATION_UNKNOWN
}

type ScoreAdjustment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ScoreAdjustment_AdjustmentType `protobuf:"varint,1,opt,name=type,proto3,enum=recordscores.ScoreAdjustment_AdjustmentType" json:"type,omitempty"`
	ValueChange float32                        `protobuf:"fixed32,2,opt,name=value_change,json=valueChange,proto3" json:"value_change,omitempty"`
}

func (x *ScoreAdjustment) Reset() {
	*x = ScoreAdjustment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreAdjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreAdjustment) ProtoMessage() {}

func (x *ScoreAdjustment) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ScoreAdjustment.ProtoReflect.Descriptor instead.
func (*ScoreAdjustment) Descriptor() ([]byte, []int) {
	return file_recordscores_proto_rawDescGZIP(), []int{3}
}

func (x *ScoreAdjustment) GetType() ScoreAdjustment_AdjustmentType {
	if x != nil {
		return x.Type
	}
	return ScoreAdjustment_IGNORED_ADJUSTMENT
}

func (x *ScoreAdjustment) GetValueChange() float32 {
	if x != nil {
		return x.ValueChange
	}
	return 0
}

type GetScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32                          `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Category   proto.ReleaseMetadata_Category `protobuf:"varint,2,opt,name=category,proto3,enum=recordcollection.ReleaseMetadata_Category" json:"category,omitempty"`
}

func (x *GetScoreRequest) Reset() {
	*x = GetScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreRequest) ProtoMessage() {}

func (x *GetScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[4]
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
	return file_recordscores_proto_rawDescGZIP(), []int{4}
}

func (x *GetScoreRequest) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *GetScoreRequest) GetCategory() proto.ReleaseMetadata_Category {
	if x != nil {
		return x.Category
	}
	return proto.ReleaseMetadata_UNKNOWN
}

type GetScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores        []*Score       `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	ComputedScore *ComputedScore `protobuf:"bytes,2,opt,name=computed_score,json=computedScore,proto3" json:"computed_score,omitempty"`
}

func (x *GetScoreResponse) Reset() {
	*x = GetScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordscores_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreResponse) ProtoMessage() {}

func (x *GetScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recordscores_proto_msgTypes[5]
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
	return file_recordscores_proto_rawDescGZIP(), []int{5}
}

func (x *GetScoreResponse) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *GetScoreResponse) GetComputedScore() *ComputedScore {
	if x != nil {
		return x.ComputedScore
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
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x06, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x42, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x73, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x59, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
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
	0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a,
	0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x72, 0x72, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x02, 0x0a, 0x0f, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x44,
	0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x19, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x4f, 0x57, 0x4e, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53,
	0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x45, 0x56, 0x49,
	0x4f, 0x55, 0x53, 0x59, 0x5f, 0x53, 0x4f, 0x4c, 0x44, 0x5f, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x41, 0x44, 0x4a, 0x55,
	0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x22, 0x7a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x61, 0x0a, 0x12, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_recordscores_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_recordscores_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_recordscores_proto_goTypes = []interface{}{
	(ScoreAdjustment_AdjustmentType)(0), // 0: recordscores.ScoreAdjustment.AdjustmentType
	(*Scores)(nil),                      // 1: recordscores.Scores
	(*Score)(nil),                       // 2: recordscores.Score
	(*ComputedScore)(nil),               // 3: recordscores.ComputedScore
	(*ScoreAdjustment)(nil),             // 4: recordscores.ScoreAdjustment
	(*GetScoreRequest)(nil),             // 5: recordscores.GetScoreRequest
	(*GetScoreResponse)(nil),            // 6: recordscores.GetScoreResponse
	nil,                                 // 7: recordscores.Scores.LastScoreEntry
	(proto.ReleaseMetadata_Category)(0), // 8: recordcollection.ReleaseMetadata.Category
	(proto.PurchaseLocation)(0),         // 9: recordcollection.PurchaseLocation
}
var file_recordscores_proto_depIdxs = []int32{
	2,  // 0: recordscores.Scores.scores:type_name -> recordscores.Score
	7,  // 1: recordscores.Scores.last_score:type_name -> recordscores.Scores.LastScoreEntry
	8,  // 2: recordscores.Score.category:type_name -> recordcollection.ReleaseMetadata.Category
	4,  // 3: recordscores.ComputedScore.adjustments:type_name -> recordscores.ScoreAdjustment
	9,  // 4: recordscores.ComputedScore.location:type_name -> recordcollection.PurchaseLocation
	0,  // 5: recordscores.ScoreAdjustment.type:type_name -> recordscores.ScoreAdjustment.AdjustmentType
	8,  // 6: recordscores.GetScoreRequest.category:type_name -> recordcollection.ReleaseMetadata.Category
	2,  // 7: recordscores.GetScoreResponse.scores:type_name -> recordscores.Score
	3,  // 8: recordscores.GetScoreResponse.computed_score:type_name -> recordscores.ComputedScore
	3,  // 9: recordscores.Scores.LastScoreEntry.value:type_name -> recordscores.ComputedScore
	5,  // 10: recordscores.RecordScoreService.GetScore:input_type -> recordscores.GetScoreRequest
	6,  // 11: recordscores.RecordScoreService.GetScore:output_type -> recordscores.GetScoreResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
			switch v := v.(*ComputedScore); i {
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
			switch v := v.(*ScoreAdjustment); i {
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
		file_recordscores_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_recordscores_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recordscores_proto_goTypes,
		DependencyIndexes: file_recordscores_proto_depIdxs,
		EnumInfos:         file_recordscores_proto_enumTypes,
		MessageInfos:      file_recordscores_proto_msgTypes,
	}.Build()
	File_recordscores_proto = out.File
	file_recordscores_proto_rawDesc = nil
	file_recordscores_proto_goTypes = nil
	file_recordscores_proto_depIdxs = nil
}
