package types

type Secretsmanager_getSecretRotationRotationRule struct {
	//
	AutomaticallyAfterDays int `json:"automaticallyAfterDays,omitempty" yaml:"automaticallyAfterDays,omitempty"`

	//
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	//
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`
}
