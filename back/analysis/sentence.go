package analysis

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/PoCInnovation/Elivia/locales"
	"github.com/PoCInnovation/Elivia/plugins"
	"github.com/PoCInnovation/Elivia/plugins/bridge"

	"github.com/PoCInnovation/Elivia/network"
	"github.com/PoCInnovation/Elivia/util"
	"github.com/gookit/color"
	gocache "github.com/patrickmn/go-cache"
)

// Document is any sentence from the intents' patterns linked with its tag
type Document struct {
	Sentence Sentence
	Tag      string
}

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

func randomizeResponse(pack plugins.Package, tag, locale string) (r bridge.Response) {
	for _, reponse := range pack.Responses {
		if reponse.Tag != tag {
			continue
		}
		len := len(reponse.Messages)
		if len >= 1 {
			rand.Seed(time.Now().UnixNano())
			response := reponse.Messages[rand.Intn(len)]
			return r.Init(pack.Name, response)
		}
	}
	return r.Init(DontUnderstand, util.GetMessage(locale, DontUnderstand))
}

func (sentence Sentence) extractEntries(entries []plugins.Entries) map[string]string {
	return bridge.ExtractEntries(entries, sentence.Content)
}

// Calculate send the sentence content to the neural network and returns a response with the matching tag
func (sentence Sentence) Calculate(_ gocache.Cache, neuralNetwork network.Network, token string) (r bridge.Response) {
	// tag, found := cache.Get(sentence.Content)
	// Todo check if caching is still possible once reformat is done
	// Predict tag with the neural network if the sentence isn't in the cache
	// if !found {
	// cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	// }

	tag := sentence.PredictTag(neuralNetwork)
	splited := strings.Split(tag, "_")
	if len(splited) != 4 {
		return r.Init(DontUnderstand, util.GetMessage(sentence.Locale, DontUnderstand))
	}

	for _, pack := range plugins.GetPackage(sentence.Locale) {
		if pack.Name != splited[0] {
			continue
		}
		module, ok := pack.Modules[splited[2]]
		if !ok {
			continue
		}
		responseTag, json := module.Func(sentence.Locale, sentence.extractEntries(module.Triggers[tag].Entries))
		r = randomizeResponse(pack, responseTag, sentence.Locale)
		r.AppendData(json)
		return r.Format()
	}
	return r.Init(DontUnderstand, util.GetMessage(sentence.Locale, DontUnderstand))
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
