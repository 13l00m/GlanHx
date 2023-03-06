package plugins

//
//import (
//	"GlanHx/config"
//	"GlanHx/plugins/portscan"
//	"GlanHx/plugins/portscan/protocols"
//	"GlanHx/plugins/portscan/protocols/protocol_http"
//	"strconv"
//	"strings"
//)
//
//func Init() {
//	//对端口进行设置
//	portlist_str := strings.Split(config.GlobalConfig.Plugins.Portscan.DefaultPortList, ",")
//	portlist_int := make([]int, len(portlist_str))
//
//	for i, port := range portlist_str {
//		port_int, err := strconv.Atoi(port)
//		if err != nil {
//			panic("PortScan load port error")
//			return
//		}
//
//		portlist_int[i] = port_int
//	}
//	portscan.ScanThread = config.GlobalConfig.Plugins.Portscan.PortScanThread
//	portscan.DefaltScanPort = portlist_int
//	protocols.Protocol_support = []protocols.AnalysisProtocol{protocol_http.Protocol_HTTP{}}
//
//}
