package types

type Organizations_PolicyListPolicy struct {
	// or `deny` - (Optional) One or the other must be set.
	Allow Organizations_PolicyListPolicyAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	// One or the other must be set.
	Deny Organizations_PolicyListPolicyDeny `json:"deny,omitempty" yaml:"deny,omitempty"`

	/*
	   If set to true, the values from the effective Policy of the parent resource
	   are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.

	   The `allow` or `deny` blocks support:
	*/
	InheritFromParent bool `json:"inheritFromParent,omitempty" yaml:"inheritFromParent,omitempty"`

	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	SuggestedValue string `json:"suggestedValue,omitempty" yaml:"suggestedValue,omitempty"`
}
