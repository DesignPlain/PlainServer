package types

type Ecs_ServiceDeploymentController struct {
	// Type of deployment controller. Valid values: `CODE_DEPLOY`, `ECS`, `EXTERNAL`. Default: `ECS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
