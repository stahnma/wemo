package main

import "github.com/lnguyen/wemo"
import "os"
//import "fmt"
import "strings"

func main() {
	switches, err := wemo.Switches()
	if err != nil {

	}
	if strings.Contains(os.Args[0], "off") {
		switches[0].Off()
	} else {
		switches[0].On()
	}
}
