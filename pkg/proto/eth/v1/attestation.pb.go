// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: pkg/proto/eth/v1/attestation.proto

// Note: largely inspired by
// https://github.com/prysmaticlabs/prysm/tree/develop/proto/eth/v1

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregationBits string           `protobuf:"bytes,1,opt,name=aggregation_bits,proto3" json:"aggregation_bits,omitempty"`
	Signature       string           `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Data            *AttestationData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetAggregationBits() string {
	if x != nil {
		return x.AggregationBits
	}
	return ""
}

func (x *Attestation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Attestation) GetData() *AttestationData {
	if x != nil {
		return x.Data
	}
	return nil
}

type AttestationV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregationBits string             `protobuf:"bytes,1,opt,name=aggregation_bits,proto3" json:"aggregation_bits,omitempty"`
	Signature       string             `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Data            *AttestationDataV2 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AttestationV2) Reset() {
	*x = AttestationV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationV2) ProtoMessage() {}

func (x *AttestationV2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationV2.ProtoReflect.Descriptor instead.
func (*AttestationV2) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *AttestationV2) GetAggregationBits() string {
	if x != nil {
		return x.AggregationBits
	}
	return ""
}

func (x *AttestationV2) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AttestationV2) GetData() *AttestationDataV2 {
	if x != nil {
		return x.Data
	}
	return nil
}

type AttestationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot            uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Index           uint64      `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	BeaconBlockRoot string      `protobuf:"bytes,3,opt,name=beacon_block_root,proto3" json:"beacon_block_root,omitempty"`
	Source          *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target          *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AttestationData) Reset() {
	*x = AttestationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationData) ProtoMessage() {}

func (x *AttestationData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationData.ProtoReflect.Descriptor instead.
func (*AttestationData) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{2}
}

func (x *AttestationData) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *AttestationData) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *AttestationData) GetBeaconBlockRoot() string {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return ""
}

func (x *AttestationData) GetSource() *Checkpoint {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *AttestationData) GetTarget() *Checkpoint {
	if x != nil {
		return x.Target
	}
	return nil
}

type AttestationDataV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot            *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Index           *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	BeaconBlockRoot string                  `protobuf:"bytes,3,opt,name=beacon_block_root,proto3" json:"beacon_block_root,omitempty"`
	Source          *CheckpointV2           `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target          *CheckpointV2           `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AttestationDataV2) Reset() {
	*x = AttestationDataV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationDataV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationDataV2) ProtoMessage() {}

func (x *AttestationDataV2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationDataV2.ProtoReflect.Descriptor instead.
func (*AttestationDataV2) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{3}
}

func (x *AttestationDataV2) GetSlot() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Slot
	}
	return nil
}

func (x *AttestationDataV2) GetIndex() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *AttestationDataV2) GetBeaconBlockRoot() string {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return ""
}

func (x *AttestationDataV2) GetSource() *CheckpointV2 {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *AttestationDataV2) GetTarget() *CheckpointV2 {
	if x != nil {
		return x.Target
	}
	return nil
}

type AggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregatorIndex uint64       `protobuf:"varint,1,opt,name=aggregator_index,proto3" json:"aggregator_index,omitempty"`
	Aggregate       *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	SelectionProof  string       `protobuf:"bytes,2,opt,name=selection_proof,proto3" json:"selection_proof,omitempty"`
}

func (x *AggregateAttestationAndProof) Reset() {
	*x = AggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateAttestationAndProof) ProtoMessage() {}

func (x *AggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{4}
}

func (x *AggregateAttestationAndProof) GetAggregatorIndex() uint64 {
	if x != nil {
		return x.AggregatorIndex
	}
	return 0
}

func (x *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *AggregateAttestationAndProof) GetSelectionProof() string {
	if x != nil {
		return x.SelectionProof
	}
	return ""
}

type SignedAggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *AggregateAttestationAndProof `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature string                        `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedAggregateAttestationAndProof) Reset() {
	*x = SignedAggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedAggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedAggregateAttestationAndProof) ProtoMessage() {}

func (x *SignedAggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedAggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*SignedAggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{5}
}

func (x *SignedAggregateAttestationAndProof) GetMessage() *AggregateAttestationAndProof {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedAggregateAttestationAndProof) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type ElaboratedAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature        string                    `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Data             *AttestationDataV2        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ValidatorIndexes []*wrapperspb.UInt64Value `protobuf:"bytes,3,rep,name=validator_indexes,proto3" json:"validator_indexes,omitempty"`
}

func (x *ElaboratedAttestation) Reset() {
	*x = ElaboratedAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElaboratedAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElaboratedAttestation) ProtoMessage() {}

func (x *ElaboratedAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v1_attestation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElaboratedAttestation.ProtoReflect.Descriptor instead.
func (*ElaboratedAttestation) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP(), []int{6}
}

func (x *ElaboratedAttestation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ElaboratedAttestation) GetData() *AttestationDataV2 {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ElaboratedAttestation) GetValidatorIndexes() []*wrapperspb.UInt64Value {
	if x != nil {
		return x.ValidatorIndexes
	}
	return nil
}

var File_pkg_proto_eth_v1_attestation_proto protoreflect.FileDescriptor

var file_pkg_proto_eth_v1_attestation_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69,
	0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x32, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x32,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x78,
	0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x56, 0x32, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x2c, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x2f,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0x8d, 0x02, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x32, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x11,
	0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x61, 0x74,
	0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x56, 0x32, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x32, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0xac, 0x01, 0x0a, 0x1c, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x36, 0x0a,
	0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22,
	0x87, 0x01, 0x0a, 0x22, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x43, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x15, 0x45, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x56, 0x32, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4a, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x74, 0x68, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x70, 0x73, 0x2f, 0x78, 0x61, 0x74, 0x75,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_eth_v1_attestation_proto_rawDescOnce sync.Once
	file_pkg_proto_eth_v1_attestation_proto_rawDescData = file_pkg_proto_eth_v1_attestation_proto_rawDesc
)

func file_pkg_proto_eth_v1_attestation_proto_rawDescGZIP() []byte {
	file_pkg_proto_eth_v1_attestation_proto_rawDescOnce.Do(func() {
		file_pkg_proto_eth_v1_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_eth_v1_attestation_proto_rawDescData)
	})
	return file_pkg_proto_eth_v1_attestation_proto_rawDescData
}

var file_pkg_proto_eth_v1_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_eth_v1_attestation_proto_goTypes = []interface{}{
	(*Attestation)(nil),                        // 0: xatu.eth.v1.Attestation
	(*AttestationV2)(nil),                      // 1: xatu.eth.v1.AttestationV2
	(*AttestationData)(nil),                    // 2: xatu.eth.v1.AttestationData
	(*AttestationDataV2)(nil),                  // 3: xatu.eth.v1.AttestationDataV2
	(*AggregateAttestationAndProof)(nil),       // 4: xatu.eth.v1.AggregateAttestationAndProof
	(*SignedAggregateAttestationAndProof)(nil), // 5: xatu.eth.v1.SignedAggregateAttestationAndProof
	(*ElaboratedAttestation)(nil),              // 6: xatu.eth.v1.ElaboratedAttestation
	(*Checkpoint)(nil),                         // 7: xatu.eth.v1.Checkpoint
	(*wrapperspb.UInt64Value)(nil),             // 8: google.protobuf.UInt64Value
	(*CheckpointV2)(nil),                       // 9: xatu.eth.v1.CheckpointV2
}
var file_pkg_proto_eth_v1_attestation_proto_depIdxs = []int32{
	2,  // 0: xatu.eth.v1.Attestation.data:type_name -> xatu.eth.v1.AttestationData
	3,  // 1: xatu.eth.v1.AttestationV2.data:type_name -> xatu.eth.v1.AttestationDataV2
	7,  // 2: xatu.eth.v1.AttestationData.source:type_name -> xatu.eth.v1.Checkpoint
	7,  // 3: xatu.eth.v1.AttestationData.target:type_name -> xatu.eth.v1.Checkpoint
	8,  // 4: xatu.eth.v1.AttestationDataV2.slot:type_name -> google.protobuf.UInt64Value
	8,  // 5: xatu.eth.v1.AttestationDataV2.index:type_name -> google.protobuf.UInt64Value
	9,  // 6: xatu.eth.v1.AttestationDataV2.source:type_name -> xatu.eth.v1.CheckpointV2
	9,  // 7: xatu.eth.v1.AttestationDataV2.target:type_name -> xatu.eth.v1.CheckpointV2
	0,  // 8: xatu.eth.v1.AggregateAttestationAndProof.aggregate:type_name -> xatu.eth.v1.Attestation
	4,  // 9: xatu.eth.v1.SignedAggregateAttestationAndProof.message:type_name -> xatu.eth.v1.AggregateAttestationAndProof
	3,  // 10: xatu.eth.v1.ElaboratedAttestation.data:type_name -> xatu.eth.v1.AttestationDataV2
	8,  // 11: xatu.eth.v1.ElaboratedAttestation.validator_indexes:type_name -> google.protobuf.UInt64Value
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_pkg_proto_eth_v1_attestation_proto_init() }
func file_pkg_proto_eth_v1_attestation_proto_init() {
	if File_pkg_proto_eth_v1_attestation_proto != nil {
		return
	}
	file_pkg_proto_eth_v1_checkpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationV2); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationData); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDataV2); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateAttestationAndProof); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedAggregateAttestationAndProof); i {
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
		file_pkg_proto_eth_v1_attestation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElaboratedAttestation); i {
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
			RawDescriptor: file_pkg_proto_eth_v1_attestation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_eth_v1_attestation_proto_goTypes,
		DependencyIndexes: file_pkg_proto_eth_v1_attestation_proto_depIdxs,
		MessageInfos:      file_pkg_proto_eth_v1_attestation_proto_msgTypes,
	}.Build()
	File_pkg_proto_eth_v1_attestation_proto = out.File
	file_pkg_proto_eth_v1_attestation_proto_rawDesc = nil
	file_pkg_proto_eth_v1_attestation_proto_goTypes = nil
	file_pkg_proto_eth_v1_attestation_proto_depIdxs = nil
}
