package hostcolide

import (
	"GlanHx/plugins/portscan"
	"GlanHx/plugins/portscan/protocols"
	"GlanHx/plugins/portscan/protocols/protocol_http"
	"GlanHx/utils"
	"bufio"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var result []scanData

func Scan() {
	protocolList := make(map[string][]protocols.Protocol)
	semaphore := make(chan struct{}, thread)
	var wg sync.WaitGroup
	for _, ip := range ipList {
		protocol := portscan.Run(ip, portList)
		protocolList[ip] = protocol
		ipTask := make(map[string][]scanData)
		for _, p := range protocol {
			if p.Protocol == "HTTP" {
				if p.ProtocolDetail != nil {
					hash := generateHash(p.ProtocolDetail.(protocol_http.HTTP_Response).Title, p.ProtocolDetail.(protocol_http.HTTP_Response).StatusCode)
					ipTask[p.GetHost()] = append(ipTask[p.GetHost()], scanData{url: generateUrl(p.GetHost(), p.GetPort(), p.ProtocolDetail.(protocol_http.HTTP_Response).TLS), hash: hash})
				}
			}

		}
		for _, tasks := range ipTask {
			for _, task := range tasks {
				for _, h := range hostList {
					u := task.url
					semaphore <- struct{}{}
					wg.Add(1)
					go func(host, url string) {
						defer wg.Done()
						defer func() { <-semaphore }()
						scdata, err := doRequest(url, host)
						if err != nil {
							return
						}
						if debug == true {
							if scdata.hash != task.hash {
								saveData(scdata, 1, "[debug][+]")
							} else {
								saveData(scdata, 1, "[debug]")
							}
						}
						if debug == false {
							if scdata.hash != task.hash {
								saveData(scdata, 1, "[+]")
							}
						}

					}(h, u)
				}
				wg.Wait()
			}
		}

	}

	saveData(scanData{}, 2, "")

}

func saveData(scdata scanData, mod int, prefix string) {
	if mod == 1 {
		result = append(result, scdata)
		fmt.Println(prefix, scdata.url, scdata.host, scdata.length, scdata.title, scdata.status_code)
	}

	if mod == 2 {
		file, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for _, data := range result {
			writer.WriteString(fmt.Sprintf("%s %s %d %s %d\n", data.url, data.host, data.length, data.title, data.status_code))
		}
		writer.Flush()
	}
}

//这里就直接生成hash得了
func doRequest(Url, Host string) (scanData, error) {

	req, err := http.NewRequest("GET", Url, nil)

	if err != nil {
		return scanData{}, err
	}
	req.Host = Host
	//proxy, _ := url.Parse("http://127.0.0.1:8080")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//Proxy:           http.ProxyURL(proxy),
	}

	var client *http.Client

	if protocol_http.SupportRedirect == true {
		client = &http.Client{
			Timeout:   5 * time.Second,
			Transport: tr,
			//CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//	return http.ErrUseLastResponse
			//},
		}
	} else {
		client = &http.Client{
			Timeout:   5 * time.Second,
			Transport: tr,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	http_resp, err := client.Do(req)
	if err != nil {
		return scanData{}, err
	}
	if http_resp != nil {
		defer http_resp.Body.Close()

		title, length, _ := utils.GetTitleAndLength(http_resp.Body)

		hash := generateHash(title, http_resp.StatusCode)

		data := scanData{}
		data.title = title
		data.url = Url
		data.host = Host
		data.status_code = http_resp.StatusCode
		data.hash = hash
		data.length = length
		return data, nil
	}
	return scanData{}, errors.New("noresponse")

}

func getTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return strings.TrimSpace(n.FirstChild.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := getTitle(c)
		if title != "" {
			return title
		}
	}
	return ""
}

func Scan_API(iplist []string, portList []int) {

}

func generateHash(title string, status_code int) string {
	str := fmt.Sprintf("%s:%d", title, status_code)
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func generateUrl(host string, port int, tls bool) string {
	if port == 80 && tls == false {
		return "http://" + host + "/"
	} else if port == 443 && tls == true {
		return "https://" + host + "/"
	} else {
		if tls == true {
			return fmt.Sprintf("https://%s:%d/", host, port)
		}
		if tls == false {
			return fmt.Sprintf("http://%s:%d/", host, port)
		}
	}
	return ""

}

type scanData struct {
	hash        string
	url         string
	host        string
	title       string
	length      int64
	status_code int
}
