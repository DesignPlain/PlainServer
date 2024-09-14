package types

type Ssm_getPatchBaselineApprovalRule struct {
	// Boolean enabling the application of non-security updates.
	EnableNonSecurity bool `json:"enableNonSecurity,omitempty" yaml:"enableNonSecurity,omitempty"`

	// Patch filter group that defines the criteria for the rule.
	PatchFilters []Ssm_getPatchBaselineApprovalRulePatchFilter `json:"patchFilters,omitempty" yaml:"patchFilters,omitempty"`

	// Number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline.
	ApproveAfterDays int `json:"approveAfterDays,omitempty" yaml:"approveAfterDays,omitempty"`

	// Cutoff date for auto approval of released patches. Any patches released on or before this date are installed automatically. Date is formatted as `YYYY-MM-DD`. Conflicts with `approve_after_days`
	ApproveUntilDate string `json:"approveUntilDate,omitempty" yaml:"approveUntilDate,omitempty"`

	// Compliance level for patches approved by this rule.
	ComplianceLevel string `json:"complianceLevel,omitempty" yaml:"complianceLevel,omitempty"`
}
