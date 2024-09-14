package gamelift

type MatchmakingRuleSet struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the matchmaking rule set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// JSON encoded string containing rule set data.
	RuleSetBody string `json:"ruleSetBody,omitempty" yaml:"ruleSetBody,omitempty"`
}
