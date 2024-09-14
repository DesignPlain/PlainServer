package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesBotControlRuleSet struct {
	// Applies only to the targeted inspection level. Determines whether to use machine learning (ML) to analyze your web traffic for bot-related activity. Defaults to `true`.
	EnableMachineLearning bool `json:"enableMachineLearning,omitempty" yaml:"enableMachineLearning,omitempty"`

	// The inspection level to use for the Bot Control rule group.
	InspectionLevel string `json:"inspectionLevel,omitempty" yaml:"inspectionLevel,omitempty"`
}
