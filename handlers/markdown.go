package handlers

import (
	"github.com/theplant/blackfriday"
)

var (
	mdRenderer   blackfriday.Renderer
	mdExtensions int
)

func initMarkdownFlags() {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_ESCAPE_HTML
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	mdRenderer = renderer

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_NO_SPACE_LISTS
	extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK
	extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK
	extensions |= blackfriday.EXTENSION_ONE_SPACE_INDENT
	extensions |= blackfriday.EXTENSION_UNICODE_LIST_ITEM
	mdExtensions = extensions
}

func renderMD(input string) (r string) {
	r = string(blackfriday.Markdown([]byte(input), mdRenderer, mdExtensions))
	return
}
