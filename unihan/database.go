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
 * @file database.go
 * @package unihan
 * @author Dr.NP <np@herewe.tech>
 * @since 04/01/2025
 */

package unihan

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Structs
type Dictionary struct {
	Name string `json:"name"`
}

type Han struct {
	CodePoint  rune   `json:"code_point"`
	Unicode    string `json:"unicode"`
	Value      string `json:"value"`
	Properties struct {
		IRGSources          map[string][]string `json:"IRGSources"`
		OtherMappings       map[string][]string `json:"OtherMappings"`
		DictionaryIndices   map[string][]string `json:"DictionaryIndices"`
		Readings            map[string]string   `json:"readings"`
		DictionaryLikeData  map[string][]string `json:"DictionaryLikeData"`
		RadicalStrokeCounts map[string][]string `json:"RadicalStrokeCounts"`
		Variants            map[string][]string `json:"variants"`
		NumericValues       map[string][]string `json:"numeric_values"`
	} `json:"properties"`
}

var (
	Database     = make(map[rune]*Han)
	DatabaseLock sync.RWMutex
)

func DumpDatabase() {
	DatabaseLock.Lock()
	defer DatabaseLock.Unlock()

	b, _ := json.MarshalIndent(Database, "", "  ")

	fmt.Println(string(b))
}

func CountDatabase() int {
	DatabaseLock.Lock()
	defer DatabaseLock.Unlock()

	return len(Database)
}

/* {{{ [Han struct] */
func (h *Han) Dump() string {
	b, _ := json.MarshalIndent(h, "", "  ")

	return string(b)
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
