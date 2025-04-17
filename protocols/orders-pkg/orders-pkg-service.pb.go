// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: orders-pkg-service.proto

package orders_pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderUuid     []byte                 `protobuf:"bytes,1,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrUpdateResponse) Reset() {
	*x = CreateOrUpdateResponse{}
	mi := &file_orders_pkg_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateResponse) ProtoMessage() {}

func (x *CreateOrUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateResponse) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateResponse) GetOrderUuid() []byte {
	if x != nil {
		return x.OrderUuid
	}
	return nil
}

type OrderDeleteReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderUuid     []byte                 `protobuf:"bytes,1,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderDeleteReq) Reset() {
	*x = OrderDeleteReq{}
	mi := &file_orders_pkg_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDeleteReq) ProtoMessage() {}

func (x *OrderDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDeleteReq.ProtoReflect.Descriptor instead.
func (*OrderDeleteReq) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{1}
}

func (x *OrderDeleteReq) GetOrderUuid() []byte {
	if x != nil {
		return x.OrderUuid
	}
	return nil
}

type OrdersEmpty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrdersEmpty) Reset() {
	*x = OrdersEmpty{}
	mi := &file_orders_pkg_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrdersEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersEmpty) ProtoMessage() {}

func (x *OrdersEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersEmpty.ProtoReflect.Descriptor instead.
func (*OrdersEmpty) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{2}
}

type ByOrderUuidReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderUuid     []byte                 `protobuf:"bytes,1,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ByOrderUuidReq) Reset() {
	*x = ByOrderUuidReq{}
	mi := &file_orders_pkg_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ByOrderUuidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByOrderUuidReq) ProtoMessage() {}

func (x *ByOrderUuidReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByOrderUuidReq.ProtoReflect.Descriptor instead.
func (*ByOrderUuidReq) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{3}
}

func (x *ByOrderUuidReq) GetOrderUuid() []byte {
	if x != nil {
		return x.OrderUuid
	}
	return nil
}

type ByUserUuidReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserUuid      []byte                 `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ByUserUuidReq) Reset() {
	*x = ByUserUuidReq{}
	mi := &file_orders_pkg_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ByUserUuidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByUserUuidReq) ProtoMessage() {}

func (x *ByUserUuidReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByUserUuidReq.ProtoReflect.Descriptor instead.
func (*ByUserUuidReq) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{4}
}

func (x *ByUserUuidReq) GetUserUuid() []byte {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

type ByMealsUuidReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MealsUuid     []byte                 `protobuf:"bytes,1,opt,name=meals_uuid,json=mealsUuid,proto3" json:"meals_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ByMealsUuidReq) Reset() {
	*x = ByMealsUuidReq{}
	mi := &file_orders_pkg_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ByMealsUuidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByMealsUuidReq) ProtoMessage() {}

func (x *ByMealsUuidReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_pkg_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByMealsUuidReq.ProtoReflect.Descriptor instead.
func (*ByMealsUuidReq) Descriptor() ([]byte, []int) {
	return file_orders_pkg_service_proto_rawDescGZIP(), []int{5}
}

func (x *ByMealsUuidReq) GetMealsUuid() []byte {
	if x != nil {
		return x.MealsUuid
	}
	return nil
}

var File_orders_pkg_service_proto protoreflect.FileDescriptor

const file_orders_pkg_service_proto_rawDesc = "" +
	"\n" +
	"\x18orders-pkg-service.proto\x12\x06models\x1a\x17orders-pkg-models.proto\"7\n" +
	"\x16CreateOrUpdateResponse\x12\x1d\n" +
	"\n" +
	"order_uuid\x18\x01 \x01(\fR\torderUuid\"/\n" +
	"\x0eOrderDeleteReq\x12\x1d\n" +
	"\n" +
	"order_uuid\x18\x01 \x01(\fR\torderUuid\"\r\n" +
	"\vOrdersEmpty\"/\n" +
	"\x0eByOrderUuidReq\x12\x1d\n" +
	"\n" +
	"order_uuid\x18\x01 \x01(\fR\torderUuid\",\n" +
	"\rByUserUuidReq\x12\x1b\n" +
	"\tuser_uuid\x18\x01 \x01(\fR\buserUuid\"/\n" +
	"\x0eByMealsUuidReq\x12\x1d\n" +
	"\n" +
	"meals_uuid\x18\x01 \x01(\fR\tmealsUuid2\xff\x02\n" +
	"\rOrdersService\x12D\n" +
	"\x13CreateOrUpdateOrder\x12\r.models.Order\x1a\x1e.models.CreateOrUpdateResponse\x120\n" +
	"\tGetOrders\x12\x13.models.OrdersEmpty\x1a\x0e.models.Orders\x129\n" +
	"\x10OrdersByUserUuid\x12\x15.models.ByUserUuidReq\x1a\x0e.models.Orders\x129\n" +
	"\x10OrderByOrderUuid\x12\x16.models.ByOrderUuidReq\x1a\r.models.Order\x12D\n" +
	"\x15DeleteMealsFromParams\x12\x16.models.ByMealsUuidReq\x1a\x13.models.OrdersEmpty\x12:\n" +
	"\vDeleteOrder\x12\x16.models.OrderDeleteReq\x1a\x13.models.OrdersEmptyB\x16Z\x14protocols/orders-pkgb\x06proto3"

var (
	file_orders_pkg_service_proto_rawDescOnce sync.Once
	file_orders_pkg_service_proto_rawDescData []byte
)

func file_orders_pkg_service_proto_rawDescGZIP() []byte {
	file_orders_pkg_service_proto_rawDescOnce.Do(func() {
		file_orders_pkg_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orders_pkg_service_proto_rawDesc), len(file_orders_pkg_service_proto_rawDesc)))
	})
	return file_orders_pkg_service_proto_rawDescData
}

var file_orders_pkg_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orders_pkg_service_proto_goTypes = []any{
	(*CreateOrUpdateResponse)(nil), // 0: models.CreateOrUpdateResponse
	(*OrderDeleteReq)(nil),         // 1: models.OrderDeleteReq
	(*OrdersEmpty)(nil),            // 2: models.OrdersEmpty
	(*ByOrderUuidReq)(nil),         // 3: models.ByOrderUuidReq
	(*ByUserUuidReq)(nil),          // 4: models.ByUserUuidReq
	(*ByMealsUuidReq)(nil),         // 5: models.ByMealsUuidReq
	(*Order)(nil),                  // 6: models.Order
	(*Orders)(nil),                 // 7: models.Orders
}
var file_orders_pkg_service_proto_depIdxs = []int32{
	6, // 0: models.OrdersService.CreateOrUpdateOrder:input_type -> models.Order
	2, // 1: models.OrdersService.GetOrders:input_type -> models.OrdersEmpty
	4, // 2: models.OrdersService.OrdersByUserUuid:input_type -> models.ByUserUuidReq
	3, // 3: models.OrdersService.OrderByOrderUuid:input_type -> models.ByOrderUuidReq
	5, // 4: models.OrdersService.DeleteMealsFromParams:input_type -> models.ByMealsUuidReq
	1, // 5: models.OrdersService.DeleteOrder:input_type -> models.OrderDeleteReq
	0, // 6: models.OrdersService.CreateOrUpdateOrder:output_type -> models.CreateOrUpdateResponse
	7, // 7: models.OrdersService.GetOrders:output_type -> models.Orders
	7, // 8: models.OrdersService.OrdersByUserUuid:output_type -> models.Orders
	6, // 9: models.OrdersService.OrderByOrderUuid:output_type -> models.Order
	2, // 10: models.OrdersService.DeleteMealsFromParams:output_type -> models.OrdersEmpty
	2, // 11: models.OrdersService.DeleteOrder:output_type -> models.OrdersEmpty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_pkg_service_proto_init() }
func file_orders_pkg_service_proto_init() {
	if File_orders_pkg_service_proto != nil {
		return
	}
	file_orders_pkg_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orders_pkg_service_proto_rawDesc), len(file_orders_pkg_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_pkg_service_proto_goTypes,
		DependencyIndexes: file_orders_pkg_service_proto_depIdxs,
		MessageInfos:      file_orders_pkg_service_proto_msgTypes,
	}.Build()
	File_orders_pkg_service_proto = out.File
	file_orders_pkg_service_proto_goTypes = nil
	file_orders_pkg_service_proto_depIdxs = nil
}
