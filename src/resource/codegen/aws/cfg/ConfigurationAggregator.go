package cfg

import types "libds/aws/types"

type ConfigurationAggregator struct {
	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   Either `account_aggregation_source` or `organization_aggregation_source` must be specified.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource types.Cfg_ConfigurationAggregatorAccountAggregationSource `json:"accountAggregationSource,omitempty" yaml:"accountAggregationSource,omitempty"`

	// The name of the configuration aggregator.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource types.Cfg_ConfigurationAggregatorOrganizationAggregationSource `json:"organizationAggregationSource,omitempty" yaml:"organizationAggregationSource,omitempty"`
}
