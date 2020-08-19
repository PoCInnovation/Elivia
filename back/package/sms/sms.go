package main

import "fmt"

// Module test package
func Module(entries map[string]string) {
	fmt.Println("entries :")
	for n, e := range entries {
		fmt.Println("\"", n, "\"\t: \"", e, "\"")
	}

	fmt.Println("sms loaded")
}
