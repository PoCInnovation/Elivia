package main

// Module test package
func Module(_ string, entries map[string]string) (string, map[string]interface{}) {
	message := entries["message"]
	contact := entries["contact"]
	return "success", map[string]interface{}{
		"actions": []string{"sms", message, contact},
		"message": message,
		"contact": contact,
	}
}
