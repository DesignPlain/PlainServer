package securityhub

import types "DesignSphere_Server/src/resource/aws/types"

type Insight struct {
	// A configuration block including one or more (up to 10 distinct) attributes used to filter the findings included in the insight. The insight only includes findings that match criteria defined in the filters. See filters below for more details.
	Filters types.Securityhub_InsightFilters `json:"filters,omitempty" yaml:"filters,omitempty"`

	// The attribute used to group the findings for the insight e.g., if an insight is grouped by `ResourceId`, then the insight produces a list of resource identifiers.
	GroupByAttribute string `json:"groupByAttribute,omitempty" yaml:"groupByAttribute,omitempty"`

	// The name of the custom insight.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
