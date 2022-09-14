package main

import (
	"example/src/pages"

	"github.com/danecwalker/goact"
)

func App() goact.El {
	return pages.Home()
}
