// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/core/l2tp.api.json

// Package l2tp contains generated bindings for API file l2tp.api.
//
// Contents:
//   1 enum
//  10 messages
//
package l2tp

import (
	api "git.fd.io/govpp.git/api"
	_ "git.fd.io/govpp.git/binapi/ethernet_types"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "l2tp"
	APIVersion = "2.0.0"
	VersionCrc = 0xb018cef7
)

// L2tLookupKey defines enum 'l2t_lookup_key'.
type L2tLookupKey uint8

const (
	L2T_LOOKUP_KEY_API_SRC_ADDR   L2tLookupKey = 0
	L2T_LOOKUP_KEY_API_DST_ADDR   L2tLookupKey = 1
	L2T_LOOKUP_KEY_API_SESSION_ID L2tLookupKey = 2
)

var (
	L2tLookupKey_name = map[uint8]string{
		0: "L2T_LOOKUP_KEY_API_SRC_ADDR",
		1: "L2T_LOOKUP_KEY_API_DST_ADDR",
		2: "L2T_LOOKUP_KEY_API_SESSION_ID",
	}
	L2tLookupKey_value = map[string]uint8{
		"L2T_LOOKUP_KEY_API_SRC_ADDR":   0,
		"L2T_LOOKUP_KEY_API_DST_ADDR":   1,
		"L2T_LOOKUP_KEY_API_SESSION_ID": 2,
	}
)

func (x L2tLookupKey) String() string {
	s, ok := L2tLookupKey_name[uint8(x)]
	if ok {
		return s
	}
	return "L2tLookupKey(" + strconv.Itoa(int(x)) + ")"
}

// L2tpv3CreateTunnel defines message 'l2tpv3_create_tunnel'.
type L2tpv3CreateTunnel struct {
	ClientAddress     ip_types.Address `binapi:"address,name=client_address" json:"client_address,omitempty"`
	OurAddress        ip_types.Address `binapi:"address,name=our_address" json:"our_address,omitempty"`
	LocalSessionID    uint32           `binapi:"u32,name=local_session_id" json:"local_session_id,omitempty"`
	RemoteSessionID   uint32           `binapi:"u32,name=remote_session_id" json:"remote_session_id,omitempty"`
	LocalCookie       uint64           `binapi:"u64,name=local_cookie" json:"local_cookie,omitempty"`
	RemoteCookie      uint64           `binapi:"u64,name=remote_cookie" json:"remote_cookie,omitempty"`
	L2SublayerPresent bool             `binapi:"bool,name=l2_sublayer_present" json:"l2_sublayer_present,omitempty"`
	EncapVrfID        uint32           `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
}

func (m *L2tpv3CreateTunnel) Reset()               { *m = L2tpv3CreateTunnel{} }
func (*L2tpv3CreateTunnel) GetMessageName() string { return "l2tpv3_create_tunnel" }
func (*L2tpv3CreateTunnel) GetCrcString() string   { return "596892cb" }
func (*L2tpv3CreateTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3CreateTunnel) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1      // m.ClientAddress.Af
	size += 1 * 16 // m.ClientAddress.Un
	size += 1      // m.OurAddress.Af
	size += 1 * 16 // m.OurAddress.Un
	size += 4      // m.LocalSessionID
	size += 4      // m.RemoteSessionID
	size += 8      // m.LocalCookie
	size += 8      // m.RemoteCookie
	size += 1      // m.L2SublayerPresent
	size += 4      // m.EncapVrfID
	return size
}
func (m *L2tpv3CreateTunnel) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint8(uint8(m.ClientAddress.Af))
	buf.EncodeBytes(m.ClientAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint8(uint8(m.OurAddress.Af))
	buf.EncodeBytes(m.OurAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint32(uint32(m.LocalSessionID))
	buf.EncodeUint32(uint32(m.RemoteSessionID))
	buf.EncodeUint64(uint64(m.LocalCookie))
	buf.EncodeUint64(uint64(m.RemoteCookie))
	buf.EncodeBool(m.L2SublayerPresent)
	buf.EncodeUint32(uint32(m.EncapVrfID))
	return buf.Bytes(), nil
}
func (m *L2tpv3CreateTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ClientAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.ClientAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.OurAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.OurAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalSessionID = buf.DecodeUint32()
	m.RemoteSessionID = buf.DecodeUint32()
	m.LocalCookie = buf.DecodeUint64()
	m.RemoteCookie = buf.DecodeUint64()
	m.L2SublayerPresent = buf.DecodeBool()
	m.EncapVrfID = buf.DecodeUint32()
	return nil
}

// L2tpv3CreateTunnelReply defines message 'l2tpv3_create_tunnel_reply'.
type L2tpv3CreateTunnelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *L2tpv3CreateTunnelReply) Reset()               { *m = L2tpv3CreateTunnelReply{} }
func (*L2tpv3CreateTunnelReply) GetMessageName() string { return "l2tpv3_create_tunnel_reply" }
func (*L2tpv3CreateTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*L2tpv3CreateTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3CreateTunnelReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *L2tpv3CreateTunnelReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *L2tpv3CreateTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// L2tpv3InterfaceEnableDisable defines message 'l2tpv3_interface_enable_disable'.
type L2tpv3InterfaceEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *L2tpv3InterfaceEnableDisable) Reset() { *m = L2tpv3InterfaceEnableDisable{} }
func (*L2tpv3InterfaceEnableDisable) GetMessageName() string {
	return "l2tpv3_interface_enable_disable"
}
func (*L2tpv3InterfaceEnableDisable) GetCrcString() string { return "3865946c" }
func (*L2tpv3InterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3InterfaceEnableDisable) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *L2tpv3InterfaceEnableDisable) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *L2tpv3InterfaceEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// L2tpv3InterfaceEnableDisableReply defines message 'l2tpv3_interface_enable_disable_reply'.
type L2tpv3InterfaceEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3InterfaceEnableDisableReply) Reset() { *m = L2tpv3InterfaceEnableDisableReply{} }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageName() string {
	return "l2tpv3_interface_enable_disable_reply"
}
func (*L2tpv3InterfaceEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3InterfaceEnableDisableReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3InterfaceEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *L2tpv3InterfaceEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// L2tpv3SetLookupKey defines message 'l2tpv3_set_lookup_key'.
type L2tpv3SetLookupKey struct {
	Key L2tLookupKey `binapi:"l2t_lookup_key,name=key" json:"key,omitempty"`
}

func (m *L2tpv3SetLookupKey) Reset()               { *m = L2tpv3SetLookupKey{} }
func (*L2tpv3SetLookupKey) GetMessageName() string { return "l2tpv3_set_lookup_key" }
func (*L2tpv3SetLookupKey) GetCrcString() string   { return "c9892c86" }
func (*L2tpv3SetLookupKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3SetLookupKey) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.Key
	return size
}
func (m *L2tpv3SetLookupKey) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint8(uint8(m.Key))
	return buf.Bytes(), nil
}
func (m *L2tpv3SetLookupKey) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Key = L2tLookupKey(buf.DecodeUint8())
	return nil
}

// L2tpv3SetLookupKeyReply defines message 'l2tpv3_set_lookup_key_reply'.
type L2tpv3SetLookupKeyReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3SetLookupKeyReply) Reset()               { *m = L2tpv3SetLookupKeyReply{} }
func (*L2tpv3SetLookupKeyReply) GetMessageName() string { return "l2tpv3_set_lookup_key_reply" }
func (*L2tpv3SetLookupKeyReply) GetCrcString() string   { return "e8d4e804" }
func (*L2tpv3SetLookupKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3SetLookupKeyReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3SetLookupKeyReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *L2tpv3SetLookupKeyReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// L2tpv3SetTunnelCookies defines message 'l2tpv3_set_tunnel_cookies'.
type L2tpv3SetTunnelCookies struct {
	SwIfIndex       interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	NewLocalCookie  uint64                         `binapi:"u64,name=new_local_cookie" json:"new_local_cookie,omitempty"`
	NewRemoteCookie uint64                         `binapi:"u64,name=new_remote_cookie" json:"new_remote_cookie,omitempty"`
}

func (m *L2tpv3SetTunnelCookies) Reset()               { *m = L2tpv3SetTunnelCookies{} }
func (*L2tpv3SetTunnelCookies) GetMessageName() string { return "l2tpv3_set_tunnel_cookies" }
func (*L2tpv3SetTunnelCookies) GetCrcString() string   { return "b3f4faf7" }
func (*L2tpv3SetTunnelCookies) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2tpv3SetTunnelCookies) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.SwIfIndex
	size += 8 // m.NewLocalCookie
	size += 8 // m.NewRemoteCookie
	return size
}
func (m *L2tpv3SetTunnelCookies) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint64(uint64(m.NewLocalCookie))
	buf.EncodeUint64(uint64(m.NewRemoteCookie))
	return buf.Bytes(), nil
}
func (m *L2tpv3SetTunnelCookies) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.NewLocalCookie = buf.DecodeUint64()
	m.NewRemoteCookie = buf.DecodeUint64()
	return nil
}

// L2tpv3SetTunnelCookiesReply defines message 'l2tpv3_set_tunnel_cookies_reply'.
type L2tpv3SetTunnelCookiesReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2tpv3SetTunnelCookiesReply) Reset()               { *m = L2tpv3SetTunnelCookiesReply{} }
func (*L2tpv3SetTunnelCookiesReply) GetMessageName() string { return "l2tpv3_set_tunnel_cookies_reply" }
func (*L2tpv3SetTunnelCookiesReply) GetCrcString() string   { return "e8d4e804" }
func (*L2tpv3SetTunnelCookiesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2tpv3SetTunnelCookiesReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *L2tpv3SetTunnelCookiesReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *L2tpv3SetTunnelCookiesReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// SwIfL2tpv3TunnelDetails defines message 'sw_if_l2tpv3_tunnel_details'.
type SwIfL2tpv3TunnelDetails struct {
	SwIfIndex         interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName     string                         `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	ClientAddress     ip_types.Address               `binapi:"address,name=client_address" json:"client_address,omitempty"`
	OurAddress        ip_types.Address               `binapi:"address,name=our_address" json:"our_address,omitempty"`
	LocalSessionID    uint32                         `binapi:"u32,name=local_session_id" json:"local_session_id,omitempty"`
	RemoteSessionID   uint32                         `binapi:"u32,name=remote_session_id" json:"remote_session_id,omitempty"`
	LocalCookie       []uint64                       `binapi:"u64[2],name=local_cookie" json:"local_cookie,omitempty"`
	RemoteCookie      uint64                         `binapi:"u64,name=remote_cookie" json:"remote_cookie,omitempty"`
	L2SublayerPresent bool                           `binapi:"bool,name=l2_sublayer_present" json:"l2_sublayer_present,omitempty"`
}

func (m *SwIfL2tpv3TunnelDetails) Reset()               { *m = SwIfL2tpv3TunnelDetails{} }
func (*SwIfL2tpv3TunnelDetails) GetMessageName() string { return "sw_if_l2tpv3_tunnel_details" }
func (*SwIfL2tpv3TunnelDetails) GetCrcString() string   { return "1dab5c7e" }
func (*SwIfL2tpv3TunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwIfL2tpv3TunnelDetails) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4      // m.SwIfIndex
	size += 64     // m.InterfaceName
	size += 1      // m.ClientAddress.Af
	size += 1 * 16 // m.ClientAddress.Un
	size += 1      // m.OurAddress.Af
	size += 1 * 16 // m.OurAddress.Un
	size += 4      // m.LocalSessionID
	size += 4      // m.RemoteSessionID
	size += 8 * 2  // m.LocalCookie
	size += 8      // m.RemoteCookie
	size += 1      // m.L2SublayerPresent
	return size
}
func (m *SwIfL2tpv3TunnelDetails) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeUint8(uint8(m.ClientAddress.Af))
	buf.EncodeBytes(m.ClientAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint8(uint8(m.OurAddress.Af))
	buf.EncodeBytes(m.OurAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint32(uint32(m.LocalSessionID))
	buf.EncodeUint32(uint32(m.RemoteSessionID))
	for i := 0; i < 2; i++ {
		var x uint64
		if i < len(m.LocalCookie) {
			x = uint64(m.LocalCookie[i])
		}
		buf.EncodeUint64(uint64(x))
	}
	buf.EncodeUint64(uint64(m.RemoteCookie))
	buf.EncodeBool(m.L2SublayerPresent)
	return buf.Bytes(), nil
}
func (m *SwIfL2tpv3TunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.ClientAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.ClientAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.OurAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.OurAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalSessionID = buf.DecodeUint32()
	m.RemoteSessionID = buf.DecodeUint32()
	m.LocalCookie = make([]uint64, 2)
	for i := 0; i < len(m.LocalCookie); i++ {
		m.LocalCookie[i] = buf.DecodeUint64()
	}
	m.RemoteCookie = buf.DecodeUint64()
	m.L2SublayerPresent = buf.DecodeBool()
	return nil
}

// SwIfL2tpv3TunnelDump defines message 'sw_if_l2tpv3_tunnel_dump'.
type SwIfL2tpv3TunnelDump struct{}

func (m *SwIfL2tpv3TunnelDump) Reset()               { *m = SwIfL2tpv3TunnelDump{} }
func (*SwIfL2tpv3TunnelDump) GetMessageName() string { return "sw_if_l2tpv3_tunnel_dump" }
func (*SwIfL2tpv3TunnelDump) GetCrcString() string   { return "51077d14" }
func (*SwIfL2tpv3TunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwIfL2tpv3TunnelDump) Size() int {
	if m == nil {
		return 0
	}
	var size int
	return size
}
func (m *SwIfL2tpv3TunnelDump) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	return buf.Bytes(), nil
}
func (m *SwIfL2tpv3TunnelDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_l2tp_binapi_init() }
func file_l2tp_binapi_init() {
	api.RegisterMessage((*L2tpv3CreateTunnel)(nil), "l2tpv3_create_tunnel_596892cb")
	api.RegisterMessage((*L2tpv3CreateTunnelReply)(nil), "l2tpv3_create_tunnel_reply_5383d31f")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisable)(nil), "l2tpv3_interface_enable_disable_3865946c")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisableReply)(nil), "l2tpv3_interface_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*L2tpv3SetLookupKey)(nil), "l2tpv3_set_lookup_key_c9892c86")
	api.RegisterMessage((*L2tpv3SetLookupKeyReply)(nil), "l2tpv3_set_lookup_key_reply_e8d4e804")
	api.RegisterMessage((*L2tpv3SetTunnelCookies)(nil), "l2tpv3_set_tunnel_cookies_b3f4faf7")
	api.RegisterMessage((*L2tpv3SetTunnelCookiesReply)(nil), "l2tpv3_set_tunnel_cookies_reply_e8d4e804")
	api.RegisterMessage((*SwIfL2tpv3TunnelDetails)(nil), "sw_if_l2tpv3_tunnel_details_1dab5c7e")
	api.RegisterMessage((*SwIfL2tpv3TunnelDump)(nil), "sw_if_l2tpv3_tunnel_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L2tpv3CreateTunnel)(nil),
		(*L2tpv3CreateTunnelReply)(nil),
		(*L2tpv3InterfaceEnableDisable)(nil),
		(*L2tpv3InterfaceEnableDisableReply)(nil),
		(*L2tpv3SetLookupKey)(nil),
		(*L2tpv3SetLookupKeyReply)(nil),
		(*L2tpv3SetTunnelCookies)(nil),
		(*L2tpv3SetTunnelCookiesReply)(nil),
		(*SwIfL2tpv3TunnelDetails)(nil),
		(*SwIfL2tpv3TunnelDump)(nil),
	}
}