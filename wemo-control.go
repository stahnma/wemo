package main

import "github.com/lnguyen/wemo"
import "os"
import "fmt"
import "strings"
import "flag"

func debug(message string) {
	if os.Getenv("DEBUG") == "1" || os.Getenv("DEBUG") == "true" {
		fmt.Printf("%v\n", message)
	}
}

func display_status() {
	switches, err := wemo.Switches()
	if err != nil {
		fmt.Println(err)
	}
	for _, device := range switches {
		on_off := device.Status()
		on_off_s := ""
		if on_off == 1 {
			on_off_s = "On"
		} else {
			on_off_s = "Off"
		}
		fmt.Println(device.Name + ": " + on_off_s)
	}
}

func main() {
	var status = flag.Bool("status", false, "List out wemo devices and status")
	flag.Parse()

	if *status {
		display_status()
		os.Exit(0)
	}

	switches, err := wemo.Switches()
	if err != nil {
		fmt.Println(err)
	}
	// next if it doesn't contain office
	for _, device := range switches {
		if strings.Contains(device.Name, "Office") {
			if strings.Contains(os.Args[0], "off") {
				device.Off()
			} else {
				device.On()
			}
		}
	}
}
