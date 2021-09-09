package install

import (
	"os"
	"os/exec"

	"github.com/bigopenworld/fastset/errmsg"
	"github.com/bigopenworld/fastset/logger"
	"github.com/bigopenworld/fastset/tools"
	ct "github.com/daviddengcn/go-colortext"
)

func Node_Check() bool {
	var npm_exist bool = tools.CommandExists("npm")
	var node_exist bool = tools.CommandExists("node")
	return node_exist || npm_exist
}

func Nvm_Check() bool {
	var nvm_exist bool = tools.CommandExists(os.Getenv("HOME") + "/.nvm/nvm.sh")
	return nvm_exist
}

func Node_Install() error {
	logger.PrintNewLine(logger.Color_Green, "Installing node package")
	// check os
	switch {
	case tools.OS_EQ(tools.OS_Linux):
		return Node_Install_Linux()
	default:
		return errmsg.OS_ERROR_INSTALL_PKG("linux", "node.js")
	}
}

func Node_Install_Linux() error {
	logger.PrintNewLine(ct.Green, "Detected plateform linux")
	if !Nvm_Check() {
		Nvm_Install_Linux()
		// reload terminal
		err_reload_term := tools.Reload()
		if err_reload_term != nil {
			logger.PrintNewLine(logger.Color_Red, "Error : Unable to reload terminal, please close your terminal and restart fastset if you think this is an issue : https://github.com/bigopenworld/fastset/issues")
			return errmsg.EXEC_ERROR(err_reload_term)
		}
	}

	// check permission 
	scriptpath := os.Getenv("HOME") + "/.nvm/nvm.sh"
	println(scriptpath)
	err_file_test_exec := tools.FileExecutable(scriptpath)
	switch err_file_test_exec {
	case tools.ERR_NOT_OWNER:
		os.Chmod(scriptpath, 0701)
	case tools.ERR_NOT_EXEC_CUR_USER:
		os.Chmod(scriptpath, 0700)
	case nil:
		break
	default:
		return errmsg.EXEC_ERROR(err_file_test_exec)
	}
	// exec script
	err_install_node := tools.Exec_SH(scriptpath, "nvm", "install", "node")
	if err_install_node != nil {
			return errmsg.EXEC_ERROR(err_install_node)
	}
	logger.PrintNewLine(ct.Blue, "Node.JS installation success")
	return nil
}

func Nvm_Install_Linux() error {
	// download NVM
	err_download_deps_nvm := exec.Command("curl", "-sL", "https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh", "-o", "install_nvm.sh").Run()
	defer os.Remove("install_nvm.sh")
	if err_download_deps_nvm != nil {
		return errmsg.EXEC_ERROR(err_download_deps_nvm)
	}
	// execute NVM
	err_install_deps_nvm := exec.Command("bash", "./install_nvm.sh").Run()
	if err_install_deps_nvm != nil {
		return errmsg.EXEC_ERROR(err_install_deps_nvm)
	}
	return nil
}
