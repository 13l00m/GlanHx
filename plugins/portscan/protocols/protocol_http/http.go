package protocol_http

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

type HTTP_Response struct {
	Title      string
	Body       string
	Length     int64
	StatusCode int
	TLS        bool
}

type Protocol_HTTP struct {
}

func (r *HTTP_Response) ParseResponse(response *http.Response, TLS bool) {
	r.Length = response.ContentLength
	r.TLS = TLS
	r.StatusCode = response.StatusCode
	//r.Title = getTitle(response.Body)
	parser, err := html.Parse(response.Body)
	if err != nil {
		r.Title = ""
	}
	r.Title = getTitle(parser)
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

func (r Protocol_HTTP) Analysis(host string, port int) (string, any, error) {
	isHTTP := false
	isHTTPS := false
	isWEB := false
	response := HTTP_Response{}

	// 尝试发送 HTTP 请求

	http_req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/", host, port), nil)
	if err != nil {

	}
	https_req, err := http.NewRequest("GET", fmt.Sprintf("https://%s:%d/", host, port), nil)
	if err != nil {

	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   1 * time.Second,
		Transport: tr,
	}
	http_resp, _ := client.Do(http_req)
	if http_resp != nil {
		isWEB = true
		isHTTP = true
		defer http_resp.Body.Close()
	}
	https_resp, _ := client.Do(https_req)

	if https_resp != nil {
		isWEB = true
		isHTTPS = true
		isHTTP = false

		defer https_resp.Body.Close()
	}

	if isWEB {
		if isHTTP {
			response.ParseResponse(http_resp, false)
		} else if isHTTPS {
			response.ParseResponse(https_resp, true)
		}
		return "HTTP", response, nil
	}

	return "", nil, fmt.Errorf(host, port, "not http protocol")

}
