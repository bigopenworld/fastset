package tools

import "runtime"

const (
	OS_Linux = "linux"
	OS_Windows = "windows"
	OS_MacOS = "darwin"
)

func OS_EQ(name string) bool {
	return runtime.GOOS == name
}
func OS_NEQ(name string) bool {
	return runtime.GOOS != name
}