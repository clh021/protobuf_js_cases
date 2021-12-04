// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: message/window.proto

package message

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

type WindowId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChromiumId int32 `protobuf:"varint,1,opt,name=chromium_id,json=chromiumId,proto3" json:"chromium_id,omitempty"`
	GeckoId    int32 `protobuf:"varint,2,opt,name=gecko_id,json=geckoId,proto3" json:"gecko_id,omitempty"`
}

func (x *WindowId) Reset() {
	*x = WindowId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_window_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowId) ProtoMessage() {}

func (x *WindowId) ProtoReflect() protoreflect.Message {
	mi := &file_message_window_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowId.ProtoReflect.Descriptor instead.
func (*WindowId) Descriptor() ([]byte, []int) {
	return file_message_window_proto_rawDescGZIP(), []int{0}
}

func (x *WindowId) GetChromiumId() int32 {
	if x != nil {
		return x.ChromiumId
	}
	return 0
}

func (x *WindowId) GetGeckoId() int32 {
	if x != nil {
		return x.GeckoId
	}
	return 0
}

type MainWindowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool        `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Windows []*WindowId `protobuf:"bytes,2,rep,name=windows,proto3" json:"windows,omitempty"`
}

func (x *MainWindowList) Reset() {
	*x = MainWindowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_window_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainWindowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainWindowList) ProtoMessage() {}

func (x *MainWindowList) ProtoReflect() protoreflect.Message {
	mi := &file_message_window_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainWindowList.ProtoReflect.Descriptor instead.
func (*MainWindowList) Descriptor() ([]byte, []int) {
	return file_message_window_proto_rawDescGZIP(), []int{1}
}

func (x *MainWindowList) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *MainWindowList) GetWindows() []*WindowId {
	if x != nil {
		return x.Windows
	}
	return nil
}

type MainWindowProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *WindowId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ok            bool      `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Title         string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Visibility    bool      `protobuf:"varint,4,opt,name=visibility,proto3" json:"visibility,omitempty"`
	X             int32     `protobuf:"varint,5,opt,name=x,proto3" json:"x,omitempty"`
	Y             int32     `protobuf:"varint,6,opt,name=y,proto3" json:"y,omitempty"`
	Width         int32     `protobuf:"varint,7,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32     `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	WindowState   string    `protobuf:"bytes,9,opt,name=window_state,json=windowState,proto3" json:"window_state,omitempty"`
	TopLevelFocus bool      `protobuf:"varint,10,opt,name=top_level_focus,json=topLevelFocus,proto3" json:"top_level_focus,omitempty"`
	Icon          string    `protobuf:"bytes,11,opt,name=icon,proto3" json:"icon,omitempty"`
	Url           string    `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
	AppId         string    `protobuf:"bytes,13,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	SkipBar       bool      `protobuf:"varint,14,opt,name=skip_bar,json=skipBar,proto3" json:"skip_bar,omitempty"`
	AlwaysOnTop   bool      `protobuf:"varint,15,opt,name=always_on_top,json=alwaysOnTop,proto3" json:"always_on_top,omitempty"`
}

func (x *MainWindowProps) Reset() {
	*x = MainWindowProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_window_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainWindowProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainWindowProps) ProtoMessage() {}

func (x *MainWindowProps) ProtoReflect() protoreflect.Message {
	mi := &file_message_window_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainWindowProps.ProtoReflect.Descriptor instead.
func (*MainWindowProps) Descriptor() ([]byte, []int) {
	return file_message_window_proto_rawDescGZIP(), []int{2}
}

func (x *MainWindowProps) GetId() *WindowId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MainWindowProps) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *MainWindowProps) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MainWindowProps) GetVisibility() bool {
	if x != nil {
		return x.Visibility
	}
	return false
}

func (x *MainWindowProps) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MainWindowProps) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *MainWindowProps) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *MainWindowProps) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *MainWindowProps) GetWindowState() string {
	if x != nil {
		return x.WindowState
	}
	return ""
}

func (x *MainWindowProps) GetTopLevelFocus() bool {
	if x != nil {
		return x.TopLevelFocus
	}
	return false
}

func (x *MainWindowProps) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MainWindowProps) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MainWindowProps) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *MainWindowProps) GetSkipBar() bool {
	if x != nil {
		return x.SkipBar
	}
	return false
}

func (x *MainWindowProps) GetAlwaysOnTop() bool {
	if x != nil {
		return x.AlwaysOnTop
	}
	return false
}

var File_message_window_proto protoreflect.FileDescriptor

var file_message_window_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6e, 0x6b, 0x73, 0x22, 0x46, 0x0a, 0x08,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x28, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6e, 0x6b, 0x73, 0x2e, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x64, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73,
	0x22, 0x88, 0x03, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6c, 0x6e, 0x6b, 0x73, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x70, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x74, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x6b, 0x69, 0x70, 0x42, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6c, 0x77, 0x61, 0x79,
	0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x4f, 0x6e, 0x54, 0x6f, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_window_proto_rawDescOnce sync.Once
	file_message_window_proto_rawDescData = file_message_window_proto_rawDesc
)

func file_message_window_proto_rawDescGZIP() []byte {
	file_message_window_proto_rawDescOnce.Do(func() {
		file_message_window_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_window_proto_rawDescData)
	})
	return file_message_window_proto_rawDescData
}

var file_message_window_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_window_proto_goTypes = []interface{}{
	(*WindowId)(nil),        // 0: lnks.WindowId
	(*MainWindowList)(nil),  // 1: lnks.MainWindowList
	(*MainWindowProps)(nil), // 2: lnks.MainWindowProps
}
var file_message_window_proto_depIdxs = []int32{
	0, // 0: lnks.MainWindowList.windows:type_name -> lnks.WindowId
	0, // 1: lnks.MainWindowProps.id:type_name -> lnks.WindowId
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_window_proto_init() }
func file_message_window_proto_init() {
	if File_message_window_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_window_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowId); i {
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
		file_message_window_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainWindowList); i {
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
		file_message_window_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainWindowProps); i {
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
			RawDescriptor: file_message_window_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_window_proto_goTypes,
		DependencyIndexes: file_message_window_proto_depIdxs,
		MessageInfos:      file_message_window_proto_msgTypes,
	}.Build()
	File_message_window_proto = out.File
	file_message_window_proto_rawDesc = nil
	file_message_window_proto_goTypes = nil
	file_message_window_proto_depIdxs = nil
}
