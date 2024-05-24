package types

type Ecs_ServiceAlarms struct {
	// One or more CloudWatch alarm names.
	AlarmNames []string `json:"alarmNames,omitempty" yaml:"alarmNames,omitempty"`

	// Determines whether to use the CloudWatch alarm option in the service deployment process.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails. If rollback is used, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	Rollback bool `json:"rollback,omitempty" yaml:"rollback,omitempty"`
}
