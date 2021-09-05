package tools

import (
	"os"

	"github.com/progrium/go-basher"
)

func Exec_SH(path string, args ...string) error {
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.Source(path, nil)
	bash.Stdin = os.Stdin
	bash.Stdout = os.Stdout
	_, err := bash.Run(args[0], args[1:])
	return err 
	/*cmd := &exec.Cmd{
		Path:        	"bash",
		Args:        []string{"./test.sh"},
	}*/
	//cmd := exec.Command("/bin/bash", "/home/paul/.nvm/nvm.sh","&&","nvm","install")
	// cmd := exec.Command("/home/paul/.nvm/nvm.sh", "nvm")
	// println(string(cmd.String()))
	// return cmd
}