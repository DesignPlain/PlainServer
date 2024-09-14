package types

type Alloydb_getLocationsLocation struct {
	// The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Cross-service attributes for the location. For example `{"cloud.googleapis.com/region": "us-east1"}`.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The canonical id for this location. For example: "us-east1"..
	LocationId string `json:"locationId,omitempty" yaml:"locationId,omitempty"`

	// Service-specific metadata. For example the available capacity at the given location.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Resource name for the location, which may vary between implementations. For example: "projects/example-project/locations/us-east1".
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
