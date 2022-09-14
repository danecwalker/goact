package components

import "github.com/danecwalker/goact"

func Nav() goact.El {
	goact.UseStyle(`
		.link {
			text-decoration: none;
			color: #000;
			transition: all 0.3s;
		}

		.link:hover {
			color: #aaa;
			text-decoration: underline;
		}
	`)
	return goact.El{
		Tag:  "nav",
		Attr: nil,
		Children: goact.Els{
			goact.El{
				Tag: "ul",
				Attr: &map[string]interface{}{
					"style": `
						margin: 0;
						padding: 0;
						display: flex;
						flex-direction: row;
						list-style: none;
						justify-content: space-around;
						align-items: center;
						width: 100%;
						position: fixed;
						background-color: #fff;
						z-index: 100;`,
				},
				Children: goact.Els{
					goact.El{
						Tag:  "li",
						Attr: &map[string]interface{}{},
						Children: goact.El{
							Tag: "a",
							Attr: &map[string]interface{}{
								"href":  "#home",
								"class": "link",
								"style": `
									font-family: sans-serif;`,
							},
							Children: goact.Text("Home"),
						},
					},
					goact.El{
						Tag:  "li",
						Attr: &map[string]interface{}{},
						Children: goact.El{
							Tag: "a",
							Attr: &map[string]interface{}{
								"href": "/",
								"style": `
									text-decoration: none;
									color: black;
									font-family: sans-serif;`,
							},
							Children: goact.El{
								Tag:      "h1",
								Attr:     &map[string]interface{}{},
								Children: goact.Text("Goact"),
							},
						},
					},
					goact.El{
						Tag:  "li",
						Attr: &map[string]interface{}{},
						Children: goact.El{
							Tag: "a",
							Attr: &map[string]interface{}{
								"href":  "#about",
								"class": "link",
								"style": `
									font-family: sans-serif;`,
							},
							Children: goact.Text("About"),
						},
					},
				},
			},
		},
	}
}
