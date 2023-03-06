package portscan

import (
	"GlanHx/plugins/portscan/protocols"
	"fmt"
	"net"
	"sync"
	"time"
)

var ScanThread int
var DefaltScanPort []int

func Run(target string, portlist []int) []protocols.Protocol {
	var protocolList []protocols.Protocol
	host := target
	// 去重
	topPorts := removeDuplicate(portlist)

	// 控制协程数
	//concurrency := 100
	semaphore := make(chan struct{}, ScanThread)

	var wg sync.WaitGroup

	for _, port := range topPorts {

		semaphore <- struct{}{}
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			// 建立 TCP 连接
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 1*time.Second)

			if err != nil {
				return
			}
			defer conn.Close()

			protocol := protocols.Protocol{}
			protocol.ParseProtocol(conn)
			protocolList = append(protocolList, protocol)

		}(port)
	}

	wg.Wait()
	return protocolList
}

func removeDuplicate(elements []int) []int {
	// 使用 map 去重
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			continue
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
