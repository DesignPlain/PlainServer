package types

type Gamelift_FleetResourceCreationLimitPolicy struct {
	// Maximum number of game sessions that an individual can create during the policy period.
	NewGameSessionsPerCreator int `json:"newGameSessionsPerCreator,omitempty" yaml:"newGameSessionsPerCreator,omitempty"`

	// Time span used in evaluating the resource creation limit policy.
	PolicyPeriodInMinutes int `json:"policyPeriodInMinutes,omitempty" yaml:"policyPeriodInMinutes,omitempty"`
}
