package threatflag

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"threatmatrix/consts"
)

type cvsMap []map[string]string

type cvsMapList []cvsMap

func LoadWordlists(lang string) (cvsMapList, error) {
	fileroot := "tm_wordlists/"
	if !slices.Contains(consts.SupportedLanguages, lang) {
		return nil, fmt.Errorf("language not supported")
	}

	var wordlists cvsMapList

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

			var records cvsMap
			for {
				row, err := reader.Read()
				if err != nil {
					break
				}

				// Create a map for each row
				rowMap := make(map[string]string)
				for i, value := range row {
					rowMap[strings.ToLower(header[i])] = strings.ToLower(value)
				}

				// Append the map to the slice
				records = append(records, rowMap)
			}

			file.Close() // Ensure the file is closed after processing

			wordlists = append(wordlists, records)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return wordlists, nil
}
