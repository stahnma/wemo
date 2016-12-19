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
	if err != nil {
		fmt.Println(err)
	}
	// next if it doesn't contain office
	for _, device := range switches {
		//fmt.Println(device.Name)
		if strings.Contains(device.Name, "Office") {
			if strings.Contains(os.Args[0], "off") {
				device.Off()
			} else {
				device.On()
			}
		}
	}
}
