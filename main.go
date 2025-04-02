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
 * @file main.go
 * @package main
 * @author Dr.NP <np@herewe.tech>
 * @since 04/01/2025
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/drnp/go-xuan/unihan"
)

func main() {
	err := unihan.Load("./data/Unihan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Load unihan database success. Total character :", unihan.CountDatabase())
	han := unihan.GetHanByValue("发") // 发
	fmt.Println(han.Dump())
	fmt.Println(han.TotalStrokes())
	fmt.Println(han.WuXing())
	fmt.Println(han.Pinyin())

	b, _ := json.MarshalIndent(han.SemanticVariants(), "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(han.SimplifiedVariants(), "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(han.TraditionalVariants(), "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(han.Readings(), "", "  ")
	fmt.Println(string(b))

	try := "我爱北京天安门，老板来个手抓饼。周潤發踢足球！！！"
	fmt.Println(try)
	fmt.Println(unihan.StrToSimplified(try))
	fmt.Println(unihan.StrToTraditional(try))
	fmt.Println(unihan.StrToPinyin(try))
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
