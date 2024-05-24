package types

type Ecs_ServiceDeploymentCircuitBreaker struct {
	// Whether to enable the deployment circuit breaker logic for the service.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// Whether to enable Amazon ECS to roll back the service if a service deployment fails. If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	Rollback bool `json:"rollback,omitempty" yaml:"rollback,omitempty"`
}
