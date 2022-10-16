package examples

import _ "embed" // Needed to support go:embed directive

//go:embed service.md
var service string

// ResourceExample contains Markdown-formatted examples corresponding to a resource's type token.
var ResourceExample = map[string]string{
	"render:services:StaticSite": service,
}
