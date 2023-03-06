////package main
////
////import (
////	"crypto/tls"
////	"fmt"
////	"net"
////	"net/http"
////	"time"
////)
////
////func main() {
////	ports := []string{"80", "443"} // 可能的端口号
////
////	for _, port := range ports {
////		conn, err := net.DialTimeout("tcp", "www.baidu.com:"+port, 2*time.Second)
////		if err != nil {
////			fmt.Printf("Port %s closed\n", port)
////			continue
////		}
////		conn.Close()
////
////		url := "http://www.bilibili.com:" + port // 指定请求协议和主机地址
////		req, err := http.NewRequest("GET", url, nil)
////		if err != nil {
////			fmt.Printf("Error creating request: %s\n", err)
////			continue
////		}
////		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
////
////		client := &http.Client{
////			Timeout: 5 * time.Second,
////			Transport: &http.Transport{
////				TLSHandshakeTimeout: 2 * time.Second,
////				Proxy:               http.ProxyFromEnvironment,
////				DialContext: (&net.Dialer{
////					Timeout: 2 * time.Second,
////				}).DialContext,
////				TLSClientConfig: &tls.Config{
////					InsecureSkipVerify: true, // 取消证书验证
////				},
////			},
////		}
////
////		resp, err := client.Do(req)
////		if err != nil {
////			fmt.Printf("Port %s is open but not HTTP or HTTPS\n", port)
////			continue
////		}
////		defer resp.Body.Close()
////
////		if resp.StatusCode == 200 {
////			fmt.Printf("Port %s is open and using HTTP\n", port)
////		} else if resp.TLS != nil {
////			fmt.Printf("Port %s is open and using HTTPS\n", port)
////		} else {
////			fmt.Printf("Port %s is open but not HTTP or HTTPS\n", port)
////		}
////	}
////}
//
//package main
//
//import (
//	"GlanHx/plugins/portscan/protocols"
//	"fmt"
//	"io/ioutil"
//	"net"
//)
//
//func main() {
//
//	host := "www.baidu.com"
//	portList := []int{80, 443, 22, 21, 9901}
//
//	for _, port := range portList {
//		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
//		if err != nil {
//			continue
//		}
//		defer conn.Close()
//		protocol := protocols.Protocol{}
//		protocol.ParseProtocol(conn)
//		//remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
//		//fmt.Println(remoteAddr.IP, remoteAddr.Port)
//		//protocols := ""
//		//
//		//if tryHTTP(conn) {
//		//	fmt.Println(123123)
//		//	protocols = "HTTP"
//		//} else if trySSH(conn) {
//		//	fmt.Println(22323)
//		//	protocols = "SSH"
//		//}
//		//
//		//fmt.Println(host, port, protocols)
//	}
//}
//
//func tryHTTP(conn net.Conn) bool {
//	_, err := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
//	if err != nil {
//		return false
//	}
//
//	//var buf []byte
//	fmt.Println("succ")
//	buf, err := ioutil.ReadAll(conn)
//	if err != nil {
//		return false
//	}
//
//	fmt.Sprintln("http", string(buf))
//
//	return string(buf) == "HTTP/1.0 200 OK\r\n"
//}
//
//func tryFTP(conn net.Conn) bool {
//	buf := make([]byte, 1024)
//	n, err := conn.Read(buf)
//	if err != nil {
//		return false
//	}
//
//	return string(buf[:n]) == "220 FTP Service Ready\r\n"
//}
//
//func trySMTP(conn net.Conn) bool {
//	buf := make([]byte, 1024)
//	n, err := conn.Read(buf)
//	if err != nil {
//		return false
//	}
//	return string(buf[:n]) == "220 smtp.example.com ESMTP Postfix\r\n"
//}
//
//func trySSH(conn net.Conn) bool {
//
//	fmt.Println("xxx")
//	return false
//	//n, err := conn.Read(buf)
//	//if err != nil {
//	//	return false
//	//}
//	//fmt.Println("ssh", string(buf[:n]))
//	//return string(buf[:n]) == "SSH-2.0-OpenSSH_7.4p1 Debian-10+deb9u6\r\n"
//}

package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "http://205.189.160.142:9901", nil)
	//proxy, _ := url.Parse("http://127.0.0.1:8080")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//Proxy:           http.ProxyURL(proxy),
	}

	req.Host = "evilHost"

	client := &http.Client{
		Timeout:   1 * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	http_resp, _ := client.Do(req)
	if http_resp != nil {
		defer http_resp.Body.Close()
		//fmt.Println(Url, Host)
	}

}
