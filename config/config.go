package config

import (
	"GlanHx/plugins/portscan"
	"GlanHx/plugins/portscan/protocols"
	"GlanHx/plugins/portscan/protocols/protocol_http"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var GlobalConfig config

type portScanConfig struct {
	DefaultPortList string
	PortScanThread  int
}

type hostCollideConfig struct {
	ImpactThread int
}

type pluginsConfig struct {
	Portscan    portScanConfig
	HostCollide hostCollideConfig
}

type config struct {
	Plugins pluginsConfig
}

func GetConfig() config {
	return config{}
}

func Init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic("config file not exists")
	}
	var config config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err.Error())
	}

	GlobalConfig = config

	//对plugin/portscan 进行初始化
	portlist_str := strings.Split(config.Plugins.Portscan.DefaultPortList, ",")
	portlist_int := make([]int, len(portlist_str))

	for i, port := range portlist_str {
		port_int, err := strconv.Atoi(port)
		if err != nil {
			panic("PortScan load port error")
			return
		}

		portlist_int[i] = port_int
	}
	portscan.ScanThread = config.Plugins.Portscan.PortScanThread
	portscan.DefaltScanPort = portlist_int
	protocols.Protocol_support = []protocols.AnalysisProtocol{protocol_http.Protocol_HTTP{}}
	//plugins.Init()
}

func GenerateConfig(config config) {
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将 YAML 格式的数据写入本地文件
	err = ioutil.WriteFile("config.yaml", yamlData, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
