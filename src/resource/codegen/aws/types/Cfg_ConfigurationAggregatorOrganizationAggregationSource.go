package types

type Cfg_ConfigurationAggregatorOrganizationAggregationSource struct {
	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions bool `json:"allRegions,omitempty" yaml:"allRegions,omitempty"`

	// List of source regions being aggregated.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	/*
	   ARN of the IAM role used to retrieve AWS Organization details associated with the aggregator account.

	   Either `regions` or `all_regions` (as true) must be specified.
	*/
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
