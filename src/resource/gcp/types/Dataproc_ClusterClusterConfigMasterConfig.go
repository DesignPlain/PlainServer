package types

type Dataproc_ClusterClusterConfigMasterConfig struct {
	// List of master instance names which have been assigned to the cluster.
	InstanceNames []string `json:"instanceNames,omitempty" yaml:"instanceNames,omitempty"`

	/*
	   The name of a Google Compute Engine machine type
	   to create for the master. If not specified, GCP will default to a predetermined
	   computed value (currently `n1-standard-4`).
	*/
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   The name of a minimum generation of CPU family
	   for the master. If not specified, GCP will default to a predetermined computed value
	   for each zone. See [the guide](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
	   for details about which CPU families are available (and defaulted) for each zone.
	*/
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	/*
	   Specifies the number of master nodes to create.
	   If not specified, GCP will default to a predetermined computed value (currently 1).
	*/
	NumInstances int `json:"numInstances,omitempty" yaml:"numInstances,omitempty"`

	// The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.
	Accelerators []Dataproc_ClusterClusterConfigMasterConfigAccelerator `json:"accelerators,omitempty" yaml:"accelerators,omitempty"`

	// Disk Config
	DiskConfig Dataproc_ClusterClusterConfigMasterConfigDiskConfig `json:"diskConfig,omitempty" yaml:"diskConfig,omitempty"`

	/*
	   The URI for the image to use for this worker.  See [the guide](https://cloud.google.com/dataproc/docs/guides/dataproc-images)
	   for more information.
	*/
	ImageUri string `json:"imageUri,omitempty" yaml:"imageUri,omitempty"`
}
