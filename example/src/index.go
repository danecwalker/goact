package main

import (
	"fmt"
	"syscall/js"

	"github.com/danecwalker/goact"
)

func main() {
	ch := make(chan struct{}, 0)
	fmt.Println("Hello, playground")
	// fmt.Println(js.Global().Get("window").Get("location").Get("pathname"))
	js.Global().Get("document").Call("getElementById", "app").Set("innerHTML",
		goact.Render(
			goact.El{
				Tag: "div",
				Attr: &map[string]interface{}{
					"style": `
						margin: 0;
						padding: 0;`,
				},
				Children: App(),
			},
		))
	<-ch
}
