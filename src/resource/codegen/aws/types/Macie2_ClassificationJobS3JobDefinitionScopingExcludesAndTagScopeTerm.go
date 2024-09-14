package types

type Macie2_ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm struct {
	// The operator to use in the condition.
	Comparator string `json:"comparator,omitempty" yaml:"comparator,omitempty"`

	// The tag key to use in the condition. The only valid value is `TAG`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The tag keys or tag key and value pairs to use in the condition.
	TagValues []Macie2_ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermTagValue `json:"tagValues,omitempty" yaml:"tagValues,omitempty"`

	// The type of object to apply the condition to. The only valid value is `S3_OBJECT`.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
