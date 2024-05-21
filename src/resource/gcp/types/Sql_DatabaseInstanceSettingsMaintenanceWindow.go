package types

type Sql_DatabaseInstanceSettingsMaintenanceWindow struct {
	// Day of week (`1-7`), starting on Monday
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// Hour of day (`0-23`), ignored if `day` not set
	Hour int `json:"hour,omitempty" yaml:"hour,omitempty"`

	/*
	   Receive updates earlier (`canary`) or later
	   (`stable`)

	   The optional `settings.insights_config` subblock for instances declares Query Insights([MySQL](https://cloud.google.com/sql/docs/mysql/using-query-insights), [PostgreSQL](https://cloud.google.com/sql/docs/postgres/using-query-insights)) configuration. It contains:
	*/
	UpdateTrack string `json:"updateTrack,omitempty" yaml:"updateTrack,omitempty"`
}
