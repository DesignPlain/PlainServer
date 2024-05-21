package iam

type WorkloadIdentityPool struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	   existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	   access again.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The ID to use for the pool, which becomes the final component of the resource name. This
	   value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	   `gcp-` is reserved for use by Google, and may not be specified.


	   - - -
	*/
	WorkloadIdentityPoolId string `json:"workloadIdentityPoolId,omitempty" yaml:"workloadIdentityPoolId,omitempty"`
}
