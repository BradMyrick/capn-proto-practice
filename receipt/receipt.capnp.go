// Code generated by capnpc-go. DO NOT EDIT.

package receipt

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Receipt struct{ capnp.Struct }

// Receipt_TypeID is the unique identifier for the type Receipt.
const Receipt_TypeID = 0xa43270a6b4fa26e1

func NewReceipt(s *capnp.Segment) (Receipt, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Receipt{st}, err
}

func NewRootReceipt(s *capnp.Segment) (Receipt, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Receipt{st}, err
}

func ReadRootReceipt(msg *capnp.Message) (Receipt, error) {
	root, err := msg.RootPtr()
	return Receipt{root.Struct()}, err
}

func (s Receipt) String() string {
	str, _ := text.Marshal(0xa43270a6b4fa26e1, s.Struct)
	return str
}

func (s Receipt) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Receipt) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Receipt) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Receipt) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Receipt) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Receipt) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Receipt) HasSignature() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Receipt) SetSignature(v []byte) error {
	return s.Struct.SetData(1, v)
}

// Receipt_List is a list of Receipt.
type Receipt_List struct{ capnp.List }

// NewReceipt creates a new list of Receipt.
func NewReceipt_List(s *capnp.Segment, sz int32) (Receipt_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Receipt_List{l}, err
}

func (s Receipt_List) At(i int) Receipt { return Receipt{s.List.Struct(i)} }

func (s Receipt_List) Set(i int, v Receipt) error { return s.List.SetStruct(i, v.Struct) }

func (s Receipt_List) String() string {
	str, _ := text.MarshalList(0xa43270a6b4fa26e1, s.List)
	return str
}

// Receipt_Promise is a wrapper for a Receipt promised by a client call.
type Receipt_Promise struct{ *capnp.Pipeline }

func (p Receipt_Promise) Struct() (Receipt, error) {
	s, err := p.Pipeline.Struct()
	return Receipt{s}, err
}

const schema_dd30180cd0e36f1f = "x\xda4\xc81J\xc6@\x14\xc4\xf1\x99\xb7\x89\xf2A" +
	"\x82.D\x11A\x82\x8d\x85\x85\x8a\xa5U@-\x84\x08" +
	"y\xa9\xecdI\x16I\x13\x97\x18o$\xb66\x1e\xc4" +
	"\x0bX\x887\xb0\xb4Y\x89h5\xf3\xffm\x9eWb" +
	"\xd3[@\x93t-~\x1c|\xbf>\x87\xd3'\xe86" +
	"\x19\xcb\xfb\xcf\xb7l\xe7\xe4\x1d\xa9\xac\x03v\xeb\xcb\xee" +
	"/\xbb\xf7\x02\xc6\xc9w~\x08\xf3\xb1\xfc\x9d\xa3\xce\x85" +
	"1\x9c\xb5\xbe\xfc\xcd\x86\xd4\xcc$@B\xc0^\xee\x02" +
	"Z\x19j-$\x0b.vu\x08\xe8\x85\xa16B+" +
	",(\x80\xbdn\x01\xad\x0d\xf5Fh\x86\x9e+\x08W" +
	"\xe0F\xeff\xc7\x1c\xc2\x1c\x8c\x0f\xc3\xdd\xe8\xe6\xc7\x09" +
	"\xf4\xff\xf6\x13\x00\x00\xff\xffS\x9e(\x90"

func init() {
	schemas.Register(schema_dd30180cd0e36f1f,
		0xa43270a6b4fa26e1)
}
