package main

import "fmt"

func SetupDevice() string {
	var device string
	fmt.Println("Enter your boot device:")
	fmt.Print("> ")
	fmt.Scan(&device)
	return device
}
