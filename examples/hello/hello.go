package main

import (
	"fmt"

	"github.com/muesli/termenv"
	"github.com/trashhalo/tsx"
)

func main() {
	a := app(
		hello(),
		world(),
	)
	fmt.Print(a.String(style, nil))
}

func style(path []tsx.Path, text string) termenv.Style {
	str := termenv.String(text)
	node := path[len(path)-1]

	switch node.Name {
	case "hello":
		return str.Bold()
	case "world":
		return str.Italic()
	}

	return str
}

func app(children ...tsx.T) tsx.T {
	return tsx.T{
		Name:     "app",
		Children: children,
	}
}

func hello() tsx.T {
	return tsx.T{
		Name: "hello",
		Text: "Hello ",
	}
}

func world() tsx.T {
	return tsx.T{
		Name: "world",
		Text: "World",
	}
}
