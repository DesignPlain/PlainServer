package dataplex

import types "libds/gcp/types"

type Zone struct {
	// Optional. Description of the zone.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec types.Dataplex_ZoneDiscoverySpec `json:"discoverySpec,omitempty" yaml:"discoverySpec,omitempty"`

	/*
	   Optional. User defined labels for the zone.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The lake for the resource
	Lake string `json:"lake,omitempty" yaml:"lake,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the zone.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec types.Dataplex_ZoneResourceSpec `json:"resourceSpec,omitempty" yaml:"resourceSpec,omitempty"`
}
