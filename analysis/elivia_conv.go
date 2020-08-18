package analysis

import (
	"sort"

	"github.com/PoCFrance/e/plugins"
	"github.com/PoCFrance/e/util"
)

func IntentTagFormat(locale string, pack_i, trig_i int) string {
	return plugins.GetPackage(locale)[pack_i].Name + "_" + locale + "_" + plugins.GetPackage(locale)[pack_i].IO.Triggers[trig_i].CallBack
}

// IntentOlivia - conv plugin of elivia to olivia's form.
func IntentOlivia(locale string) []Intent {
	var intents []Intent

	for pi, pack := range plugins.GetPackage(locale) {
		for ti, trigger := range pack.IO.Triggers {
			intents = append(intents, Intent{
				Tag:       IntentTagFormat(locale, pi, ti),
				Patterns:  trigger.OliviaPatterns,
				Responses: []string{},
				Context:   ""})
		}
	}
	return intents
}

// Organize intents with an array of all words, an array with a representative word of each tag
// and an array of Documents which contains a word list associated with a tag
func Organize(locale string) (words, classes []string, documents []Document) {
	intents := IntentOlivia(locale)

	for _, intent := range intents {
		for _, pattern := range intent.Patterns {
			// Tokenize the pattern's sentence
			patternSentence := Sentence{locale, pattern}
			patternSentence.arrange()

			// Add each word to response
			for _, word := range patternSentence.stem() {

				if !util.Contains(words, word) {
					words = append(words, word)
				}
			}

			// Add a new document
			documents = append(documents, Document{
				patternSentence,
				intent.Tag,
			})
		}

		// Add the intent tag to classes
		classes = append(classes, intent.Tag)
	}

	sort.Strings(words)
	sort.Strings(classes)

	return words, classes, documents
}
