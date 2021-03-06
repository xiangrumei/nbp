// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/policy.proto

package iam

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The type of action performed on a Binding in a policy.
type BindingDelta_Action int32

const (
	// Unspecified.
	BindingDelta_ACTION_UNSPECIFIED BindingDelta_Action = 0
	// Addition of a Binding.
	BindingDelta_ADD BindingDelta_Action = 1
	// Removal of a Binding.
	BindingDelta_REMOVE BindingDelta_Action = 2
)

var BindingDelta_Action_name = map[int32]string{
	0: "ACTION_UNSPECIFIED",
	1: "ADD",
	2: "REMOVE",
}
var BindingDelta_Action_value = map[string]int32{
	"ACTION_UNSPECIFIED": 0,
	"ADD":                1,
	"REMOVE":             2,
}

func (x BindingDelta_Action) String() string {
	return proto.EnumName(BindingDelta_Action_name, int32(x))
}
func (BindingDelta_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

// Defines an Identity and Access Management (IAM) policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list of
// `members` to a `role`, where the members can be user accounts, Google groups,
// Google domains, and service accounts. A `role` is a named list of permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//             "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// Version of the `Policy`. The default version is 0.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings" json:"bindings,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform policy updates in order to avoid race
	// conditions: An `etag` is returned in the response to `getIamPolicy`, and
	// systems are expected to put that etag in the request to `setIamPolicy` to
	// ensure that their change will be applied to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the existing
	// policy is overwritten blindly.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Policy) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *Policy) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

// Associates `members` with a `role`.
type Binding struct {
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// Specifies the identities requesting access for a Cloud Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents anyone
	//    who is authenticated with a Google account or a service account.
	//
	// * `user:{emailid}`: An email address that represents a specific Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a service
	//    account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all the
	//    users of that domain. For example, `google.com` or `example.com`.
	//
	//
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *Binding) Reset()                    { *m = Binding{} }
func (m *Binding) String() string            { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()               {}
func (*Binding) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Binding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Binding) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

// The difference delta between two policies.
type PolicyDelta struct {
	// The delta for Bindings between two policies.
	BindingDeltas []*BindingDelta `protobuf:"bytes,1,rep,name=binding_deltas,json=bindingDeltas" json:"binding_deltas,omitempty"`
}

func (m *PolicyDelta) Reset()                    { *m = PolicyDelta{} }
func (m *PolicyDelta) String() string            { return proto.CompactTextString(m) }
func (*PolicyDelta) ProtoMessage()               {}
func (*PolicyDelta) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PolicyDelta) GetBindingDeltas() []*BindingDelta {
	if m != nil {
		return m.BindingDeltas
	}
	return nil
}

// One delta entry for Binding. Each individual change (only one member in each
// entry) to a binding will be a separate entry.
type BindingDelta struct {
	// The action that was performed on a Binding.
	// Required
	Action BindingDelta_Action `protobuf:"varint,1,opt,name=action,enum=google.iam.v1.BindingDelta_Action" json:"action,omitempty"`
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	// A single identity requesting access for a Cloud Platform resource.
	// Follows the same format of Binding.members.
	// Required
	Member string `protobuf:"bytes,3,opt,name=member" json:"member,omitempty"`
}

func (m *BindingDelta) Reset()                    { *m = BindingDelta{} }
func (m *BindingDelta) String() string            { return proto.CompactTextString(m) }
func (*BindingDelta) ProtoMessage()               {}
func (*BindingDelta) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BindingDelta) GetAction() BindingDelta_Action {
	if m != nil {
		return m.Action
	}
	return BindingDelta_ACTION_UNSPECIFIED
}

func (m *BindingDelta) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *BindingDelta) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func init() {
	proto.RegisterType((*Policy)(nil), "google.iam.v1.Policy")
	proto.RegisterType((*Binding)(nil), "google.iam.v1.Binding")
	proto.RegisterType((*PolicyDelta)(nil), "google.iam.v1.PolicyDelta")
	proto.RegisterType((*BindingDelta)(nil), "google.iam.v1.BindingDelta")
	proto.RegisterEnum("google.iam.v1.BindingDelta_Action", BindingDelta_Action_name, BindingDelta_Action_value)
}

func init() { proto.RegisterFile("google/iam/v1/policy.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0xed, 0x92, 0xd2, 0xd9, 0x0f, 0x15, 0x23, 0x55, 0xd1, 0xc2, 0xa1, 0xca, 0x29, 0x27,
	0x87, 0x16, 0x21, 0x24, 0x38, 0x35, 0x4d, 0x40, 0x39, 0xb0, 0x1b, 0x0c, 0xec, 0x81, 0xcb, 0xca,
	0x69, 0x2d, 0xcb, 0x28, 0xb6, 0xa3, 0x24, 0x54, 0xe2, 0x2f, 0x21, 0xf1, 0xff, 0x38, 0xa2, 0xd8,
	0xee, 0xaa, 0x95, 0x10, 0xb7, 0x79, 0x79, 0xef, 0x65, 0xde, 0xcc, 0x18, 0xae, 0x85, 0x31, 0xa2,
	0xe6, 0x89, 0x64, 0x2a, 0xd9, 0x2f, 0x93, 0xc6, 0xd4, 0x72, 0xfb, 0x93, 0x34, 0xad, 0xe9, 0x0d,
	0xbe, 0x74, 0x1c, 0x91, 0x4c, 0x91, 0xfd, 0xf2, 0xfa, 0x85, 0x97, 0xb2, 0x46, 0x26, 0x4c, 0x6b,
	0xd3, 0xb3, 0x5e, 0x1a, 0xdd, 0x39, 0x71, 0xf4, 0x1d, 0x82, 0xd2, 0x9a, 0x71, 0x08, 0x93, 0x3d,
	0x6f, 0x3b, 0x69, 0x74, 0x88, 0x16, 0x28, 0x7e, 0x4c, 0x0f, 0x10, 0xaf, 0xe0, 0x49, 0x25, 0xf5,
	0x4e, 0x6a, 0xd1, 0x85, 0x67, 0x8b, 0x71, 0x7c, 0xbe, 0x9a, 0x93, 0x93, 0x1e, 0x24, 0x75, 0x34,
	0x7d, 0xd0, 0x61, 0x0c, 0x67, 0xbc, 0x67, 0x22, 0x1c, 0x2f, 0x50, 0x7c, 0x41, 0x6d, 0x1d, 0xbd,
	0x81, 0x89, 0x17, 0x0e, 0x74, 0x6b, 0x6a, 0x6e, 0x3b, 0x4d, 0xa9, 0xad, 0x87, 0x00, 0x8a, 0xab,
	0x8a, 0xb7, 0x5d, 0x38, 0x5a, 0x8c, 0xe3, 0x29, 0x3d, 0xc0, 0xe8, 0x13, 0x9c, 0xbb, 0x90, 0x19,
	0xaf, 0x7b, 0x86, 0x53, 0xb8, 0xf2, 0x7d, 0xee, 0x77, 0xc3, 0x87, 0x2e, 0x44, 0x36, 0xd5, 0xf3,
	0x7f, 0xa7, 0xb2, 0x26, 0x7a, 0x59, 0x1d, 0xa1, 0x2e, 0xfa, 0x8d, 0xe0, 0xe2, 0x98, 0xc7, 0x6f,
	0x21, 0x60, 0xdb, 0xfe, 0x30, 0xfd, 0xd5, 0x2a, 0xfa, 0xcf, 0xcf, 0xc8, 0xda, 0x2a, 0xa9, 0x77,
	0x3c, 0x4c, 0x33, 0x3a, 0x9a, 0x66, 0x0e, 0x81, 0x8b, 0x6f, 0x57, 0x30, 0xa5, 0x1e, 0x45, 0xaf,
	0x21, 0x70, 0x6e, 0x3c, 0x07, 0xbc, 0xde, 0x7c, 0x29, 0x6e, 0x6f, 0xee, 0xbf, 0xde, 0x7c, 0x2e,
	0xf3, 0x4d, 0xf1, 0xbe, 0xc8, 0xb3, 0xd9, 0x23, 0x3c, 0x81, 0xf1, 0x3a, 0xcb, 0x66, 0x08, 0x03,
	0x04, 0x34, 0xff, 0x78, 0x7b, 0x97, 0xcf, 0x46, 0xa9, 0x82, 0xa7, 0x5b, 0xa3, 0x4e, 0x33, 0xa5,
	0x7e, 0x2b, 0xe5, 0x70, 0xc9, 0x12, 0x7d, 0x7b, 0xe9, 0x59, 0x61, 0x6a, 0xa6, 0x05, 0x31, 0xad,
	0x48, 0x04, 0xd7, 0xf6, 0xce, 0x89, 0xa3, 0x58, 0x23, 0x3b, 0xff, 0x66, 0xde, 0x49, 0xa6, 0xfe,
	0x20, 0xf4, 0x6b, 0xf4, 0xec, 0x83, 0x73, 0x6d, 0x6a, 0xf3, 0x63, 0x47, 0x0a, 0xa6, 0xc8, 0xdd,
	0xb2, 0x0a, 0xac, 0xeb, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x4a, 0x85, 0x10, 0x68,
	0x02, 0x00, 0x00,
}
