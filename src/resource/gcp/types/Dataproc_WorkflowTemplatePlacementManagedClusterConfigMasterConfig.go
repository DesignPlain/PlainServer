package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfig struct {
	// The Compute Engine accelerator configuration for these instances.
	Accelerators []Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerator `json:"accelerators,omitempty" yaml:"accelerators,omitempty"`

	// The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: - `https://www.googleapis.com/compute/beta/projects/` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.
	InstanceNames []string `json:"instanceNames,omitempty" yaml:"instanceNames,omitempty"`

	// Output only. Specifies that this instance group contains preemptible instances.
	IsPreemptible bool `json:"isPreemptible,omitempty" yaml:"isPreemptible,omitempty"`

	// The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: - `https://www.googleapis.com/compute/v1/projects/(https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Specifies the minimum cpu platform for the Instance Group. See (https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// Disk option config settings.
	DiskConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig `json:"diskConfig,omitempty" yaml:"diskConfig,omitempty"`

	// Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.
	ManagedGroupConfigs []Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig `json:"managedGroupConfigs,omitempty" yaml:"managedGroupConfigs,omitempty"`

	// The number of VM instances in the instance group. For master instance groups, must be set to 1.
	NumInstances int `json:"numInstances,omitempty" yaml:"numInstances,omitempty"`

	// Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE
	Preemptibility string `json:"preemptibility,omitempty" yaml:"preemptibility,omitempty"`
}
