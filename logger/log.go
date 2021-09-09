package logger

import (
	"fmt"
	"os"

	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
)

const (
	Color_Black = ct.Black
	Color_Blue = ct.Blue
	Color_Cyan = ct.Cyan
	Color_Green = ct.Green
	Color_Magenta = ct.Magenta
	Color_None = ct.None
	Color_Red = ct.Red
	Color_White = ct.White
	Color_Yellow = ct.Yellow

	
)

var writer *goterminal.Writer 

func Init() {
	writer = goterminal.New(os.Stdout)
}

func PrintSameLine(color ct.Color, text string) {
	writer.Clear()
	ct.Foreground(color, false)
	fmt.Fprintln(writer, text)
	writer.Print()
	ct.ResetColor()
}
func PrintNewLine(color ct.Color, text string) {
	ct.Foreground(color, false)
	fmt.Fprintln(writer, text)
	writer.Print()
	ct.ResetColor()
}