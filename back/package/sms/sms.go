package main

import "fmt"

// Module test package
func Module(_ string, entries map[string]string) (string, map[string]interface{}) {
	fmt.Println("entries :")
	for n, e := range entries {
		fmt.Println("\"", n, "\"\t: \"", e, "\"")
	}

	fmt.Println("sms loaded")
	return "success", map[string]interface{}{
		"actions": []string{"sms", "haha", "Theo"},
	}
}
