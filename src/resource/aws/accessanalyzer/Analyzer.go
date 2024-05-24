package accessanalyzer

import types "DesignSphere_Server/src/resource/aws/types"

type Analyzer struct {
	/*
	   Name of the Analyzer.

	   The following arguments are optional:
	*/
	AnalyzerName string `json:"analyzerName,omitempty" yaml:"analyzerName,omitempty"`

	// A block that specifies the configuration of the analyzer. Documented below
	Configuration types.Accessanalyzer_AnalyzerConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Type of Analyzer. Valid values are `ACCOUNT`, `ORGANIZATION`, `ACCOUNT_UNUSED_ACCESS `, `ORGANIZATION_UNUSED_ACCESS`. Defaults to `ACCOUNT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
