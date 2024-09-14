package types

type Dataloss_PreventionInspectTemplateInspectConfig struct {
	/*
	   Only returns findings equal or above this threshold. See https://cloud.google.com/dlp/docs/likelihood for more info
	   Default value is `POSSIBLE`.
	   Possible values are: `VERY_UNLIKELY`, `UNLIKELY`, `POSSIBLE`, `LIKELY`, `VERY_LIKELY`.
	*/
	MinLikelihood string `json:"minLikelihood,omitempty" yaml:"minLikelihood,omitempty"`

	/*
	   Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end,
	   other rules are executed in the order they are specified for each info type.
	   Structure is documented below.
	*/
	RuleSets []Dataloss_PreventionInspectTemplateInspectConfigRuleSet `json:"ruleSets,omitempty" yaml:"ruleSets,omitempty"`

	/*
	   List of options defining data content to scan. If empty, text, images, and other content will be included.
	   Each value may be one of: `CONTENT_TEXT`, `CONTENT_IMAGE`.
	*/
	ContentOptions []string `json:"contentOptions,omitempty" yaml:"contentOptions,omitempty"`

	/*
	   Custom info types to be used. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.
	   Structure is documented below.
	*/
	CustomInfoTypes []Dataloss_PreventionInspectTemplateInspectConfigCustomInfoType `json:"customInfoTypes,omitempty" yaml:"customInfoTypes,omitempty"`

	// When true, excludes type information of the findings.
	ExcludeInfoTypes bool `json:"excludeInfoTypes,omitempty" yaml:"excludeInfoTypes,omitempty"`

	// When true, a contextual quote from the data that triggered a finding is included in the response.
	IncludeQuote bool `json:"includeQuote,omitempty" yaml:"includeQuote,omitempty"`

	/*
	   Restricts what infoTypes to look for. The values must correspond to InfoType values returned by infoTypes.list
	   or listed at https://cloud.google.com/dlp/docs/infotypes-reference.
	   When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run.
	   By default this may be all types, but may change over time as detectors are updated.
	   Structure is documented below.
	*/
	InfoTypes []Dataloss_PreventionInspectTemplateInspectConfigInfoType `json:"infoTypes,omitempty" yaml:"infoTypes,omitempty"`

	/*
	   Configuration to control the number of findings returned.
	   Structure is documented below.
	*/
	Limits Dataloss_PreventionInspectTemplateInspectConfigLimits `json:"limits,omitempty" yaml:"limits,omitempty"`
}
