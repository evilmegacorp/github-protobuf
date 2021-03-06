// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/appengine/v1/operation.proto
// DO NOT EDIT!

package google_appengine_v1 // import "google.golang.org/genproto/googleapis/appengine/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata for the given [google.longrunning.Operation][google.longrunning.Operation].
type OperationMetadataV1 struct {
	// API method that initiated this operation. Example:
	// `google.appengine.v1.Versions.CreateVersion`.
	//
	// @OutputOnly
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	// Time that this operation was created.
	//
	// @OutputOnly
	InsertTime *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=insert_time,json=insertTime" json:"insert_time,omitempty"`
	// Time that this operation completed.
	//
	// @OutputOnly
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// User who requested this operation.
	//
	// @OutputOnly
	User string `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	// Name of the resource that this operation is acting on. Example:
	// `apps/myapp/services/default`.
	//
	// @OutputOnly
	Target string `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
}

func (m *OperationMetadataV1) Reset()                    { *m = OperationMetadataV1{} }
func (m *OperationMetadataV1) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadataV1) ProtoMessage()               {}
func (*OperationMetadataV1) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *OperationMetadataV1) GetInsertTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.InsertTime
	}
	return nil
}

func (m *OperationMetadataV1) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationMetadataV1)(nil), "google.appengine.v1.OperationMetadataV1")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/appengine/v1/operation.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xe9, 0xff, 0xbf, 0xae, 0x9a, 0x05, 0x0f, 0x5d, 0xd0, 0xd2, 0xd3, 0xe2, 0x69, 0x41,
	0x48, 0xa8, 0xe2, 0x69, 0x6f, 0xbd, 0x79, 0x10, 0x97, 0x45, 0xbc, 0x4a, 0xda, 0xce, 0x66, 0x03,
	0xdb, 0x4c, 0x48, 0xa6, 0x05, 0x3f, 0xa4, 0xdf, 0x49, 0x9a, 0xb4, 0xeb, 0x45, 0xd0, 0xdb, 0x4c,
	0x66, 0xde, 0xef, 0xbd, 0x09, 0x2b, 0x15, 0xa2, 0x3a, 0x02, 0x57, 0x78, 0x94, 0x46, 0x71, 0x74,
	0x4a, 0x28, 0x30, 0xd6, 0x21, 0xa1, 0x88, 0x23, 0x69, 0xb5, 0x17, 0xd2, 0x5a, 0x30, 0x4a, 0x1b,
	0x10, 0x7d, 0x21, 0xd0, 0x82, 0x93, 0xa4, 0xd1, 0xf0, 0xb0, 0x97, 0x2e, 0x47, 0xc6, 0x69, 0x89,
	0xf7, 0x45, 0xfe, 0xf4, 0x57, 0xb0, 0x16, 0x1e, 0x5c, 0xaf, 0x6b, 0xa8, 0xd1, 0xec, 0xb5, 0x12,
	0xd2, 0x18, 0xa4, 0x80, 0xf7, 0x91, 0x9f, 0x6f, 0x94, 0xa6, 0x43, 0x57, 0xf1, 0x1a, 0x5b, 0x11,
	0x71, 0x22, 0x0c, 0xaa, 0x6e, 0x2f, 0x2c, 0x7d, 0x58, 0xf0, 0x82, 0x74, 0x0b, 0x9e, 0x64, 0x6b,
	0xbf, 0xab, 0x28, 0xbe, 0xfd, 0x4c, 0xd8, 0xf2, 0x65, 0x0a, 0xfc, 0x0c, 0x24, 0x1b, 0x49, 0xf2,
	0xad, 0x48, 0xaf, 0xd9, 0xbc, 0x05, 0x3a, 0x60, 0x93, 0x25, 0xab, 0x64, 0x7d, 0xb9, 0x1b, 0xbb,
	0x74, 0xc3, 0x16, 0xda, 0x78, 0x70, 0xf4, 0x3e, 0x90, 0xb2, 0x7f, 0xab, 0x64, 0xbd, 0xb8, 0xcf,
	0xf9, 0x78, 0xcd, 0xe4, 0xcb, 0x5f, 0x27, 0x9b, 0x1d, 0x8b, 0xeb, 0xc3, 0x43, 0xfa, 0xc8, 0x2e,
	0xc0, 0x34, 0x51, 0xf9, 0xff, 0x57, 0xe5, 0x39, 0x98, 0x26, 0xc8, 0x52, 0x36, 0xeb, 0x3c, 0xb8,
	0x6c, 0x16, 0x92, 0x84, 0x7a, 0xc8, 0x47, 0xd2, 0x29, 0xa0, 0xec, 0x2c, 0xe6, 0x8b, 0x5d, 0x79,
	0xc7, 0x6e, 0x6a, 0x6c, 0xf9, 0x0f, 0x5f, 0x5e, 0x5e, 0x9d, 0xee, 0xdc, 0x0e, 0x66, 0xdb, 0xa4,
	0x9a, 0x07, 0xd7, 0x87, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x6e, 0x06, 0x0e, 0xe6, 0x01,
	0x00, 0x00,
}
