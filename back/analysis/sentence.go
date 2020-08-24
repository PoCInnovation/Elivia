package analysis

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/PoCFrance/e/myutil"

	"github.com/PoCFrance/e/locales"
	"github.com/PoCFrance/e/plugins"

	"github.com/PoCFrance/e/network"
	"github.com/PoCFrance/e/util"
	"github.com/gookit/color"
	gocache "github.com/patrickmn/go-cache"
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

func randomizeResponse(p int, tag, locale string) (r plugins.Response) {
	for _, res := range plugins.GetPackage(locale)[p].IO.Response {
		if res.Tag == tag {
			len := len(res.Messages)
			if len >= 1 {
				rand.Seed(time.Now().UnixNano())
				response := res.Messages[rand.Intn(len)]
				return r.Init(tag, response)
			}
			break
		}
	}
	return r.Init(DontUnderstand, util.GetMessage(locale, DontUnderstand))
}

func seekModule(tag, locale string) (int, int, error) {
	splited := strings.Split(tag, "_")
	if len(splited) != 3 {
		return 0, 0, errors.New("invalid tag")
	}
	for packi, pack := range plugins.GetPackage(locale) {
		if pack.Name == splited[0] {
			for triggeri, t := range pack.IO.Triggers {
				if t.CallBack != splited[2] {
					continue
				}
				return packi, triggeri, nil
			}
		}
	}

	return 0, 0, errors.New("module not found")
}

type Module (func(string, map[string]string) (string, map[string]interface{}))

func retrieveModule(p, t int, locale string) (Module, error) {
	plug := plugins.GetPackage(locale)[p]
	symb, err := plug.Plug.Lookup(plug.IO.Triggers[t].CallBack)
	if err != nil {
		return nil, err
	}
	module, ok := symb.(func(string, map[string]string) (string, map[string]interface{}))
	_, _ = symb.(Module)
	if !ok {
		return nil, errors.New("Module has invalid format")
	}
	return module, nil
}

func (sentence Sentence) extractEntries(entries []myutil.Entries) map[string]string {
	return myutil.ExtractEntries(entries, sentence.Content)
}

// Calculate send the sentence content to the neural network and returns a response with the matching tag
func (sentence Sentence) Calculate(_ gocache.Cache, neuralNetwork network.Network, token string) (r plugins.Response) {
	// tag, found := cache.Get(sentence.Content)
	// Todo check if caching is still possible once reformat is done
	// Predict tag with the neural network if the sentence isn't in the cache
	// if !found {
	// cache.Set(sentence.Content, tag, gocache.DefaultExpiration)
	// }

	tag := sentence.PredictTag(neuralNetwork)
	p, t, err := seekModule(tag, sentence.Locale)
	if err != nil {
		fmt.Println(err)
		return r.Init(DontUnderstand, util.GetMessage(sentence.Locale, DontUnderstand))
	}

	module, err := retrieveModule(p, t, sentence.Locale)
	if err != nil {
		fmt.Println(err)
		return r.Init(DontUnderstand, util.GetMessage(sentence.Locale, DontUnderstand))
	}

	responseTag, json := module(sentence.Locale, sentence.extractEntries(plugins.GetPackage(sentence.Locale)[p].IO.Triggers[t].Entries))
	r = randomizeResponse(p, responseTag, sentence.Locale)
	r.AppendData(json)
	return r.Format()
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
