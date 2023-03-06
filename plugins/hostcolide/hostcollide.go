package hostcolide

import (
	"GlanHx/utils"
	"flag"
	"strings"
)

var (
	ip         string
	ipFile     string
	host       string
	hostFile   string
	thread     int
	port       string
	output     string
	factor     string
	debug      bool
	ipList     []string
	hostList   []string
	factorList []string
	portList   []int
)

func ParseFlag(flagset *flag.FlagSet, args []string) {

	flagset.Usage = func() {
		flagset.PrintDefaults()
	}

	flagset.StringVar(&ip, "I", "", "Nginx Ip")
	flagset.StringVar(&ipFile, "IF", "", "Nginx Ip in file")
	flagset.StringVar(&host, "H", "", "Host")
	flagset.StringVar(&hostFile, "HF", "", "Host in file")
	flagset.IntVar(&thread, "T", 10, "Thread default 10")
	flagset.StringVar(&port, "P", "80,443", "Port default 80,443")
	flagset.StringVar(&output, "O", "result.txt", "output ")
	flagset.StringVar(&factor, "F", "title,status_code", "Generate hash based on factors to filter junk data,Default title,status_code. Supported title,status_code,length")
	flagset.BoolVar(&debug, "D", false, "Debug Mod if Open output All info")
	flagset.Parse(args)

	if ip == "" && ipFile == "" {
		flagset.Usage()
		return
	}

	if host == "" && hostFile == "" {
		flagset.Usage()
		return
	}

	if ip != "" {
		ip = strings.ReplaceAll(ip, " ", "")
		ipList = append(ipList, ip)
	}
	if ipFile != "" {
		ipFile = strings.ReplaceAll(ipFile, " ", "")
		ipList = append(ipList, utils.IpFileToStringArray(ipFile)...)
	}
	if host != "" {
		host = strings.ReplaceAll(host, " ", "")
		hostList = append(hostList, host)
	}
	if hostFile != "" {
		hostFile = strings.ReplaceAll(hostFile, " ", "")
		hostList = append(hostList, utils.HostFileToStringArray(hostFile)...)
	}

	port = strings.ReplaceAll(port, " ", "")
	portList = utils.StringArray2IntArray(strings.Split(port, ","))
	ipList = utils.RemoveDuplicate_String(ipList)
	hostList = utils.RemoveDuplicate_String(hostList)
	factor = strings.ReplaceAll(factor, " ", "")
	factorList = strings.Split(factor, ",")

	Scan()

}
