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
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/drnp/go-xuan/common"
)

// Structs
type SemanticVariant struct {
	Target   string `json:"target"`
	Property string `json:"property"`
	Tag      string `json:"tag"`
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
		WuXing              string              `json:"wu_xing"`
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
	if h == nil {
		return ""
	}

	b, _ := json.MarshalIndent(h, "", "  ")

	return string(b)
}

func (h *Han) Readings(properties ...string) map[string]string {
	if h == nil {
		// Nothing
		return nil
	}

	r := h.Properties.Readings
	if r == nil {
		return nil
	}

	if len(properties) == 0 {
		// All
		return r
	}

	ret := make(map[string]string)
	for _, rp := range properties {
		if r[rp] != "" {
			ret[rp] = r[rp]
		}
	}

	return ret
}

func (h *Han) TotalStrokes() int {
	if h == nil {
		return 0
	}

	i := h.Properties.IRGSources
	if i == nil {
		return 0
	}

	ikts := i["kTotalStrokes"]
	if ikts == nil {
		return 0
	}

	ts, _ := strconv.Atoi(ikts[0])

	return ts
}

func (h *Han) WuXing() int {
	if h == nil {
		return 0
	}

	wx := 0

	switch strings.ToUpper(h.Properties.WuXing) {
	case "WOOD":
		wx = common.Wood
	case "FIRE":
		wx = common.Fire
	case "EARTH":
		wx = common.Earth
	case "METAL":
		wx = common.Metal
	case "WATER":
		wx = common.Water
	}

	return wx
}

func (h *Han) SemanticVariants() []*SemanticVariant {
	if h == nil {
		return nil
	}

	v := h.Properties.Variants
	if v == nil {
		return nil
	}

	vs := v["kSemanticVariant"]
	if vs == nil {
		return nil
	}

	ret := make([]*SemanticVariant, 0)
	rev := regexp.MustCompile(`(U\+[0-9A-F]+)<(.+)`)
	for _, vsv := range vs {
		sv := new(SemanticVariant)
		matches := rev.FindStringSubmatch(vsv)
		if len(matches) == 3 {
			// With tags
			tMatches := strings.Split(matches[2], ":")
			sv.Target = matches[1]
			sv.Property = tMatches[0]
			if len(tMatches) > 1 {
				sv.Tag = tMatches[1]
			}
		} else {
			// Only code
			sv.Target = vsv
		}

		ret = append(ret, sv)
	}

	return ret
}

func (h *Han) SimplifiedVariants() []string {
	if h == nil {
		return nil
	}

	v := h.Properties.Variants
	if v == nil {
		return nil
	}

	vs := v["kSimplifiedVariant"]

	return vs
}

func (h *Han) TraditionalVariants() []string {
	if h == nil {
		return nil
	}

	v := h.Properties.Variants
	if v == nil {
		return nil
	}

	vs := v["kTraditionalVariant"]

	return vs
}

// Helpers
func (h *Han) Pinyin() (string, []string) {
	if h == nil {
		return "", nil
	}

	v := h.Properties.Readings
	if v == nil {
		return "", nil
	}

	readings := make([]string, 0)

	// kMandarin
	readings = append(readings, strings.Fields(v["kMandarin"])...)

	// kTGHZ2013
	rTGHZs := strings.Fields(v["kTGHZ2013"])
	for _, rTGHZ := range rTGHZs {
		parts := strings.SplitN(rTGHZ, ":", 2)
		if len(parts) == 2 {
			readings = append(readings, parts[1])
		}
	}

	// kXHC1983
	rXHCs := strings.Fields(v["kXHC1983"])
	for _, rXHC := range rXHCs {
		parts := strings.SplitN(rXHC, ":", 2)
		if len(parts) == 2 {
			readings = append(readings, parts[1])
		}
	}

	// kHanyuPinyin
	// -- TODO --

	// kHanyuPinlu
	// -- TODO --

	if len(readings) == 0 {
		return "", nil
	}

	unqMap := make(map[string]bool)
	unq := make([]string, 0)
	for _, r := range readings {
		if unqMap[r] == true {
			continue
		}

		unq = append(unq, r)
		unqMap[r] = true
	}

	return readings[0], unq
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
