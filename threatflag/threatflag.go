package threatflag

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"threatmatrix/consts"
)

// func threatFlag(lang string, text string) (s string, p string) {
// 	//  CHECK IF IN LANGUAGES SUPPORTED
// 	if !(slices.Contains(consts.IsoLanguages, lang) || slices.Contains(consts.Utf8Languages, lang)) {
// 		return "", ""
// 	}

// 	// # produce strings of ngrams
// 	tokenizer := tokenizer.NewTweetTokenizer()
// 	text_tokens := tokenizer.Tokenize(text)

// 	// # FIND MAX NGRAM SIZE FROM BAD WORD
// 	max_tokens := 908

// 	for i := 1; i <= max_tokens; i++ {

// 	}
// 	return
// }

func LoadWordlists(lang string) ([]map[string][]string, error) {
	fileroot := "tm_wordlists/"
	if !slices.Contains(consts.SupportedLanguages, lang) {
		return nil, fmt.Errorf("language not supported")
	}

	var wordlists []map[string][]string

	err := filepath.Walk(fileroot, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !info.IsDir() && strings.Contains(info.Name(), "_"+lang) {
			file, openErr := os.Open(path)
			if openErr != nil {
				return openErr
			}

			reader := csv.NewReader(file)
			header, readErr := reader.Read()
			if readErr != nil {
				file.Close() // Close the file if header read fails
				return readErr
			}

			dataMap := make(map[string][]string)
			for {
				row, readErr := reader.Read()
				if readErr == io.EOF {
					break // End of file reached
				}
				if readErr != nil {
					file.Close() // Close the file if row read fails
					return readErr
				}

				for i, value := range row {
					columnName := strings.ToLower(header[i])
					dataMap[columnName] = append(dataMap[columnName], value)
				}
			}
			file.Close() // Ensure the file is closed after processing

			wordlists = append(wordlists, dataMap)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return wordlists, nil
}
