package analysis

import (
	"sort"

	"github.com/PoCInnovation/Elivia/plugins"
	"github.com/PoCInnovation/Elivia/util"
)

// IntentTagFormat formats the tag associated with a function of the package
func IntentTagFormat(packName, locale, predicat string) string {
	return packName + "_" + locale + "_" + predicat
}

// Organize intents with an array of all words, an array with a representative word of each tag
// and an array of Documents which contains a word list associated with a tag
func Organize(locale string) (words, classes []string, documents []Document) {
	packages := plugins.GetPackage(locale)

	for _, pack := range packages {
		for _, module := range pack.Modules {
			for tag, trigger := range module.Triggers {
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
						tag,
					})
				}
				// Add the intent tag to classes
				classes = append(classes, tag)
			}
		}
	}

	sort.Strings(words)
	sort.Strings(classes)

	return words, classes, documents
}
