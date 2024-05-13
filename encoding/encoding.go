package encoding

import (
	// Importing necessary packages for encoding, decoding, and string manipulation.
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/forPelevin/gomoji" // Third-party package for emoji handling.
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// NormalizeText ensures that strings with different Unicode representations but identical visual appearances are treated as equal.
// NFKD stands for "Normalization Form Compatibility Decomposition," a Unicode normalization form defined by the Unicode Consortium.
func NormalizeText(text string) string {
	n := norm.NFKD        // Create a Normalizer with NFKD normalization.
	return n.String(text) // Normalize the text and return it.
}

// EncodeAndNormalizeAndDemojize takes a string and a language code, normalizes the string, removes URLs and emojis, and returns the processed string.
func EncodeAndNormalizeAndDemojize(text string, lang string) (string, error) {
	text = strings.ReplaceAll(text, "\n", " ") // Replace newlines with spaces.
	text = NormalizeText(text)                 // Normalize the text using the previously defined function.

	re := regexp.MustCompile(`http[s]?://[a-zA-Z0-9./-]*`) // Compile a regex to match URLs.
	text = re.ReplaceAllString(text, "")                   // Remove URLs from the text.

	text = strings.ToLower(text)   // Convert text to lowercase.
	text = strings.TrimSpace(text) // Trim leading and trailing white spaces.
	text = RemoveEmojies(text)     // Convert emojis to text using the RemoveEmojies function.

	return text, nil // Return the processed text. The error is always nil, consider handling or removing the error return value.
}

// EncodeDecodeISO8859_1 encodes a string to ISO8859-1 and then decodes it back, returning the decoded string or an error if encoding fails.
func EncodeDecodeISO8859_1(str string) (string, error) {
	encoded := make([]byte, 0) // Initialize an empty byte slice for the encoded data.
	for _, r := range str {
		if e, ok := charmap.ISO8859_1.EncodeRune(r); ok {
			encoded = append(encoded, e) // Append the encoded rune to the slice.
		} else {
			return "", fmt.Errorf("rune %q cannot be encoded in ISO8859-1", r) // Handle unencodable runes.
		}
	}

	decoder := charmap.ISO8859_1.NewDecoder()                                               // Create a new decoder for ISO8859-1.
	decodedBytes, err := io.ReadAll(transform.NewReader(bytes.NewReader(encoded), decoder)) // Decode the encoded bytes.
	if err != nil {
		return "", err // Handle decoding errors.
	}

	return string(decodedBytes), nil // Return the decoded string.
}

// EncodeDecodeUTF16 encodes a string to UTF-16 and then decodes it back, returning the decoded string or an error if the process fails.
func EncodeDecodeUTF16(str string) (string, error) {
	encoded := utf16.Encode([]rune(str)) // Encode the string to UTF-16.
	buf := new(bytes.Buffer)             // Create a new buffer to hold the binary data.
	for _, r := range encoded {
		if err := binary.Write(buf, binary.BigEndian, r); err != nil {
			return "", fmt.Errorf("binary write failed: %v", err) // Handle binary write errors.
		}
	}

	decoded := make([]uint16, buf.Len()/2) // Create a slice to hold the decoded data.
	if err := binary.Read(buf, binary.BigEndian, &decoded); err != nil {
		return "", fmt.Errorf("binary read failed: %v", err) // Handle binary read errors.
	}

	return string(utf16.Decode(decoded)), nil // Return the decoded string.
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
