package templates

import (
	"bytes"
	"context"
	"strings"

	"github.com/a-h/templ"
)

func String(render templ.Component) string {
	return strings.ReplaceAll(parseString(render), "<br>", "\n")
}

func Markdown(render templ.Component) string {
	return parseString(render)
}

func parseString(render templ.Component) string {
	buf := new(bytes.Buffer)
	render.Render(context.Background(), buf)
	return buf.String()
}
