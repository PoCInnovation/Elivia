package analysis

import (
	"sort"

	"github.com/PoCFrance/e/plugins"
	"github.com/PoCFrance/e/util"
)

// IntentTagFormat formats the tag associated with a function of the package
func IntentTagFormat(packName, locale, predicat string) string {
	return packName + "_" + locale + "_" + predicat
}

// Organize intents with an array of all words, an array with a representative word of each tag
// and an array of Documents which contains a word list associated with a tag
func Organize(locale string) (words, classes []string, documents []Document) {
	plugins := plugins.GetPackage(locale)

	for _, pack := range plugins {
		for _, trigger := range pack.IO.Triggers {
			intent := IntentTagFormat(pack.Name, locale, trigger.CallBack)
			for _, pattern := range trigger.Patterns {
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
					intent,
				})
			}
			// Add the intent tag to classes
			classes = append(classes, intent)
		}
	}

	sort.Strings(words)
	sort.Strings(classes)

	return words, classes, documents
}
