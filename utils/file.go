package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

func IpFileToStringArray(filename string) []string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {

	}

	var verified_ips []string
	nverifiedips := strings.Split(string(data), "\n")

	for _, ip := range nverifiedips {
		ip = strings.ReplaceAll(ip, "\r", "")
		ip = strings.ReplaceAll(ip, " ", "")
		if net.ParseIP(ip) == nil {
			fmt.Println(ip)
			continue
		}
		verified_ips = append(verified_ips, ip)
	}

	return verified_ips

}

func HostFileToStringArray(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {

	}

	var verfied_hosts []string
	nverifiedhosts := strings.Split(string(data), "\n")

	for _, host := range nverifiedhosts {
		host = strings.ReplaceAll(host, "\r", "")
		host = strings.ReplaceAll(host, " ", "")
		if len(host) != 0 {
			verfied_hosts = append(verfied_hosts, host)
		}
	}
	return verfied_hosts
}
