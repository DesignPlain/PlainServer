package types

type Lightsail_InstanceAddOn struct {
	// The add-on type. There is currently only one valid type `AutoSnapshot`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The daily time when an automatic snapshot will be created. Must be in HH:00 format, and in an hourly increment and specified in Coordinated Universal Time (UTC). The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTime string `json:"snapshotTime,omitempty" yaml:"snapshotTime,omitempty"`

	// The status of the add on. Valid Values: `Enabled`, `Disabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
