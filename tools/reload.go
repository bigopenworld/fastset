package tools

import (
	"os/exec"

	"github.com/bigopenworld/fastset/errmsg"
	"github.com/bigopenworld/fastset/logger"
)

func Reload() error {
	err_reload := exec.Command(".","~/.bashrc").Run()
	logger.PrintNewLine(logger.Color_Yellow, "Warning : you may need to reload and restart the script to continue the installation")
	if err_reload != nil {
		return errmsg.EXEC_ERROR(err_reload)	
	}
	return nil 
}