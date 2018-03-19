package main

const (
	EtherAddrLen = 6
	IPv4AddrLen  = 4
	IPv6AddrLen  = 16
)

// Supported EtherType for L2
const (
	IPV4Number = 0x0800
	ARPNumber  = 0x0806
	VLANNumber = 0x8100
	MPLSNumber = 0x8847
	IPV6Number = 0x86dd
)

// Supported L4 types
const (
	ICMPNumber   = 0x01
	IPNumber     = 0x04
	TCPNumber    = 0x06
	UDPNumber    = 0x11
	NoNextHeader = 0x3B
)

// Supported ICMP Types
const (
	ICMPTypeEchoRequest  uint8 = 8
	ICMPTypeEchoResponse uint8 = 0
)

// These constants keep length of supported headers in bytes.
//
// IPv6Len - minimum length of IPv6 header in bytes. It can be higher and it
// is not determined inside packet. Only default minimum size is used.
//
// IPv4MinLen and TCPMinLen are used only in packet generation functions.
//
// In parsing we take actual length of TCP header from DataOff field and length of
// IPv4 take from Ihl field.
const (
	EtherLen   = 14
	VLANLen    = 4
	MPLSLen    = 4
	IPv4MinLen = 20
	IPv6Len    = 40
	ICMPLen    = 8
	TCPMinLen  = 20
	UDPLen     = 8
	ARPLen     = 28
	GTPMinLen  = 8
)

// LogType - type of logging, used in flow package
type LogType uint8

const (
	// No - no output even after fatal errors
	No LogType = 1 << iota
	// Initialization - output during system initialization
	Initialization
	// Debug - output during execution one time per time period (scheduler ticks)
	Debug
	// Verbose - output during execution as soon as something happens. Can influence performance
	Verbose
)

// TCPFlags contains set TCP flags.
type TCPFlags uint8

// Constants for valuues of TCP flags.
const (
	TCPFlagFin = 0x01
	TCPFlagSyn = 0x02
	TCPFlagRst = 0x04
	TCPFlagPsh = 0x08
	TCPFlagAck = 0x10
	TCPFlagUrg = 0x20
	TCPFlagEce = 0x40
	TCPFlagCwr = 0x80
)
