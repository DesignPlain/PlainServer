package types

type Cfg_ConfigurationAggregatorAccountAggregationSource struct {
	// List of 12-digit account IDs of the account(s) being aggregated.
	AccountIds []string `json:"accountIds,omitempty" yaml:"accountIds,omitempty"`

	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions bool `json:"allRegions,omitempty" yaml:"allRegions,omitempty"`

	/*
	   List of source regions being aggregated.

	   Either `regions` or `all_regions` (as true) must be specified.
	*/
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
}
