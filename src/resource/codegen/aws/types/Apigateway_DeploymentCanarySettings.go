package types

type Apigateway_DeploymentCanarySettings struct {
	// Stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.
	StageVariableOverrides map[string]string `json:"stageVariableOverrides,omitempty" yaml:"stageVariableOverrides,omitempty"`

	// Boolean flag to indicate whether the canary release deployment uses the stage cache or not.
	UseStageCache bool `json:"useStageCache,omitempty" yaml:"useStageCache,omitempty"`

	// Percentage (0.0-100.0) of traffic routed to the canary deployment.
	PercentTraffic float64 `json:"percentTraffic,omitempty" yaml:"percentTraffic,omitempty"`
}
