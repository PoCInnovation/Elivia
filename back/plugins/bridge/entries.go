package bridge

import (
	"fmt"

	"github.com/PoCInnovation/Elivia/metatools"
	"github.com/PoCInnovation/Elivia/plugins"
)

// ExtractEntries finds the corresponding text in the field
func ExtractEntries(parserInfo []plugins.Entries, trigger string) map[string]string {
	entries := make(map[string]string)

	var parser metatools.Parser
	parser.Init(trigger)

	fmt.Println("extracting extries for sentence : ", trigger)
	for _, e := range parserInfo {
		res := e.Resources.(map[string]interface{})
		switch e.Parser {
		case "before":
			entries[e.Name] = parser.Before(res["key"].(string), int(res["x"].(float64)))
		case "after":
			entries[e.Name] = parser.After(res["key"].(string), int(res["x"].(float64)))
		case "between":
			entries[e.Name] = parser.Between(res["after"].(string), res["before"].(string), int(res["x"].(float64)), int(res["y"].(float64)))
		}
	}
	return entries
}
