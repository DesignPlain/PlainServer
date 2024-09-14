package types

type Compute_RegionInstanceGroupManagerAllInstancesConfig struct {
	// , The metadata key-value pairs that you want to patch onto the instance. For more information, see [Project and instance metadata](https://cloud.google.com/compute/docs/metadata#project_and_instance_metadata).
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   , The label key-value pairs that you want to patch onto the instance.

	   - - -
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
