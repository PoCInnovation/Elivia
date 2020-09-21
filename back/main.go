package main

// import (
// 	"fmt"

// 	"github.com/PoCInnovation/Elivia/myutil"
// 	"github.com/PoCInnovation/Elivia/plugins"
// )

// func main() {
// 	packages := plugins.LoadPlugins()

// 	for _, pack := range packages {
// 		for _, r := range pack.IO.Response {
// 			fmt.Println(pack.Name, "response > ", r.Tag, r.Messages)
// 		}
// 		fmt.Println()
// 		for _, r := range pack.IO.Triggers {
// 			fmt.Println(pack.Name, "triggers > ", r.Patterns, r.CallBack)

// 			module, err := pack.Plug.Lookup(r.CallBack)
// 			if err != nil {
// 				panic(err)
// 			}

// 			//test
// 			fmt.Println()
// 			sentence := "send \"hello dad how are you ?\" to dad"
// 			if r.OliviaPatterns != nil && len(r.OliviaPatterns) >= 1 {
// 				module.(func(map[string]string))(myutil.ExtractEntries(r.Entries, sentence))
// 			}

// 		}
// 	}
// }

import (
	"flag"
	"fmt"
	"os"

	"github.com/PoCInnovation/Elivia/locales"
	"github.com/PoCInnovation/Elivia/plugins"
	"github.com/PoCInnovation/Elivia/training"

	"github.com/PoCInnovation/Elivia/util"

	"github.com/gookit/color"

	"github.com/PoCInnovation/Elivia/network"

	"github.com/PoCInnovation/Elivia/server"
)

var neuralNetworks = map[string]network.Network{}

func main() {
	port := flag.String("port", "8080", "The port for the API and WebSocket.")
	//localesFlag := flag.String("re-train", "", "The locale(s) to re-train.")
	flag.Parse()

	// Print the Olivia ascii text
	oliviaASCII := string(util.ReadFile("res/olivia-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(oliviaASCII))

	reTrainModels()

	for _, locale := range locales.Locales {
		plugins.LoadPackage(locale.Tag)
		n, err := training.CreateNeuralNetwork(
			locale.Tag,
			false,
		)
		if err != nil {
			continue
		}
		neuralNetworks[locale.Tag] = n
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetworks, *port)
}

// reTrainModels retrain the given locales
func reTrainModels() {
	for _, locale := range locales.Locales {
		path := fmt.Sprintf("res/locales/%s/training.json", locale.Tag)
		os.Remove(path)
	}
}
