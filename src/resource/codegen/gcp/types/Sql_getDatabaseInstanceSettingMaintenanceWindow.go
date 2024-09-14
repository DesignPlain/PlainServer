package types

type Sql_getDatabaseInstanceSettingMaintenanceWindow struct {
	// Receive updates earlier (canary) or later (stable)
	UpdateTrack string `json:"updateTrack,omitempty" yaml:"updateTrack,omitempty"`

	// Day of week (1-7), starting on Monday
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// Hour of day (0-23), ignored if day not set
	Hour int `json:"hour,omitempty" yaml:"hour,omitempty"`
}
