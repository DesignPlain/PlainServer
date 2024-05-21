package types

type Sql_getDatabaseInstanceSettingDenyMaintenancePeriod struct {
	// End date before which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01
	EndDate string `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	// Start date after which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01
	StartDate string `json:"startDate,omitempty" yaml:"startDate,omitempty"`

	// Time in UTC when the "deny maintenance period" starts on start_date and ends on end_date. The time is in format: HH:mm:SS, i.e., 00:00:00
	Time string `json:"time,omitempty" yaml:"time,omitempty"`
}
