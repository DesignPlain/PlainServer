package types

type Compute_SecurityPolicyRuleRedirectOptions struct {
	// External redirection target when `EXTERNAL_302` is set in `type`.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// Type of redirect action.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
