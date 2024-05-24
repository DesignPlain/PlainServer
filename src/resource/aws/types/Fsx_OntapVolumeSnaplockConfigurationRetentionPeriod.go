package types

type Fsx_OntapVolumeSnaplockConfigurationRetentionPeriod struct {
	// The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an FSx for ONTAP SnapLock volume. The default retention period must be greater than or equal to the minimum retention period and less than or equal to the maximum retention period. See Retention Period below.
	DefaultRetention Fsx_OntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention `json:"defaultRetention,omitempty" yaml:"defaultRetention,omitempty"`

	// The longest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume. See Retention Period below.
	MaximumRetention Fsx_OntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention `json:"maximumRetention,omitempty" yaml:"maximumRetention,omitempty"`

	// The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume. See Retention Period below.
	MinimumRetention Fsx_OntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention `json:"minimumRetention,omitempty" yaml:"minimumRetention,omitempty"`
}
