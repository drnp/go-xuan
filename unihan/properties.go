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
 * @file properties.go
 * @package unihan
 * @author Dr.NP <np@herewe.tech>
 * @since 04/02/2025
 */

package unihan

// DictionaryIndices
const (
	kCheungBauerIndexw = iota
	kCihaiT
	kCowles
	kDaeJaweon
	kFennIndex
	kGSR
	kHanYu
	kIRGDaeJaweon
	kIRGHanyuDaZidian
	kIRGKangXi
	kKangXi
	kKarlgren
	kLau
	kMatthews
	kMeyerWempe
	kMorohashi
	kNelson
	kSBGY
	kSMSZD2003Index
)

// DictionaryLikeData
const (
	kAlternateTotalStrokes = iota
	kCangjie
	kCheungBauer
	kFenn
	kFourCornerCode
	kGradeLevel
	kHDZRadBreak
	kHKGlyph
	kMojiJoho
	kPhonetic
	kStrange
	kUnihanCore2020
)

// IRGSources
const (
	kCompatibilityVariant = iota
	kIICore
	kIRG_GSource
	kIRG_HSource
	kIRG_JSource
	kIRG_KPSource
	kIRG_KSource
	kIRG_MSource
	kIRG_SSource
	kIRG_TSource
	kIRG_UKSource
	kIRG_USource
	kIRG_VSource
	kRSUnicode
	kTotalStrokes
)

// NumericValues
const (
	kAccountingNumeric = iota
	kOtherNumeric
	kPrimaryNumeric
	kVietnameseNumeric
	kZhuangNumeric
)

// OtherMappings
const (
	kBigFive = iota
	kCCCII
	kCNS1986
	kCNS1992
	kEACC
	kGB0
	kGB1
	kGB3
	kGB5
	kGB7
	kGB8
	kIBMJapan
	kJa
	kJinmeiyoKanji
	kJis0
	kJis1
	kJIS0213
	kJoyoKanji
	kKoreanEducationHanja
	kKoreanName
	kMainlandTelegraph
	kPseudoGB1
	kTaiwanTelegraph
	kTGH
	kXerox
)

// RadicalStrokeCounts
const (
	kRSAdobe_Japan1_6 = iota
)

// Readings
const (
	kCantonese = iota
	kDefinition
	kFanqie
	kHangul
	kHanyuPinlu
	kHanyuPinyin
	kJapanese
	kJapaneseKun
	kJapaneseOn
	kKorean
	kMandarin
	kSMSZD2003Readings
	kTang
	kTGHZ2013
	kVietnamese
	kXHC1983
	kZhuang
)

// Variants
const (
	kSemanticVariant = iota
	kSimplifiedVariant
	kSpecializedSemanticVariant
	kSpoofingVariant
	kTraditionalVariant
	kZVariant
)

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
