package types

type Fsx_OntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention struct {
	// The type of time for the retention period of an FSx for ONTAP SnapLock volume. Set it to one of the valid types. If you set it to `INFINITE`, the files are retained forever. If you set it to `UNSPECIFIED`, the files are retained until you set an explicit retention period. Valid values: `SECONDS`, `MINUTES`, `HOURS`, `DAYS`, `MONTHS`, `YEARS`, `INFINITE`, `UNSPECIFIED`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
