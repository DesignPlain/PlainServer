package types

type Dataplex_DatascanExecutionSpecTriggerSchedule struct {
	/*
	   Cron schedule for running scans periodically. This field is required for Schedule scans.

	   - - -
	*/
	Cron string `json:"cron,omitempty" yaml:"cron,omitempty"`
}
