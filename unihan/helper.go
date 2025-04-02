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
 * @file helper.go
 * @package unihan
 * @author Dr.NP <np@herewe.tech>
 * @since 04/01/2025
 */

package unihan

import "strings"

func StrToSimplified(input string) string {
	var sb strings.Builder
	for _, r := range input {
		han := GetHanByCodePoint(r)
		if han != nil {
			ss := han.SimplifiedVariants()
			if len(ss) > 0 {
				sHan := GetHanByUnicode(ss[0])
				if sHan != nil {
					sb.WriteRune(sHan.CodePoint)
				} else {
					sb.WriteRune(han.CodePoint)
				}
			} else {
				sb.WriteRune(han.CodePoint)
			}
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func StrToTraditional(input string) string {
	var sb strings.Builder
	for _, r := range input {
		han := GetHanByCodePoint(r)
		if han != nil {
			ts := han.TraditionalVariants()
			if len(ts) > 0 {
				tHan := GetHanByUnicode(ts[0])
				if tHan != nil {
					sb.WriteRune(tHan.CodePoint)
				} else {
					sb.WriteRune(han.CodePoint)
				}
			} else {
				sb.WriteRune(han.CodePoint)
			}
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func StrToPinyin(input string) string {
	var sb strings.Builder
	for _, r := range input {
		han := GetHanByCodePoint(r)
		if han != nil {
			rPinyin, _ := han.Pinyin()
			sb.WriteString(rPinyin)
		}
	}

	return sb.String()
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
