// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: spacex/api/device/wifi_util.proto

package device

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IfaceType int32

const (
	IfaceType_IFACE_TYPE_UNKNOWN      IfaceType = 0
	IfaceType_IFACE_TYPE_ETH          IfaceType = 1
	IfaceType_IFACE_TYPE_RF_2GHZ      IfaceType = 2
	IfaceType_IFACE_TYPE_RF_5GHZ      IfaceType = 5
	IfaceType_IFACE_TYPE_RF_5GHZ_HIGH IfaceType = 6
)

// Enum value maps for IfaceType.
var (
	IfaceType_name = map[int32]string{
		0: "IFACE_TYPE_UNKNOWN",
		1: "IFACE_TYPE_ETH",
		2: "IFACE_TYPE_RF_2GHZ",
		5: "IFACE_TYPE_RF_5GHZ",
		6: "IFACE_TYPE_RF_5GHZ_HIGH",
	}
	IfaceType_value = map[string]int32{
		"IFACE_TYPE_UNKNOWN":      0,
		"IFACE_TYPE_ETH":          1,
		"IFACE_TYPE_RF_2GHZ":      2,
		"IFACE_TYPE_RF_5GHZ":      5,
		"IFACE_TYPE_RF_5GHZ_HIGH": 6,
	}
)

func (x IfaceType) Enum() *IfaceType {
	p := new(IfaceType)
	*p = x
	return p
}

func (x IfaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IfaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_wifi_util_proto_enumTypes[0].Descriptor()
}

func (IfaceType) Type() protoreflect.EnumType {
	return &file_spacex_api_device_wifi_util_proto_enumTypes[0]
}

func (x IfaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IfaceType.Descriptor instead.
func (IfaceType) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_wifi_util_proto_rawDescGZIP(), []int{0}
}

type TxPowerLevel int32

const (
	TxPowerLevel_TX_POWER_LEVEL_100 TxPowerLevel = 0
	TxPowerLevel_TX_POWER_LEVEL_80  TxPowerLevel = 1
	TxPowerLevel_TX_POWER_LEVEL_50  TxPowerLevel = 2
	TxPowerLevel_TX_POWER_LEVEL_25  TxPowerLevel = 3
	TxPowerLevel_TX_POWER_LEVEL_12  TxPowerLevel = 4
	TxPowerLevel_TX_POWER_LEVEL_6   TxPowerLevel = 5
)

// Enum value maps for TxPowerLevel.
var (
	TxPowerLevel_name = map[int32]string{
		0: "TX_POWER_LEVEL_100",
		1: "TX_POWER_LEVEL_80",
		2: "TX_POWER_LEVEL_50",
		3: "TX_POWER_LEVEL_25",
		4: "TX_POWER_LEVEL_12",
		5: "TX_POWER_LEVEL_6",
	}
	TxPowerLevel_value = map[string]int32{
		"TX_POWER_LEVEL_100": 0,
		"TX_POWER_LEVEL_80":  1,
		"TX_POWER_LEVEL_50":  2,
		"TX_POWER_LEVEL_25":  3,
		"TX_POWER_LEVEL_12":  4,
		"TX_POWER_LEVEL_6":   5,
	}
)

func (x TxPowerLevel) Enum() *TxPowerLevel {
	p := new(TxPowerLevel)
	*p = x
	return p
}

func (x TxPowerLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxPowerLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_wifi_util_proto_enumTypes[1].Descriptor()
}

func (TxPowerLevel) Type() protoreflect.EnumType {
	return &file_spacex_api_device_wifi_util_proto_enumTypes[1]
}

func (x TxPowerLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxPowerLevel.Descriptor instead.
func (TxPowerLevel) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_wifi_util_proto_rawDescGZIP(), []int{1}
}

type InflatedBasicServiceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bssid      string    `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Ssid       string    `protobuf:"bytes,2,opt,name=ssid,proto3" json:"ssid,omitempty"`
	MacLan     string    `protobuf:"bytes,3,opt,name=mac_lan,json=macLan,proto3" json:"mac_lan,omitempty"`
	IfaceName  string    `protobuf:"bytes,4,opt,name=iface_name,json=ifaceName,proto3" json:"iface_name,omitempty"`
	IfaceType  IfaceType `protobuf:"varint,5,opt,name=iface_type,json=ifaceType,proto3,enum=SpaceX.API.Device.IfaceType" json:"iface_type,omitempty"`
	Channel    uint32    `protobuf:"varint,6,opt,name=channel,proto3" json:"channel,omitempty"`
	Preference uint32    `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
}

func (x *InflatedBasicServiceSet) Reset() {
	*x = InflatedBasicServiceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_wifi_util_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InflatedBasicServiceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InflatedBasicServiceSet) ProtoMessage() {}

func (x *InflatedBasicServiceSet) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_wifi_util_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InflatedBasicServiceSet.ProtoReflect.Descriptor instead.
func (*InflatedBasicServiceSet) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_wifi_util_proto_rawDescGZIP(), []int{0}
}

func (x *InflatedBasicServiceSet) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *InflatedBasicServiceSet) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *InflatedBasicServiceSet) GetMacLan() string {
	if x != nil {
		return x.MacLan
	}
	return ""
}

func (x *InflatedBasicServiceSet) GetIfaceName() string {
	if x != nil {
		return x.IfaceName
	}
	return ""
}

func (x *InflatedBasicServiceSet) GetIfaceType() IfaceType {
	if x != nil {
		return x.IfaceType
	}
	return IfaceType_IFACE_TYPE_UNKNOWN
}

func (x *InflatedBasicServiceSet) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *InflatedBasicServiceSet) GetPreference() uint32 {
	if x != nil {
		return x.Preference
	}
	return 0
}

type DhcpLease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress   string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	MacAddress  string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Hostname    string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ExpiresTime string `protobuf:"bytes,4,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"`
	Active      bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	ClientId    uint32 `protobuf:"varint,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *DhcpLease) Reset() {
	*x = DhcpLease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_wifi_util_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DhcpLease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DhcpLease) ProtoMessage() {}

func (x *DhcpLease) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_wifi_util_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DhcpLease.ProtoReflect.Descriptor instead.
func (*DhcpLease) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_wifi_util_proto_rawDescGZIP(), []int{1}
}

func (x *DhcpLease) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *DhcpLease) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *DhcpLease) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DhcpLease) GetExpiresTime() string {
	if x != nil {
		return x.ExpiresTime
	}
	return ""
}

func (x *DhcpLease) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *DhcpLease) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type DhcpServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string       `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Subnet string       `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Leases []*DhcpLease `protobuf:"bytes,3,rep,name=leases,proto3" json:"leases,omitempty"`
}

func (x *DhcpServer) Reset() {
	*x = DhcpServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_wifi_util_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DhcpServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DhcpServer) ProtoMessage() {}

func (x *DhcpServer) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_wifi_util_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DhcpServer.ProtoReflect.Descriptor instead.
func (*DhcpServer) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_wifi_util_proto_rawDescGZIP(), []int{2}
}

func (x *DhcpServer) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DhcpServer) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *DhcpServer) GetLeases() []*DhcpLease {
	if x != nil {
		return x.Leases
	}
	return nil
}

var File_spacex_api_device_wifi_util_proto protoreflect.FileDescriptor

var file_spacex_api_device_wifi_util_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x77, 0x69, 0x66, 0x69, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x66, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x61, 0x63, 0x5f, 0x6c, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x61, 0x63, 0x4c, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x66, 0x61, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x69, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x66, 0x61,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x69, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x09,
	0x44, 0x68, 0x63, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x72, 0x0a,
	0x0a, 0x44, 0x68, 0x63, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x68, 0x63, 0x70, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x2a, 0x84, 0x01, 0x0a, 0x09, 0x49, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x49, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x46, 0x41, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x54, 0x48, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49,
	0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x46, 0x5f, 0x32, 0x47, 0x48,
	0x5a, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x46, 0x5f, 0x35, 0x47, 0x48, 0x5a, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x49,
	0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x46, 0x5f, 0x35, 0x47, 0x48,
	0x5a, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x06, 0x2a, 0x98, 0x01, 0x0a, 0x0c, 0x54, 0x78, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x58, 0x5f,
	0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x31, 0x30, 0x30, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x38, 0x30, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x50,
	0x4f, 0x57, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x35, 0x30, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x32, 0x35, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x50, 0x4f, 0x57,
	0x45, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x31, 0x32, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x58, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x36, 0x10, 0x05, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_device_wifi_util_proto_rawDescOnce sync.Once
	file_spacex_api_device_wifi_util_proto_rawDescData = file_spacex_api_device_wifi_util_proto_rawDesc
)

func file_spacex_api_device_wifi_util_proto_rawDescGZIP() []byte {
	file_spacex_api_device_wifi_util_proto_rawDescOnce.Do(func() {
		file_spacex_api_device_wifi_util_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_device_wifi_util_proto_rawDescData)
	})
	return file_spacex_api_device_wifi_util_proto_rawDescData
}

var file_spacex_api_device_wifi_util_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spacex_api_device_wifi_util_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spacex_api_device_wifi_util_proto_goTypes = []interface{}{
	(IfaceType)(0),                  // 0: SpaceX.API.Device.IfaceType
	(TxPowerLevel)(0),               // 1: SpaceX.API.Device.TxPowerLevel
	(*InflatedBasicServiceSet)(nil), // 2: SpaceX.API.Device.InflatedBasicServiceSet
	(*DhcpLease)(nil),               // 3: SpaceX.API.Device.DhcpLease
	(*DhcpServer)(nil),              // 4: SpaceX.API.Device.DhcpServer
}
var file_spacex_api_device_wifi_util_proto_depIdxs = []int32{
	0, // 0: SpaceX.API.Device.InflatedBasicServiceSet.iface_type:type_name -> SpaceX.API.Device.IfaceType
	3, // 1: SpaceX.API.Device.DhcpServer.leases:type_name -> SpaceX.API.Device.DhcpLease
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spacex_api_device_wifi_util_proto_init() }
func file_spacex_api_device_wifi_util_proto_init() {
	if File_spacex_api_device_wifi_util_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacex_api_device_wifi_util_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InflatedBasicServiceSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_wifi_util_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DhcpLease); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spacex_api_device_wifi_util_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DhcpServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacex_api_device_wifi_util_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_device_wifi_util_proto_goTypes,
		DependencyIndexes: file_spacex_api_device_wifi_util_proto_depIdxs,
		EnumInfos:         file_spacex_api_device_wifi_util_proto_enumTypes,
		MessageInfos:      file_spacex_api_device_wifi_util_proto_msgTypes,
	}.Build()
	File_spacex_api_device_wifi_util_proto = out.File
	file_spacex_api_device_wifi_util_proto_rawDesc = nil
	file_spacex_api_device_wifi_util_proto_goTypes = nil
	file_spacex_api_device_wifi_util_proto_depIdxs = nil
}
