package types

type Elasticsearch_getDomainAutoTuneOptionMaintenanceSchedule struct {
	// Cron expression for an Auto-Tune maintenance schedule.
	CronExpressionForRecurrence string `json:"cronExpressionForRecurrence,omitempty" yaml:"cronExpressionForRecurrence,omitempty"`

	// Configuration block for the duration of the Auto-Tune maintenance window.
	Durations []Elasticsearch_getDomainAutoTuneOptionMaintenanceScheduleDuration `json:"durations,omitempty" yaml:"durations,omitempty"`

	// Date and time at which the Auto-Tune maintenance schedule starts in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	StartAt string `json:"startAt,omitempty" yaml:"startAt,omitempty"`
}
