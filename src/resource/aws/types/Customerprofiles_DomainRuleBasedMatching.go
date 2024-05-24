package types

type Customerprofiles_DomainRuleBasedMatching struct {
	// A block that configures how the rule-based matching process should match profiles. You can have up to 15 `rule` in the `natching_rules`. Documented below.
	MatchingRules []Customerprofiles_DomainRuleBasedMatchingMatchingRule `json:"matchingRules,omitempty" yaml:"matchingRules,omitempty"`

	// Indicates the maximum allowed rule level for matching.
	MaxAllowedRuleLevelForMatching int `json:"maxAllowedRuleLevelForMatching,omitempty" yaml:"maxAllowedRuleLevelForMatching,omitempty"`

	// Indicates the maximum allowed rule level for merging.
	MaxAllowedRuleLevelForMerging int `json:"maxAllowedRuleLevelForMerging,omitempty" yaml:"maxAllowedRuleLevelForMerging,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A block that configures information about the `AttributeTypesSelector` where the rule-based identity resolution uses to match profiles. Documented below.
	AttributeTypesSelector Customerprofiles_DomainRuleBasedMatchingAttributeTypesSelector `json:"attributeTypesSelector,omitempty" yaml:"attributeTypesSelector,omitempty"`

	// A block that specifies how the auto-merging process should resolve conflicts between different profiles. Documented below.
	ConflictResolution Customerprofiles_DomainRuleBasedMatchingConflictResolution `json:"conflictResolution,omitempty" yaml:"conflictResolution,omitempty"`

	// The flag that enables the rule-based matching process of duplicate profiles.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A block that specifies the configuration for exporting Identity Resolution results. Documented below.
	ExportingConfig Customerprofiles_DomainRuleBasedMatchingExportingConfig `json:"exportingConfig,omitempty" yaml:"exportingConfig,omitempty"`
}
