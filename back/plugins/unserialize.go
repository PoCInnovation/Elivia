package plugins

import (
	"encoding/json"
	"strconv"

	"github.com/PoCInnovation/Elivia/util"
)

// TMP TO UPDATE

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

func tagTrigger(pname, mname, locale string, id int) string {
	return pname + "_" + locale + "_" + mname + "_" + strconv.Itoa(id)
}

func tagTriggers(triggers []Trigger, pname, mname, locale string) map[string]Trigger {
	m := make(map[string]Trigger)
	for i, elm := range triggers {
		m[tagTrigger(pname, mname, locale, i)] = elm
	}
	return m
}

func (pack *Package) loadTriggers(locale string) error {
	var jsonTrigger [](map[string]([]Trigger))
	content, err := util.ReadFileErr("package/" + pack.Name + "/res/locales/" + locale + "/triggers.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &jsonTrigger)
	if err != nil {
		return err
	}
	for _, elm := range jsonTrigger {
		for key, value := range elm {
			if f, err := pack.Plug.Lookup(key); err == nil {
				if mf, ok := f.(ModuleFunc); ok {
					pack.Modules[key] = Module{
						Triggers: tagTriggers(value, pack.Name, key, locale),
						Func:     mf,
					}
				}
			}
		}
	}
	return nil
}

func (pack *Package) loadResponses(locale string) error {
	var jsonResponse []Response
	content, err := util.ReadFileErr("package/" + pack.Name + "/res/locales/" + locale + "/response.json")
	if err != nil {
		return nil
	}
	err = json.Unmarshal(content, &jsonResponse)
	if err != nil {
		return nil
	}
	pack.Responses = jsonResponse
	return nil
}

func (pack *Package) loadModules(locale string) error {
	if err := pack.loadTriggers(locale); err != nil {
		return err
	}
	if err := pack.loadResponses(locale); err != nil {
		return err
	}
	return nil
}
