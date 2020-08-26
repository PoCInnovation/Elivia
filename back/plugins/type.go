package plugins

import "plugin"

// A package handle a specific theme
// To do so, it is composed of various Module
// Each module has it own set of trigger and response
// All the module of a package share the same UserData

type ModuleFunc = func(string, map[string]string) (string, map[string]interface{})

// Package records all the necessary to load runtime package
type Package struct {
	Plug      *plugin.Plugin
	Responses []Response
	Modules   map[string]Module
	Name      string
}

// Response contains the trigger's tag and its contained matched sentences
type Response struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

// Module -
type Module struct {
	Triggers map[string]Trigger
	Func     ModuleFunc
}

// Entries -
type Entries struct {
	Name      string      `json:"name"`
	Parser    string      `json:"parser"`
	Resources interface{} `json:"resources"`
}

// Trigger contains the trigger's tag and its contained matched sentences
type Trigger struct {
	Patterns []string  `json:"patterns"`
	Entries  []Entries `json:"entries"`

	// this is a tomporary system implemented until olivia IA is revisited and can train with entries.
	OliviaPatterns []string `json:"olivia-feed"`
}
