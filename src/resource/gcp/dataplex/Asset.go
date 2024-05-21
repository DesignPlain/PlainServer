package dataplex

import types "DesignSphere_Server/src/resource/gcp/types"

type Asset struct {
	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The zone for the resource
	DataplexZone string `json:"dataplexZone,omitempty" yaml:"dataplexZone,omitempty"`

	// Optional. Description of the asset.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Optional. User defined labels for the asset.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Required. Immutable. Specification of the resource that is referenced by this asset.
	ResourceSpec types.Dataplex_AssetResourceSpec `json:"resourceSpec,omitempty" yaml:"resourceSpec,omitempty"`

	// Required. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec types.Dataplex_AssetDiscoverySpec `json:"discoverySpec,omitempty" yaml:"discoverySpec,omitempty"`

	// The lake for the resource
	Lake string `json:"lake,omitempty" yaml:"lake,omitempty"`

	// The name of the asset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
