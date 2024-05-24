package securityhub

type FindingAggregator struct {
	// Indicates whether to aggregate findings from all of the available Regions or from a specified list. The options are `ALL_REGIONS`, `ALL_REGIONS_EXCEPT_SPECIFIED` or `SPECIFIED_REGIONS`. When `ALL_REGIONS` or `ALL_REGIONS_EXCEPT_SPECIFIED` are used, Security Hub will automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	LinkingMode string `json:"linkingMode,omitempty" yaml:"linkingMode,omitempty"`

	// List of regions to include or exclude (required if `linking_mode` is set to `ALL_REGIONS_EXCEPT_SPECIFIED` or `SPECIFIED_REGIONS`)
	SpecifiedRegions []string `json:"specifiedRegions,omitempty" yaml:"specifiedRegions,omitempty"`
}
