package main

import (
	"GlanHx/config"
	"GlanHx/plugins/hostcolide"
	"flag"
	"fmt"
	"os"
)

func main() {

	// 定义基础命令
	//flag.StringVar(&configFile, "config", "", "The path to the config file")
	// 定义 flag 集合
	// 根据不同目录调用不同的 flag 集合

	config.Init()
	flag.Usage = func() {
		fmt.Println("\n  ________ .__                     ___ ___          \n /  _____/ |  |  _____     ____   /   |   \\ ___  ___\n/   \\  ___ |  |  \\__  \\   /    \\ /    ~    \\\\  \\/  /\n\\    \\_\\  \\|  |__ / __ \\_|   |  \\\\    Y    / >    < \n \\______  /|____/(____  /|___|  / \\___|_  / /__/\\_ \\\n        \\/            \\/      \\/        \\/        \\/\n")
		fmt.Println("\n相关用法如下：")
		fmt.Println("hostcolide", "-h host碰撞模块,使用-h参数查看具体用法：./GlanHx hostcolide -h")
	}

	dir1Flags := flag.NewFlagSet("hostcolide", flag.ExitOnError)
	if len(os.Args) < 2 {
		//fmt.Println("No command group used")
		flag.Usage()
	} else {
		switch os.Args[1] {
		case "hostcolide":
			hostcolide.ParseFlag(dir1Flags, os.Args[2:])

		default:
			flag.Usage()
		}
	}

}
