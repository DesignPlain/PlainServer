package efs

type MountTarget struct {
	// The ID of the file system for which the mount target is intended.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	/*
	   The address (within the address range of the specified subnet) at
	   which the file system may be mounted via the mount target.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   A list of up to 5 VPC security group IDs (that must
	   be for the same VPC as subnet specified) in effect for the mount target.
	*/
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The ID of the subnet to add the mount target in.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
