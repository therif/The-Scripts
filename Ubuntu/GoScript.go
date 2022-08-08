package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("-- MacOS Auto ServiceKU --")

	fmt.Println("Service MySQL...")
	SrvNama("mysql")

	fmt.Println("Service PHP...")
	SrvNama("php")

	fmt.Println("Service Nginx...")
	SrvNama("nginx")
}

func SrvNama(servicenya string) {
	if len(strings.TrimSpace(servicenya)) > 0 {
		cmd := exec.Command("brew", "services", "restart", servicenya)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
	}
}
