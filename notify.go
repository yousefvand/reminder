package main

import (
	"fmt"
	"os/exec"
)

// Notify : Send notify message
func Notify() {
	// notify-send -u normal -t 10000 -a "System" -i flag-yellow "Battery Health" "Battery should be replaced"
	cmd := exec.Command("notify-send", "Alert", "Somebody's birthday tomorrow")
	if err := cmd.Run(); err != nil {
		fmt.Println("Sending notification failed: ", err)
	}
}
