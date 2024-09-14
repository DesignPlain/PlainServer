package workstations

import types "libds/gcp/types"

type WorkstationCluster struct {
	// Human-readable name for this resource.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Configuration options for a custom domain.
	   Structure is documented below.
	*/
	DomainConfig types.Workstations_WorkstationClusterDomainConfig `json:"domainConfig,omitempty" yaml:"domainConfig,omitempty"`

	// The location where the workstation cluster should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The relative resource name of the VPC network on which the instance can be accessed.
	   It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	   Must be part of the subnetwork specified for this cluster.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   ID to use for the workstation cluster.


	   - - -
	*/
	WorkstationClusterId string `json:"workstationClusterId,omitempty" yaml:"workstationClusterId,omitempty"`

	/*
	   Client-specified annotations. This is distinct from labels.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Configuration for private cluster.
	   Structure is documented below.
	*/
	PrivateClusterConfig types.Workstations_WorkstationClusterPrivateClusterConfig `json:"privateClusterConfig,omitempty" yaml:"privateClusterConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
