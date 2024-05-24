package types

type Quicksight_RefreshScheduleScheduleScheduleFrequency struct {
	// The [refresh on entity](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ScheduleRefreshOnEntity.html) configuration for weekly or monthly schedules. See refresh_on_day.
	RefreshOnDay Quicksight_RefreshScheduleScheduleScheduleFrequencyRefreshOnDay `json:"refreshOnDay,omitempty" yaml:"refreshOnDay,omitempty"`

	// The time of day that you want the dataset to refresh. This value is expressed in `HH:MM` format. This field is not required for schedules that refresh hourly.
	TimeOfTheDay string `json:"timeOfTheDay,omitempty" yaml:"timeOfTheDay,omitempty"`

	// The timezone that you want the refresh schedule to use.
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	// The interval between scheduled refreshes. Valid values are `MINUTE15`, `MINUTE30`, `HOURLY`, `DAILY`, `WEEKLY` and `MONTHLY`.
	Interval string `json:"interval,omitempty" yaml:"interval,omitempty"`
}
