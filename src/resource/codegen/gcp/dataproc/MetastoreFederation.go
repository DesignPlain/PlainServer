package dataproc

import types "libds/gcp/types"

type MetastoreFederation struct {
	/*
	   A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
	   Structure is documented below.
	*/
	BackendMetastores []types.Dataproc_MetastoreFederationBackendMetastore `json:"backendMetastores,omitempty" yaml:"backendMetastores,omitempty"`

	/*
	   The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	   and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	   3 and 63 characters.
	*/
	FederationId string `json:"federationId,omitempty" yaml:"federationId,omitempty"`

	/*
	   User-defined labels for the metastore federation.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location where the metastore federation should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
