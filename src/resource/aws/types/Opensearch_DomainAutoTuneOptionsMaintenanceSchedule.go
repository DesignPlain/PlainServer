package types

type Opensearch_DomainAutoTuneOptionsMaintenanceSchedule struct {
	// A cron expression specifying the recurrence pattern for an Auto-Tune maintenance schedule.
	CronExpressionForRecurrence string `json:"cronExpressionForRecurrence,omitempty" yaml:"cronExpressionForRecurrence,omitempty"`

	// Configuration block for the duration of the Auto-Tune maintenance window. Detailed below.
	Duration Opensearch_DomainAutoTuneOptionsMaintenanceScheduleDuration `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Date and time at which to start the Auto-Tune maintenance schedule in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	StartAt string `json:"startAt,omitempty" yaml:"startAt,omitempty"`
}
