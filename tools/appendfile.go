package tools

import (
	"bufio"
	"os"
)

func Appendfilebeg(file string, line string) {
	var temp = "temp.txt"
	addline := line
    // make a temporary outfile
    outfile, err := os.Create(temp)

    if err != nil {
        panic(err)
    }

    defer outfile.Close()

    // open the file to be appended to for read
    f, err := os.Open(file)

    if err != nil {
        panic(err)
    }

    defer f.Close()

    // append at the start
    _, err = outfile.WriteString(addline)
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)

    // read the file to be appended to and output all of it
    for scanner.Scan() {

        _, err = outfile.WriteString(scanner.Text())
        _, err = outfile.WriteString("\n")
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    // ensure all lines are written
    outfile.Sync()
    // over write the old file with the new one
    err = os.Rename(temp, file)
    if err != nil {
        panic(err)
    }
} 