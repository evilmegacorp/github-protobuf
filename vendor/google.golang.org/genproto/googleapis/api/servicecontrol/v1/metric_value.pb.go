// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/servicecontrol/v1/metric_value.proto
// DO NOT EDIT!

package google_api_servicecontrol_v1 // import "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/type/money"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a single metric value.
type MetricValue struct {
	// The labels describing the metric value.
	// See comments on [google.api.servicecontrol.v1.Operation.labels][google.api.servicecontrol.v1.Operation.labels] for
	// the overriding relationship.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The start of the time period over which this metric value's measurement
	// applies. The time period has different semantics for different metric
	// types (cumulative, delta, and gauge). See the metric definition
	// documentation in the service configuration for details.
	StartTime *google_protobuf3.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The end of the time period over which this metric value's measurement
	// applies.
	EndTime *google_protobuf3.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// The value. The type of value used in the request must
	// agree with the metric definition in the service configuration, otherwise
	// the MetricValue is rejected.
	//
	// Types that are valid to be assigned to Value:
	//	*MetricValue_BoolValue
	//	*MetricValue_Int64Value
	//	*MetricValue_DoubleValue
	//	*MetricValue_StringValue
	//	*MetricValue_DistributionValue
	Value isMetricValue_Value `protobuf_oneof:"value"`
}

func (m *MetricValue) Reset()                    { *m = MetricValue{} }
func (m *MetricValue) String() string            { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()               {}
func (*MetricValue) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isMetricValue_Value interface {
	isMetricValue_Value()
}

type MetricValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,oneof"`
}
type MetricValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,5,opt,name=int64_value,json=int64Value,oneof"`
}
type MetricValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,oneof"`
}
type MetricValue_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=string_value,json=stringValue,oneof"`
}
type MetricValue_DistributionValue struct {
	DistributionValue *Distribution `protobuf:"bytes,8,opt,name=distribution_value,json=distributionValue,oneof"`
}

func (*MetricValue_BoolValue) isMetricValue_Value()         {}
func (*MetricValue_Int64Value) isMetricValue_Value()        {}
func (*MetricValue_DoubleValue) isMetricValue_Value()       {}
func (*MetricValue_StringValue) isMetricValue_Value()       {}
func (*MetricValue_DistributionValue) isMetricValue_Value() {}

func (m *MetricValue) GetValue() isMetricValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricValue) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricValue) GetStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *MetricValue) GetEndTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *MetricValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*MetricValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *MetricValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*MetricValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *MetricValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*MetricValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *MetricValue) GetStringValue() string {
	if x, ok := m.GetValue().(*MetricValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *MetricValue) GetDistributionValue() *Distribution {
	if x, ok := m.GetValue().(*MetricValue_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricValue_OneofMarshaler, _MetricValue_OneofUnmarshaler, _MetricValue_OneofSizer, []interface{}{
		(*MetricValue_BoolValue)(nil),
		(*MetricValue_Int64Value)(nil),
		(*MetricValue_DoubleValue)(nil),
		(*MetricValue_StringValue)(nil),
		(*MetricValue_DistributionValue)(nil),
	}
}

func _MetricValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *MetricValue_Int64Value:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *MetricValue_DoubleValue:
		b.EncodeVarint(6<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *MetricValue_StringValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *MetricValue_DistributionValue:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DistributionValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MetricValue.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricValue)
	switch tag {
	case 4: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricValue_BoolValue{x != 0}
		return true, err
	case 5: // value.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricValue_Int64Value{int64(x)}
		return true, err
	case 6: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 7: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &MetricValue_StringValue{x}
		return true, err
	case 8: // value.distribution_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution)
		err := b.DecodeMessage(msg)
		m.Value = &MetricValue_DistributionValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MetricValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_BoolValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *MetricValue_Int64Value:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *MetricValue_DoubleValue:
		n += proto.SizeVarint(6<<3 | proto.WireFixed64)
		n += 8
	case *MetricValue_StringValue:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *MetricValue_DistributionValue:
		s := proto.Size(x.DistributionValue)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Represents a set of metric values in the same metric.
// Each metric value in the set should have a unique combination of start time,
// end time, and label values.
type MetricValueSet struct {
	// The metric name defined in the service configuration.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	// The values in this metric.
	MetricValues []*MetricValue `protobuf:"bytes,2,rep,name=metric_values,json=metricValues" json:"metric_values,omitempty"`
}

func (m *MetricValueSet) Reset()                    { *m = MetricValueSet{} }
func (m *MetricValueSet) String() string            { return proto.CompactTextString(m) }
func (*MetricValueSet) ProtoMessage()               {}
func (*MetricValueSet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MetricValueSet) GetMetricValues() []*MetricValue {
	if m != nil {
		return m.MetricValues
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricValue)(nil), "google.api.servicecontrol.v1.MetricValue")
	proto.RegisterType((*MetricValueSet)(nil), "google.api.servicecontrol.v1.MetricValueSet")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/servicecontrol/v1/metric_value.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6f, 0xd3, 0x30,
	0x1c, 0xad, 0x1b, 0xd6, 0x8f, 0x5f, 0x06, 0x82, 0xc0, 0x21, 0xaa, 0x90, 0x16, 0xc6, 0x25, 0x70,
	0xb0, 0xb5, 0xc1, 0x10, 0x83, 0x9d, 0x2a, 0x90, 0x8a, 0xc4, 0xaa, 0x29, 0x20, 0x2e, 0x1c, 0x26,
	0x27, 0xf5, 0x82, 0x45, 0x62, 0x47, 0xb1, 0x53, 0xa9, 0x47, 0xfe, 0x6b, 0x2e, 0x48, 0xc8, 0x1f,
	0x65, 0xd9, 0xa5, 0x4c, 0x70, 0x89, 0xe2, 0xe7, 0xf7, 0xde, 0xef, 0xd3, 0xb0, 0x2c, 0xa5, 0x2c,
	0x2b, 0x86, 0x4b, 0x59, 0x51, 0x51, 0x62, 0xd9, 0x96, 0xa4, 0x64, 0xa2, 0x69, 0xa5, 0x96, 0xc4,
	0x5d, 0xd1, 0x86, 0x2b, 0x42, 0x1b, 0x4e, 0x14, 0x6b, 0xd7, 0xbc, 0x60, 0x85, 0x14, 0xba, 0x95,
	0x15, 0x59, 0x1f, 0x91, 0x9a, 0xe9, 0x96, 0x17, 0x97, 0x6b, 0x5a, 0x75, 0x0c, 0x5b, 0x4d, 0xf4,
	0xd8, 0xfb, 0xd1, 0x86, 0xe3, 0x9b, 0x02, 0xbc, 0x3e, 0x9a, 0x7d, 0xf8, 0x97, 0x68, 0x57, 0xbc,
	0x24, 0x54, 0x08, 0xa9, 0xa9, 0xe6, 0x52, 0x28, 0x17, 0x68, 0xf6, 0x5f, 0x89, 0xaf, 0xb8, 0xd2,
	0x2d, 0xcf, 0x3b, 0x63, 0xe8, 0xfd, 0xde, 0x96, 0x5c, 0x7f, 0xeb, 0x72, 0x5c, 0xc8, 0x9a, 0x38,
	0x4f, 0x62, 0x2f, 0xf2, 0xee, 0x8a, 0x34, 0x7a, 0xd3, 0x30, 0x45, 0x34, 0xaf, 0x99, 0xd2, 0xb4,
	0x6e, 0xae, 0xff, 0xbc, 0xf8, 0xec, 0x76, 0xc9, 0x18, 0x1b, 0x52, 0x4b, 0xc1, 0x36, 0xee, 0xeb,
	0xd4, 0x87, 0xbf, 0x02, 0x08, 0xcf, 0x6d, 0x2b, 0xbf, 0x98, 0x4e, 0x46, 0xe7, 0x30, 0xaa, 0x68,
	0xce, 0x2a, 0x15, 0xa3, 0x24, 0x48, 0xc3, 0xe3, 0x13, 0xbc, 0xab, 0xa9, 0xb8, 0x27, 0xc5, 0x1f,
	0xad, 0xee, 0xbd, 0xd0, 0xed, 0x26, 0xf3, 0x26, 0xd1, 0x29, 0x80, 0xd2, 0xb4, 0xd5, 0x97, 0x26,
	0xeb, 0x78, 0x98, 0xa0, 0x34, 0x3c, 0x9e, 0x6d, 0x2d, 0xb7, 0x35, 0xe2, 0xcf, 0xdb, 0x92, 0xb2,
	0xa9, 0x65, 0x9b, 0x73, 0x74, 0x02, 0x13, 0x26, 0x56, 0x4e, 0x18, 0xfc, 0x55, 0x38, 0x66, 0x62,
	0x65, 0x65, 0x07, 0x00, 0xb9, 0x94, 0x95, 0x5b, 0x8c, 0xf8, 0x4e, 0x82, 0xd2, 0xc9, 0x62, 0x90,
	0x4d, 0x0d, 0xe6, 0x2a, 0x7c, 0x02, 0x21, 0x17, 0xfa, 0xd5, 0x4b, 0xcf, 0xd8, 0x4b, 0x50, 0x1a,
	0x2c, 0x06, 0x19, 0x58, 0xd0, 0x51, 0x9e, 0xc2, 0xfe, 0x4a, 0x76, 0x79, 0xc5, 0x3c, 0x67, 0x94,
	0xa0, 0x14, 0x2d, 0x06, 0x59, 0xe8, 0xd0, 0x3f, 0x24, 0x33, 0x48, 0x51, 0x7a, 0xd2, 0x38, 0x41,
	0xe9, 0xd4, 0x90, 0x1c, 0xea, 0x48, 0x5f, 0x21, 0xea, 0xcf, 0xdb, 0x53, 0x27, 0xb6, 0x9c, 0xe7,
	0xbb, 0x5b, 0xfb, 0xae, 0xa7, 0x5b, 0x0c, 0xb2, 0x07, 0x7d, 0x1f, 0x6b, 0x3e, 0x3b, 0x85, 0xb0,
	0xd7, 0xf3, 0xe8, 0x3e, 0x04, 0xdf, 0xd9, 0x26, 0x46, 0x26, 0x8f, 0xcc, 0xfc, 0x46, 0x8f, 0x60,
	0xcf, 0x05, 0x1c, 0x5a, 0xcc, 0x1d, 0xde, 0x0c, 0x5f, 0xa3, 0xf9, 0xd8, 0xdf, 0x1c, 0xfe, 0x40,
	0x70, 0xaf, 0x37, 0xc4, 0x4f, 0x4c, 0x47, 0x07, 0x10, 0xfa, 0xc7, 0x25, 0x68, 0xcd, 0xbc, 0x1f,
	0x38, 0x68, 0x49, 0x6b, 0x16, 0x2d, 0xe1, 0x6e, 0xff, 0xf5, 0xa9, 0x78, 0x68, 0x57, 0xe5, 0xd9,
	0xad, 0x57, 0x25, 0xdb, 0xaf, 0xaf, 0x0f, 0x6a, 0x7e, 0x06, 0x49, 0x21, 0xeb, 0x9d, 0xea, 0xf9,
	0xc3, 0x9b, 0x49, 0x5e, 0x98, 0x0d, 0xb8, 0x40, 0x3f, 0x11, 0xca, 0x47, 0x76, 0x1b, 0x5e, 0xfc,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xde, 0xbe, 0xad, 0x24, 0x4e, 0x04, 0x00, 0x00,
}