package types

type Dataproc_ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfigDiskConfig struct {
	/*
	   The disk type of the primary disk attached to each node.
	   One of `"pd-ssd"` or `"pd-standard"`. Defaults to `"pd-standard"`.
	*/
	BootDiskType string `json:"bootDiskType,omitempty" yaml:"bootDiskType,omitempty"`

	/*
	   The amount of local SSD disks that will be attached to each master cluster node.
	   Defaults to 0.
	*/
	NumLocalSsds int `json:"numLocalSsds,omitempty" yaml:"numLocalSsds,omitempty"`

	/*
	   Size of the primary disk attached to each node, specified
	   in GB. The primary disk contains the boot volume and system libraries, and the
	   smallest allowed disk size is 10GB. GCP will default to a predetermined
	   computed value if not set (currently 500GB). Note: If SSDs are not
	   attached, it also contains the HDFS data blocks and Hadoop working directories.
	*/
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" yaml:"bootDiskSizeGb,omitempty"`
}
