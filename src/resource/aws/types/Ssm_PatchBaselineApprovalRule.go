package types

type Ssm_PatchBaselineApprovalRule struct {
	// Patch filter group that defines the criteria for the rule. Up to 5 patch filters can be specified per approval rule using Key/Value pairs. Valid combinations of these Keys and the `operating_system` value can be found in the [SSM DescribePatchProperties API Reference](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribePatchProperties.html). Valid Values are exact values for the patch property given as the key, or a wildcard `-`, which matches all values. `PATCH_SET` defaults to `OS` if unspecified
	PatchFilters []Ssm_PatchBaselineApprovalRulePatchFilter `json:"patchFilters,omitempty" yaml:"patchFilters,omitempty"`

	// Number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline. Valid Range: 0 to 100. Conflicts with `approve_until_date`.
	ApproveAfterDays int `json:"approveAfterDays,omitempty" yaml:"approveAfterDays,omitempty"`

	// Cutoff date for auto approval of released patches. Any patches released on or before this date are installed automatically. Date is formatted as `YYYY-MM-DD`. Conflicts with `approve_after_days`
	ApproveUntilDate string `json:"approveUntilDate,omitempty" yaml:"approveUntilDate,omitempty"`

	// Compliance level for patches approved by this rule. Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, and `UNSPECIFIED`. The default value is `UNSPECIFIED`.
	ComplianceLevel string `json:"complianceLevel,omitempty" yaml:"complianceLevel,omitempty"`

	// Boolean enabling the application of non-security updates. The default value is `false`. Valid for Linux instances only.
	EnableNonSecurity bool `json:"enableNonSecurity,omitempty" yaml:"enableNonSecurity,omitempty"`
}
