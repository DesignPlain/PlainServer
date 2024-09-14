package compute

import types "libds/gcp/types"

type SecurityPolicy struct {
	// The type indicates the intended use of the security policy. This field can be set only at resource creation time.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
	AdaptiveProtectionConfig types.Compute_SecurityPolicyAdaptiveProtectionConfig `json:"adaptiveProtectionConfig,omitempty" yaml:"adaptiveProtectionConfig,omitempty"`

	/*
	   [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
	   Structure is documented below.
	*/
	AdvancedOptionsConfig types.Compute_SecurityPolicyAdvancedOptionsConfig `json:"advancedOptionsConfig,omitempty" yaml:"advancedOptionsConfig,omitempty"`

	// An optional description of this security policy. Max size is 2048.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the security policy.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
	RecaptchaOptionsConfig types.Compute_SecurityPolicyRecaptchaOptionsConfig `json:"recaptchaOptionsConfig,omitempty" yaml:"recaptchaOptionsConfig,omitempty"`

	/*
	   The set of rules that belong to this policy. There must always be a default
	   rule (rule with priority 2147483647 and match "\-"). If no rules are provided when creating a
	   security policy, a default rule with action "allow" will be added. Structure is documented below.
	*/
	Rules []types.Compute_SecurityPolicyRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
