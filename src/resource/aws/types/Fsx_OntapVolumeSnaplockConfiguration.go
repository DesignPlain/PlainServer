package types

type Fsx_OntapVolumeSnaplockConfiguration struct {
	// Specifies the retention mode of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed. Valid values: `COMPLIANCE`, `ENTERPRISE`.
	SnaplockType string `json:"snaplockType,omitempty" yaml:"snaplockType,omitempty"`

	// Enables or disables volume-append mode on an FSx for ONTAP SnapLock volume. The default value is `false`.
	VolumeAppendModeEnabled bool `json:"volumeAppendModeEnabled,omitempty" yaml:"volumeAppendModeEnabled,omitempty"`

	// Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume. The default value is `false`.
	AuditLogVolume bool `json:"auditLogVolume,omitempty" yaml:"auditLogVolume,omitempty"`

	// The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume. See Autocommit Period below.
	AutocommitPeriod Fsx_OntapVolumeSnaplockConfigurationAutocommitPeriod `json:"autocommitPeriod,omitempty" yaml:"autocommitPeriod,omitempty"`

	// Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock Enterprise volume. Valid values: `DISABLED`, `ENABLED`, `PERMANENTLY_DISABLED`. The default value is `DISABLED`.
	PrivilegedDelete string `json:"privilegedDelete,omitempty" yaml:"privilegedDelete,omitempty"`

	// The retention period of an FSx for ONTAP SnapLock volume. See SnapLock Retention Period below.
	RetentionPeriod Fsx_OntapVolumeSnaplockConfigurationRetentionPeriod `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`
}
