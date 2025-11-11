// Package embed types defined in the Go file after the parsed file in the same package
package embed

type EmbedBase struct {
	ID string `json:"id"`
}

// Reference struct type, defined after embed.go, the same pkg but not the same file
type Reference struct {
	Foo string `json:"foo"`
}
