package types

type Fsx_OpenZfsVolumeUserAndGroupQuota struct {
	// The ID of the user or group. Valid values between `0` and `2147483647`
	Id int `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   The amount of storage that the user or group can use in gibibytes (GiB). Valid values between `0` and `2147483647`
	   - `Type` - (Required) - A value that specifies whether the quota applies to a user or group. Valid values are `USER` or `GROUP`.
	*/
	StorageCapacityQuotaGib int `json:"storageCapacityQuotaGib,omitempty" yaml:"storageCapacityQuotaGib,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
