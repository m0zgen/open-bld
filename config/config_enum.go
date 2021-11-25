// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package config

import (
	"fmt"
	"strings"
)

const (
	// NetProtocolUdp is a NetProtocol of type Udp.
	// Deprecated: use tcp+udp instead
	NetProtocolUdp NetProtocol = iota
	// NetProtocolTcp is a NetProtocol of type Tcp.
	// Deprecated: use tcp+udp instead
	NetProtocolTcp
	// NetProtocolTcpUdp is a NetProtocol of type Tcp+Udp.
	// TCP and UDP protocols
	NetProtocolTcpUdp
	// NetProtocolTcpTls is a NetProtocol of type Tcp-Tls.
	// TCP-TLS protocol
	NetProtocolTcpTls
	// NetProtocolHttps is a NetProtocol of type Https.
	// HTTPS protocol
	NetProtocolHttps
)

const _NetProtocolName = "udptcptcp+udptcp-tlshttps"

var _NetProtocolNames = []string{
	_NetProtocolName[0:3],
	_NetProtocolName[3:6],
	_NetProtocolName[6:13],
	_NetProtocolName[13:20],
	_NetProtocolName[20:25],
}

// NetProtocolNames returns a list of possible string values of NetProtocol.
func NetProtocolNames() []string {
	tmp := make([]string, len(_NetProtocolNames))
	copy(tmp, _NetProtocolNames)
	return tmp
}

var _NetProtocolMap = map[NetProtocol]string{
	0: _NetProtocolName[0:3],
	1: _NetProtocolName[3:6],
	2: _NetProtocolName[6:13],
	3: _NetProtocolName[13:20],
	4: _NetProtocolName[20:25],
}

// String implements the Stringer interface.
func (x NetProtocol) String() string {
	if str, ok := _NetProtocolMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NetProtocol(%d)", x)
}

var _NetProtocolValue = map[string]NetProtocol{
	_NetProtocolName[0:3]:   0,
	_NetProtocolName[3:6]:   1,
	_NetProtocolName[6:13]:  2,
	_NetProtocolName[13:20]: 3,
	_NetProtocolName[20:25]: 4,
}

// ParseNetProtocol attempts to convert a string to a NetProtocol
func ParseNetProtocol(name string) (NetProtocol, error) {
	if x, ok := _NetProtocolValue[name]; ok {
		return x, nil
	}
	return NetProtocol(0), fmt.Errorf("%s is not a valid NetProtocol, try [%s]", name, strings.Join(_NetProtocolNames, ", "))
}

// MarshalText implements the text marshaller method
func (x NetProtocol) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *NetProtocol) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNetProtocol(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// QueryLogTypeConsole is a QueryLogType of type Console.
	// use logger as fallback
	QueryLogTypeConsole QueryLogType = iota
	// QueryLogTypeNone is a QueryLogType of type None.
	// no logging
	QueryLogTypeNone
	// QueryLogTypeMysql is a QueryLogType of type Mysql.
	// MySQL or MariaDB database
	QueryLogTypeMysql
	// QueryLogTypeCsv is a QueryLogType of type Csv.
	// CSV file per day
	QueryLogTypeCsv
	// QueryLogTypeCsvClient is a QueryLogType of type Csv-Client.
	// CSV file per day and client
	QueryLogTypeCsvClient
)

const _QueryLogTypeName = "consolenonemysqlcsvcsv-client"

var _QueryLogTypeNames = []string{
	_QueryLogTypeName[0:7],
	_QueryLogTypeName[7:11],
	_QueryLogTypeName[11:16],
	_QueryLogTypeName[16:19],
	_QueryLogTypeName[19:29],
}

// QueryLogTypeNames returns a list of possible string values of QueryLogType.
func QueryLogTypeNames() []string {
	tmp := make([]string, len(_QueryLogTypeNames))
	copy(tmp, _QueryLogTypeNames)
	return tmp
}

var _QueryLogTypeMap = map[QueryLogType]string{
	0: _QueryLogTypeName[0:7],
	1: _QueryLogTypeName[7:11],
	2: _QueryLogTypeName[11:16],
	3: _QueryLogTypeName[16:19],
	4: _QueryLogTypeName[19:29],
}

// String implements the Stringer interface.
func (x QueryLogType) String() string {
	if str, ok := _QueryLogTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("QueryLogType(%d)", x)
}

var _QueryLogTypeValue = map[string]QueryLogType{
	_QueryLogTypeName[0:7]:   0,
	_QueryLogTypeName[7:11]:  1,
	_QueryLogTypeName[11:16]: 2,
	_QueryLogTypeName[16:19]: 3,
	_QueryLogTypeName[19:29]: 4,
}

// ParseQueryLogType attempts to convert a string to a QueryLogType
func ParseQueryLogType(name string) (QueryLogType, error) {
	if x, ok := _QueryLogTypeValue[name]; ok {
		return x, nil
	}
	return QueryLogType(0), fmt.Errorf("%s is not a valid QueryLogType, try [%s]", name, strings.Join(_QueryLogTypeNames, ", "))
}

// MarshalText implements the text marshaller method
func (x QueryLogType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *QueryLogType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseQueryLogType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
