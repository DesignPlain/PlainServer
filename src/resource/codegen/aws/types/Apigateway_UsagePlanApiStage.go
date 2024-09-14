package types

type Apigateway_UsagePlanApiStage struct {
	// API Id of the associated API stage in a usage plan.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// API stage name of the associated API stage in a usage plan.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`

	// The throttling limits of the usage plan.
	Throttles []Apigateway_UsagePlanApiStageThrottle `json:"throttles,omitempty" yaml:"throttles,omitempty"`
}
