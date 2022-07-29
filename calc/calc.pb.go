// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: calc/calc.proto

package calc

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

type TwoNumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *TwoNumRequest) Reset() {
	*x = TwoNumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoNumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoNumRequest) ProtoMessage() {}

func (x *TwoNumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoNumRequest.ProtoReflect.Descriptor instead.
func (*TwoNumRequest) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{0}
}

func (x *TwoNumRequest) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *TwoNumRequest) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type NumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *NumRequest) Reset() {
	*x = NumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumRequest) ProtoMessage() {}

func (x *NumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumRequest.ProtoReflect.Descriptor instead.
func (*NumRequest) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{2}
}

func (x *NumRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AllPrimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *AllPrimesResponse) Reset() {
	*x = AllPrimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPrimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPrimesResponse) ProtoMessage() {}

func (x *AllPrimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPrimesResponse.ProtoReflect.Descriptor instead.
func (*AllPrimesResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{3}
}

func (x *AllPrimesResponse) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float32 `protobuf:"fixed32,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{4}
}

func (x *AverageResponse) GetNum() float32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type MaxNumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *MaxNumResponse) Reset() {
	*x = MaxNumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxNumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxNumResponse) ProtoMessage() {}

func (x *MaxNumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxNumResponse.ProtoReflect.Descriptor instead.
func (*MaxNumResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_rawDescGZIP(), []int{5}
}

func (x *MaxNumResponse) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_calc_calc_proto protoreflect.FileDescriptor

var file_calc_calc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x37, 0x0a, 0x0d, 0x54, 0x77, 0x6f, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32,
	0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75,
	0x6d, 0x22, 0x1e, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x22, 0x25, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x22, 0x0a,
	0x0e, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x32, 0xe9, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x32, 0x0a, 0x06, 0x54, 0x77,
	0x6f, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x54, 0x77, 0x6f, 0x4e,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4e, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x39, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_calc_calc_proto_rawDescOnce sync.Once
	file_calc_calc_proto_rawDescData = file_calc_calc_proto_rawDesc
)

func file_calc_calc_proto_rawDescGZIP() []byte {
	file_calc_calc_proto_rawDescOnce.Do(func() {
		file_calc_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_calc_proto_rawDescData)
	})
	return file_calc_calc_proto_rawDescData
}

var file_calc_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_calc_calc_proto_goTypes = []interface{}{
	(*TwoNumRequest)(nil),     // 0: calc.TwoNumRequest
	(*SumResponse)(nil),       // 1: calc.SumResponse
	(*NumRequest)(nil),        // 2: calc.NumRequest
	(*AllPrimesResponse)(nil), // 3: calc.AllPrimesResponse
	(*AverageResponse)(nil),   // 4: calc.AverageResponse
	(*MaxNumResponse)(nil),    // 5: calc.MaxNumResponse
}
var file_calc_calc_proto_depIdxs = []int32{
	0, // 0: calc.Calc.TwoSum:input_type -> calc.TwoNumRequest
	2, // 1: calc.Calc.PrimeNums:input_type -> calc.NumRequest
	2, // 2: calc.Calc.Average:input_type -> calc.NumRequest
	2, // 3: calc.Calc.MaxNumber:input_type -> calc.NumRequest
	1, // 4: calc.Calc.TwoSum:output_type -> calc.SumResponse
	3, // 5: calc.Calc.PrimeNums:output_type -> calc.AllPrimesResponse
	4, // 6: calc.Calc.Average:output_type -> calc.AverageResponse
	5, // 7: calc.Calc.MaxNumber:output_type -> calc.MaxNumResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_calc_proto_init() }
func file_calc_calc_proto_init() {
	if File_calc_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoNumRequest); i {
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
		file_calc_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calc_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumRequest); i {
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
		file_calc_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPrimesResponse); i {
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
		file_calc_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
		file_calc_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxNumResponse); i {
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
			RawDescriptor: file_calc_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_calc_proto_goTypes,
		DependencyIndexes: file_calc_calc_proto_depIdxs,
		MessageInfos:      file_calc_calc_proto_msgTypes,
	}.Build()
	File_calc_calc_proto = out.File
	file_calc_calc_proto_rawDesc = nil
	file_calc_calc_proto_goTypes = nil
	file_calc_calc_proto_depIdxs = nil
}