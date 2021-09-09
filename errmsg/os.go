package errmsg

import "errors"

func OS_ERROR(os string) error {
	return errors.New("OS_ERROR : sorry this cli function only support " + os)
}
func OS_ERROR_INSTALL_PKG(ossupported string, pkgname string) error {
	return errors.New("OS_ERROR_INSTALL_PKG : sorry the cli only support "+pkgname+" install on "+ossupported)
}