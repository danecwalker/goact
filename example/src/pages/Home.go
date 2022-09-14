package pages

import (
	"example/src/components"

	"github.com/danecwalker/goact"
)

func Home() goact.El {
	return goact.El{
		Tag: "div",
		Attr: &map[string]interface{}{
			"class": "test",
		},
		Children: goact.Els{
			components.Nav(),
			components.Section("Hello, World!", "This is a simple example of Goact.", "home"),
			components.Section("About", "Goact is a simple, lightweight, and fast virtual DOM library for Go.", "about"),
		},
	}
}
