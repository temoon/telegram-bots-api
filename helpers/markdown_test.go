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

var testsMarkdown = []CaseMarkdown{
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
	for _, test := range testsMarkdown {
		if value := Markdown(test.Text); value != test.Return {
			t.Error("Text:", test.Text, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Markdown(testsMarkdown[i%len(testsMarkdown)].Text)
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

var testsMarkdownBold = []CaseMarkdownBold{
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
	for _, test := range testsMarkdownBold {
		if value := MarkdownBold(test.Text); value != test.Return {
			t.Error("Text:", test.Text, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownBold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownBold(testsMarkdownBold[i%len(testsMarkdownBold)].Text)
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

var testsMarkdownItalic = []CaseMarkdownItalic{
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
	for _, test := range testsMarkdownItalic {
		if value := MarkdownItalic(test.Text); value != test.Return {
			t.Error("Text:", test.Text, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownItalic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownItalic(testsMarkdownItalic[i%len(testsMarkdownItalic)].Text)
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

var testsMarkdownCode = []CaseMarkdownCode{
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
	for _, test := range testsMarkdownCode {
		if value := MarkdownCode(test.Text); value != test.Return {
			t.Error("Text:", test.Text, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownCode(testsMarkdownCode[i%len(testsMarkdownCode)].Text)
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

var testsMarkdownCodeBlock = []CaseMarkdownCodeBlock{
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
	for _, test := range testsMarkdownCodeBlock {
		if value := MarkdownCodeBlock(test.Text); value != test.Return {
			t.Error("Text:", test.Text, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownCodeBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarkdownCodeBlock(testsMarkdownCodeBlock[i%len(testsMarkdownCodeBlock)].Text)
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

var testsMarkdownUserMention = []CaseMarkdownUserMention{
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
	for _, test := range testsMarkdownUserMention {
		if value := MarkdownUserMention(test.Name, test.Id); value != test.Return {
			t.Error("Name:", test.Name, "Id:", test.Id, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownUserMention(b *testing.B) {
	var test CaseMarkdownUserMention
	for i := 0; i < b.N; i++ {
		test = testsMarkdownUserMention[i%len(testsMarkdownUserMention)]
		MarkdownUserMention(test.Name, test.Id)
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

var testsMarkdownUrl = []CaseMarkdownUrl{
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
	for _, test := range testsMarkdownUrl {
		if value := MarkdownUrl(test.Name, test.Url); value != test.Return {
			t.Error("Name:", test.Name, "Url:", test.Url, "Expected:", test.Return, "Got:", value)
		}
	}
}

func BenchmarkMarkdownUrl(b *testing.B) {
	var test CaseMarkdownUrl
	for i := 0; i < b.N; i++ {
		test = testsMarkdownUrl[i%len(testsMarkdownUrl)]
		MarkdownUrl(test.Name, test.Url)
	}
}

func ExampleMarkdownUrl() {
	fmt.Println(MarkdownUrl("a_bc", "https://example.com"))
	// Output: [a\_bc](https://example.com)
}

// endregion
