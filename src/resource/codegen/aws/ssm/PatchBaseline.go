package ssm

import types "libds/aws/types"

type PatchBaseline struct {
	// Compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel string `json:"approvedPatchesComplianceLevel,omitempty" yaml:"approvedPatchesComplianceLevel,omitempty"`

	// Whether the list of approved patches includes non-security updates that should be applied to the instances. Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity bool `json:"approvedPatchesEnableNonSecurity,omitempty" yaml:"approvedPatchesEnableNonSecurity,omitempty"`

	/*
	   Name of the patch baseline.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block with alternate sources for patches. Applies to Linux instances only. See `source` below.
	Sources []types.Ssm_PatchBaselineSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// Set of rules used to include patches in the baseline. Up to 10 approval rules can be specified. See `approval_rule` below.
	ApprovalRules []types.Ssm_PatchBaselineApprovalRule `json:"approvalRules,omitempty" yaml:"approvalRules,omitempty"`

	// List of explicitly approved patches for the baseline. Cannot be specified with `approval_rule`.
	ApprovedPatches []string `json:"approvedPatches,omitempty" yaml:"approvedPatches,omitempty"`

	// Operating system the patch baseline applies to. Valid values are `ALMA_LINUX`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `AMAZON_LINUX_2022`, `AMAZON_LINUX_2023`, `CENTOS`, `DEBIAN`, `MACOS`, `ORACLE_LINUX`, `RASPBIAN`, `REDHAT_ENTERPRISE_LINUX`, `ROCKY_LINUX`, `SUSE`, `UBUNTU`, and `WINDOWS`. The default value is `WINDOWS`.
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`

	// List of rejected patches.
	RejectedPatches []string `json:"rejectedPatches,omitempty" yaml:"rejectedPatches,omitempty"`

	// Action for Patch Manager to take on patches included in the `rejected_patches` list. Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction string `json:"rejectedPatchesAction,omitempty" yaml:"rejectedPatchesAction,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the patch baseline.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters []types.Ssm_PatchBaselineGlobalFilter `json:"globalFilters,omitempty" yaml:"globalFilters,omitempty"`
}
