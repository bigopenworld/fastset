package framework

import (
	"fmt"
	"os/exec"
	"runtime"
)

func React(name string) {
	err := exec.Command("delta").Run()
	fmt.Println(err.Error())
	if err != nil {
		if runtime.GOOS != "linux" {
			fmt.Println("Sorry the CLI only support linux npm installation at this time \n please install node")
		}
		err := exec.Command("curl", "-sL","https://raw.githubusercontent.com/nvm-sh/nvm/v0.36.0/install.sh","-o", "install_nvm.sh").Run()
		if err != nil {
			fmt.Println("Error occured : ", err.Error())
		}
		err = exec.Command("bash","./install_nvm.sh").Run()
		if err != nil {
			fmt.Println("Error occured : ", err.Error())
		}
	}
}