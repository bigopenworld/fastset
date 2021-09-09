package framework

import (
	"github.com/bigopenworld/fastset/tools"
	"github.com/thatisuday/commando"
)

func Select(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

	tools.Reload()
	switch (args["framework"].Value) {
	case "react" : 
		React(args["name"].Value)
	}
}