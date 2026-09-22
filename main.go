package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	device := SetupDevice()
	var input string
	fmt.Println("Welcome to rpi-status-cli! To show available commands, enter \"help\"")
	path := fmt.Sprintf("mount | grep -E '%s|boot'", device)
	for {
		fmt.Print("~> ")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("checknet - check the internet connection")
			fmt.Println("showcfg - display the Raspberry Pi boot configuration")
			fmt.Println("t - check temperature")
		case "showcfg":
			cmd := exec.Command(path)
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			result := string(output)
			pathread := fmt.Sprintf("%s/config.txt", result)
			read, err := os.ReadFile(pathread)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			fmt.Println(string(read))
		case "checknet":
			cmd := exec.Command("ping", "-c", "4", "8.8.8.8")
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println(string(output))
		case "t":
			cmd := exec.Command("sh", "-c", "paste <(cat /sys/class/thermal/thermal_zone*/type) <(cat /sys/class/thermal/thermal_zone*/temp) | awk '{print $1 \" : \" $2/1000 \"°C\"}'")
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			fmt.Println(string(output))
		}
	}
}
