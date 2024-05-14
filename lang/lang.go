package lang

import (
	"errors"

	"github.com/pemistahl/lingua-go"
)

var langMap = map[string]string{
	"Arabic":         "ar",
	"Persian":        "fa",
	"Japanese":       "ja",
	"Korean":         "ko",
	"Russian":        "ru",
	"Chinese":        "zh",
	"Vietnamese":     "vi",
	"Hebrew":         "he",
	"English":        "en",
	"Spanish":        "es",
	"Croatian":       "hr",
	"Polish":         "pl",
	"Portuguese":     "pt",
	"Serbian":        "sr",
	"Turkish":        "tr",
	"French":         "fr",
	"Italian":        "it",
	"German":         "de",
	"Zulu":           "zu",
	"Slovak":         "sk",
	"Filipino":       "fil",
	"Afrikaans":      "af",
	"Dutch":          "nl",
	"Swedish":        "se",
	"Norwegian":      "no",
	"Haitian Creole": "ht",
	"Danish":         "da",
}

func GetLanguage(text string) (lang string, err error) {
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	language, exists := detector.DetectLanguageOf(text)
	if !exists {
		return "", errors.New("language detection failed")
	}

	if lang, ok := langMap[language.String()]; ok {
		return lang, nil
	}

	return "", errors.New("language not supported")
}
