package myutil

import (
	"encoding/json"
	"io/ioutil"
)

// TMP TO UPDATE

// ReadFile returns the bytes of a file searched in the path and beyond it
func ReadFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		bytes, err = ioutil.ReadFile("../" + path)
	}

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Response contains the trigger's tag and its contained matched sentences
type Response struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

// AButl - After-Before Utils
type abutl struct {
	Key string
	X   int
}

// BWutl - Between Utils
type bwutl struct {
	After string
	X     int

	Before string
	Y      int
}

// Entries -
type Entries struct {
	Name      string      `json:"name"`
	Parser    string      `json:"parser"`
	Resources interface{} `json:"resources"`
}

// Trigger contains the trigger's tag and its contained matched sentences
type Trigger struct {
	Patterns []string `json:"patterns"`

	// this is a tomporary system implemented until olivia IA is revisited and can train with entries.
	OliviaPatterns []string  `json:"olivia-feed"`
	CallBack       string    `json:"callback"`
	Entries        []Entries `json:"entries"`
}

// IOMod pairs up the response and trigger for each modules individualy
type IOMod struct {
	Triggers []Trigger
	Response []Response
}

var iomodList = map[string](map[string]IOMod){}

func getTrigger(name string, locale string) ([]Trigger, error) {
	var ModuleTriggers []Trigger

	content, err := ReadFile("package/" + name + "/res/locales/" + locale + "/triggers.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &ModuleTriggers)
	if err != nil {
		return nil, err
	}
	return ModuleTriggers, nil
}

func getResponse(name string, locale string) ([]Response, error) {
	var ModuleResponse []Response
	content, err := ReadFile("package/" + name + "/res/locales/" + locale + "/response.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &ModuleResponse)
	if err != nil {
		return nil, err
	}
	return ModuleResponse, nil
}

// SerializeIO -
func SerializeIO(name string, locale string) (IOMod, error) {
	var err error
	var triggers []Trigger
	var responses []Response
	var iomod IOMod

	triggers, err = getTrigger(name, locale)
	if err != nil {
		return iomod, err
	}
	responses, err = getResponse(name, locale)
	if err != nil {
		return iomod, err
	}
	iomod = IOMod{
		Triggers: triggers,
		Response: responses}

	//iomodList[locale][name] = iomod

	return iomod, nil
}
