package iam

import types "libds/gcp/types"

type WorkforcePool struct {
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The location for the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Immutable. The resource name of the parent. Format: `organizations/{org-id}`.


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Duration that the Google Cloud access tokens, console sign-in sessions,
	   and `gcloud` sign-in sessions from this pool are valid.
	   Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	   If `sessionDuration` is not configured, minted credentials have a default duration of one hour (3600s).
	   A duration in seconds with up to nine fractional digits, ending with '`s`'. Example: "`3.5s`".
	*/
	SessionDuration string `json:"sessionDuration,omitempty" yaml:"sessionDuration,omitempty"`

	/*
	   The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
	   digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	   The prefix `gcp-` is reserved for use by Google, and may not be specified.
	*/
	WorkforcePoolId string `json:"workforcePoolId,omitempty" yaml:"workforcePoolId,omitempty"`

	/*
	   Configure access restrictions on the workforce pool users. This is an optional field. If specified web
	   sign-in can be restricted to given set of services or programmatic sign-in can be disabled for pool users.
	   Structure is documented below.
	*/
	AccessRestrictions types.Iam_WorkforcePoolAccessRestrictions `json:"accessRestrictions,omitempty" yaml:"accessRestrictions,omitempty"`

	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
	   or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
