package install

import (
	"github.com/bigopenworld/fastset/logger"
	"github.com/thatisuday/commando"
)

func Select(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	switch args["packagename"].Value {
	case "node":
		err := Node_Install()
		if err != nil {
			logger.PrintNewLine(logger.Color_Red, err.Error())
		}
	
	case "all":
		logger.PrintNewLine(logger.Color_Blue, "node ; ")
	}
}
