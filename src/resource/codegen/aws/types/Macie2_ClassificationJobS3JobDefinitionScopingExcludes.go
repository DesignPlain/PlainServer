package types

type Macie2_ClassificationJobS3JobDefinitionScopingExcludes struct {
	// An array of conditions, one for each condition that determines which objects to include or exclude from the job. (documented below)
	Ands []Macie2_ClassificationJobS3JobDefinitionScopingExcludesAnd `json:"ands,omitempty" yaml:"ands,omitempty"`
}
