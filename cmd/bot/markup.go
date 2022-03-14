package main

import (
	"fmt"
	"strings"
)

func (app *application) Bold(s string) string {
	return fmt.Sprintf("<strong>%s</strong>", s)
}

func (app *application) Code(s string) string {
	return fmt.Sprintf("<code>%s</code>", s)
}

func (app *application) Italic(s string) string {
	return fmt.Sprintf("<em>%s</em>", s)
}

func (app *application) Underline(s string) string {
	return fmt.Sprintf("<ins>%s</ins>", s)
}

func (app *application) Strikethrough(s string) string {
	return fmt.Sprintf("<strike>%s</strike>", s)
}

func (app *application) Spoiler(s string) string {
	return fmt.Sprintf("<tg-spoiler>%s</tg-spoiler>", s)
}

func (app *application) InlineUrl(url, title string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, title)
}

func (app *application) Pre(s string) string {
	return fmt.Sprintf("<pre>%s</pre>", s)
}

func (app *application) NewLine(count int) string {
	return strings.Repeat("\n", count)
}

func (app *application) Clean(s, del string) string {
	return strings.ReplaceAll(s, del, "")
}
