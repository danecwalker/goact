package components

import "github.com/danecwalker/goact"

func Section(title, description, id string) goact.El {
	goact.UseStyle(`
		.interact {
			transition: all 0.3s;
			scale: 1;
		}
		.interact:hover {
			scale: 1.1;
		}`,
	)
	return goact.El{
		Tag: "section",
		Attr: &map[string]interface{}{
			"id": id,
			"style": `
				margin: 0;
				padding: 0;
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
				width: 100%;
				height: 100vh;
				height: 100dvh;
				text-align: center;
				padding: 1rem;
				box-sizing: border-box;
			`,
		},
		Children: goact.Els{
			goact.El{
				Tag: "h1",
				Attr: &map[string]interface{}{
					"class": "interact",
					"style": `
						margin: 0;
						padding: 0;
						font-family: sans-serif;`,
				},
				Children: goact.Text(title),
			},
			goact.El{
				Tag: "p",
				Attr: &map[string]interface{}{
					"class": "interact",
					"style": `
						margin: 0;
						padding: 0;
						font-family: sans-serif;`,
				},
				Children: goact.Text(description),
			},
		},
	}
}
