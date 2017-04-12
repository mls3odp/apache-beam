// Code generated by protoc-gen-go.
// source: v1.proto
// DO NOT EDIT!

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	v1.proto

It has these top-level messages:
	Type
	FunctionRef
	CustomCoder
	Window
	Coder
	Node
	MultiEdge
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Kind is mostly identical to reflect.TypeKind, expect we handle certain
// types specially, such as "error".
type Type_Kind int32

const (
	Type_INVALID Type_Kind = 0
	// Primitive.
	Type_BOOL   Type_Kind = 1
	Type_INT    Type_Kind = 2
	Type_INT8   Type_Kind = 3
	Type_INT16  Type_Kind = 4
	Type_INT32  Type_Kind = 5
	Type_INT64  Type_Kind = 6
	Type_UINT   Type_Kind = 7
	Type_UINT8  Type_Kind = 8
	Type_UINT16 Type_Kind = 9
	Type_UINT32 Type_Kind = 10
	Type_UINT64 Type_Kind = 11
	Type_STRING Type_Kind = 12
	// Aggregate.
	Type_SLICE  Type_Kind = 20
	Type_STRUCT Type_Kind = 21
	// Special.
	Type_FUNC      Type_Kind = 30
	Type_CHAN      Type_Kind = 31
	Type_ERROR     Type_Kind = 32
	Type_UNIVERSAL Type_Kind = 33
	// Data
	Type_FNVALUE Type_Kind = 40
	Type_TYPE    Type_Kind = 41
)

var Type_Kind_name = map[int32]string{
	0:  "INVALID",
	1:  "BOOL",
	2:  "INT",
	3:  "INT8",
	4:  "INT16",
	5:  "INT32",
	6:  "INT64",
	7:  "UINT",
	8:  "UINT8",
	9:  "UINT16",
	10: "UINT32",
	11: "UINT64",
	12: "STRING",
	20: "SLICE",
	21: "STRUCT",
	30: "FUNC",
	31: "CHAN",
	32: "ERROR",
	33: "UNIVERSAL",
	40: "FNVALUE",
	41: "TYPE",
}
var Type_Kind_value = map[string]int32{
	"INVALID":   0,
	"BOOL":      1,
	"INT":       2,
	"INT8":      3,
	"INT16":     4,
	"INT32":     5,
	"INT64":     6,
	"UINT":      7,
	"UINT8":     8,
	"UINT16":    9,
	"UINT32":    10,
	"UINT64":    11,
	"STRING":    12,
	"SLICE":     20,
	"STRUCT":    21,
	"FUNC":      30,
	"CHAN":      31,
	"ERROR":     32,
	"UNIVERSAL": 33,
	"FNVALUE":   40,
	"TYPE":      41,
}

func (x Type_Kind) String() string {
	return proto.EnumName(Type_Kind_name, int32(x))
}
func (Type_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// ChanDir matches reflect.ChanDir.
type Type_ChanDir int32

const (
	Type_RECV Type_ChanDir = 0
	Type_SEND Type_ChanDir = 1
	Type_BOTH Type_ChanDir = 2
)

var Type_ChanDir_name = map[int32]string{
	0: "RECV",
	1: "SEND",
	2: "BOTH",
}
var Type_ChanDir_value = map[string]int32{
	"RECV": 0,
	"SEND": 1,
	"BOTH": 2,
}

func (x Type_ChanDir) String() string {
	return proto.EnumName(Type_ChanDir_name, int32(x))
}
func (Type_ChanDir) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Window_Kind int32

const (
	Window_GLOBAL Window_Kind = 0
)

var Window_Kind_name = map[int32]string{
	0: "GLOBAL",
}
var Window_Kind_value = map[string]int32{
	"GLOBAL": 0,
}

func (x Window_Kind) String() string {
	return proto.EnumName(Window_Kind_name, int32(x))
}
func (Window_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Coder_Kind int32

const (
	Coder_CUSTOM        Coder_Kind = 0
	Coder_VARINT        Coder_Kind = 1
	Coder_BYTES         Coder_Kind = 2
	Coder_PARI          Coder_Kind = 3
	Coder_LENGTHPREFIX  Coder_Kind = 4
	Coder_WINDOWEDVALUE Coder_Kind = 5
	Coder_STREAM        Coder_Kind = 6
)

var Coder_Kind_name = map[int32]string{
	0: "CUSTOM",
	1: "VARINT",
	2: "BYTES",
	3: "PARI",
	4: "LENGTHPREFIX",
	5: "WINDOWEDVALUE",
	6: "STREAM",
}
var Coder_Kind_value = map[string]int32{
	"CUSTOM":        0,
	"VARINT":        1,
	"BYTES":         2,
	"PARI":          3,
	"LENGTHPREFIX":  4,
	"WINDOWEDVALUE": 5,
	"STREAM":        6,
}

func (x Coder_Kind) String() string {
	return proto.EnumName(Coder_Kind_name, int32(x))
}
func (Coder_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// Type represents a serializable reflect.Type.
type Type struct {
	// (Required) Type kind.
	Kind Type_Kind `protobuf:"varint,1,opt,name=kind,enum=v1.Type_Kind" json:"kind,omitempty"`
	// (Optional) Element type (if SLICE or CHAN)
	Element *Type `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	// (Optional) Fields (if STRUCT).
	Fields []*Type_StructField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
	// (Optional) Parameter types (if FUNC).
	ParameterTypes []*Type `protobuf:"bytes,4,rep,name=parameter_types,json=parameterTypes" json:"parameter_types,omitempty"`
	// (Optional) Return types (if FUNC).
	ReturnTypes []*Type `protobuf:"bytes,5,rep,name=return_types,json=returnTypes" json:"return_types,omitempty"`
	// (Optional) Is variadic (if FUNC).
	IsVariadic bool `protobuf:"varint,6,opt,name=is_variadic,json=isVariadic" json:"is_variadic,omitempty"`
	// (Optional) Channel direction (if CHAN).
	ChanDir Type_ChanDir `protobuf:"varint,7,opt,name=chan_dir,json=chanDir,enum=v1.Type_ChanDir" json:"chan_dir,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Type) GetKind() Type_Kind {
	if m != nil {
		return m.Kind
	}
	return Type_INVALID
}

func (m *Type) GetElement() *Type {
	if m != nil {
		return m.Element
	}
	return nil
}

func (m *Type) GetFields() []*Type_StructField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetParameterTypes() []*Type {
	if m != nil {
		return m.ParameterTypes
	}
	return nil
}

func (m *Type) GetReturnTypes() []*Type {
	if m != nil {
		return m.ReturnTypes
	}
	return nil
}

func (m *Type) GetIsVariadic() bool {
	if m != nil {
		return m.IsVariadic
	}
	return false
}

func (m *Type) GetChanDir() Type_ChanDir {
	if m != nil {
		return m.ChanDir
	}
	return Type_RECV
}

// StructField matches reflect.StructField.
type Type_StructField struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PkgPath   string  `protobuf:"bytes,2,opt,name=pkg_path,json=pkgPath" json:"pkg_path,omitempty"`
	Type      *Type   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tag       string  `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Offset    int64   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Index     []int32 `protobuf:"varint,6,rep,packed,name=index" json:"index,omitempty"`
	Anonymous bool    `protobuf:"varint,7,opt,name=anonymous" json:"anonymous,omitempty"`
}

func (m *Type_StructField) Reset()                    { *m = Type_StructField{} }
func (m *Type_StructField) String() string            { return proto.CompactTextString(m) }
func (*Type_StructField) ProtoMessage()               {}
func (*Type_StructField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Type_StructField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type_StructField) GetPkgPath() string {
	if m != nil {
		return m.PkgPath
	}
	return ""
}

func (m *Type_StructField) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type_StructField) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Type_StructField) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Type_StructField) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Type_StructField) GetAnonymous() bool {
	if m != nil {
		return m.Anonymous
	}
	return false
}

// FunctionRef represents a serialized function reference. The
// implementation is notably not serialized and must be present (and
// somehow discoverable from the symbol name) on the decoding side.
type FunctionRef struct {
	// (Required) Symbol name of function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Function type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FunctionRef) Reset()                    { *m = FunctionRef{} }
func (m *FunctionRef) String() string            { return proto.CompactTextString(m) }
func (*FunctionRef) ProtoMessage()               {}
func (*FunctionRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FunctionRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionRef) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// CustomCoder
type CustomCoder struct {
	// (Required) Name of the coder. For informational purposes only.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Concrete type being coded.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Required) Encoding function.
	Enc *FunctionRef `protobuf:"bytes,3,opt,name=enc" json:"enc,omitempty"`
	// (Required) Decoding function.
	Dec *FunctionRef `protobuf:"bytes,4,opt,name=dec" json:"dec,omitempty"`
	// (Optional) JSON-serialized data.
	Data string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (m *CustomCoder) Reset()                    { *m = CustomCoder{} }
func (m *CustomCoder) String() string            { return proto.CompactTextString(m) }
func (*CustomCoder) ProtoMessage()               {}
func (*CustomCoder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomCoder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomCoder) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomCoder) GetEnc() *FunctionRef {
	if m != nil {
		return m.Enc
	}
	return nil
}

func (m *CustomCoder) GetDec() *FunctionRef {
	if m != nil {
		return m.Dec
	}
	return nil
}

func (m *CustomCoder) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Window struct {
	Kind Window_Kind `protobuf:"varint,1,opt,name=kind,enum=v1.Window_Kind" json:"kind,omitempty"`
}

func (m *Window) Reset()                    { *m = Window{} }
func (m *Window) String() string            { return proto.CompactTextString(m) }
func (*Window) ProtoMessage()               {}
func (*Window) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Window) GetKind() Window_Kind {
	if m != nil {
		return m.Kind
	}
	return Window_GLOBAL
}

// Coder represents a serialized Coder. Unused, for now.
type Coder struct {
	Kind       Coder_Kind   `protobuf:"varint,1,opt,name=kind,enum=v1.Coder_Kind" json:"kind,omitempty"`
	Components []*Coder     `protobuf:"bytes,2,rep,name=components" json:"components,omitempty"`
	Custom     *CustomCoder `protobuf:"bytes,3,opt,name=custom" json:"custom,omitempty"`
	Window     *Window      `protobuf:"bytes,4,opt,name=window" json:"window,omitempty"`
}

func (m *Coder) Reset()                    { *m = Coder{} }
func (m *Coder) String() string            { return proto.CompactTextString(m) }
func (*Coder) ProtoMessage()               {}
func (*Coder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Coder) GetKind() Coder_Kind {
	if m != nil {
		return m.Kind
	}
	return Coder_CUSTOM
}

func (m *Coder) GetComponents() []*Coder {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *Coder) GetCustom() *CustomCoder {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *Coder) GetWindow() *Window {
	if m != nil {
		return m.Window
	}
	return nil
}

type Node struct {
	Type *Type `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Node) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// MultiEdge prepresents a serialized MultiEdge. It includes the output
// nodes, but not the input nodes.
type MultiEdge struct {
	// UserFn
	UserFn *FunctionRef `protobuf:"bytes,1,opt,name=user_fn,json=userFn" json:"user_fn,omitempty"`
	// (Optional) JSON-serialized data.
	Data     string                `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Inbound  []*MultiEdge_Inbound  `protobuf:"bytes,3,rep,name=inbound" json:"inbound,omitempty"`
	Outbound []*MultiEdge_Outbound `protobuf:"bytes,4,rep,name=outbound" json:"outbound,omitempty"`
}

func (m *MultiEdge) Reset()                    { *m = MultiEdge{} }
func (m *MultiEdge) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge) ProtoMessage()               {}
func (*MultiEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MultiEdge) GetUserFn() *FunctionRef {
	if m != nil {
		return m.UserFn
	}
	return nil
}

func (m *MultiEdge) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MultiEdge) GetInbound() []*MultiEdge_Inbound {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *MultiEdge) GetOutbound() []*MultiEdge_Outbound {
	if m != nil {
		return m.Outbound
	}
	return nil
}

type MultiEdge_Inbound struct {
	Type *Type `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Inbound) Reset()                    { *m = MultiEdge_Inbound{} }
func (m *MultiEdge_Inbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Inbound) ProtoMessage()               {}
func (*MultiEdge_Inbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *MultiEdge_Inbound) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type MultiEdge_Outbound struct {
	Type *Type `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Node *Node `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *MultiEdge_Outbound) Reset()                    { *m = MultiEdge_Outbound{} }
func (m *MultiEdge_Outbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Outbound) ProtoMessage()               {}
func (*MultiEdge_Outbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *MultiEdge_Outbound) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MultiEdge_Outbound) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "v1.Type")
	proto.RegisterType((*Type_StructField)(nil), "v1.Type.StructField")
	proto.RegisterType((*FunctionRef)(nil), "v1.FunctionRef")
	proto.RegisterType((*CustomCoder)(nil), "v1.CustomCoder")
	proto.RegisterType((*Window)(nil), "v1.Window")
	proto.RegisterType((*Coder)(nil), "v1.Coder")
	proto.RegisterType((*Node)(nil), "v1.Node")
	proto.RegisterType((*MultiEdge)(nil), "v1.MultiEdge")
	proto.RegisterType((*MultiEdge_Inbound)(nil), "v1.MultiEdge.Inbound")
	proto.RegisterType((*MultiEdge_Outbound)(nil), "v1.MultiEdge.Outbound")
	proto.RegisterEnum("v1.Type_Kind", Type_Kind_name, Type_Kind_value)
	proto.RegisterEnum("v1.Type_ChanDir", Type_ChanDir_name, Type_ChanDir_value)
	proto.RegisterEnum("v1.Window_Kind", Window_Kind_name, Window_Kind_value)
	proto.RegisterEnum("v1.Coder_Kind", Coder_Kind_name, Coder_Kind_value)
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x45, 0x8a, 0xa4, 0x46, 0xb6, 0xb3, 0x5d, 0x38, 0x01, 0x6b, 0x04, 0x8d, 0xcc, 0x16,
	0xb0, 0x82, 0x14, 0x2a, 0x24, 0x07, 0x46, 0x6f, 0x85, 0x2c, 0x51, 0x36, 0x51, 0x99, 0x32, 0x56,
	0x94, 0xdc, 0x9c, 0x04, 0x86, 0x5c, 0xd9, 0x84, 0xad, 0xa5, 0xc0, 0x1f, 0xa7, 0xbe, 0xf6, 0xd6,
	0x47, 0xe8, 0xa9, 0xaf, 0xd1, 0xe7, 0xea, 0x13, 0x14, 0xbb, 0x5c, 0xfd, 0x38, 0x35, 0x02, 0xe4,
	0x36, 0x9c, 0xef, 0x9b, 0xd1, 0x7c, 0x33, 0x1f, 0x29, 0x30, 0x1f, 0xda, 0xad, 0x65, 0x9a, 0xe4,
	0x09, 0xae, 0x3c, 0xb4, 0xed, 0x3f, 0x74, 0xd0, 0xfc, 0xc7, 0x25, 0xc5, 0x47, 0xa0, 0xdd, 0xc5,
	0x2c, 0xb2, 0x94, 0x86, 0xd2, 0xdc, 0xef, 0xec, 0xb5, 0x1e, 0xda, 0x2d, 0x9e, 0x6f, 0xfd, 0x1a,
	0xb3, 0x88, 0x08, 0x08, 0xdb, 0x60, 0xd0, 0x7b, 0xba, 0xa0, 0x2c, 0xb7, 0x2a, 0x0d, 0xa5, 0x59,
	0xef, 0x98, 0x2b, 0x16, 0x59, 0x01, 0xf8, 0x47, 0xd0, 0xe7, 0x31, 0xbd, 0x8f, 0x32, 0x4b, 0x6d,
	0xa8, 0xcd, 0x7a, 0xe7, 0x60, 0xdd, 0x68, 0x9c, 0xa7, 0x45, 0x98, 0x0f, 0x38, 0x48, 0x24, 0x07,
	0xb7, 0xe1, 0xc5, 0x32, 0x48, 0x83, 0x05, 0xcd, 0x69, 0x3a, 0xcb, 0x1f, 0x97, 0x34, 0xb3, 0x34,
	0x51, 0xb6, 0xe9, 0xbc, 0xbf, 0x26, 0xf0, 0xc7, 0x0c, 0xbf, 0x83, 0xdd, 0x94, 0xe6, 0x45, 0xca,
	0x24, 0xbf, 0xfa, 0x19, 0xbf, 0x5e, 0xa2, 0x25, 0xf9, 0x0d, 0xd4, 0xe3, 0x6c, 0xf6, 0x10, 0xa4,
	0x71, 0x10, 0xc5, 0xa1, 0xa5, 0x37, 0x94, 0xa6, 0x49, 0x20, 0xce, 0xa6, 0x32, 0x83, 0xdf, 0x81,
	0x19, 0xde, 0x06, 0x6c, 0x16, 0xc5, 0xa9, 0x65, 0x08, 0xe5, 0x68, 0x3d, 0x70, 0xef, 0x36, 0x60,
	0xfd, 0x38, 0x25, 0x46, 0x58, 0x06, 0x87, 0xff, 0x28, 0x50, 0xdf, 0x52, 0x81, 0x31, 0x68, 0x2c,
	0x58, 0x50, 0xb1, 0xb2, 0x1a, 0x11, 0x31, 0xfe, 0x16, 0xcc, 0xe5, 0xdd, 0xcd, 0x6c, 0x19, 0xe4,
	0xb7, 0x62, 0x49, 0x35, 0x62, 0x2c, 0xef, 0x6e, 0xae, 0x82, 0xfc, 0x16, 0xbf, 0x06, 0x8d, 0x8f,
	0x6c, 0xa9, 0x9f, 0xed, 0x4e, 0x64, 0x31, 0x02, 0x35, 0x0f, 0x6e, 0x2c, 0x4d, 0xd4, 0xf0, 0x10,
	0xbf, 0x02, 0x3d, 0x99, 0xcf, 0x33, 0x9a, 0x5b, 0xd5, 0x86, 0xd2, 0x54, 0x89, 0x7c, 0xc2, 0x07,
	0x50, 0x8d, 0x59, 0x44, 0x7f, 0xb7, 0xf4, 0x86, 0xda, 0xac, 0x92, 0xf2, 0x01, 0xbf, 0x86, 0x5a,
	0xc0, 0x12, 0xf6, 0xb8, 0x48, 0x8a, 0x4c, 0x48, 0x31, 0xc9, 0x26, 0x61, 0xff, 0xab, 0x80, 0xc6,
	0x2f, 0x89, 0xeb, 0x60, 0xb8, 0xde, 0xb4, 0x3b, 0x74, 0xfb, 0x68, 0x07, 0x9b, 0xa0, 0x9d, 0x8d,
	0x46, 0x43, 0xa4, 0x60, 0x03, 0x54, 0xd7, 0xf3, 0x51, 0x85, 0xa7, 0x5c, 0xcf, 0xff, 0x19, 0xa9,
	0xb8, 0x06, 0x55, 0xd7, 0xf3, 0xdb, 0xa7, 0x48, 0x93, 0xe1, 0x49, 0x07, 0x55, 0x65, 0x78, 0xfa,
	0x1e, 0xe9, 0x9c, 0x3a, 0xe1, 0x45, 0x06, 0x4f, 0x4e, 0x44, 0x95, 0x89, 0x01, 0xf4, 0x49, 0x59,
	0x56, 0x5b, 0xc5, 0x27, 0x1d, 0x04, 0xab, 0xf8, 0xf4, 0x3d, 0xaa, 0xf3, 0x78, 0xec, 0x13, 0xd7,
	0x3b, 0x47, 0xbb, 0xbc, 0x74, 0x3c, 0x74, 0x7b, 0x0e, 0x3a, 0x90, 0xe9, 0x49, 0xcf, 0x47, 0x2f,
	0x79, 0xef, 0xc1, 0xc4, 0xeb, 0xa1, 0xef, 0x78, 0xd4, 0xbb, 0xe8, 0x7a, 0xe8, 0x0d, 0xa7, 0x3a,
	0x84, 0x8c, 0x08, 0x6a, 0xe0, 0x3d, 0xa8, 0x4d, 0x3c, 0x77, 0xea, 0x90, 0x71, 0x77, 0x88, 0x8e,
	0xb8, 0xa8, 0x01, 0x17, 0x35, 0x71, 0x50, 0x93, 0x17, 0xf8, 0x1f, 0xae, 0x1c, 0xf4, 0xd6, 0x3e,
	0x06, 0x43, 0xde, 0x90, 0x27, 0x89, 0xd3, 0x9b, 0x96, 0x9a, 0xc7, 0x8e, 0xd7, 0x47, 0x4a, 0xa9,
	0xde, 0xbf, 0x40, 0x15, 0xfb, 0x17, 0xa8, 0x0f, 0x0a, 0x16, 0xe6, 0x71, 0xc2, 0x08, 0x9d, 0x3f,
	0x7b, 0xd7, 0xd5, 0xf1, 0x2a, 0xcf, 0x1d, 0xcf, 0xfe, 0x4b, 0x81, 0x7a, 0xaf, 0xc8, 0xf2, 0x64,
	0xd1, 0x4b, 0x22, 0x9a, 0x7e, 0x7d, 0x07, 0x7c, 0x04, 0x2a, 0x65, 0xa1, 0xf4, 0xc6, 0x0b, 0x0e,
	0x6e, 0x4d, 0x44, 0x38, 0xc6, 0x29, 0x11, 0x0d, 0x85, 0x43, 0x9e, 0xa3, 0x44, 0x34, 0xe4, 0xbf,
	0x1b, 0x05, 0x79, 0x20, 0x0c, 0x53, 0x23, 0x22, 0xb6, 0xbb, 0xa0, 0x5f, 0xc7, 0x2c, 0x4a, 0x3e,
	0xe1, 0xef, 0x9f, 0xbc, 0xe2, 0xa2, 0x43, 0x89, 0x6c, 0xbd, 0xe4, 0x36, 0x96, 0x46, 0x01, 0xd0,
	0xcf, 0x87, 0xa3, 0xb3, 0xee, 0x10, 0xed, 0xd8, 0x7f, 0x56, 0xa0, 0x5a, 0x0a, 0xb3, 0x9f, 0xb4,
	0xd8, 0xe7, 0x2d, 0x04, 0xb0, 0xfd, 0x99, 0x78, 0x0b, 0x10, 0x26, 0x8b, 0x65, 0xc2, 0x28, 0xcb,
	0x33, 0xab, 0x22, 0xde, 0xcf, 0xda, 0x9a, 0x49, 0xb6, 0x40, 0x7c, 0x0c, 0x7a, 0x28, 0xd6, 0xb6,
	0x2d, 0x7c, 0x6b, 0x91, 0x44, 0xc2, 0xd8, 0x06, 0xfd, 0x93, 0x18, 0x55, 0xca, 0x87, 0xcd, 0xf0,
	0x44, 0x22, 0x76, 0xb4, 0x99, 0xbc, 0x37, 0x19, 0xfb, 0xa3, 0x4b, 0xb4, 0xc3, 0xe3, 0x69, 0x97,
	0x70, 0x97, 0x2a, 0xdc, 0x3f, 0x67, 0x1f, 0x7c, 0x67, 0x5c, 0xba, 0xfc, 0xaa, 0x4b, 0x5c, 0xa4,
	0x62, 0x04, 0xbb, 0x43, 0xc7, 0x3b, 0xf7, 0x2f, 0xae, 0x88, 0x33, 0x70, 0x7f, 0x43, 0x1a, 0xfe,
	0x06, 0xf6, 0xae, 0x5d, 0xaf, 0x3f, 0xba, 0x76, 0xfa, 0xa5, 0xa5, 0xaa, 0xd2, 0x99, 0x4e, 0xf7,
	0x12, 0xe9, 0xf6, 0x0f, 0xa0, 0x79, 0x49, 0xb4, 0x39, 0xa7, 0xf2, 0xac, 0x21, 0xfe, 0xae, 0x40,
	0xed, 0xb2, 0xb8, 0xcf, 0x63, 0x27, 0xba, 0xa1, 0xb8, 0x09, 0x46, 0x91, 0xd1, 0x74, 0x36, 0x67,
	0x92, 0xfe, 0xbf, 0xeb, 0xe9, 0x1c, 0x1f, 0xb0, 0xf5, 0x01, 0x2b, 0x9b, 0x03, 0xe2, 0x9f, 0xc0,
	0x88, 0xd9, 0xc7, 0xa4, 0x60, 0x91, 0xfc, 0xa6, 0xbe, 0xe4, 0xd5, 0xeb, 0xee, 0x2d, 0xb7, 0x04,
	0xc9, 0x8a, 0x85, 0x3b, 0x60, 0x26, 0x45, 0x5e, 0x56, 0x94, 0x9f, 0xd3, 0x57, 0x4f, 0x2b, 0x46,
	0x12, 0x25, 0x6b, 0xde, 0xe1, 0x31, 0x18, 0xb2, 0xcf, 0x97, 0x95, 0x1d, 0x0e, 0xc0, 0x5c, 0x95,
	0x7f, 0x99, 0xc9, 0x51, 0x96, 0x44, 0x4f, 0x0c, 0xcf, 0x37, 0x47, 0x44, 0xf6, 0xa3, 0x2e, 0xfe,
	0x83, 0x4e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa7, 0xd7, 0x4e, 0x8f, 0x06, 0x00, 0x00,
}
