package helpers

import (
	"strconv"
	"strings"
)

func Markdown(text string) string {
	return markdownEscape("", text)
}

func MarkdownBold(text string) string {
	value := markdownEscape("*", text)
	if value == "" {
		return ""
	}

	return "*" + value + "*"
}

func MarkdownItalic(text string) string {
	value := markdownEscape("_", text)
	if value == "" {
		return ""
	}

	return "_" + value + "_"
}

func MarkdownCode(text string) string {
	value := markdownEscape("`", text)
	if value == "" {
		return ""
	}

	return "`" + value + "`"
}

func MarkdownCodeBlock(text string) string {
	if text == "" {
		return ""
	}

	return "```" + text + "```"
}

func MarkdownUserMention(name string, id int64) string {
	return MarkdownUrl(name, "tg://user?id="+strconv.FormatInt(id, 10))
}

func MarkdownUrl(text string, url string) string {
	value := markdownEscape("", text)
	if value == "" {
		return ""
	}

	return "[" + value + "](" + url + ")"
}

func markdownEscape(entity string, text string) string {
	if text == "" {
		return ""
	}

	builder := strings.Builder{}

	for i := 0; i < len(text); i++ {
		switch string(text[i]) {
		case "\\":
			builder.WriteByte('\\')
			i++
		case entity:
			builder.WriteString(entity + "\\" + entity)
		case "_", "*", "`", "[":
			builder.WriteByte('\\')
		}

		if i < len(text) {
			builder.WriteByte(text[i])
		}
	}

	return builder.String()
}
