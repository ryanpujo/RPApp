// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: product.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreatedProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     int64                  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       string                 `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string                 `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Stock       uint32                 `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	CategoryId  int64                  `protobuf:"varint,8,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreatedProduct) Reset() {
	*x = CreatedProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedProduct) ProtoMessage() {}

func (x *CreatedProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedProduct.ProtoReflect.Descriptor instead.
func (*CreatedProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedProduct) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreatedProduct) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *CreatedProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatedProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatedProduct) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreatedProduct) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreatedProduct) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreatedProduct) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreatedProduct) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ProductPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     int64  `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Stock       uint32 `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	CategoryId  int64  `protobuf:"varint,8,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
}

func (x *ProductPayload) Reset() {
	*x = ProductPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPayload) ProtoMessage() {}

func (x *ProductPayload) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPayload.ProtoReflect.Descriptor instead.
func (*ProductPayload) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductPayload) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductPayload) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ProductPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductPayload) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductPayload) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductPayload) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *ProductPayload) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductPayload) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreName   string                 `protobuf:"bytes,2,opt,name=storeName,proto3" json:"storeName,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       string                 `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string                 `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Stock       uint32                 `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	Category    string                 `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ProductID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductID) Reset() {
	*x = ProductID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductID) ProtoMessage() {}

func (x *ProductID) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductID.ProtoReflect.Descriptor instead.
func (*ProductID) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *Products) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x8b, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x38, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2c, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xc4, 0x02, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x40, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_proto_goTypes = []interface{}{
	(*CreatedProduct)(nil),        // 0: product.CreatedProduct
	(*ProductPayload)(nil),        // 1: product.ProductPayload
	(*Product)(nil),               // 2: product.Product
	(*ProductID)(nil),             // 3: product.ProductID
	(*Products)(nil),              // 4: product.Products
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_product_proto_depIdxs = []int32{
	5, // 0: product.CreatedProduct.createdAt:type_name -> google.protobuf.Timestamp
	5, // 1: product.Product.createdAt:type_name -> google.protobuf.Timestamp
	2, // 2: product.Products.products:type_name -> product.Product
	1, // 3: product.ProductService.CreateProduct:input_type -> product.ProductPayload
	3, // 4: product.ProductService.GetProductById:input_type -> product.ProductID
	6, // 5: product.ProductService.GetProducts:input_type -> google.protobuf.Empty
	3, // 6: product.ProductService.DeleteProduct:input_type -> product.ProductID
	1, // 7: product.ProductService.UpdateProduct:input_type -> product.ProductPayload
	0, // 8: product.ProductService.CreateProduct:output_type -> product.CreatedProduct
	2, // 9: product.ProductService.GetProductById:output_type -> product.Product
	4, // 10: product.ProductService.GetProducts:output_type -> product.Products
	6, // 11: product.ProductService.DeleteProduct:output_type -> google.protobuf.Empty
	6, // 12: product.ProductService.UpdateProduct:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedProduct); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPayload); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductID); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
