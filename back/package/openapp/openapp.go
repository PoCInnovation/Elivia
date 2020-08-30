package main

// Module test package
func Module(_ string, entries map[string]string) (string, map[string]interface{}) {
	appname := entries["appname"]
	if appname == "" {
		return "no app", nil
	}
	return "success", map[string]interface{}{
		"actions": []string{"openapp", appname},
		"appname": appname,
	}
}
