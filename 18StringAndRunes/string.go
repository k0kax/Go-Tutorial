package main

import (
	"fmt"
	"unicode/utf8"
)

// byte 表示一个字节，rune 表示四个字节
func main() {
	const s = "สวัสดี"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune conut:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	//iteration 迭代
	fmt.Println("\nUsing DecodeRuneInString")
	//从字符串的指定位置解码一个完整的 UTF-8 符文（rune），并返回该符文及其占用的字节数
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)

		w = width

		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
