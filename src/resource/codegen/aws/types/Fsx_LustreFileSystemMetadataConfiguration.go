package types

type Fsx_LustreFileSystemMetadataConfiguration struct {
	// Amount of IOPS provisioned for metadata. This parameter should only be used when the mode is set to `USER_PROVISIONED`. Valid Values are `1500`,`3000`,`6000` and `12000` through `192000` in increments of `12000`.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	/*
	   Mode for the metadata configuration of the file system. Valid values are `AUTOMATIC`, and `USER_PROVISIONED`.

	   !> --WARNING:-- Updating the value of `iops` from a higher to a lower value will force a recreation of the resource. Any data on the file system will be lost when recreating.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
