package examples

import _ "embed" // Needed to support go:embed directive

//go:embed service.md
var service string

// ApiVersionToExample contains Markdown-formatted examples corresponding to a k8s apiVersion.
var ApiVersionToExample = map[string]string{
	"render:services:Service": service,
}
