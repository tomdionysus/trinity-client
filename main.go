package main

import (
	"bufio"
	"fmt"
	"github.com/tomdionysus/trinity-client/parser"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("# trinity-client v0.1.0\n")
	fmt.Print("# type `exit` to quit\n")
	for {
		fmt.Print("trinity>")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n")

		switch text {
		case "exit":
			goto stop
		default:
			term, err := parser.Parse(&text)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			} else {

			}
		}

	}

stop:
	os.Exit(0)
}
