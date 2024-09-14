package types

type Synthetics_CanarySchedule struct {
	// Rate expression or cron expression that defines how often the canary is to run. For rate expression, the syntax is `rate(number unit)`. _unit_ can be `minute`, `minutes`, or `hour`. For cron expression, the syntax is `cron(expression)`. For more information about the syntax for cron expressions, see [Scheduling canary runs using cron](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html).
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// Duration in seconds, for the canary to continue making regular runs according to the schedule in the Expression value.
	DurationInSeconds int `json:"durationInSeconds,omitempty" yaml:"durationInSeconds,omitempty"`
}
