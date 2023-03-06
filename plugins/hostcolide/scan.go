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
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var result []scanData

var mu sync.Mutex

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
					hash := generateHash(p.ProtocolDetail.(protocol_http.HTTP_Response).Title, p.ProtocolDetail.(protocol_http.HTTP_Response).StatusCode, p.ProtocolDetail.(protocol_http.HTTP_Response).Length)
					ipTask[p.GetHost()] = append(ipTask[p.GetHost()], scanData{url: generateUrl(p.GetHost(), p.GetPort(), p.ProtocolDetail.(protocol_http.HTTP_Response).TLS), hash: hash})
				}
			}

		}
		for _, tasks := range ipTask {
			for _, task := range tasks {
				hashMap := make(map[string]int)
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
							if checkHashCount(hashMap, scdata.hash) {
								mu.Unlock()
								if scdata.hash != task.hash {
									saveData(scdata, 1, "[+]")
								}
							} else {
								mu.Unlock()
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

func checkHashCount(hashmap map[string]int, hash string) bool {
	mu.Lock()
	count := hashmap[hash]
	if count == 0 {
		hashmap[hash] += 1
		return true
	}
	if count >= 5 {
		return false
	} else {
		hashmap[hash] += 1
		return true
	}

}

func saveData(scdata scanData, mod int, prefix string) {
	if mod == 1 {
		result = append(result, scdata)
		fmt.Println(prefix, strings.Split(scdata.url, "/")[2], "--", scdata.host, "--", scdata.url, "title:", scdata.title, "status_code:", scdata.status_code, "length:", scdata.length)
	}

	if mod == 2 {
		file, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for _, data := range result {
			writer.WriteString(fmt.Sprintf("%s -- %s -- %s status_code: %d title: %s length: %d\n", strings.Split(data.url, "/")[2], data.host, data.url, data.title, data.status_code, data.length))
		}
		writer.Flush()
	}
}

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

		hash := generateHash(title, http_resp.StatusCode, length)

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

func Scan_API(iplist []string, portList []int) {

}

func generateHash(title string, status_code int, length int64) string {
	str := fmt.Sprintf("%s:%d:%d", title, status_code, length)
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
