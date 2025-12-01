package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alireza0/s-ui/config"
	"github.com/alireza0/s-ui/database"
	"github.com/alireza0/s-ui/service"

	"github.com/shirou/gopsutil/v4/net"
)

func resetSetting() {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	settingService := service.SettingService{}
	err = settingService.ResetSettings()
	if err != nil {
		fmt.Println("reset setting failed:", err)
	} else {
		fmt.Println("reset setting success")
	}
}

func updateSetting(port int, path string, subPort int, subPath string) {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	settingService := service.SettingService{}

	if port > 0 {
		err := settingService.SetPort(port)
		if err != nil {
			fmt.Println("set port failed:", err)
		} else {
			fmt.Println("set port success")
		}
	}
	if path != "" {
		err := settingService.SetWebPath(path)
		if err != nil {
			fmt.Println("set path failed:", err)
		} else {
			fmt.Println("set path success")
		}
	}
	if subPort > 0 {
		err := settingService.SetSubPort(subPort)
		if err != nil {
			fmt.Println("set sub port failed:", err)
		} else {
			fmt.Println("set sub port success")
		}
	}
	if subPath != "" {
		err := settingService.SetSubPath(subPath)
		if err != nil {
			fmt.Println("set sub path failed:", err)
		} else {
			fmt.Println("set sub path success")
		}
	}
}

func showSetting() {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	settingService := service.SettingService{}
	allSetting, err := settingService.GetAllSetting()
	if err != nil {
		fmt.Println("get current port failed,error info:", err)
	}
	fmt.Println("Current panel settings:")
	fmt.Println("\tPanel port:\t", (*allSetting)["webPort"])
	fmt.Println("\tPanel path:\t", (*allSetting)["webPath"])
	if (*allSetting)["webListen"] != "" {
		fmt.Println("\tPanel IP:\t", (*allSetting)["webListen"])
	}
	if (*allSetting)["webDomain"] != "" {
		fmt.Println("\tPanel Domain:\t", (*allSetting)["webDomain"])
	}
	if (*allSetting)["webURI"] != "" {
		fmt.Println("\tPanel URI:\t", (*allSetting)["webURI"])
	}
	fmt.Println()
	fmt.Println("Current subscription settings:")
	fmt.Println("\tSub port:\t", (*allSetting)["subPort"])
	fmt.Println("\tSub path:\t", (*allSetting)["subPath"])
	if (*allSetting)["subListen"] != "" {
		fmt.Println("\tSub IP:\t", (*allSetting)["subListen"])
	}
	if (*allSetting)["subDomain"] != "" {
		fmt.Println("\tSub Domain:\t", (*allSetting)["subDomain"])
	}
	if (*allSetting)["subURI"] != "" {
		fmt.Println("\tSub URI:\t", (*allSetting)["subURI"])
	}
}

func getPanelURI() {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	settingService := service.SettingService{}
	Port, _ := settingService.GetPort()
	BasePath, _ := settingService.GetWebPath()
	Listen, _ := settingService.GetListen()
	Domain, _ := settingService.GetWebDomain()
	KeyFile, _ := settingService.GetKeyFile()
	CertFile, _ := settingService.GetCertFile()
	if Listen != "127.0.0.1" {
		fmt.Println("Your panel is listening on ", Listen)
		fmt.Println("It is suggested to listen on 127.0.0.1")
		fmt.Println("It is not recommended to expose the panel to the public network ON YOUR IP ADDRESS.")
		fmt.Println("It is VERY LIKELY your panel will be exploited AND your IP will be GFWed because of the panel.")
	}
	TLS := false
	if KeyFile != "" && CertFile != "" {
		TLS = true
	}
	if TLS {
		fmt.Println("Good! Your panel is using TLS.")
		fmt.Println("However, it is still not recommended to expose the panel to the public network as it may be scanned and marked and GFWed.")
		fmt.Printf("You may access the panel via: \nhttps://%s:%d%s\n", Domain, Port, BasePath)
	}
	fmt.Println("Please follow the best practice to access the panel: https://github.com/kunori-kiku/s-ui-modified")
	fmt.Printf("Current local address: http://localhost:%d%s\n", Port, BasePath)
	fmt.Println("If urgent, you may use SSH port forwarding to access the panel:")
	// get ip address
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err == nil {
		defer resp.Body.Close()
		ip, err := io.ReadAll(resp.Body)
		if err == nil {
			fmt.Printf("ssh -L %d:localhost:%d -p 22 %s\n(or replace `-p 22 %s` with SSH alias to your machine)\n", Port, Port, ip, ip)
		}
	}
}
