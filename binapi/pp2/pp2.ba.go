// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0
//  VPP:              22.06-release
// source: /usr/share/vpp/api/plugins/pp2.api.json

// Package pp2 contains generated bindings for API file pp2.api.
//
// Contents:
// -  4 messages
package pp2

import (
	api "go.fd.io/govpp/api"
	interface_types "go.fd.io/govpp/binapi/interface_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "pp2"
	APIVersion = "1.0.0"
	VersionCrc = 0xd7ab5bd7
)

// MrvlPp2Create defines message 'mrvl_pp2_create'.
type MrvlPp2Create struct {
	IfName string `binapi:"string[64],name=if_name" json:"if_name,omitempty"`
	RxQSz  uint16 `binapi:"u16,name=rx_q_sz" json:"rx_q_sz,omitempty"`
	TxQSz  uint16 `binapi:"u16,name=tx_q_sz" json:"tx_q_sz,omitempty"`
}

func (m *MrvlPp2Create) Reset()               { *m = MrvlPp2Create{} }
func (*MrvlPp2Create) GetMessageName() string { return "mrvl_pp2_create" }
func (*MrvlPp2Create) GetCrcString() string   { return "3a108396" }
func (*MrvlPp2Create) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MrvlPp2Create) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.IfName
	size += 2  // m.RxQSz
	size += 2  // m.TxQSz
	return size
}
func (m *MrvlPp2Create) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.IfName, 64)
	buf.EncodeUint16(m.RxQSz)
	buf.EncodeUint16(m.TxQSz)
	return buf.Bytes(), nil
}
func (m *MrvlPp2Create) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IfName = buf.DecodeString(64)
	m.RxQSz = buf.DecodeUint16()
	m.TxQSz = buf.DecodeUint16()
	return nil
}

// MrvlPp2CreateReply defines message 'mrvl_pp2_create_reply'.
type MrvlPp2CreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MrvlPp2CreateReply) Reset()               { *m = MrvlPp2CreateReply{} }
func (*MrvlPp2CreateReply) GetMessageName() string { return "mrvl_pp2_create_reply" }
func (*MrvlPp2CreateReply) GetCrcString() string   { return "5383d31f" }
func (*MrvlPp2CreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MrvlPp2CreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *MrvlPp2CreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MrvlPp2CreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// MrvlPp2Delete defines message 'mrvl_pp2_delete'.
type MrvlPp2Delete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MrvlPp2Delete) Reset()               { *m = MrvlPp2Delete{} }
func (*MrvlPp2Delete) GetMessageName() string { return "mrvl_pp2_delete" }
func (*MrvlPp2Delete) GetCrcString() string   { return "f9e6675e" }
func (*MrvlPp2Delete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MrvlPp2Delete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *MrvlPp2Delete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MrvlPp2Delete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// MrvlPp2DeleteReply defines message 'mrvl_pp2_delete_reply'.
type MrvlPp2DeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MrvlPp2DeleteReply) Reset()               { *m = MrvlPp2DeleteReply{} }
func (*MrvlPp2DeleteReply) GetMessageName() string { return "mrvl_pp2_delete_reply" }
func (*MrvlPp2DeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*MrvlPp2DeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MrvlPp2DeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MrvlPp2DeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MrvlPp2DeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_pp2_binapi_init() }
func file_pp2_binapi_init() {
	api.RegisterMessage((*MrvlPp2Create)(nil), "mrvl_pp2_create_3a108396")
	api.RegisterMessage((*MrvlPp2CreateReply)(nil), "mrvl_pp2_create_reply_5383d31f")
	api.RegisterMessage((*MrvlPp2Delete)(nil), "mrvl_pp2_delete_f9e6675e")
	api.RegisterMessage((*MrvlPp2DeleteReply)(nil), "mrvl_pp2_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MrvlPp2Create)(nil),
		(*MrvlPp2CreateReply)(nil),
		(*MrvlPp2Delete)(nil),
		(*MrvlPp2DeleteReply)(nil),
	}
}