package goact

import (
	"fmt"
	"syscall/js"
)

type El struct {
	Tag      string
	Attr     *map[string]interface{}
	Children goactEl
}

type Text string

type goactEl interface {
	toString() string
}

type Els []goactEl

func (els Els) toString() string {
	var str string
	for _, el := range els {
		str += el.toString()
	}
	return str
}

func (text Text) toString() string {
	return string(text)
}

func (el El) toString() string {
	attr := ""
	if el.Attr != nil {
		for key, value := range *el.Attr {
			attr += fmt.Sprintf(" %s=\"%v\"", key, value)
		}
	}
	return fmt.Sprintf("<%s%s>%s</%s>", el.Tag, attr, el.Children.toString(), el.Tag)
}

func Render(el goactEl) string {
	return el.toString()
}

func UseStyle(style string) {
	prevStyle := js.Global().Get("document").Call("querySelector", "style").Get("innerHTML").String()
	js.Global().Get("document").Call("querySelector", "style").Set("innerHTML", prevStyle+style)
}
