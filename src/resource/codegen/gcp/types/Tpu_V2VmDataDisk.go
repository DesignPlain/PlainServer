package types

type Tpu_V2VmDataDisk struct {
	/*
	   The mode in which to attach this disk. If not specified, the default is READ_WRITE
	   mode. Only applicable to dataDisks.
	   Default value is `READ_WRITE`.
	   Possible values are: `READ_WRITE`, `READ_ONLY`.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   Specifies the full path to an existing disk. For example:
	   "projects/my-project/zones/us-central1-c/disks/my-disk".
	*/
	SourceDisk string `json:"sourceDisk,omitempty" yaml:"sourceDisk,omitempty"`
}
