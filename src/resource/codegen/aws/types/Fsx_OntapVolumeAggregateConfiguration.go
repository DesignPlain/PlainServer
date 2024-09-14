package types

type Fsx_OntapVolumeAggregateConfiguration struct {
	// Used to specify the names of the aggregates on which the volume will be created. Each aggregate needs to be in the format aggrX where X is the number of the aggregate.
	Aggregates []string `json:"aggregates,omitempty" yaml:"aggregates,omitempty"`

	// Used to explicitly set the number of constituents within the FlexGroup per storage aggregate. the default value is `8`.
	ConstituentsPerAggregate int `json:"constituentsPerAggregate,omitempty" yaml:"constituentsPerAggregate,omitempty"`

	// The total amount of constituents for a `FLEXGROUP` volume. This would equal constituents_per_aggregate x aggregates.
	TotalConstituents int `json:"totalConstituents,omitempty" yaml:"totalConstituents,omitempty"`
}
