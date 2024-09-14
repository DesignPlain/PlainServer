package types

type Opensearch_DomainClusterConfigZoneAwarenessConfig struct {
	// Number of Availability Zones for the domain to use with `zone_awareness_enabled`. Defaults to `2`. Valid values: `2` or `3`.
	AvailabilityZoneCount int `json:"availabilityZoneCount,omitempty" yaml:"availabilityZoneCount,omitempty"`
}
