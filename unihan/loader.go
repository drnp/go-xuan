/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2025 BS.Group
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file loader.go
 * @package unihan
 * @author Dr.NP <np@herewe.tech>
 * @since 04/01/2025
 */

package unihan

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Unihan database files
const (
	DictionaryIndices   = "Unihan_DictionaryIndices.txt"
	DictionaryLikeData  = "Unihan_DictionaryLikeData.txt"
	IRGSources          = "Unihan_IRGSources.txt"
	NumericValues       = "Unihan_NumericValues.txt"
	OtherMappings       = "Unihan_OtherMappings.txt"
	RadicalStrokeCounts = "Unihan_RadicalStrokeCounts.txt"
	Readings            = "Unihan_Readings.txt"
	Variants            = "Unihan_Variants.txt"
)

// Load unihan database from source files
func Load(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		// File path failed
		return err
	}

	path = filepath.Clean(abs)

	err = loadDictionaryIndices(path + "/" + DictionaryIndices)
	if err != nil {
		return err
	}

	err = loadDictionaryLikeData(path + "/" + DictionaryLikeData)
	if err != nil {
		return err
	}

	err = loadIRGSources(path + "/" + IRGSources)
	if err != nil {
		return err
	}

	err = loadNumericValues(path + "/" + NumericValues)
	if err != nil {
		return err
	}

	err = loadOtherMappings(path + "/" + OtherMappings)
	if err != nil {
		return err
	}

	err = loadRadicalStrokeCounts(path + "/" + RadicalStrokeCounts)
	if err != nil {
		return err
	}

	err = loadReadings(path + "/" + Readings)
	if err != nil {
		return err
	}

	err = loadVariants(path + "/" + Variants)
	if err != nil {
		return err
	}

	return nil
}

func loadDictionaryIndices(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.DictionaryIndices == nil {
					Database[codePoint].Properties.DictionaryIndices = make(map[string][]string)
				}

				Database[codePoint].Properties.DictionaryIndices[matches[2]] = append(Database[codePoint].Properties.DictionaryIndices[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadDictionaryLikeData(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.DictionaryLikeData == nil {
					Database[codePoint].Properties.DictionaryLikeData = make(map[string][]string)
				}

				Database[codePoint].Properties.DictionaryLikeData[matches[2]] = append(Database[codePoint].Properties.DictionaryLikeData[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadIRGSources(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.IRGSources == nil {
					Database[codePoint].Properties.IRGSources = make(map[string][]string)
				}

				Database[codePoint].Properties.IRGSources[matches[2]] = append(Database[codePoint].Properties.IRGSources[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	return nil
}

func loadNumericValues(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.NumericValues == nil {
					Database[codePoint].Properties.NumericValues = make(map[string][]string)
				}

				Database[codePoint].Properties.NumericValues[matches[2]] = append(Database[codePoint].Properties.NumericValues[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadOtherMappings(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.OtherMappings == nil {
					Database[codePoint].Properties.OtherMappings = make(map[string][]string)
				}

				Database[codePoint].Properties.OtherMappings[matches[2]] = append(Database[codePoint].Properties.OtherMappings[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadRadicalStrokeCounts(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.RadicalStrokeCounts == nil {
					Database[codePoint].Properties.RadicalStrokeCounts = make(map[string][]string)
				}

				Database[codePoint].Properties.RadicalStrokeCounts[matches[2]] = append(Database[codePoint].Properties.RadicalStrokeCounts[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadReadings(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.Readings == nil {
					Database[codePoint].Properties.Readings = make(map[string]string)
				}

				Database[codePoint].Properties.Readings[matches[2]] = matches[3]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func loadVariants(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments out
		if strings.HasPrefix(line, "#") {
			continue
		}

		reLine := regexp.MustCompile(`(U\+.+?)\s+(k.+?)\s+(.+)`)
		matches := reLine.FindStringSubmatch(line)
		if len(matches) == 4 {
			codePoint := UnicodeToRune(matches[1])
			if codePoint > 0 {
				// Create new item
				if Database[codePoint] == nil {
					Database[codePoint] = &Han{
						CodePoint: codePoint,
						Unicode:   matches[1],
						Value:     string(codePoint),
					}
				}

				if Database[codePoint].Properties.Variants == nil {
					Database[codePoint].Properties.Variants = make(map[string][]string)
				}

				Database[codePoint].Properties.Variants[matches[2]] = append(Database[codePoint].Properties.Variants[matches[2]], strings.Fields(matches[3])...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
