package types

type Fsx_OntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention struct {
	// The amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`

	// The type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume. Setting this value to `NONE` disables autocommit. Valid values: `MINUTES`, `HOURS`, `DAYS`, `MONTHS`, `YEARS`, `NONE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
