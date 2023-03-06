package protocol_http

import (
	"GlanHx/utils"
	"crypto/tls"
	"fmt"
	"net/http"
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

var SupportRedirect bool

func (r *HTTP_Response) ParseResponse(response *http.Response, TLS bool) {
	r.TLS = TLS
	r.StatusCode = response.StatusCode
	r.Length = response.ContentLength
	//r.Title = getTitle(response.Body)
	r.Title, r.Length, _ = utils.GetTitleAndLength(response.Body)
}

func (r Protocol_HTTP) Analysis(host string, port int) (string, any, error) {
	var http_resp *http.Response
	var https_resp *http.Response
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

	var client *http.Client

	if SupportRedirect == true {

		client = &http.Client{
			Timeout:   10 * time.Second,
			Transport: tr,
			//CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//	return http.ErrUseLastResponse
			//},
		}
	}

	if SupportRedirect == false {

		client = &http.Client{
			Timeout:   10 * time.Second,
			Transport: tr,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	https_resp, err = client.Do(https_req)

	if err == nil {
		if https_resp != nil {
			isWEB = true
			isHTTPS = true
			isHTTP = false

			defer https_resp.Body.Close()
		}
	}

	if isWEB == false {

		http_resp, err = client.Do(http_req)

		if err == nil {
			if http_resp != nil {
				if http_resp.Status != "400 Bad Request" {
					isWEB = true
					isHTTP = true
					defer http_resp.Body.Close()
				}
			}
		}
	}

	if isWEB {
		if isHTTPS {
			response.ParseResponse(https_resp, true)
		} else if isHTTP {
			response.ParseResponse(http_resp, false)
		}
		return "HTTP", response, nil
	}

	return "", nil, fmt.Errorf(host, port, "not http protocol")

}
