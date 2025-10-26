package markdown

import (
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	md "github.com/nao1215/markdown"
)

func BuildMarkdown(fn func(m *md.Markdown)) *string {
	builder := md.NewMarkdown(nil)
	fn(builder)

	return pointer.New(builder.String())
}
