package types

type Finspace_KxClusterSavedownStorageConfiguration struct {
	// Size of temporary storage in gigabytes. Must be between 10 and 16000.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   Type of writeable storage space for temporarily storing your savedown data. The valid values are:
	   - SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The name of the kdb volume that you want to use as writeable save-down storage for clusters.
	VolumeName string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
