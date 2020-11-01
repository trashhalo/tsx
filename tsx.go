package tsx

import (
	"strings"

	"github.com/muesli/termenv"
)

// T jsx/h inspired structure to hold terminal data
type T struct {
	Name       string
	ClassNames []string
	Text       string

	Children []T
}

// Path holds style targeting information of parents/self used by StyleFn
type Path struct {
	Name       string
	ClassNames []string
}

// StyleFn is meant to provide a hook to write simple css like effects
type StyleFn = func(path []Path, text string) termenv.Style

// NoStyle applies no styles to any string
func NoStyle(text string) termenv.Style {
	return termenv.String(text)
}

func (t T) String(style StyleFn, path []Path) string {
	here := t.Path(path)
	str := strings.Builder{}
	for _, n := range t.Children {
		str.WriteString(n.String(style, here))
	}
	str.WriteString(t.Text)

	return style(here, str.String()).String()
}

func (t T) Path(path []Path) []Path {
	return append(path, Path{
		Name:       t.Name,
		ClassNames: t.ClassNames,
	})
}
