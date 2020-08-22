package plugins

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PoCFrance/e/util"
)

type Predicat func(string, map[string]string) (string, map[string]interface{})

// MData represent a Json object that will be returned by the Modules
type MData struct {
	Tag      string                 `json:"tag"`
	Response string                 `json:"Response"`
	Data     map[string]interface{} `json:"data"`
}

// Init is a way to simulate the MData constructor
func (md *MData) Init(Tag, Response string, data ...map[string]interface{}) MData {
	md.Tag = Tag
	md.Response = Response
	md.Data = make(map[string]interface{})
	for _, elm := range data {
		for key, value := range elm {
			md.Data[key] = value
		}
	}
	return *md
}

// Module is a structure for dynamic intents with a Tag, some Patterns and Responses and
// a Replacer function to execute the dynamic changes.
type Module struct {
	Tag       string
	Patterns  []string
	Responses []string
	Replacer  interface{}
	Context   string
}

var modules = map[string][]Module{}

// RegisterModule registers a module into the map
func RegisterModule(locale string, module Module) {
	modules[locale] = append(modules[locale], module)
}

// RegisterModules registers an array of modules into the map
func RegisterModules(locale string, _modules []Module) {
	modules[locale] = append(modules[locale], _modules...)
}

// GetModules returns all the registered modules
func GetModules(locale string) []Module {
	return modules[locale]
}

// GetModuleByTag returns a module found by the given tag and locale
func GetModuleByTag(tag, locale string) Module {
	for _, module := range modules[locale] {
		if tag != module.Tag {
			continue
		}

		return module
	}

	return Module{}
}

// CallPredicat -
func CallPredicat(pack Package, callback, locale string, entries map[string]string) (md MData) {
	module, err := pack.Plug.Lookup(callback)
	if err != nil {
		md.Init("don't understand", util.GetMessage(locale, "don't understand"))
		return
	}

	tag, json := module.(func(string, map[string]string) (string, map[string]interface{}))(locale, entries)

	for _, r := range pack.IO.Response {
		fmt.Println("TAGS :", r.Tag, tag)
		if r.Tag == tag {
			len := len(r.Messages)
			if len >= 1 {
				rand.Seed(time.Now().UnixNano())
				response := r.Messages[rand.Intn(len)]
				md.Init(tag, response, json)
				return
			}
		}
	}
	md.Init("don't understand", util.GetMessage(locale, "don't understand"))
	return
}
