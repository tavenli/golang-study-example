package main

import (
	"GoBase/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**

https://www.w3school.com.cn/charsets/ref_emoji_smileys.asp

https://tool.chinaz.com/Tools/unicode.aspx

https://github.com/matrix-org/gomatrixserverlib/blob/14ee7615d6041e851658f2f46057f81738a49d69/json_test.go

其它语言使用的方式，如Java、NodeJs
\u6211\u7231\u4e2d\u56fd\ud83d\udd25\u6C38\u65E0\u6B62\u5883\uD83D\uDD25

这里的表情符号为🔥

对应十进制：128293
对应十六进制：0x1F525
对应Unicode编码：\uD83D\uDD25

Go语言
\u6211\u7231\u4e2d\u56fd\U0001f525\u6c38\u65e0\u6b62\u5883\U0001f525

*/

func EmojiUnicode_main() {
	//
	NormalUnicode()
	//
	QuoteDemo2()
}

func NormalUnicode() {
	input := "Hi,☝ 我爱中国🔥永无止境🔥"
	//
	fmt.Println(strconv.Quote(input))

	//
	fmt.Println(strconv.QuoteToASCII(input))
	fmt.Println(strconv.QuoteToGraphic(input))

	toUnicode := strconv.QuoteToASCII(input)

	content, err := strconv.Unquote(toUnicode)
	fmt.Println(content, err)

	fmt.Println("\u6211\u7231\u4e2d\u56fd\u261d")
	content2, err := strconv.Unquote("`\u6211\u7231\u4e2d\u56fd\u261d`")
	fmt.Println(content2, err)

	fmt.Println(strconv.QuotedPrefix(toUnicode))

	fmt.Println(UnicodeEmojiCode("🔥永无止境🔥"))

	var jsonMap map[string]interface{}
	json := []byte(`{"Emoji":"Hi,\u261d \u6211\u7231\u4e2d\u56fd\ud83d\udd25\u6C38\u65E0\u6B62\u5883\uD83D\uDD25"}`)
	utils.FromJson(json, &jsonMap)
	fmt.Println(jsonMap)
	for k, v := range jsonMap {
		fmt.Println(k, v)
	}

	fmt.Println("-------------")
}

// 表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

// 表情转换
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

func QuoteDemo2() {
	// 转意字符串
	fmt.Println(strconv.Quote("hello world your's are good"))

	// 将转意字符追加到 dst 并返回缓冲区
	fmt.Println(strconv.AppendQuote([]byte(""), "hello world your's are good"))

	// 转义字符串为ASCII
	fmt.Println(strconv.QuoteToASCII("hello world your's are good"))

	// 将rune转义为ASCII并添加
	fmt.Println(strconv.AppendQuoteRuneToASCII([]byte("rune (ascii):"), '☺'))

	// 转义字符串为Unicode图形字符
	fmt.Println(strconv.QuoteToGraphic("jsdjfkh"))

	// 将rune转义为Unicode图形字符并添加
	fmt.Println(strconv.AppendQuoteToGraphic([]byte("rune (ascii):"), "☺"))

	// rune 转 string
	fmt.Println(strconv.QuoteRune('☺'))

	// rune 转义 并添加
	fmt.Println(strconv.AppendQuoteRune([]byte(""), '☺'))

	// rune转ASCII
	fmt.Println(strconv.QuoteRuneToASCII('☺'))

	// rune转ASCII 并添加
	fmt.Println(strconv.AppendQuoteRuneToASCII([]byte(""), '☺'))

	// rune 转 Unicode
	fmt.Println(strconv.QuoteRuneToGraphic('☺'))

	// rune 转 Unicode 并添加
	fmt.Println(strconv.AppendQuoteRuneToGraphic([]byte(""), '☺'))

	// 判断是否可以转义
	fmt.Println(strconv.CanBackquote(`"Fran & Freddie's Diner	☺"`))

	// 反转义
	fmt.Println(strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"'))

	// 返回s前缀处的带引号的字符串
	fmt.Println(strconv.QuotedPrefix(`"Fran & Freddie's Diner	☺"`))

	// 返回s引号中的字符串值
	fmt.Println(strconv.Unquote(`"Fran & Freddie's Diner	☺"`))

	// 判断rune是否可以转义
	fmt.Println(strconv.IsPrint('☺'))

	// 判断rune是否可以转义为 Unicode
	fmt.Println(strconv.IsGraphic('☺'))
}
