package types

type Gamelift_AliasRoutingStrategy struct {
	// Type of routing strategyE.g., `SIMPLE` or `TERMINAL`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// ID of the GameLift Fleet to point the alias to.
	FleetId string `json:"fleetId,omitempty" yaml:"fleetId,omitempty"`

	// Message text to be used with the `TERMINAL` routing strategy.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}
