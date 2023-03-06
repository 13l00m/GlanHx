package protocols

import (
	"net"
)

type Protocol struct {
	Protocol       string
	ProtocolDetail any
	port           int
	host           string
}

var Protocol_support []AnalysisProtocol
var Protocol_scan []AnalysisProtocol

type AnalysisProtocol interface {
	Analysis(host string, port int) (string, any, error)
}

func (protocol *Protocol) ParseProtocol(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
	protocol.host = remoteAddr.IP.String()
	protocol.port = remoteAddr.Port
	if len(Protocol_scan) != 0 {
		for _, scanInterface := range Protocol_scan {
			p, pd, err := scanInterface.Analysis(protocol.host, protocol.port)
			if err != nil {
				protocol.Protocol = "UNKNOWN"
			} else {
				protocol.Protocol = p
				protocol.ProtocolDetail = pd
			}
		}
	} else {
		for _, scanInterface := range Protocol_support {
			p, pd, err := scanInterface.Analysis(protocol.host, protocol.port)
			if err != nil {
				protocol.Protocol = "UNKNOWN"
			} else {
				protocol.Protocol = p
				protocol.ProtocolDetail = pd
			}
		}
	}
}

func (protocol Protocol) GetHost() string {
	return protocol.host
}

func (protocol Protocol) GetPort() int {
	return protocol.port
}

func Register(a AnalysisProtocol) {
	Protocol_scan = append(Protocol_scan, a)
}
