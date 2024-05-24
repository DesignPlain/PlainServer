package appsync

type Type struct {
	// The type format: `SDL` or `JSON`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// GraphQL API ID.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The type definition.
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`
}
