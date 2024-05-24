package types

type Quicksight_RefreshScheduleSchedule struct {
	// The type of refresh that the dataset undergoes. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
	RefreshType string `json:"refreshType,omitempty" yaml:"refreshType,omitempty"`

	// The configuration of the [schedule frequency](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RefreshFrequency.html). See schedule_frequency.
	ScheduleFrequency Quicksight_RefreshScheduleScheduleScheduleFrequency `json:"scheduleFrequency,omitempty" yaml:"scheduleFrequency,omitempty"`

	// Time after which the refresh schedule can be started, expressed in `YYYY-MM-DDTHH:MM:SS` format.
	StartAfterDateTime string `json:"startAfterDateTime,omitempty" yaml:"startAfterDateTime,omitempty"`
}
