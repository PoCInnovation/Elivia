package analysis

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/PoCFrance/e/locales"
	"github.com/PoCFrance/e/plugins"

	"github.com/gookit/color"
	gocache "github.com/patrickmn/go-cache"
	"github.com/PoCFrance/e/network"
	"github.com/PoCFrance/e/util"
)

// A Sentence represents simply a sentence with its content as a string
type Sentence struct {
	Locale  string
	Content string
}

// Result contains a predicted value with its tag and its value
type Result struct {
	Tag   string
	Value float64
}

var userCache = gocache.New(5*time.Minute, 5*time.Minute)

// DontUnderstand contains the tag for the don't understand messages
const DontUnderstand = "don't understand"

// NewSentence returns a Sentence object where the content has been arranged
func NewSentence(locale, content string) (sentence Sentence) {
	sentence = Sentence{
		Locale:  locale,
		Content: content,
	}
	sentence.arrange()

	return
}

// PredictTag classifies the sentence with the model
func (sentence Sentence) PredictTag(neuralNetwork network.Network) string {
	words, classes, _ := Organize(sentence.Locale)

	// Predict with the model
	predict := neuralNetwork.Predict(sentence.WordsBag(words))

	// Enumerate the results with the intent tags
	var resultsTag []Result
	for i, result := range predict {
		resultsTag = append(resultsTag, Result{classes[i], result})
	}

	// Sort the results in ascending order
	sort.Slice(resultsTag, func(i, j int) bool {
		return resultsTag[i].Value > resultsTag[j].Value
	})

	LogResults(sentence.Locale, sentence.Content, resultsTag)

	return resultsTag[0].Tag
}

// RandomizeResponse takes the entry message, the response tag and the token and returns a random
// message from res/datasets/intents.json where the triggers are applied
func RandomizeResponse(locale, entry, tag, token string) plugins.MData {
	var md plugins.MData
	if tag == DontUnderstand {
		return md.Init(DontUnderstand, util.GetMessage(locale, tag))
	}

	for _, pack := range plugins.GetPackage(locale) {
		for _, resp := range pack.IO.Response {
			if resp.Tag != tag {
				continue
			}

			// Reply a "don't understand" message if the context isn't correct

			// TODO understand this
			// cacheTag, _ := userCache.Get(token)
			// if intent.Context != "" && cacheTag != intent.Context {
			// 	return md.Init(DontUnderstand, util.GetMessage(locale, DontUnderstand))
			// }

			// Set the actual context
			// userCache.Set(token, tag, gocache.DefaultExpiration)

			// Choose a random response in intents
			response := ""
			len := len(resp.Messages)
			if len > 1 {
				rand.Seed(time.Now().UnixNano())
				response = resp.Messages[rand.Intn(len)]
			}

			// And then apply the triggers on the message
			return plugins.ReplaceContent(locale, tag, entry, response, token)
		}
	}

	return md.Init(DontUnderstand, util.GetMessage(locale, DontUnderstand))
}

// Calculate send the sentence content to the neural network and returns a response with the matching tag
func (sentence Sentence) Calculate(_ gocache.Cache, neuralNetwork network.Network, token string) plugins.MData {
	// tag, found := cache.Get(sentence.Content)
	// Todo check if caching is still possible once reformat is done

	tag := sentence.PredictTag(neuralNetwork)

	// Predict tag with the neural network if the sentence isn't in the cache
	// if !found {
	// cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	// }

	return RandomizeResponse(sentence.Locale, sentence.Content, tag, token)
}

// LogResults print in the console the sentence and its tags sorted by prediction
func LogResults(locale, entry string, results []Result) {
	// If NO_LOGS is present, then don't print the given messages
	if os.Getenv("NO_LOGS") == "1" {
		return
	}

	green := color.FgGreen.Render
	yellow := color.FgYellow.Render

	fmt.Printf(
		"\n“%s” - %s\n",
		color.FgCyan.Render(entry),
		color.FgRed.Render(locales.GetNameByTag(locale)),
	)
	for _, result := range results {
		// Arbitrary choice of 0.004 to have less tags to show
		if result.Value < 0.004 {
			continue
		}

		fmt.Printf("  %s %s - %s\n", green("▫︎"), result.Tag, yellow(result.Value))
	}
}
