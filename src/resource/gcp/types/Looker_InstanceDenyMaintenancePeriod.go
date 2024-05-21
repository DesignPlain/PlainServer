package types

type Looker_InstanceDenyMaintenancePeriod struct {
	/*
	   Required. Start date of the deny maintenance period
	   Structure is documented below.
	*/
	EndDate Looker_InstanceDenyMaintenancePeriodEndDate `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	/*
	   Required. Start date of the deny maintenance period
	   Structure is documented below.
	*/
	StartDate Looker_InstanceDenyMaintenancePeriodStartDate `json:"startDate,omitempty" yaml:"startDate,omitempty"`

	/*
	   Required. Start time of the window in UTC time.
	   Structure is documented below.
	*/
	Time Looker_InstanceDenyMaintenancePeriodTime `json:"time,omitempty" yaml:"time,omitempty"`
}
