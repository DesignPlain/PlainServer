package types

type Macie2_ClassificationJobS3JobDefinitionScopingIncludes struct {
	// An array of conditions, one for each condition that determines which objects to include or exclude from the job. (documented below)
	Ands []Macie2_ClassificationJobS3JobDefinitionScopingIncludesAnd `json:"ands,omitempty" yaml:"ands,omitempty"`
}
