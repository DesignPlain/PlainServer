package iam

import types "libds/gcp/types"

type DenyPolicy struct {
	// The display name of the rule.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The attachment point is identified by its URL-encoded full resource name.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Rules to be applied.
	   Structure is documented below.
	*/
	Rules []types.Iam_DenyPolicyRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
