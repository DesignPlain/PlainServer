package types

type Macie2_ClassificationJobS3JobDefinitionScoping struct {
	// The property- or tag-based conditions that determine which objects to exclude from the analysis. (documented below)
	Excludes Macie2_ClassificationJobS3JobDefinitionScopingExcludes `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	// The property- or tag-based conditions that determine which objects to include in the analysis. (documented below)
	Includes Macie2_ClassificationJobS3JobDefinitionScopingIncludes `json:"includes,omitempty" yaml:"includes,omitempty"`
}
