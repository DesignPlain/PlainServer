package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordProximity struct {
	// Number of characters after the finding to consider.
	WindowAfter int `json:"windowAfter,omitempty" yaml:"windowAfter,omitempty"`

	// Number of characters before the finding to consider.
	WindowBefore int `json:"windowBefore,omitempty" yaml:"windowBefore,omitempty"`
}
