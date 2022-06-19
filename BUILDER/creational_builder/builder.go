package creationalbuilder

import (
	"fmt"
	"strings"
)

const (
	IndentSize = 2
)

type HtmlBuilder struct {
	RootName string
	Root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName,
		HtmlElement{rootName, "", []HtmlElement{}}}
}

func (b *HtmlBuilder) String() string {
	return b.Root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.Root.Elements = append(b.Root.Elements, e)
}

// for chaining the calls
// Fluent interfaces return poiners to the receivers
// It is just a convinience, to keep on calling different build methods
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.Root.Elements = append(b.Root.Elements, e)
	return b
}

type HtmlElement struct {
	Name, Text string
	Elements   []HtmlElement
}

func (h *HtmlElement) String() string {
	return h.string(0)
}

func (h *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", IndentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, h.Name))
	if len(h.Text) > 0 {
		sb.WriteString(strings.Repeat(" ", IndentSize*(indent+1)))
		sb.WriteString(h.Text)
		sb.WriteString("\n")
	}
	for _, el := range h.Elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, h.Name))
	return sb.String()
}
