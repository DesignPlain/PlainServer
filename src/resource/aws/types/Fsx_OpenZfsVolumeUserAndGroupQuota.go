package types

type Fsx_OpenZfsVolumeUserAndGroupQuota struct {
	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The ID of the user or group. Valid values between `0` and `2147483647`
	Id int `json:"id,omitempty" yaml:"id,omitempty"`

	// The amount of storage that the user or group can use in gibibytes (GiB). Valid values between `0` and `2147483647`
	StorageCapacityQuotaGib int `json:"storageCapacityQuotaGib,omitempty" yaml:"storageCapacityQuotaGib,omitempty"`
}
