package types

type Vpclattice_ListenerRuleActionForward struct {
	/*
	   The target groups. Traffic matching the rule is forwarded to the specified target groups. With forward actions, you can assign a weight that controls the prioritization and selection of each target group. This means that requests are distributed to individual target groups based on their weights. For example, if two target groups have the same weight, each target group receives half of the traffic.

	   The default value is 1 with maximum number of 2. If only one target group is provided, there is no need to set the weight; 100%!o(MISSING)f traffic will go to that target group.
	*/
	TargetGroups []Vpclattice_ListenerRuleActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
