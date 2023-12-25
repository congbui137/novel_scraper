package main

import (
	"fmt"
	"strings"
	"time"
)

import "runtime/debug"

func PrintBuildInformation() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to retrieve build information!")
		return
	}

	var settings = map[string]string{}

	//flatten
	for _, setting := range info.Settings {
		settings[strings.ToLower(setting.Key)] = setting.Value
	}

	parsedTime, _ := time.Parse(time.RFC3339, settings["vcs.time"]) //it's always valid?

	fmt.Printf("Program built using: %s | Commit ID: %s | Commit Date: %s\n",
		info.GoVersion,
		settings["vcs.revision"][:7],
		parsedTime.In(time.Now().Location()).Format(time.RFC822Z),
	)

}

func main() {
	PrintBuildInformation()
}
