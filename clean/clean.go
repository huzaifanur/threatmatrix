package clean

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/forPelevin/gomoji" // Third-party package for emoji handling.
	"golang.org/x/text/unicode/norm"
)

// EncodeAndNormalizeAndDemojize takes a string and a language code, normalizes the string, removes URLs and emojis, and returns the processed string.
func NormalizeAndDemojize(text string) (string, error) {

	// Trim leading and trailing white spaces, new lines
	wordSlice := strings.Fields(text)
	text = strings.Join(wordSlice, " ")
	text = strings.TrimSpace(text)

	// Normalize the text using the previously defined function.
	text = NormalizeText(text)

	re := regexp.MustCompile(`http[s]?://[a-zA-Z0-9./-]*`) // Compile a regex to match URLs.
	text = re.ReplaceAllString(text, "")                   // Remove URLs from the text.

	text = strings.ToLower(text)
	text = RemoveEmojies(text) // Convert emojis to text using the RemoveEmojies function.

	return text, nil // Return the processed text. The error is always nil, consider handling or removing the error return value.
}

// NormalizeText ensures that strings with different Unicode representations but identical visual appearances are treated as equal.
// NFKD stands for "Normalization Form Compatibility Decomposition," a Unicode normalization form defined by the Unicode Consortium.
func NormalizeText(text string) string {
	n := norm.NFKD        // Create a Normalizer with NFKD normalization.
	return n.String(text) // Normalize the text and return it.
}

// RemoveEmojies uses the gomoji package to remove emojis from a string and return the result.
func RemoveEmojies(text string) string {
	return gomoji.RemoveEmojis(text)
}

// ReplaceEmojiWithSlug replaces emojis in a string with their corresponding slugs using the gomoji package.
func ReplaceEmojiWithSlug(text string) string {
	var result []string             // Initialize a slice to hold the result.
	var wordBuilder strings.Builder // Use a Builder to efficiently build strings.

	for len(text) > 0 {
		r, size := utf8.DecodeRuneInString(text) // Decode the next rune in the string.
		if unicode.IsSpace(r) {
			if wordBuilder.Len() > 0 {
				result = append(result, wordBuilder.String()) // Append the built word to the result.
				wordBuilder.Reset()                           // Reset the builder for the next word.
			}
			result = append(result, string(r)) // Append the space to the result.
		} else {
			info, err := gomoji.GetInfo(string(r)) // Get emoji info for the rune.
			if err == nil {
				if wordBuilder.Len() > 0 {
					result = append(result, wordBuilder.String()) // Append the built word to the result.
					wordBuilder.Reset()                           // Reset the builder for the next word.
				}
				result = append(result, ":", info.Slug) // Append the emoji slug to the result.
			} else {
				wordBuilder.WriteRune(r) // Add the rune to the current word.
			}
		}
		text = text[size:] // Move to the next rune in the string.
	}

	if wordBuilder.Len() > 0 {
		result = append(result, wordBuilder.String()) // Append the last word if there is one.
	}

	return strings.Join(result, "") // Join the result slice into a single string and return it.
}
