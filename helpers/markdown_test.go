package helpers

import (
	"fmt"
	"testing"
)

// region Markdown

type CaseMarkdown struct {
	Text   string
	Return string
}

var testCasesMarkdown = []CaseMarkdown{
	{Text: "", Return: ""},
	{Text: "abc", Return: "abc"},
	{Text: "*abc*", Return: "\\*abc\\*"},
	{Text: "**abc**", Return: "\\*\\*abc\\*\\*"},
	{Text: "_abc_", Return: "\\_abc\\_"},
	{Text: "__abc__", Return: "\\_\\_abc\\_\\_"},
	{Text: "`abc`", Return: "\\`abc\\`"},
	{Text: "```abc```", Return: "\\`\\`\\`abc\\`\\`\\`"},
	{Text: "[abc]", Return: "\\[abc]"},
	{Text: "[[abc]]", Return: "\\[\\[abc]]"},
	{Text: "*a*b*c*", Return: "\\*a\\*b\\*c\\*"},
	{Text: "_a_b_c_", Return: "\\_a\\_b\\_c\\_"},
	{Text: "`a`b`c`", Return: "\\`a\\`b\\`c\\`"},
	{Text: "[a][b][c]", Return: "\\[a]\\[b]\\[c]"},
	{Text: "\\*abc\\*", Return: "\\*abc\\*"},
	{Text: "\\_abc\\_", Return: "\\_abc\\_"},
	{Text: "\\`abc\\`", Return: "\\`abc\\`"},
	{Text: "\\[abc\\]", Return: "\\[abc\\]"},
	{Text: "a\\\\b\\\\c", Return: "a\\\\b\\\\c"},
	{Text: "\\", Return: "\\"},
}

func TestMarkdown(t *testing.T) {
	for _, testCase := range testCasesMarkdown {
		if value := Markdown(testCase.Text); value != testCase.Return {
			t.Error("Text:", testCase.Text, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Markdown(testCasesMarkdown[i%len(testCasesMarkdown)].Text)
	}
}

func ExampleMarkdown() {
	fmt.Println(Markdown("a*bc"))
	// Output: a\*bc
}

// endregion

// region MarkdownBold

type CaseMarkdownBold struct {
	Text   string
	Return string
}

var testCasesMarkdownBold = []CaseMarkdownBold{
	{Text: "", Return: ""},
	{Text: "abc", Return: "*abc*"},
	{Text: "a*bc", Return: "*a*\\**bc*"},
	{Text: "a_bc", Return: "*a\\_bc*"},
	{Text: "a`bc", Return: "*a\\`bc*"},
	{Text: "a[b]c", Return: "*a\\[b]c*"},
	{Text: "a\\*bc", Return: "*a\\*bc*"},
	{Text: "a\\_bc", Return: "*a\\_bc*"},
	{Text: "a\\`bc", Return: "*a\\`bc*"},
	{Text: "a\\[b]c", Return: "*a\\[b]c*"},
	{Text: "a\\[b\\]c", Return: "*a\\[b\\]c*"},
	{Text: "a\\b\\c", Return: "*a\\b\\c*"},
	{Text: "\\", Return: "*\\*"},
}

func TestMarkdownBold(t *testing.T) {
	for _, testCase := range testCasesMarkdownBold {
		if value := MarkdownBold(testCase.Text); value != testCase.Return {
			t.Error("Text:", testCase.Text, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownBold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownBold(testCasesMarkdownBold[i%len(testCasesMarkdownBold)].Text)
	}
}

func ExampleMarkdownBold() {
	fmt.Println(MarkdownBold("a*bc"))
	// Output: *a*\**bc*
}

// endregion

// region MarkdownItalic

type CaseMarkdownItalic struct {
	Text   string
	Return string
}

var testCasesMarkdownItalic = []CaseMarkdownItalic{
	{Text: "", Return: ""},
	{Text: "abc", Return: "_abc_"},
	{Text: "a*bc", Return: "_a\\*bc_"},
	{Text: "a_bc", Return: "_a_\\__bc_"},
	{Text: "a`bc", Return: "_a\\`bc_"},
	{Text: "a[b]c", Return: "_a\\[b]c_"},
	{Text: "a\\*bc", Return: "_a\\*bc_"},
	{Text: "a\\_bc", Return: "_a\\_bc_"},
	{Text: "a\\`bc", Return: "_a\\`bc_"},
	{Text: "a\\[b]c", Return: "_a\\[b]c_"},
	{Text: "a\\[b\\]c", Return: "_a\\[b\\]c_"},
	{Text: "a\\b\\c", Return: "_a\\b\\c_"},
	{Text: "\\", Return: "_\\_"},
}

func TestMarkdownItalic(t *testing.T) {
	for _, testCase := range testCasesMarkdownItalic {
		if value := MarkdownItalic(testCase.Text); value != testCase.Return {
			t.Error("Text:", testCase.Text, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownItalic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownItalic(testCasesMarkdownItalic[i%len(testCasesMarkdownItalic)].Text)
	}
}

func ExampleMarkdownItalic() {
	fmt.Println(MarkdownItalic("a_bc"))
	// Output: _a_\__bc_
}

// endregion

// region MarkdownCode

type CaseMarkdownCode struct {
	Text   string
	Return string
}

var testCasesMarkdownCode = []CaseMarkdownCode{
	{Text: "", Return: ""},
	{Text: "abc", Return: "`abc`"},
	{Text: "a*bc", Return: "`a\\*bc`"},
	{Text: "a_bc", Return: "`a\\_bc`"},
	{Text: "a`bc", Return: "`a`\\``bc`"},
	{Text: "a[b]c", Return: "`a\\[b]c`"},
	{Text: "a\\*bc", Return: "`a\\*bc`"},
	{Text: "a\\_bc", Return: "`a\\_bc`"},
	{Text: "a\\`bc", Return: "`a\\`bc`"},
	{Text: "a\\[b]c", Return: "`a\\[b]c`"},
	{Text: "a\\[b\\]c", Return: "`a\\[b\\]c`"},
	{Text: "a\\b\\c", Return: "`a\\b\\c`"},
	{Text: "\\", Return: "`\\`"},
}

func TestMarkdownCode(t *testing.T) {
	for _, testCase := range testCasesMarkdownCode {
		if value := MarkdownCode(testCase.Text); value != testCase.Return {
			t.Error("Text:", testCase.Text, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownCode(testCasesMarkdownCode[i%len(testCasesMarkdownCode)].Text)
	}
}

func ExampleMarkdownCode() {
	fmt.Println(MarkdownCode("a`bc"))
	// Output: `a`\``bc`
}

// endregion

// region MarkdownCodeBlock

type CaseMarkdownCodeBlock struct {
	Text   string
	Return string
}

var testCasesMarkdownCodeBlock = []CaseMarkdownCodeBlock{
	{Text: "", Return: ""},
	{Text: "abc", Return: "```abc```"},
	{Text: "a*bc", Return: "```a*bc```"},
	{Text: "a_bc", Return: "```a_bc```"},
	{Text: "a`bc", Return: "```a`bc```"},
	{Text: "a[b]c", Return: "```a[b]c```"},
	{Text: "a\\*bc", Return: "```a\\*bc```"},
	{Text: "a\\_bc", Return: "```a\\_bc```"},
	{Text: "a\\`bc", Return: "```a\\`bc```"},
	{Text: "a\\[b]c", Return: "```a\\[b]c```"},
	{Text: "a\\[b\\]c", Return: "```a\\[b\\]c```"},
	{Text: "a\\b\\c", Return: "```a\\b\\c```"},
	{Text: "\\", Return: "```\\```"},
}

func TestMarkdownCodeBlock(t *testing.T) {
	for _, testCase := range testCasesMarkdownCodeBlock {
		if value := MarkdownCodeBlock(testCase.Text); value != testCase.Return {
			t.Error("Text:", testCase.Text, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownCodeBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownCodeBlock(testCasesMarkdownCodeBlock[i%len(testCasesMarkdownCodeBlock)].Text)
	}
}

func ExampleMarkdownCodeBlock() {
	fmt.Println(MarkdownCodeBlock("a`bc"))
	// Output: ```a`bc```
}

// endregion

// region MarkdownUserMention

type CaseMarkdownUserMention struct {
	Name   string
	Id     int64
	Return string
}

var testCasesMarkdownUserMention = []CaseMarkdownUserMention{
	{Name: "", Id: 0, Return: ""},
	{Name: "abc", Id: 0, Return: "[abc](tg://user?id=0)"},
	{Name: "abc", Id: 42, Return: "[abc](tg://user?id=42)"},
	{Name: "abc", Id: -42, Return: "[abc](tg://user?id=-42)"},
	{Name: "a*bc", Id: 42, Return: "[a\\*bc](tg://user?id=42)"},
	{Name: "a_bc", Id: 42, Return: "[a\\_bc](tg://user?id=42)"},
	{Name: "a`bc", Id: 42, Return: "[a\\`bc](tg://user?id=42)"},
	{Name: "a[b]c", Id: 42, Return: "[a\\[b]c](tg://user?id=42)"},
	{Name: "a\\*bc", Id: 42, Return: "[a\\*bc](tg://user?id=42)"},
	{Name: "a\\_bc", Id: 42, Return: "[a\\_bc](tg://user?id=42)"},
	{Name: "a\\`bc", Id: 42, Return: "[a\\`bc](tg://user?id=42)"},
	{Name: "a\\[b]c", Id: 42, Return: "[a\\[b]c](tg://user?id=42)"},
	{Name: "a\\[b\\]c", Id: 42, Return: "[a\\[b\\]c](tg://user?id=42)"},
	{Name: "a\\b\\c", Id: 42, Return: "[a\\b\\c](tg://user?id=42)"},
	{Name: "\\", Id: 42, Return: "[\\](tg://user?id=42)"},
}

func TestMarkdownUserMention(t *testing.T) {
	for _, testCase := range testCasesMarkdownUserMention {
		if value := MarkdownUserMention(testCase.Name, testCase.Id); value != testCase.Return {
			t.Error("Name:", testCase.Name, "Id:", testCase.Id, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownUserMention(b *testing.B) {
	var testCase CaseMarkdownUserMention
	for i := 0; i < b.N; i++ {
		testCase = testCasesMarkdownUserMention[i%len(testCasesMarkdownUserMention)]
		MarkdownUserMention(testCase.Name, testCase.Id)
	}
}

func ExampleMarkdownUserMention() {
	fmt.Println(MarkdownUserMention("a_bc", 42))
	// Output: [a\_bc](tg://user?id=42)
}

// endregion

// region MarkdownUrl

type CaseMarkdownUrl struct {
	Name   string
	Url    string
	Return string
}

var testCasesMarkdownUrl = []CaseMarkdownUrl{
	{Name: "", Url: "", Return: ""},
	{Name: "abc", Url: "https://example.com", Return: "[abc](https://example.com)"},
	{Name: "abc", Url: "https://example.com", Return: "[abc](https://example.com)"},
	{Name: "abc", Url: "https://example.com", Return: "[abc](https://example.com)"},
	{Name: "a*bc", Url: "https://example.com", Return: "[a\\*bc](https://example.com)"},
	{Name: "a_bc", Url: "https://example.com", Return: "[a\\_bc](https://example.com)"},
	{Name: "a`bc", Url: "https://example.com", Return: "[a\\`bc](https://example.com)"},
	{Name: "a[b]c", Url: "https://example.com", Return: "[a\\[b]c](https://example.com)"},
	{Name: "a\\*bc", Url: "https://example.com", Return: "[a\\*bc](https://example.com)"},
	{Name: "a\\_bc", Url: "https://example.com", Return: "[a\\_bc](https://example.com)"},
	{Name: "a\\`bc", Url: "https://example.com", Return: "[a\\`bc](https://example.com)"},
	{Name: "a\\[b]c", Url: "https://example.com", Return: "[a\\[b]c](https://example.com)"},
	{Name: "a\\[b\\]c", Url: "https://example.com", Return: "[a\\[b\\]c](https://example.com)"},
	{Name: "a\\b\\c", Url: "https://example.com", Return: "[a\\b\\c](https://example.com)"},
	{Name: "\\", Url: "https://example.com", Return: "[\\](https://example.com)"},
}

func TestMarkdownUrl(t *testing.T) {
	for _, testCase := range testCasesMarkdownUrl {
		if value := MarkdownUrl(testCase.Name, testCase.Url); value != testCase.Return {
			t.Error("Name:", testCase.Name, "Url:", testCase.Url, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownUrl(b *testing.B) {
	var testCase CaseMarkdownUrl
	for i := 0; i < b.N; i++ {
		testCase = testCasesMarkdownUrl[i%len(testCasesMarkdownUrl)]
		MarkdownUrl(testCase.Name, testCase.Url)
	}
}

func ExampleMarkdownUrl() {
	fmt.Println(MarkdownUrl("a_bc", "https://example.com"))
	// Output: [a\_bc](https://example.com)
}

// endregion
