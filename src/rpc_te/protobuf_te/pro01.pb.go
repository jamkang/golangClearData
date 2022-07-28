// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pro01.proto

package protobuf_te

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//枚举类型
type Person_Corpus int32

const (
	Person_UNIVER Person_Corpus = 0
	Person_WEB    Person_Corpus = 1
	Person_IMAGE  Person_Corpus = 2
	Person_KEUAI  Person_Corpus = 2
)

// Enum value maps for Person_Corpus.
var (
	Person_Corpus_name = map[int32]string{
		0: "UNIVER",
		1: "WEB",
		2: "IMAGE",
		// Duplicate value: 2: "KEUAI",
	}
	Person_Corpus_value = map[string]int32{
		"UNIVER": 0,
		"WEB":    1,
		"IMAGE":  2,
		"KEUAI":  2,
	}
)

func (x Person_Corpus) Enum() *Person_Corpus {
	p := new(Person_Corpus)
	*p = x
	return p
}

func (x Person_Corpus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Corpus) Descriptor() protoreflect.EnumDescriptor {
	return file_pro01_proto_enumTypes[0].Descriptor()
}

func (Person_Corpus) Type() protoreflect.EnumType {
	return &file_pro01_proto_enumTypes[0]
}

func (x Person_Corpus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Corpus.Descriptor instead.
func (Person_Corpus) EnumDescriptor() ([]byte, []int) {
	return file_pro01_proto_rawDescGZIP(), []int{0, 0}
}

//message为关键词，有点像定义结构体
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    int32         `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Corpus Person_Corpus `protobuf:"varint,4,opt,name=corpus,proto3,enum=protobuf_te.Person_Corpus" json:"corpus,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro01_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_pro01_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_pro01_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetCorpus() Person_Corpus {
	if x != nil {
		return x.Corpus
	}
	return Person_UNIVER
}

//嵌套
type SearchRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchRespone) Reset() {
	*x = SearchRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro01_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRespone) ProtoMessage() {}

func (x *SearchRespone) ProtoReflect() protoreflect.Message {
	mi := &file_pro01_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRespone.ProtoReflect.Descriptor instead.
func (*SearchRespone) Descriptor() ([]byte, []int) {
	return file_pro01_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRespone) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	//Any 可以表示不在.proto 中定义任意的内置类型
	Detail []*anypb.Any `protobuf:"bytes,3,rep,name=detail,proto3" json:"detail,omitempty"`
	//oneof
	//
	// Types that are assignable to TestOneof:
	//	*Result_Name
	TestOneof isResult_TestOneof `protobuf_oneof:"test_oneof"`
	//map
	Points map[string]int32 `protobuf:"bytes,5,rep,name=points,proto3" json:"points,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro01_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pro01_proto_msgTypes[2]
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
	return file_pro01_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Result) GetDetail() []*anypb.Any {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (m *Result) GetTestOneof() isResult_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *Result) GetName() string {
	if x, ok := x.GetTestOneof().(*Result_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Result) GetPoints() map[string]int32 {
	if x != nil {
		return x.Points
	}
	return nil
}

type isResult_TestOneof interface {
	isResult_TestOneof()
}

type Result_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"` //map不能定义在这个里面
}

func (*Result_Name) isResult_TestOneof() {}

var File_pro01_proto protoreflect.FileDescriptor

var file_pro01_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x74, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x74, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x06, 0x43, 0x6f,
	0x72, 0x70, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x45, 0x55, 0x41, 0x49, 0x10, 0x02, 0x1a,
	0x02, 0x10, 0x01, 0x22, 0x3e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5f, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro01_proto_rawDescOnce sync.Once
	file_pro01_proto_rawDescData = file_pro01_proto_rawDesc
)

func file_pro01_proto_rawDescGZIP() []byte {
	file_pro01_proto_rawDescOnce.Do(func() {
		file_pro01_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro01_proto_rawDescData)
	})
	return file_pro01_proto_rawDescData
}

var file_pro01_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pro01_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pro01_proto_goTypes = []interface{}{
	(Person_Corpus)(0),    // 0: protobuf_te.Person.Corpus
	(*Person)(nil),        // 1: protobuf_te.Person
	(*SearchRespone)(nil), // 2: protobuf_te.SearchRespone
	(*Result)(nil),        // 3: protobuf_te.Result
	nil,                   // 4: protobuf_te.Result.PointsEntry
	(*anypb.Any)(nil),     // 5: google.protobuf.Any
}
var file_pro01_proto_depIdxs = []int32{
	0, // 0: protobuf_te.Person.corpus:type_name -> protobuf_te.Person.Corpus
	3, // 1: protobuf_te.SearchRespone.results:type_name -> protobuf_te.Result
	5, // 2: protobuf_te.Result.detail:type_name -> google.protobuf.Any
	4, // 3: protobuf_te.Result.points:type_name -> protobuf_te.Result.PointsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pro01_proto_init() }
func file_pro01_proto_init() {
	if File_pro01_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro01_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_pro01_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRespone); i {
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
		file_pro01_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_pro01_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Result_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pro01_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro01_proto_goTypes,
		DependencyIndexes: file_pro01_proto_depIdxs,
		EnumInfos:         file_pro01_proto_enumTypes,
		MessageInfos:      file_pro01_proto_msgTypes,
	}.Build()
	File_pro01_proto = out.File
	file_pro01_proto_rawDesc = nil
	file_pro01_proto_goTypes = nil
	file_pro01_proto_depIdxs = nil
}