package workstations

type Workstation struct {
	/*
	   The location where the workstation parent resources reside.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The ID of the parent workstation cluster config.
	WorkstationConfigId string `json:"workstationConfigId,omitempty" yaml:"workstationConfigId,omitempty"`

	// 'Client-specified environment variables passed to the workstation container's entrypoint.'
	Env map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	/*
	   Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The ID of the parent workstation cluster.
	WorkstationClusterId string `json:"workstationClusterId,omitempty" yaml:"workstationClusterId,omitempty"`

	// ID to use for the workstation.
	WorkstationId string `json:"workstationId,omitempty" yaml:"workstationId,omitempty"`

	/*
	   Client-specified annotations. This is distinct from labels.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Human-readable name for this resource.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
