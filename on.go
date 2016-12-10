package main

import "github.com/lnguyen/wemo"
import "os"
import "fmt"
import "strings"

func debug(message string) {
	if os.Getenv("DEBUG") == "1" || os.Getenv("DEBUG") == "true" {
		fmt.Printf("%v\n", message)
	}
}

/*
handle no switches
index is off
index by name maybe?
*/

func main() {
	switches, err := wemo.Switches()
	fmt.Println(switches)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(os.Args[0], "off") {
		switches[0].Off()
	} else {
		switches[0].On()
	}
}
