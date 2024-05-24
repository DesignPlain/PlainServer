package types

type Cloudsearch_DomainIndexField struct {
	// The default value for the field. This value is used when no value is specified for the field in the document data.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// You can get facet information by enabling this.
	Facet bool `json:"facet,omitempty" yaml:"facet,omitempty"`

	// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// You can enable returning the value of all searchable fields.
	Return bool `json:"return,omitempty" yaml:"return,omitempty"`

	// The field type. Valid values: `date`, `date-array`, `double`, `double-array`, `int`, `int-array`, `literal`, `literal-array`, `text`, `text-array`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The analysis scheme you want to use for a `text` field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	AnalysisScheme string `json:"analysisScheme,omitempty" yaml:"analysisScheme,omitempty"`

	// You can highlight information.
	Highlight bool `json:"highlight,omitempty" yaml:"highlight,omitempty"`

	// You can set whether this index should be searchable or not.
	Search bool `json:"search,omitempty" yaml:"search,omitempty"`

	// You can enable the property to be sortable.
	Sort bool `json:"sort,omitempty" yaml:"sort,omitempty"`

	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	SourceFields string `json:"sourceFields,omitempty" yaml:"sourceFields,omitempty"`
}
