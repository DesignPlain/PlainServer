package types

type Guardduty_DetectorDatasourcesS3Logs struct {
	// Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}
