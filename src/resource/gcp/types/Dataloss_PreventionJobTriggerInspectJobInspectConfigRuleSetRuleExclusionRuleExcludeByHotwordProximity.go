package types

type Dataloss_PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordProximity struct {
	// Number of characters after the finding to consider. Either this or window_before must be specified
	WindowAfter int `json:"windowAfter,omitempty" yaml:"windowAfter,omitempty"`

	// Number of characters before the finding to consider. Either this or window_after must be specified
	WindowBefore int `json:"windowBefore,omitempty" yaml:"windowBefore,omitempty"`
}
