package types

type Networkmanager_CoreNetworkSegment struct {
	// Regions where the edges are located.
	EdgeLocations []string `json:"edgeLocations,omitempty" yaml:"edgeLocations,omitempty"`

	// Name of a core network segment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Shared segments of a core network.
	SharedSegments []string `json:"sharedSegments,omitempty" yaml:"sharedSegments,omitempty"`
}
