package tools

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

var ( 
	ERR_NOT_OWNER = errors.New("File not owned by current user")
	ERR_NOT_EXEC_CUR_USER = errors.New("File not executable by current user")
)

func CommandExists(cmd string) bool {
	if strings.Contains(cmd, ".sh") {
		if _, err := os.Stat(cmd); os.IsNotExist(err) {
			return false
		} 
		return true 
	} else {
		_, err := exec.LookPath(cmd)
		return err == nil
	}
}
func FileExecutable(path string) error {
	fileinfo, errstat := os.Lstat(path)
	if errstat != nil {
		return errstat
	}
	statinfo := fileinfo.Sys().(*syscall.Stat_t)
	currentuser, erruser := user.Current()
	if erruser != nil {
		return erruser
	}
	if strconv.Itoa(int(statinfo.Uid)) != currentuser.Uid {
		return ERR_NOT_OWNER
	}
	if IsExecOwner(fileinfo.Mode()) {
		return nil
	}
	return ERR_NOT_EXEC_CUR_USER
}
func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}
func IsExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}
func IsExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}