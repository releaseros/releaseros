package markdown

import (
	"bytes"
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type Document struct {
	Nodes []Node
}

func NewDocument() Document {
	return Document{}
}

func (d Document) With(nodes ...Node) Document {
	d.Nodes = append(d.Nodes, nodes...)
	return d
}

func (d Document) String() string {
	var buf bytes.Buffer

	for _, node := range d.Nodes {
		buf.WriteString(node.String())
	}

	return buf.String()
}

type Header struct {
	Text  string
	Level int
}

func H2(text string) Header {
	return Header{
		Level: 2,
		Text:  text,
	}
}

func H3(text string) Header {
	return Header{
		Level: 3,
		Text:  text,
	}
}

func (h Header) String() string {
	return fmt.Sprintf("%s %s\n\n", strings.Repeat("#", h.Level), h.Text)
}

type UnorderedList struct {
	Items []ListItem
}

func UL() UnorderedList {
	return UnorderedList{}
}

func (ul UnorderedList) With(items ...ListItem) UnorderedList {
	ul.Items = append(ul.Items, items...)
	return ul
}

func (ul UnorderedList) String() string {
	var buf bytes.Buffer

	for _, item := range ul.Items {
		buf.WriteString(item.String())
	}

	return fmt.Sprintf("%s\n", &buf)
}

type ListItem struct {
	Text string
}

func LI(text string) ListItem {
	return ListItem{
		Text: text,
	}
}

func (li ListItem) String() string {
	return fmt.Sprintf("* %s\n", li.Text)
}
