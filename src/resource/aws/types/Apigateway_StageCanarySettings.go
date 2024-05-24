package types

type Apigateway_StageCanarySettings struct {
	// Percent `0.0` - `100.0` of traffic to divert to the canary deployment.
	PercentTraffic float64 `json:"percentTraffic,omitempty" yaml:"percentTraffic,omitempty"`

	// Map of overridden stage `variables` (including new variables) for the canary deployment.
	StageVariableOverrides map[string]string `json:"stageVariableOverrides,omitempty" yaml:"stageVariableOverrides,omitempty"`

	// Whether the canary deployment uses the stage cache. Defaults to false.
	UseStageCache bool `json:"useStageCache,omitempty" yaml:"useStageCache,omitempty"`
}
