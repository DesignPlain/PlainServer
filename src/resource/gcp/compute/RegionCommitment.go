package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type RegionCommitment struct {
	/*
	   A list of commitment amounts for particular resources.
	   Note that VCPU and MEMORY resource commitments must occur together.
	   Structure is documented below.
	*/
	Resources []types.Compute_RegionCommitmentResource `json:"resources,omitempty" yaml:"resources,omitempty"`

	/*
	   Specifies whether to enable automatic renewal for the commitment.
	   The default value is false if not specified.
	   If the field is set to true, the commitment will be automatically renewed for either
	   one or three years according to the terms of the existing commitment.
	*/
	AutoRenew bool `json:"autoRenew,omitempty" yaml:"autoRenew,omitempty"`

	/*
	   The category of the commitment. Category MACHINE specifies commitments composed of
	   machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE
	   specifies commitments composed of software licenses, listed in licenseResources.
	   Note that only MACHINE commitments should have a Type specified.
	   Possible values are: `LICENSE`, `MACHINE`.
	*/
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The license specification required as part of a license commitment.
	   Structure is documented below.
	*/
	LicenseResource types.Compute_RegionCommitmentLicenseResource `json:"licenseResource,omitempty" yaml:"licenseResource,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// URL of the region where this commitment may be used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The plan for this commitment, which determines duration and discount rate.
	   The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
	   Possible values are: `TWELVE_MONTH`, `THIRTY_SIX_MONTH`.


	   - - -
	*/
	Plan string `json:"plan,omitempty" yaml:"plan,omitempty"`

	/*
	   The type of commitment, which affects the discount rate and the eligible resources.
	   The type could be one of the following value: `MEMORY_OPTIMIZED`, `ACCELERATOR_OPTIMIZED`,
	   `GENERAL_PURPOSE_N1`, `GENERAL_PURPOSE_N2`, `GENERAL_PURPOSE_N2D`, `GENERAL_PURPOSE_E2`,
	   `GENERAL_PURPOSE_T2D`, `GENERAL_PURPOSE_C3`, `COMPUTE_OPTIMIZED_C2`, `COMPUTE_OPTIMIZED_C2D` and
	   `GRAPHICS_OPTIMIZED_G2`
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
