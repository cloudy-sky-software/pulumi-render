package examples

// Needed to support go:embed directive
import _ "embed"

//go:embed staticSite.md
var staticSite string

// ResourceExample contains Markdown-formatted examples
// corresponding to a resource's type token.
var ResourceExample = map[string]string{
	"render:services:StaticSite": staticSite,
}
