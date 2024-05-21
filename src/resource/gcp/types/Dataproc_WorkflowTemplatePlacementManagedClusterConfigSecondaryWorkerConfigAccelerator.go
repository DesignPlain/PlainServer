package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerator struct {
	// The number of the accelerator cards of this type exposed to this instance.
	AcceleratorCount int `json:"acceleratorCount,omitempty" yaml:"acceleratorCount,omitempty"`

	// Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`
}
