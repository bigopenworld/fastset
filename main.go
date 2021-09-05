package main

import (
	"fmt"

	"github.com/bigopenworld/fastset/framework"
	"github.com/bigopenworld/fastset/install"
	"github.com/bigopenworld/fastset/logger"
	"github.com/thatisuday/commando"
)

func main() {
	logger.Init()
	commando.
		SetExecutableName("fastset").
		SetVersion("v1.0.0").
		SetDescription("A cli that help you create project initial files faster")
	commando.
		Register(nil).
		AddArgument("target", "the name of the framework or lang that you want to use", "").
		AddFlag("verbose, V", "display log information", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	

		})
	commando.
		Register("framework").
		AddArgument("framework", "the name of the framework that you want to use", "").
		AddArgument("name", "the name of your project", "").
		AddFlag("verbose, V", "display log information", commando.Bool, nil).
		SetAction(framework.Select)
	commando.
		Register("lang").
		AddArgument("lang", "the name of the lang that you want to use", "").
		AddFlag("verbose, V", "display log information", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println(args["framework"].Value)
		})
	commando.
		Register("install").
		AddArgument("packagename", "the name of the lang that you want to use", "").
		AddFlag("verbose, V", "display log information", commando.Bool, nil).
		SetAction(install.Select)
	commando.Parse(nil)
}
