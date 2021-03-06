// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/appengine/v1/location.proto
// DO NOT EDIT!

package google_appengine_v1 // import "google.golang.org/genproto/googleapis/appengine/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata for the given [google.cloud.location.Location][google.cloud.location.Location].
type LocationMetadata struct {
	// App Engine Standard Environment is available in the given location.
	//
	// @OutputOnly
	StandardEnvironmentAvailable bool `protobuf:"varint,2,opt,name=standard_environment_available,json=standardEnvironmentAvailable" json:"standard_environment_available,omitempty"`
	// App Engine Flexible Environment is available in the given location.
	//
	// @OutputOnly
	FlexibleEnvironmentAvailable bool `protobuf:"varint,4,opt,name=flexible_environment_available,json=flexibleEnvironmentAvailable" json:"flexible_environment_available,omitempty"`
}

func (m *LocationMetadata) Reset()                    { *m = LocationMetadata{} }
func (m *LocationMetadata) String() string            { return proto.CompactTextString(m) }
func (*LocationMetadata) ProtoMessage()               {}
func (*LocationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func init() {
	proto.RegisterType((*LocationMetadata)(nil), "google.appengine.v1.LocationMetadata")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/appengine/v1/location.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0xd0, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc0, 0x71, 0x56, 0x44, 0x64, 0x41, 0x90, 0x7a, 0x50, 0x8a, 0x88, 0x78, 0x12, 0x0f, 0x09,
	0xc5, 0xbb, 0xd0, 0xa2, 0x07, 0x41, 0xa1, 0xf8, 0x02, 0x65, 0x76, 0x77, 0x1a, 0x06, 0xa6, 0x33,
	0x21, 0x3b, 0x04, 0x7d, 0x19, 0x9f, 0x55, 0xba, 0x1f, 0x7a, 0xa9, 0xb0, 0xa7, 0x1c, 0xf2, 0xcf,
	0x8f, 0x99, 0x94, 0xcb, 0xa0, 0x1a, 0x18, 0x5d, 0x50, 0x06, 0x09, 0x4e, 0x53, 0xf0, 0x01, 0x25,
	0x26, 0x35, 0xf5, 0xfd, 0x15, 0x44, 0x6a, 0x3d, 0xc4, 0x88, 0x12, 0x48, 0xd0, 0xe7, 0x85, 0x67,
	0xad, 0xc1, 0x48, 0xc5, 0x75, 0xd9, 0xec, 0x62, 0x20, 0x7e, 0x1b, 0x97, 0x17, 0xf3, 0xd7, 0xa9,
	0x2e, 0xf9, 0x16, 0x53, 0xa6, 0x1a, 0x6b, 0x95, 0x2d, 0x05, 0x0f, 0x22, 0x6a, 0x1d, 0xdf, 0xf6,
	0xfe, 0xfc, 0x69, 0x1a, 0x65, 0x5f, 0x11, 0x3d, 0x83, 0xb1, 0x84, 0xe1, 0xe8, 0xdf, 0xdf, 0x7d,
	0x17, 0xe5, 0xf9, 0xdb, 0x30, 0xf2, 0x3b, 0x1a, 0x34, 0x60, 0x30, 0x7b, 0x2e, 0x6f, 0x5a, 0x03,
	0x69, 0x20, 0x35, 0x1b, 0x94, 0x4c, 0x49, 0x65, 0x87, 0x62, 0x1b, 0xc8, 0x40, 0x0c, 0x15, 0xe3,
	0xd5, 0xd1, 0x6d, 0x71, 0x7f, 0xfa, 0x71, 0x3d, 0x56, 0x2f, 0x7f, 0xd1, 0x72, 0x6c, 0xf6, 0xca,
	0x96, 0xf1, 0x93, 0x2a, 0xc6, 0x7f, 0x94, 0xe3, 0x5e, 0x19, 0xab, 0x43, 0xca, 0xea, 0xa1, 0xbc,
	0xac, 0x75, 0xe7, 0x0e, 0x7c, 0xe3, 0xea, 0x6c, 0x1c, 0x7c, 0xbd, 0x5f, 0x65, 0x5d, 0x54, 0x27,
	0xdd, 0x4e, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xc1, 0x4f, 0xe8, 0xb8, 0x01, 0x00,
	0x00,
}
