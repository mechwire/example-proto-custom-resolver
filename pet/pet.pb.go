// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pet.proto

package pet

import (
	example_proto_types "github.com/jncmaguire/example-proto-types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Blank) Reset() {
	*x = Blank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{0}
}

type Animal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classification *example_proto_types.Classification `protobuf:"bytes,1,opt,name=classification,proto3" json:"classification,omitempty"`
	Names          []string                            `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *Animal) Reset() {
	*x = Animal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Animal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Animal) ProtoMessage() {}

func (x *Animal) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Animal.ProtoReflect.Descriptor instead.
func (*Animal) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{1}
}

func (x *Animal) GetClassification() *example_proto_types.Classification {
	if x != nil {
		return x.Classification
	}
	return nil
}

func (x *Animal) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DateOfBirth *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Animal      *Animal                `protobuf:"bytes,3,opt,name=animal,proto3" json:"animal,omitempty"`                                                                                           // Message
	Friends     []*Animal              `protobuf:"bytes,4,rep,name=friends,proto3" json:"friends,omitempty"`                                                                                         // List
	Parents     map[string]*Animal     `protobuf:"bytes,5,rep,name=parents,proto3" json:"parents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_pet_proto_rawDescGZIP(), []int{2}
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *Pet) GetAnimal() *Animal {
	if x != nil {
		return x.Animal
	}
	return nil
}

func (x *Pet) GetFriends() []*Animal {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Pet) GetParents() map[string]*Animal {
	if x != nil {
		return x.Parents
	}
	return nil
}

var File_pet_proto protoreflect.FileDescriptor

var file_pet_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x65, 0x74,
	0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6e, 0x63,
	0x6d, 0x61, 0x67, 0x75, 0x69, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x42,
	0x6c, 0x61, 0x6e, 0x6b, 0x22, 0x60, 0x0a, 0x06, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x40,
	0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x2f,
	0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x50, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x47, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6e, 0x63, 0x6d, 0x61, 0x67, 0x75, 0x69, 0x72,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pet_proto_rawDescOnce sync.Once
	file_pet_proto_rawDescData = file_pet_proto_rawDesc
)

func file_pet_proto_rawDescGZIP() []byte {
	file_pet_proto_rawDescOnce.Do(func() {
		file_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_pet_proto_rawDescData)
	})
	return file_pet_proto_rawDescData
}

var file_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pet_proto_goTypes = []interface{}{
	(*Blank)(nil),  // 0: pet.Blank
	(*Animal)(nil), // 1: pet.Animal
	(*Pet)(nil),    // 2: pet.Pet
	nil,            // 3: pet.Pet.ParentsEntry
	(*example_proto_types.Classification)(nil), // 4: taxonomy.Classification
	(*timestamppb.Timestamp)(nil),              // 5: google.protobuf.Timestamp
}
var file_pet_proto_depIdxs = []int32{
	4, // 0: pet.Animal.classification:type_name -> taxonomy.Classification
	5, // 1: pet.Pet.date_of_birth:type_name -> google.protobuf.Timestamp
	1, // 2: pet.Pet.animal:type_name -> pet.Animal
	1, // 3: pet.Pet.friends:type_name -> pet.Animal
	3, // 4: pet.Pet.parents:type_name -> pet.Pet.ParentsEntry
	1, // 5: pet.Pet.ParentsEntry.value:type_name -> pet.Animal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pet_proto_init() }
func file_pet_proto_init() {
	if File_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blank); i {
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
		file_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Animal); i {
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
		file_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
			RawDescriptor: file_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pet_proto_goTypes,
		DependencyIndexes: file_pet_proto_depIdxs,
		MessageInfos:      file_pet_proto_msgTypes,
	}.Build()
	File_pet_proto = out.File
	file_pet_proto_rawDesc = nil
	file_pet_proto_goTypes = nil
	file_pet_proto_depIdxs = nil
}