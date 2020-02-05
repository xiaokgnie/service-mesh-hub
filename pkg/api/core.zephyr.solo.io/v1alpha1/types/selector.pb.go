// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/selector.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// select resources by their labels
type LabelSelector struct {
	LabelsToMatch        map[string]string `protobuf:"bytes,1,rep,name=labels_to_match,json=labelsToMatch,proto3" json:"labels_to_match,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{0}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelector.Unmarshal(m, b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
}
func (m *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(m, src)
}
func (m *LabelSelector) XXX_Size() int {
	return xxx_messageInfo_LabelSelector.Size(m)
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetLabelsToMatch() map[string]string {
	if m != nil {
		return m.LabelsToMatch
	}
	return nil
}

// select resources individually
type ResourceSelector struct {
	// apply the selector to one or more services by adding their refs here
	Services             []*ResourceRef `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResourceSelector) Reset()         { *m = ResourceSelector{} }
func (m *ResourceSelector) String() string { return proto.CompactTextString(m) }
func (*ResourceSelector) ProtoMessage()    {}
func (*ResourceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{1}
}
func (m *ResourceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceSelector.Unmarshal(m, b)
}
func (m *ResourceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceSelector.Marshal(b, m, deterministic)
}
func (m *ResourceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceSelector.Merge(m, src)
}
func (m *ResourceSelector) XXX_Size() int {
	return xxx_messageInfo_ResourceSelector.Size(m)
}
func (m *ResourceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceSelector proto.InternalMessageInfo

func (m *ResourceSelector) GetServices() []*ResourceRef {
	if m != nil {
		return m.Services
	}
	return nil
}

// select all resources in these namespaces
type NamespaceSelector struct {
	// list of namespaces to search
	Namespaces           []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceSelector) Reset()         { *m = NamespaceSelector{} }
func (m *NamespaceSelector) String() string { return proto.CompactTextString(m) }
func (*NamespaceSelector) ProtoMessage()    {}
func (*NamespaceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{2}
}
func (m *NamespaceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceSelector.Unmarshal(m, b)
}
func (m *NamespaceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceSelector.Marshal(b, m, deterministic)
}
func (m *NamespaceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceSelector.Merge(m, src)
}
func (m *NamespaceSelector) XXX_Size() int {
	return xxx_messageInfo_NamespaceSelector.Size(m)
}
func (m *NamespaceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceSelector proto.InternalMessageInfo

func (m *NamespaceSelector) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

//
//specifies a method by which to select pods
//with in a mesh for the application of rules and policies
//
//Precedence:
//1. LabelSeclector
//2. NamespaceSelector
//3. ResourceSelector
type ServiceSelector struct {
	Labels               *LabelSelector     `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	Resources            *ResourceSelector  `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	Namespaces           *NamespaceSelector `protobuf:"bytes,3,opt,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServiceSelector) Reset()         { *m = ServiceSelector{} }
func (m *ServiceSelector) String() string { return proto.CompactTextString(m) }
func (*ServiceSelector) ProtoMessage()    {}
func (*ServiceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{3}
}
func (m *ServiceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSelector.Unmarshal(m, b)
}
func (m *ServiceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSelector.Marshal(b, m, deterministic)
}
func (m *ServiceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSelector.Merge(m, src)
}
func (m *ServiceSelector) XXX_Size() int {
	return xxx_messageInfo_ServiceSelector.Size(m)
}
func (m *ServiceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSelector proto.InternalMessageInfo

func (m *ServiceSelector) GetLabels() *LabelSelector {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceSelector) GetResources() *ResourceSelector {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *ServiceSelector) GetNamespaces() *NamespaceSelector {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

//
//Precedence:
//1. LabelSeclector
//2. ResourceSelector
type ClusterSelector struct {
	// specify the type of selector to use with selectorType
	// select clusters by their labels
	LabelSelector *LabelSelector `protobuf:"bytes,1,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	// select clusters by their names
	ClusterSelector      *ResourceSelector `protobuf:"bytes,4,opt,name=cluster_selector,json=clusterSelector,proto3" json:"cluster_selector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterSelector) Reset()         { *m = ClusterSelector{} }
func (m *ClusterSelector) String() string { return proto.CompactTextString(m) }
func (*ClusterSelector) ProtoMessage()    {}
func (*ClusterSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{4}
}
func (m *ClusterSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterSelector.Unmarshal(m, b)
}
func (m *ClusterSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterSelector.Marshal(b, m, deterministic)
}
func (m *ClusterSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterSelector.Merge(m, src)
}
func (m *ClusterSelector) XXX_Size() int {
	return xxx_messageInfo_ClusterSelector.Size(m)
}
func (m *ClusterSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterSelector proto.InternalMessageInfo

func (m *ClusterSelector) GetLabelSelector() *LabelSelector {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *ClusterSelector) GetClusterSelector() *ResourceSelector {
	if m != nil {
		return m.ClusterSelector
	}
	return nil
}

// global selector used to select resources from any set of clusters, namespaces, and/or labels
type Selector struct {
	ClusterSelector      *ClusterSelector `protobuf:"bytes,1,opt,name=cluster_selector,json=clusterSelector,proto3" json:"cluster_selector,omitempty"`
	ServiceSelector      *ServiceSelector `protobuf:"bytes,2,opt,name=service_selector,json=serviceSelector,proto3" json:"service_selector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ca43acbd26ca19e, []int{5}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetClusterSelector() *ClusterSelector {
	if m != nil {
		return m.ClusterSelector
	}
	return nil
}

func (m *Selector) GetServiceSelector() *ServiceSelector {
	if m != nil {
		return m.ServiceSelector
	}
	return nil
}

func init() {
	proto.RegisterType((*LabelSelector)(nil), "core.zephyr.solo.io.LabelSelector")
	proto.RegisterMapType((map[string]string)(nil), "core.zephyr.solo.io.LabelSelector.LabelsToMatchEntry")
	proto.RegisterType((*ResourceSelector)(nil), "core.zephyr.solo.io.ResourceSelector")
	proto.RegisterType((*NamespaceSelector)(nil), "core.zephyr.solo.io.NamespaceSelector")
	proto.RegisterType((*ServiceSelector)(nil), "core.zephyr.solo.io.ServiceSelector")
	proto.RegisterType((*ClusterSelector)(nil), "core.zephyr.solo.io.ClusterSelector")
	proto.RegisterType((*Selector)(nil), "core.zephyr.solo.io.Selector")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/core/v1alpha1/selector.proto", fileDescriptor_6ca43acbd26ca19e)
}

var fileDescriptor_6ca43acbd26ca19e = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x51, 0x6b, 0xd4, 0x40,
	0x14, 0x85, 0x99, 0x5d, 0x2d, 0xdd, 0x5b, 0x6a, 0xd6, 0xe8, 0xc3, 0xd2, 0x07, 0x59, 0x82, 0x4a,
	0x7d, 0x68, 0x42, 0x5b, 0x04, 0x29, 0x3e, 0xa8, 0x45, 0x41, 0xb4, 0x5a, 0xa6, 0x3e, 0x09, 0xb2,
	0xcc, 0x0e, 0xb7, 0x4d, 0xec, 0x6c, 0x67, 0x98, 0x99, 0x2c, 0xc4, 0xdf, 0x24, 0xe2, 0xef, 0xf1,
	0xd7, 0x48, 0x32, 0xd9, 0xc9, 0x66, 0x37, 0x74, 0xe9, 0x5b, 0xe6, 0x24, 0xe7, 0x9b, 0x73, 0xc2,
	0xbd, 0xf0, 0xf6, 0x2a, 0xb3, 0x69, 0x3e, 0x8d, 0xb9, 0x9c, 0x25, 0x46, 0x0a, 0x79, 0x90, 0xc9,
	0x64, 0x86, 0x26, 0x3d, 0x50, 0x5a, 0xfe, 0x44, 0x6e, 0x4d, 0xc2, 0x54, 0x96, 0x70, 0xa9, 0x31,
	0x99, 0x1f, 0x32, 0xa1, 0x52, 0x76, 0x98, 0x18, 0x14, 0xc8, 0xad, 0xd4, 0xb1, 0xd2, 0xd2, 0xca,
	0xf0, 0x51, 0xf9, 0x36, 0xfe, 0x85, 0x2a, 0x2d, 0x74, 0x5c, 0x32, 0xe2, 0x4c, 0xee, 0xbd, 0xd8,
	0x04, 0xd1, 0x78, 0xe9, 0xfc, 0xd1, 0x5f, 0x02, 0xbb, 0x9f, 0xd9, 0x14, 0xc5, 0x45, 0xcd, 0x0d,
	0x7f, 0x40, 0x20, 0x4a, 0xc1, 0x4c, 0xac, 0x9c, 0xcc, 0x98, 0xe5, 0xe9, 0x88, 0x8c, 0xfb, 0xfb,
	0x3b, 0x47, 0x2f, 0xe3, 0x8e, 0xbb, 0xe2, 0x96, 0xd9, 0x9d, 0xcc, 0x37, 0x79, 0x56, 0xfa, 0xde,
	0xdf, 0x58, 0x5d, 0xd0, 0x5d, 0xb1, 0xac, 0xed, 0xbd, 0x81, 0x70, 0xfd, 0xa3, 0x70, 0x08, 0xfd,
	0x6b, 0x2c, 0x46, 0x64, 0x4c, 0xf6, 0x07, 0xb4, 0x7c, 0x0c, 0x1f, 0xc3, 0xfd, 0x39, 0x13, 0x39,
	0x8e, 0x7a, 0x95, 0xe6, 0x0e, 0x27, 0xbd, 0x57, 0x24, 0x3a, 0x87, 0x21, 0x45, 0x23, 0x73, 0xcd,
	0xd1, 0x87, 0x7e, 0x0d, 0xdb, 0x06, 0xf5, 0x3c, 0xe3, 0x68, 0xea, 0xb4, 0xe3, 0xce, 0xb4, 0x0b,
	0x23, 0xc5, 0x4b, 0xea, 0x1d, 0xd1, 0x31, 0x3c, 0xfc, 0xc2, 0x66, 0x68, 0x14, 0x5b, 0x42, 0x3e,
	0x01, 0xb8, 0x59, 0x88, 0x0e, 0x3a, 0xa0, 0x4b, 0x4a, 0xf4, 0x8f, 0x40, 0x70, 0xe1, 0x08, 0xde,
	0x73, 0x02, 0x5b, 0xae, 0x6d, 0xd5, 0x64, 0xe7, 0x28, 0xda, 0xfc, 0xcb, 0x68, 0xed, 0x08, 0x4f,
	0x61, 0xa0, 0xeb, 0x74, 0x66, 0x74, 0xaf, 0xb2, 0x3f, 0xbb, 0xb5, 0x83, 0x27, 0x34, 0xbe, 0xf0,
	0x43, 0x2b, 0x74, 0xbf, 0xa2, 0x3c, 0xef, 0xa4, 0xac, 0x15, 0x6e, 0x95, 0xfb, 0x43, 0x20, 0x38,
	0x15, 0xb9, 0xb1, 0xa8, 0x7d, 0xb9, 0x8f, 0xf0, 0xa0, 0x8a, 0x3a, 0x59, 0x8c, 0xe0, 0x1d, 0x4a,
	0xba, 0x21, 0xf0, 0xa8, 0x73, 0x18, 0x72, 0x47, 0x6f, 0x60, 0x77, 0xaa, 0x1c, 0xf0, 0x76, 0xb8,
	0xe8, 0x37, 0x81, 0x6d, 0x8f, 0xff, 0xda, 0x81, 0x77, 0x59, 0x9f, 0x76, 0xe2, 0x57, 0x9a, 0xae,
	0xd1, 0x4b, 0x60, 0x3d, 0x2c, 0x0d, 0xb0, 0x77, 0x0b, 0x70, 0x65, 0x2e, 0x68, 0x60, 0xda, 0xc2,
	0xbb, 0xb3, 0xef, 0x9f, 0x36, 0xee, 0xbe, 0xba, 0xbe, 0xf2, 0xab, 0xbb, 0x72, 0x41, 0xb3, 0xc9,
	0xb6, 0x50, 0x68, 0xa6, 0x5b, 0xd5, 0x32, 0x1f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x83, 0x8a,
	0xf3, 0x9e, 0x51, 0x04, 0x00, 0x00,
}