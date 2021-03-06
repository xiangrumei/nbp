// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/runtimeconfig/v1beta1/resources.proto

/*
Package runtimeconfig is a generated protocol buffer package.

It is generated from these files:
	google/cloud/runtimeconfig/v1beta1/resources.proto
	google/cloud/runtimeconfig/v1beta1/runtimeconfig.proto

It has these top-level messages:
	RuntimeConfig
	Variable
	EndCondition
	Waiter
	ListConfigsRequest
	ListConfigsResponse
	GetConfigRequest
	CreateConfigRequest
	UpdateConfigRequest
	DeleteConfigRequest
	ListVariablesRequest
	ListVariablesResponse
	WatchVariableRequest
	GetVariableRequest
	CreateVariableRequest
	UpdateVariableRequest
	DeleteVariableRequest
	ListWaitersRequest
	ListWaitersResponse
	GetWaiterRequest
	CreateWaiterRequest
	DeleteWaiterRequest
*/
package runtimeconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The `VariableState` describes the last known state of the variable and is
// used during a `variables().watch` call to distinguish the state of the
// variable.
type VariableState int32

const (
	// Default variable state.
	VariableState_VARIABLE_STATE_UNSPECIFIED VariableState = 0
	// The variable was updated, while `variables().watch` was executing.
	VariableState_UPDATED VariableState = 1
	// The variable was deleted, while `variables().watch` was executing.
	VariableState_DELETED VariableState = 2
)

var VariableState_name = map[int32]string{
	0: "VARIABLE_STATE_UNSPECIFIED",
	1: "UPDATED",
	2: "DELETED",
}
var VariableState_value = map[string]int32{
	"VARIABLE_STATE_UNSPECIFIED": 0,
	"UPDATED":                    1,
	"DELETED":                    2,
}

func (x VariableState) String() string {
	return proto.EnumName(VariableState_name, int32(x))
}
func (VariableState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A RuntimeConfig resource is the primary resource in the Cloud RuntimeConfig
// service. A RuntimeConfig resource consists of metadata and a hierarchy of
// variables.
type RuntimeConfig struct {
	// The resource name of a runtime config. The name must have the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]
	//
	// The `[PROJECT_ID]` must be a valid project ID, and `[CONFIG_NAME]` is an
	// arbitrary name that matches RFC 1035 segment specification. The length of
	// `[CONFIG_NAME]` must be less than 64 bytes.
	//
	// You pick the RuntimeConfig resource name, but the server will validate that
	// the name adheres to this format. After you create the resource, you cannot
	// change the resource's name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// An optional description of the RuntimeConfig object.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *RuntimeConfig) Reset()                    { *m = RuntimeConfig{} }
func (m *RuntimeConfig) String() string            { return proto.CompactTextString(m) }
func (*RuntimeConfig) ProtoMessage()               {}
func (*RuntimeConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RuntimeConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuntimeConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes a single variable within a RuntimeConfig resource.
// The name denotes the hierarchical variable name. For example,
// `ports/serving_port` is a valid variable name. The variable value is an
// opaque string and only leaf variables can have values (that is, variables
// that do not have any child variables).
type Variable struct {
	// The name of the variable resource, in the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME]
	//
	// The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a
	// valid RuntimeConfig reource and `[VARIABLE_NAME]` follows Unix file system
	// file path naming.
	//
	// The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and
	// dashes. Slashes are used as path element separators and are not part of the
	// `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one
	// non-slash character. Multiple slashes are coalesced into single slash
	// character. Each path segment should follow RFC 1035 segment specification.
	// The length of a `[VARIABLE_NAME]` must be less than 256 bytes.
	//
	// Once you create a variable, you cannot change the variable name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The the value of the variable. It can be either a binary or a string
	// value. You must specify one of either `value` or `text`. Specifying both
	// will cause the server to return an error.
	//
	// Types that are valid to be assigned to Contents:
	//	*Variable_Value
	//	*Variable_Text
	Contents isVariable_Contents `protobuf_oneof:"contents"`
	// [Output Only] The time of the last variable update.
	UpdateTime *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// [Ouput only] The current state of the variable. The variable state indicates
	// the outcome of the `variables().watch` call and is visible through the
	// `get` and `list` calls.
	State VariableState `protobuf:"varint,4,opt,name=state,enum=google.cloud.runtimeconfig.v1beta1.VariableState" json:"state,omitempty"`
}

func (m *Variable) Reset()                    { *m = Variable{} }
func (m *Variable) String() string            { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()               {}
func (*Variable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isVariable_Contents interface {
	isVariable_Contents()
}

type Variable_Value struct {
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}
type Variable_Text struct {
	Text string `protobuf:"bytes,5,opt,name=text,oneof"`
}

func (*Variable_Value) isVariable_Contents() {}
func (*Variable_Text) isVariable_Contents()  {}

func (m *Variable) GetContents() isVariable_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variable) GetValue() []byte {
	if x, ok := m.GetContents().(*Variable_Value); ok {
		return x.Value
	}
	return nil
}

func (m *Variable) GetText() string {
	if x, ok := m.GetContents().(*Variable_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Variable) GetUpdateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Variable) GetState() VariableState {
	if m != nil {
		return m.State
	}
	return VariableState_VARIABLE_STATE_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Variable) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Variable_OneofMarshaler, _Variable_OneofUnmarshaler, _Variable_OneofSizer, []interface{}{
		(*Variable_Value)(nil),
		(*Variable_Text)(nil),
	}
}

func _Variable_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Variable)
	// contents
	switch x := m.Contents.(type) {
	case *Variable_Value:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Value)
	case *Variable_Text:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case nil:
	default:
		return fmt.Errorf("Variable.Contents has unexpected type %T", x)
	}
	return nil
}

func _Variable_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Variable)
	switch tag {
	case 2: // contents.value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Contents = &Variable_Value{x}
		return true, err
	case 5: // contents.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Contents = &Variable_Text{x}
		return true, err
	default:
		return false, nil
	}
}

func _Variable_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Variable)
	// contents
	switch x := m.Contents.(type) {
	case *Variable_Value:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Value)))
		n += len(x.Value)
	case *Variable_Text:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The condition that a Waiter resource is waiting for.
type EndCondition struct {
	// The condition oneof holds the available condition types for this
	// EndCondition. Currently, the only available type is Cardinality.
	//
	// Types that are valid to be assigned to Condition:
	//	*EndCondition_Cardinality_
	Condition isEndCondition_Condition `protobuf_oneof:"condition"`
}

func (m *EndCondition) Reset()                    { *m = EndCondition{} }
func (m *EndCondition) String() string            { return proto.CompactTextString(m) }
func (*EndCondition) ProtoMessage()               {}
func (*EndCondition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isEndCondition_Condition interface {
	isEndCondition_Condition()
}

type EndCondition_Cardinality_ struct {
	Cardinality *EndCondition_Cardinality `protobuf:"bytes,1,opt,name=cardinality,oneof"`
}

func (*EndCondition_Cardinality_) isEndCondition_Condition() {}

func (m *EndCondition) GetCondition() isEndCondition_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *EndCondition) GetCardinality() *EndCondition_Cardinality {
	if x, ok := m.GetCondition().(*EndCondition_Cardinality_); ok {
		return x.Cardinality
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EndCondition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EndCondition_OneofMarshaler, _EndCondition_OneofUnmarshaler, _EndCondition_OneofSizer, []interface{}{
		(*EndCondition_Cardinality_)(nil),
	}
}

func _EndCondition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EndCondition)
	// condition
	switch x := m.Condition.(type) {
	case *EndCondition_Cardinality_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cardinality); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EndCondition.Condition has unexpected type %T", x)
	}
	return nil
}

func _EndCondition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EndCondition)
	switch tag {
	case 1: // condition.cardinality
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EndCondition_Cardinality)
		err := b.DecodeMessage(msg)
		m.Condition = &EndCondition_Cardinality_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EndCondition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EndCondition)
	// condition
	switch x := m.Condition.(type) {
	case *EndCondition_Cardinality_:
		s := proto.Size(x.Cardinality)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A Cardinality condition for the Waiter resource. A cardinality condition is
// met when the number of variables under a specified path prefix reaches a
// predefined number. For example, if you set a Cardinality condition where
// the `path` is set to `/foo` and the number of paths is set to 2, the
// following variables would meet the condition in a RuntimeConfig resource:
//
// + `/foo/variable1 = "value1"`
// + `/foo/variable2 = "value2"`
// + `/bar/variable3 = "value3"`
//
// It would not would not satisify the same condition with the `number` set to
// 3, however, because there is only 2 paths that start with `/foo`.
// Cardinality conditions are recursive; all subtrees under the specific
// path prefix are counted.
type EndCondition_Cardinality struct {
	// The root of the variable subtree to monitor. For example, `/foo`.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The number variables under the `path` that must exist to meet this
	// condition. Defaults to 1 if not specified.
	Number int32 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
}

func (m *EndCondition_Cardinality) Reset()                    { *m = EndCondition_Cardinality{} }
func (m *EndCondition_Cardinality) String() string            { return proto.CompactTextString(m) }
func (*EndCondition_Cardinality) ProtoMessage()               {}
func (*EndCondition_Cardinality) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *EndCondition_Cardinality) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *EndCondition_Cardinality) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// A Waiter resource waits for some end condition within a RuntimeConfig resource
// to be met before it returns. For example, assume you have a distributed
// system where each node writes to a Variable resource indidicating the node's
// readiness as part of the startup process.
//
// You then configure a Waiter resource with the success condition set to wait
// until some number of nodes have checked in. Afterwards, your application
// runs some arbitrary code after the condition has been met and the waiter
// returns successfully.
//
// Once created, a Waiter resource is immutable.
//
// To learn more about using waiters, read the
// [Creating a Waiter](/deployment-manager/runtime-configurator/creating-a-waiter)
// documentation.
type Waiter struct {
	// The name of the Waiter resource, in the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME]
	//
	// The `[PROJECT_ID]` must be a valid Google Cloud project ID,
	// the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the
	// `[WAITER_NAME]` must match RFC 1035 segment specification, and the length
	// of `[WAITER_NAME]` must be less than 64 bytes.
	//
	// After you create a Waiter resource, you cannot change the resource name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// [Required] Specifies the timeout of the waiter in seconds, beginning from
	// the instant that `waiters().create` method is called. If this time elapses
	// before the success or failure conditions are met, the waiter fails and sets
	// the `error` code to `DEADLINE_EXCEEDED`.
	Timeout *google_protobuf1.Duration `protobuf:"bytes,2,opt,name=timeout" json:"timeout,omitempty"`
	// [Optional] The failure condition of this waiter. If this condition is met,
	// `done` will be set to `true` and the `error` code will be set to `ABORTED`.
	// The failure condition takes precedence over the success condition. If both
	// conditions are met, a failure will be indicated. This value is optional; if
	// no failure condition is set, the only failure scenario will be a timeout.
	Failure *EndCondition `protobuf:"bytes,3,opt,name=failure" json:"failure,omitempty"`
	// [Required] The success condition. If this condition is met, `done` will be
	// set to `true` and the `error` value will remain unset. The failure condition
	// takes precedence over the success condition. If both conditions are met, a
	// failure will be indicated.
	Success *EndCondition `protobuf:"bytes,4,opt,name=success" json:"success,omitempty"`
	// [Output Only] The instant at which this Waiter resource was created. Adding
	// the value of `timeout` to this instant yields the timeout deadline for the
	// waiter.
	CreateTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// [Output Only] If the value is `false`, it means the waiter is still waiting
	// for one of its conditions to be met.
	//
	// If true, the waiter has finished. If the waiter finished due to a timeout
	// or failure, `error` will be set.
	Done bool `protobuf:"varint,6,opt,name=done" json:"done,omitempty"`
	// [Output Only] If the waiter ended due to a failure or timeout, this value
	// will be set.
	Error *google_rpc.Status `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
}

func (m *Waiter) Reset()                    { *m = Waiter{} }
func (m *Waiter) String() string            { return proto.CompactTextString(m) }
func (*Waiter) ProtoMessage()               {}
func (*Waiter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Waiter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Waiter) GetTimeout() *google_protobuf1.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Waiter) GetFailure() *EndCondition {
	if m != nil {
		return m.Failure
	}
	return nil
}

func (m *Waiter) GetSuccess() *EndCondition {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *Waiter) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Waiter) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Waiter) GetError() *google_rpc.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*RuntimeConfig)(nil), "google.cloud.runtimeconfig.v1beta1.RuntimeConfig")
	proto.RegisterType((*Variable)(nil), "google.cloud.runtimeconfig.v1beta1.Variable")
	proto.RegisterType((*EndCondition)(nil), "google.cloud.runtimeconfig.v1beta1.EndCondition")
	proto.RegisterType((*EndCondition_Cardinality)(nil), "google.cloud.runtimeconfig.v1beta1.EndCondition.Cardinality")
	proto.RegisterType((*Waiter)(nil), "google.cloud.runtimeconfig.v1beta1.Waiter")
	proto.RegisterEnum("google.cloud.runtimeconfig.v1beta1.VariableState", VariableState_name, VariableState_value)
}

func init() { proto.RegisterFile("google/cloud/runtimeconfig/v1beta1/resources.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0x9b, 0x3e, 0x7d, 0xd9, 0x4e, 0xb6, 0x47, 0x93, 0x85, 0x46, 0xa8, 0xd0, 0xa8, 0x7a,
	0x81, 0x2a, 0x2e, 0x12, 0xda, 0x5d, 0xa1, 0x71, 0xd3, 0x97, 0xb0, 0x15, 0x4d, 0x30, 0xb9, 0xdd,
	0x90, 0xb8, 0x19, 0xae, 0xe3, 0x85, 0x48, 0xa9, 0x1d, 0x39, 0xce, 0x04, 0xdf, 0x86, 0x6b, 0x3e,
	0x01, 0x9f, 0x86, 0x2b, 0x3e, 0x08, 0xb2, 0xe3, 0x40, 0x0b, 0x13, 0x1b, 0xdc, 0xf9, 0xf8, 0xfc,
	0xcf, 0xef, 0xbc, 0xf8, 0x24, 0x30, 0x8c, 0x85, 0x88, 0x53, 0x16, 0xd0, 0x54, 0x14, 0x51, 0x20,
	0x0b, 0xae, 0x92, 0x15, 0xa3, 0x82, 0x5f, 0x25, 0x71, 0x70, 0x3d, 0x58, 0x32, 0x45, 0x06, 0x81,
	0x64, 0xb9, 0x28, 0x24, 0x65, 0xb9, 0x9f, 0x49, 0xa1, 0x04, 0xea, 0x95, 0x31, 0xbe, 0x89, 0xf1,
	0x37, 0x62, 0x7c, 0x1b, 0xd3, 0x79, 0x68, 0xb9, 0x24, 0x4b, 0x02, 0xc2, 0xb9, 0x50, 0x44, 0x25,
	0x82, 0x5b, 0x42, 0xe7, 0xc0, 0x7a, 0x8d, 0xb5, 0x2c, 0xae, 0x82, 0xa8, 0x90, 0x46, 0x60, 0xfd,
	0x8f, 0x7e, 0xf5, 0xeb, 0x0c, 0xb9, 0x22, 0xab, 0xcc, 0x0a, 0xee, 0x5b, 0x81, 0xcc, 0x68, 0x90,
	0x2b, 0xa2, 0x0a, 0x4b, 0xee, 0x85, 0xb0, 0x8b, 0xcb, 0x82, 0x26, 0xa6, 0x20, 0x84, 0xa0, 0xc1,
	0xc9, 0x8a, 0x79, 0x4e, 0xd7, 0xe9, 0x6f, 0x63, 0x73, 0x46, 0x5d, 0x70, 0x23, 0x96, 0x53, 0x99,
	0x64, 0x3a, 0xa7, 0x57, 0x37, 0xae, 0xf5, 0xab, 0xde, 0x57, 0x07, 0xb6, 0x2e, 0x88, 0x4c, 0xc8,
	0x32, 0x65, 0x37, 0x22, 0xf6, 0xa1, 0x79, 0x4d, 0xd2, 0x82, 0x99, 0xe0, 0x9d, 0x93, 0x1a, 0x2e,
	0x4d, 0x74, 0x0f, 0x1a, 0x8a, 0x7d, 0x50, 0x5e, 0x53, 0x6b, 0x4f, 0x6a, 0xd8, 0x58, 0xe8, 0x08,
	0xdc, 0x22, 0x8b, 0x88, 0x62, 0x97, 0xba, 0x32, 0xef, 0xbf, 0xae, 0xd3, 0x77, 0x87, 0x1d, 0xdf,
	0xce, 0xb1, 0xea, 0xd2, 0x5f, 0x54, 0x5d, 0x62, 0x28, 0xe5, 0xfa, 0x02, 0x1d, 0x43, 0x53, 0xb7,
	0xc8, 0xbc, 0x46, 0xd7, 0xe9, 0xff, 0x3f, 0x1c, 0xf8, 0xb7, 0x8f, 0xdf, 0xaf, 0x6a, 0x9f, 0xeb,
	0x40, 0x5c, 0xc6, 0x8f, 0x01, 0xb6, 0xa8, 0xe0, 0x8a, 0x71, 0x95, 0xf7, 0xbe, 0x38, 0xb0, 0x13,
	0xf2, 0x68, 0x22, 0x78, 0x94, 0xe8, 0x8e, 0xd1, 0x3b, 0x70, 0x29, 0x91, 0x51, 0xc2, 0x49, 0x9a,
	0xa8, 0x8f, 0xa6, 0x57, 0x77, 0xf8, 0xfc, 0x2e, 0xb9, 0xd6, 0x31, 0xfe, 0xe4, 0x27, 0xe3, 0xa4,
	0x86, 0xd7, 0x91, 0x9d, 0x67, 0xe0, 0xae, 0x79, 0xf5, 0x54, 0x33, 0xa2, 0xde, 0x57, 0x53, 0xd5,
	0x67, 0xb4, 0x0f, 0x2d, 0x5e, 0xac, 0x96, 0x4c, 0x9a, 0xb1, 0x36, 0xb1, 0xb5, 0xc6, 0x2e, 0x6c,
	0xd3, 0x2a, 0x45, 0xef, 0x5b, 0x1d, 0x5a, 0x6f, 0x48, 0xa2, 0x98, 0xbc, 0xf1, 0x65, 0x0e, 0xa1,
	0xad, 0x8b, 0x14, 0x85, 0x32, 0x10, 0x77, 0xf8, 0xe0, 0xb7, 0x39, 0x4f, 0xed, 0xb6, 0xe1, 0x4a,
	0x89, 0x5e, 0x42, 0xfb, 0x8a, 0x24, 0x69, 0x21, 0xab, 0xc7, 0x79, 0xfa, 0xb7, 0x9d, 0xe3, 0x0a,
	0xa0, 0x59, 0x79, 0x41, 0x29, 0xcb, 0x73, 0xf3, 0x62, 0xff, 0xc4, 0xb2, 0x00, 0xbd, 0x38, 0x54,
	0xb2, 0x1f, 0x8b, 0xd3, 0xbc, 0x7d, 0x71, 0x4a, 0xb9, 0x59, 0x1c, 0x04, 0x8d, 0x48, 0x70, 0xe6,
	0xb5, 0xba, 0x4e, 0x7f, 0x0b, 0x9b, 0x33, 0xea, 0x43, 0x93, 0x49, 0x29, 0xa4, 0xd7, 0x36, 0x28,
	0x54, 0xa1, 0x64, 0x46, 0xfd, 0xb9, 0xf9, 0x90, 0x70, 0x29, 0x78, 0x32, 0x83, 0xdd, 0x8d, 0x2d,
	0x42, 0x07, 0xd0, 0xb9, 0x18, 0xe1, 0xd9, 0x68, 0x7c, 0x1a, 0x5e, 0xce, 0x17, 0xa3, 0x45, 0x78,
	0x79, 0xfe, 0x6a, 0x7e, 0x16, 0x4e, 0x66, 0x2f, 0x66, 0xe1, 0x74, 0xaf, 0x86, 0x5c, 0x68, 0x9f,
	0x9f, 0x4d, 0x47, 0x8b, 0x70, 0xba, 0xe7, 0x68, 0x63, 0x1a, 0x9e, 0x86, 0xda, 0xa8, 0x8f, 0x3f,
	0x39, 0xf0, 0x98, 0x8a, 0xd5, 0x1d, 0xc6, 0x70, 0xe6, 0xbc, 0x7d, 0x6d, 0x55, 0xb1, 0x48, 0x09,
	0x8f, 0x7d, 0x21, 0xe3, 0x20, 0x66, 0xdc, 0xb4, 0x1a, 0x94, 0x2e, 0x92, 0x25, 0xf9, 0x9f, 0x7e,
	0x58, 0x47, 0x1b, 0xb7, 0x9f, 0xeb, 0xbd, 0xe3, 0x92, 0x38, 0x31, 0x79, 0x37, 0x7e, 0x0f, 0xfe,
	0xc5, 0x60, 0xac, 0x43, 0x96, 0x2d, 0x93, 0xe0, 0xf0, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61,
	0xe4, 0x09, 0x63, 0x10, 0x05, 0x00, 0x00,
}
