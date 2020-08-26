package training

import (
	"errors"
	"fmt"
	"os"

	"github.com/PoCFrance/e/analysis"
	"github.com/PoCFrance/e/network"
	"github.com/PoCFrance/e/util"
	"github.com/gookit/color"
)

// TrainData returns the inputs and outputs for the neural network
func TrainData(locale string) (inputs, outputs [][]float64) {
	words, classes, documents := analysis.Organize(locale)

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[util.Index(classes, document.Tag)] = 1

		// Append data to inputs and outputs
		inputs = append(inputs, bag)
		outputs = append(outputs, outputRow)
	}
	return inputs, outputs
}

// CreateNeuralNetwork returns a new neural network which is loaded from res/training.json or
// trained from TrainData() inputs and targets.
func CreateNeuralNetwork(locale string, ignoreTrainingFile bool) (network.Network, error) {
	// Decide if the network is created by the save or is a new one
	saveFile := "res/locales/" + locale + "/training.json"
	neuralNetwork := network.Network{}
	f, err := os.Open(saveFile)
	// Train the model if there is no training file
	if err != nil || ignoreTrainingFile {
		inputs, outputs := TrainData(locale)
		if len(inputs) == 0 || len(outputs) == 0 {
			return neuralNetwork, errors.New("input or output of layer is empty")
		}
		neuralNetwork = network.CreateNetwork(locale, 0.1, inputs, outputs, 50)
		neuralNetwork.Train(200)

		// Save the neural network in res/training.json
		neuralNetwork.Save(saveFile)
	} else {
		defer f.Close()
		fmt.Printf(
			"%s %s\n",
			color.FgBlue.Render("Loading the neural network from"),
			color.FgRed.Render(saveFile),
		)
		// Initialize the intents
		neuralNetwork = *network.LoadNetwork(saveFile)
	}
	return neuralNetwork, nil
}
