package types

type Sagemaker_MonitoringScheduleMonitoringScheduleConfigScheduleConfig struct {
	// A cron expression that describes details about the monitoring schedule. For example, and hourly schedule would be `cron(0 - ? - - -)`.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`
}
