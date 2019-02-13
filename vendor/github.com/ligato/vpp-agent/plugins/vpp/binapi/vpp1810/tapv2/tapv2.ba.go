// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/tapv2.api.json

/*
 Package tapv2 is a generated from VPP binary API module 'tapv2'.

 It contains following objects:
	  3 services
	  6 messages
*/
package tapv2

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpSwInterfaceTapV2(*SwInterfaceTapV2Dump) ([]*SwInterfaceTapV2Details, error)
	TapCreateV2(*TapCreateV2) (*TapCreateV2Reply, error)
	TapDeleteV2(*TapDeleteV2) (*TapDeleteV2Reply, error)
}

/* Messages */

// SwInterfaceTapV2Details represents VPP binary API message 'sw_interface_tap_v2_details':
type SwInterfaceTapV2Details struct {
	SwIfIndex        uint32
	ID               uint32
	DevName          []byte `struc:"[64]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfName       []byte `struc:"[64]byte"`
	HostNamespace    []byte `struc:"[64]byte"`
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
}

func (*SwInterfaceTapV2Details) GetMessageName() string {
	return "sw_interface_tap_v2_details"
}
func (*SwInterfaceTapV2Details) GetCrcString() string {
	return "b4c58229"
}
func (*SwInterfaceTapV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceTapV2Dump represents VPP binary API message 'sw_interface_tap_v2_dump':
type SwInterfaceTapV2Dump struct{}

func (*SwInterfaceTapV2Dump) GetMessageName() string {
	return "sw_interface_tap_v2_dump"
}
func (*SwInterfaceTapV2Dump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceTapV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapCreateV2 represents VPP binary API message 'tap_create_v2':
type TapCreateV2 struct {
	ID               uint32
	UseRandomMac     uint8
	MacAddress       []byte `struc:"[6]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostNamespaceSet uint8
	HostNamespace    []byte `struc:"[64]byte"`
	HostMacAddrSet   uint8
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfNameSet    uint8
	HostIfName       []byte `struc:"[64]byte"`
	HostBridgeSet    uint8
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4AddrSet   uint8
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6AddrSet   uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
	HostIP4GwSet     uint8
	HostIP4Gw        []byte `struc:"[4]byte"`
	HostIP6GwSet     uint8
	HostIP6Gw        []byte `struc:"[16]byte"`
	Tag              []byte `struc:"[64]byte"`
}

func (*TapCreateV2) GetMessageName() string {
	return "tap_create_v2"
}
func (*TapCreateV2) GetCrcString() string {
	return "34ce8043"
}
func (*TapCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapCreateV2Reply represents VPP binary API message 'tap_create_v2_reply':
type TapCreateV2Reply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapCreateV2Reply) GetMessageName() string {
	return "tap_create_v2_reply"
}
func (*TapCreateV2Reply) GetCrcString() string {
	return "fda5941f"
}
func (*TapCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapDeleteV2 represents VPP binary API message 'tap_delete_v2':
type TapDeleteV2 struct {
	SwIfIndex uint32
}

func (*TapDeleteV2) GetMessageName() string {
	return "tap_delete_v2"
}
func (*TapDeleteV2) GetCrcString() string {
	return "529cb13f"
}
func (*TapDeleteV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapDeleteV2Reply represents VPP binary API message 'tap_delete_v2_reply':
type TapDeleteV2Reply struct {
	Retval int32
}

func (*TapDeleteV2Reply) GetMessageName() string {
	return "tap_delete_v2_reply"
}
func (*TapDeleteV2Reply) GetCrcString() string {
	return "e8d4e804"
}
func (*TapDeleteV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SwInterfaceTapV2Details)(nil), "tapv2.SwInterfaceTapV2Details")
	api.RegisterMessage((*SwInterfaceTapV2Dump)(nil), "tapv2.SwInterfaceTapV2Dump")
	api.RegisterMessage((*TapCreateV2)(nil), "tapv2.TapCreateV2")
	api.RegisterMessage((*TapCreateV2Reply)(nil), "tapv2.TapCreateV2Reply")
	api.RegisterMessage((*TapDeleteV2)(nil), "tapv2.TapDeleteV2")
	api.RegisterMessage((*TapDeleteV2Reply)(nil), "tapv2.TapDeleteV2Reply")
}

var Messages = []api.Message{
	(*SwInterfaceTapV2Details)(nil),
	(*SwInterfaceTapV2Dump)(nil),
	(*TapCreateV2)(nil),
	(*TapCreateV2Reply)(nil),
	(*TapDeleteV2)(nil),
	(*TapDeleteV2Reply)(nil),
}
